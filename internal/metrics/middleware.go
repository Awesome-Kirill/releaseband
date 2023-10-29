package metrics

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/urfave/negroni"
)

type Middleware struct {
	client *PrometheusClient
}

func New() *Middleware {
	return &Middleware{client: NewPrometheusClient()}
}
func (m *Middleware) After(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		lrw := negroni.NewResponseWriter(w)
		h.ServeHTTP(lrw, r)

		statusCode := lrw.Status()
		m.client.IncrHTTPServerRequests(
			prometheus.Labels{
				"http_status_code": strconv.Itoa(statusCode),
			},

			time.Since(startTime),
		)
	})
}
