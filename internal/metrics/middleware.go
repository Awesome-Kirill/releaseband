package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/urfave/negroni"
	"net/http"
	"strconv"
	"time"
)

type MetricsMiddleware struct {
	client *PrometheusClient
}

func New() *MetricsMiddleware {
	return &MetricsMiddleware{client: NewPrometheusClient()}
}
func (m *MetricsMiddleware) Middleware(h http.Handler) http.Handler {
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
		m.client.TotalHttpCount.Inc()
	})
}
