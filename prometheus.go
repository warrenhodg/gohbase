package gohbase

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var operationDurationSeconds = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: "gohbase",
		Name:      "operation_duration_seconds",
		Help:      "Time in seconds for operation to complete",
		Buckets:   prometheus.ExponentialBuckets(0.001, 4, 10),
	},
	[]string{"operation", "result"},
)
