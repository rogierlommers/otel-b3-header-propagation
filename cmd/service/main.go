package main

import (
	"net/http"
	"rogierlommers/otel-b3-header-propagation/api"

	"github.com/sirupsen/logrus"
)

func main() {

	server := api.NewServer()

	// start listening
	logrus.Infof("listening on http://127.0.0.1:8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Fatalf("Could not bind on %d: %v\n", 8080, err)
	}

}
