package backend

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func CallBackendHandler(w http.ResponseWriter, r *http.Request) {

	// ##################################################################################################################
	// I would expect that the otel http package takes care of propagation of all B3 headers.
	// this doesn't happen and I cannot figure out why
	// ##################################################################################################################

	resp, err := otelhttp.Get(r.Context(), "http://127.0.0.1:8765/dump-request")
	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.Infof("respose; %s", resp.Status)
}
