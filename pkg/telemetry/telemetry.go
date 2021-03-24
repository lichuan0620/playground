package telemetry

import (
	"net/http"
	"net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Options struct {
	EnableProfiling bool
	EnableMetrics   bool
	BindAddress     string
}

func DefaultOptions() *Options {
	return &Options{
		EnableProfiling: true,
		EnableMetrics:   true,
		BindAddress:     "0.0.0.0:8081",
	}
}

func ServeTelemetry(options *Options) error {
	if options == nil {
		options = DefaultOptions()
	} else if !options.EnableProfiling && !options.EnableMetrics {
		return nil
	}
	mux := http.NewServeMux()
	if options.EnableProfiling {
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	}
	if options.EnableMetrics {
		mux.Handle("/metrics", promhttp.Handler())
	}
	return http.ListenAndServe(options.BindAddress, mux)
}
