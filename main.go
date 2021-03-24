package main

import (
	"flag"
	"net/http"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"github.com/lichuan0620/playground/cmd"
	"github.com/lichuan0620/playground/pkg/telemetry"
	"github.com/lichuan0620/playground/pkg/version"
)

var rootCmd = cobra.Command{
	Use: version.Name,
}

var telemetryOptions = telemetry.Options{
	EnableProfiling: false,
	EnableMetrics:   false,
	BindAddress:     "0.0.0.0:8081",
}

func init() {
	klog.InitFlags(flag.CommandLine)
	flags := rootCmd.PersistentFlags()
	flags.AddGoFlagSet(flag.CommandLine)
	flags.BoolVar(
		&telemetryOptions.EnableProfiling, "telemetry.enable-profiling", telemetryOptions.EnableProfiling,
		"Enable pprof and serve it on telemetry endpoint",
	)
	flags.BoolVar(
		&telemetryOptions.EnableMetrics, "telemetry.enable-metrics", telemetryOptions.EnableMetrics,
		"Enable Prometheus metrics and serve it on telemetry endpoint",
	)
	flags.StringVar(
		&telemetryOptions.BindAddress, "telemetry.bind-address", telemetryOptions.BindAddress,
		"Address to listen on for telemetry",
	)
	rootCmd.AddCommand(
		cmd.Version(),
	)
}

func main() {
	go func() {
		if err := telemetry.ServeTelemetry(&telemetryOptions); err != nil && err != http.ErrServerClosed {
			klog.Fatalf("telemetry server crashed: %v", err)
		}
	}()
	if err := rootCmd.Execute(); err != nil {
		klog.Fatalf("root command executed with an error: %v", err)
	}
}
