# Go-reverse-proxy

### Run locally:
1. Run each server individually
```bash
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
