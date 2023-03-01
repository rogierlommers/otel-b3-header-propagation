package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const hostPort = "127.0.0.1:8765"

var b3Headers = []string{"traceparent", "b3", "x-b3-traceid", "x-b3-parentspanid", "x-b3-spanid", "x-b3-sampled", "x-b3-flags"}

func main() {

	// create server
	server := newServer()

	// listen from here
	logrus.Infof("listening on http://%s/dump-request", hostPort)

	// start listening
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Fatalf("Could not bind: %v", err)
	}

}

func requestDumper(w http.ResponseWriter, r *http.Request) {
	for name, value := range r.Header {
		if isB3Header(name) {
			logrus.Infof("✅ header: %s, value: %s", name, value)
		} else {
			logrus.Infof("❌ header: %s, value: %s", name, value)
		}

	}
	w.Write([]byte("ok"))
}

func newServer() *http.Server {

	// create router
	router := mux.NewRouter().StrictSlash(true)

	// add route
	router.HandleFunc("/dump-request", requestDumper)

	return &http.Server{
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Addr:         hostPort,
		Handler:      router,
	}

}

func isB3Header(h string) bool {

	for _, b := range b3Headers {
		if strings.ToLower(h) == b {
			return true
		}
	}

	return false
}
