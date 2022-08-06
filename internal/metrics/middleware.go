package metrics

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func HTTPDurationMiddleware(next http.Handler) http.Handler {
	var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_response_duration_seconds",
		Help:    "Duration of HTTP requests.",
		Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
	}, []string{"path"})

	_ = prometheus.Register(httpDuration)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		path := getRoutePattern(r)
		httpDuration.WithLabelValues(path).Observe(duration.Seconds())
	})
}

func getRoutePattern(r *http.Request) string {
	reqContext := chi.RouteContext(r.Context())
	if pattern := reqContext.RoutePattern(); pattern != "" {
		return pattern
	}

	return "undefined"
}
