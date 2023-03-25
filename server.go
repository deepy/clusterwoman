package main

import (
	"clusterwoman/cfg"
	"clusterwoman/cloudprovider"
	"github.com/felixge/httpsnoop"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	externalv1 "github.com/deepy/externalgrpc-connect/gen/clusterautoscaler/cloudprovider/v1/externalgrpc/externalgrpcconnect"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	log.Println("Starting server!")
	conf := cfg.Conf{}
	res, err := conf.GetConfFile(getEnv("CONFIG", "/cfg/prod.yaml"))
	if err != nil {
		log.Fatal(err)
	}
	provider := &cloudprovider.CloudProviderServer{Nodes: res.Nodes}
	mux := http.NewServeMux()
	path, handler := externalv1.NewCloudProviderHandler(provider)
	mux.Handle(path, handler)
	loggedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(mux, w, r)
		log.Printf(
			"%s %s (code=%d dt=%s written=%d)",
			r.Method,
			r.URL,
			m.Code,
			m.Duration,
			m.Written,
		)
	})
	log.Println("ListenAndServe")
	err = http.ListenAndServe(
		":8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(loggedHandler, &http2.Server{}),
	)
	if err != nil {
		log.Fatal(err)
	}
}
