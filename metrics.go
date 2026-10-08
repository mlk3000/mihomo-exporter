package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var activeConnections = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "mihomo_active_connections",
		Help: "Current number of active mihomo connections",
	},
)

var uploadTotal = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "mihomo_upload_bytes_total",
		Help: "Total bytes uploaded by mihomo",
	},
)

var downloadTotal = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "mihomo_download_bytes_total",
		Help: "Total bytes downloaded by mihomo",
	},
)

var memoryUsage = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "mihomo_memory_usage",
		Help: "Memory usage measured in bytes",
	},
)

func updateConnectionMetrics(ch <-chan ConnectionsResponse) {
	for connections := range ch {
		activeConnections.Set(float64(len(connections.Connections)))
		uploadTotal.Set(float64(connections.UploadTotal))
		downloadTotal.Set(float64(connections.DownloadTotal))
	}

}

func updateMemMetrics(ch <-chan float64) {
	for MemUsage := range ch {
		memoryUsage.Set(float64(MemUsage))
	}
}
