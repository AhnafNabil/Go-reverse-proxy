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

### Access origin-server from proxy: 
1. For origin server 01 :
```bash
localhost:8080/origin1
```
2. For origin server 02:
```bash
localhost:8080/origin2
```

# Go-reverse-proxy by pulling two docker images

### 1. Pull the following docker images 
```bash
docker pull poridhi/codeserver-node:v1.1
```
```bash
docker pull poridhi/codeserver-python:v1.2
```


### 2. Run the code-server containers
```bash
docker run -p 8081:8080 poridhi/codeserver-python:v1.2
```

```bash
docker run -p 8082:8080 poridhi/codeserver-node:v1.1
```


### 3. Build a docker image of the proxy-server

```bash
cd proxy-server
docker build -t proxy .
```
### 4. Run a container based on that image

```bash
docker run -p 8080:8080 proxy
```

### 5. Check the codeserver

i. Visit http://localhost:8081 to check the python code-server </br>
ii. Visit http://localhost:8082 to check the node code-server

### 6. Finally visit the following links to check the proxy server

Visit: http://localhost:8080/origin1/
</br>
This will redirect to python code-server
</br>
Visit: http://localhost:8080/origin2/ 
</br>
This will redirect to node code-server


