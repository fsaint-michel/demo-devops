// Package metrics ...
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// HelloOK represent the total number of success hello
	HelloOK = promauto.NewCounter(prometheus.CounterOpts{
		Name: "demo_devops_hello_success_total",
		Help: "The total number of success hello",
	})
)

var (
	// HelloNOK represent the total number of failed hello
	HelloNOK = promauto.NewCounter(prometheus.CounterOpts{
		Name: "demo_devops_hello_failed_total",
		Help: "The total number of failed hello",
	})
)
