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
1. Docker build and run docker images:
```bash
cd origin-server-01
docker build -t origin1 .
docker run -p 8081:8081 -t origin1
```
```bash
cd origin-server-02
docker build -t origin2 .
docker run -p 8082:8082 -t origin2
```
```bash
cd proxy-server
docker build -t proxy .
docker run -p 8080:8080 -t proxy
```
