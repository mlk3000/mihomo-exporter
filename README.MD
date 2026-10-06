# mihomo-exporter

A tiny Prometheus exporter for mihomo written in Go.

## Run

go build -o mihomo-exporter .

./mihomo-exporter \
    -mihomo=http://127.0.0.1:9090 \
    -listen=:9101 \
    -interval=3s

## Metrics
- mihomo_active_connections
- mihomo_download_bytes_total 
- mihomo_upload_bytes_total

## Endpoints
- /metrics
- /healthz