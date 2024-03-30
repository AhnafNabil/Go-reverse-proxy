# Go-reverse-proxy

### Run locally:
1. Run each server individually
```bash
cd origin-server-01
go run main.go
```
```bash
cd origin-server-02
go run main.go
```
```bash
cd proxy-server
go run main.go
```

### Run using Docker file:
1. Docker build:
```bash
docker build -t origin1 .
docker build -t origin2 .
docker build -t proxy .
```
2. Run docker images:
```bash
docker run -p 8081:8081 -t origin1
docker run -p 8082:8082 -t origin2
docker run -p 8080:8080 -t proxy
```
