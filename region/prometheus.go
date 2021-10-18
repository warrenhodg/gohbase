// Copyright (C) 2021  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package region

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	batchWaitDurationSeconds = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "gohbase",
			Name:      "batch_wait_duration_seconds",
			Help:      "Time in seconds before batch is sent",
			Buckets:   prometheus.ExponentialBuckets(0.001, 2, 10),
		},
	)

	flushReasonCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "gohbase",
			Name:      "batch_flush_count",
			Help:      "Number of times a gohbase batch was flushed",
		},
		[]string{"reason"},
	)
)
