package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	targets := map[string]string{
		"origin1": "http://host.docker.internal:8081",
		"origin2": "http://host.docker.internal:8082",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		splitPaths := strings.SplitN(r.URL.Path, "/", 3)
		namespace := splitPaths[1]
		target, _ := targets[namespace]

		remote, err := url.Parse(target)
		if err != nil {
			log.Fatal(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		director := proxy.Director
		proxy.Director = func(req *http.Request) {
			director(req)
			log.Println(req.URL.Path)

			splitPath := strings.SplitN(req.URL.Path, "/", 3)
			log.Println(splitPath)
			if len(splitPath) > 2 {
				req.URL.Path = "/" + splitPath[2]
			} else {
				req.URL.Path = "/"
			}

			req.Header.Set("Host", req.Host)
			req.Header.Set("X-Forwarded-Host", req.Host)
			req.Header.Set("X-Forwarded-For", req.RemoteAddr)
			req.Header.Set("X-Forwarded-Proto", req.URL.Scheme)
			req.Header.Set("X-Real-IP", req.RemoteAddr)
		}

		r.URL.Host = remote.Host
		r.URL.Scheme = remote.Scheme
		proxy.ServeHTTP(w, r)
	})
	log.Println("Server is running on port 8080...")

	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
