package main

import (
	"encoding/json"
	"fmt"
)

type ConnectionsResponse struct {
	DownloadTotal int64        `json:"downloadTotal"`
	UploadTotal   int64        `json:"uploadTotal"`
	Connections   []Connection `json:"connections"`
}
type Connection struct {
	ID       string `json:"id"`
	Upload   int64  `json:"upload"`
	Download int64  `json:"download"`
}

type VersionResponse struct {
	Meta    bool   `json:"meta"`
	Version string `json:"version"`
}

func (v VersionResponse) PrintVer() {
	fmt.Println("mihomo version: v", v.Version)
	fmt.Println("meta: ", v.Meta)
}

func (c ConnectionsResponse) PrintStats() {
	fmt.Println("connections:", len(c.Connections))
	fmt.Println("upload:", c.UploadTotal)
	fmt.Println("download:", c.DownloadTotal)
	// for i, conn := range c.Connections[0:3] {
	// 	fmt.Printf("%d\nID:%s\nUpload:%d\nDownload:%d\n\n", i, conn.Id, conn.Upload, conn.Download)
	// }
}

func getConnections(baseURL string) (ConnectionsResponse, error) {
	resp, err := customClient.Get(baseURL + "/connections")
	if err != nil {
		return ConnectionsResponse{}, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		return ConnectionsResponse{}, fmt.Errorf("mihomo returned status %s", resp.Status)
	}
	defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	connections := ConnectionsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&connections)
	if err != nil {
		return ConnectionsResponse{}, err
	}
	return connections, nil
}

func getVersion(baseURL string) (VersionResponse, error) {
	resp, err := customClient.Get(baseURL + "/version")
	if err != nil {
		return VersionResponse{}, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		return VersionResponse{}, fmt.Errorf("mihomo returned status %s", resp.Status)
	}
	defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	version := VersionResponse{}
	err = json.NewDecoder(resp.Body).Decode(&version)
	if err != nil {
		return VersionResponse{}, err
	}
	return version, nil
}
