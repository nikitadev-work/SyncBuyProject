package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	grpcRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_request_total",
			Help: "Total number of gRPC requests",
		},
		[]string{"method"},
	)
)

func InitMetrics() {
	prometheus.MustRegister(grpcRequestsTotal)
}

func IncGRPCRequestsTotal(method string) {
	grpcRequestsTotal.WithLabelValues(method).Inc()
}
