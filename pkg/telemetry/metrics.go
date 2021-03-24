package telemetry

import (
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/lichuan0620/playground/pkg/version"
)

var (
	buildInfo = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: version.Name,
			Name:      "build_info",
			Help:      "Build and version information",
			ConstLabels: prometheus.Labels{
				"version":    version.Version,
				"go_version": runtime.Version(),
				"branch":     version.Branch,
				"commit":     version.Commit,
				"build_date": version.BuildDate,
			},
		},
	)
	startTime = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace:   version.Name,
			Name:        "start_time_seconds",
			Help:        "Start time of the service in unix timestamp",
		},
	)
)

func init() {
	prometheus.DefaultRegisterer.MustRegister(
		buildInfo,
		startTime,
	)
	buildInfo.Set(1)
	startTime.Set(float64(time.Now().Unix()))
}
