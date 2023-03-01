package api

import (
	"net/http"
	"rogierlommers/otel-b3-header-propagation/api/backend"
	"time"

	"github.com/gorilla/mux"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// NewServer creates the API server
func NewServer() *http.Server {

	// Register the B3 propagator globally.
	p := b3.New()

	// instrument init
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			p,
			propagation.TraceContext{},
			propagation.Baggage{}),
	)

	// create router
	router := mux.NewRouter().StrictSlash(true)

	// add API endpoints
	router.Handle("/call-downstream/",
		otelhttp.NewHandler(
			http.HandlerFunc(backend.CallBackendHandler),
			"call-downstream",
		))

	return &http.Server{
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Addr:         "127.0.0.1:8080",
		Handler:      router,
	}

}
