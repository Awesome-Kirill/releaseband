package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type prometheusClientConfig struct {
	Buckets []float64
}

type PrometheusClient struct {
	httpServerRequest         *prometheus.CounterVec
	httpServerRequestDuration *prometheus.HistogramVec
	TotalHTTPCount            prometheus.Counter
}

func (p *PrometheusClient) IncrHTTPServerRequests(labels map[string]string, duration time.Duration) {
	p.httpServerRequest.With(labels).Inc()

	p.httpServerRequestDuration.With(labels).Observe(duration.Seconds())
}

var httpServerLabels = []string{
	"http_status_code",
}

func NewPrometheusClient() *PrometheusClient {
	cfg := prometheusClientConfig{Buckets: prometheus.DefBuckets}

	serverRequest := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "http_server",
			Subsystem:   "",
			Name:        "request",
			Help:        "The total number of server requests",
			ConstLabels: nil,
		},
		httpServerLabels,
	)

	serverRequestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace:   "http_server",
			Subsystem:   "",
			Name:        "duration",
			Help:        "The duration of server requests in seconds",
			ConstLabels: nil,
			Buckets:     cfg.Buckets,
		},
		httpServerLabels,
	)

	totalHTTPCount := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_request_total",
	})
	pc := &PrometheusClient{
		httpServerRequest:         serverRequest,
		httpServerRequestDuration: serverRequestDuration,
		TotalHTTPCount:            totalHTTPCount,
	}

	prometheus.MustRegister(
		pc.httpServerRequest,
		pc.httpServerRequestDuration,
		totalHTTPCount,
	)

	return pc
}
