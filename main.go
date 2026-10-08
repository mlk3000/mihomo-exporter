package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func collectConnections(baseURL string, interval time.Duration, ch chan<- ConnectionsResponse) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		connections, err := getConnections(baseURL)
		if err != nil {
			fmt.Println("Collection error! ", err)
			continue
		}
		ch <- connections
	}
}

func collectMemUsage(baseURL string, interval time.Duration, ch chan<- float64) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		memUsage, err := getMemoryUsage(baseURL)
		if err != nil {
			fmt.Println("Collection error! ", err)
			continue
		}
		ch <- memUsage
	}
}

var customClient = &http.Client{
	Timeout: 5 * time.Second,
}

// func worker(ch chan int) {
// 	for i := 0; ; i++ {
// 		ch <- i
// 		time.Sleep(time.Second)
// 	}
// }

var mihomoAddr = flag.String("mihomo", "http://127.0.0.1:9090", "mihomo API address")
var listenAddr = flag.String("listen", ":9101", "exporter listen address")
var interval = flag.Duration("interval", 3*time.Second, "collection interval")

func main() {

	flag.Parse()

	version, err := getVersion(*mihomoAddr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("meta: %v\nversion: %s\n", version.Meta, version.Version)

	ch := make(chan ConnectionsResponse)
	chMem := make(chan float64)

	go collectConnections(*mihomoAddr, *interval, ch)
	go updateConnectionMetrics(ch)

	go collectMemUsage(*mihomoAddr, *interval, chMem)
	go updateMemMetrics(chMem)

	prometheus.MustRegister(
		activeConnections,
		uploadTotal,
		downloadTotal,
		memoryUsage,
	)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("ok"))
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})
	err = http.ListenAndServe(*listenAddr, nil)
	if err != nil {
		panic(err)
	}
	// for connections := range ch {
	// 	time.Sleep(10 * time.Second)
	// 	connections.PrintStats()
	// }
}
