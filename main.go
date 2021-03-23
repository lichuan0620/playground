package main

import (
    "log"
	_ "net/http/pprof"

    "github.com/spf13/cobra"
    
	"github.com/lichuan0620/playground/pkg/version"
	"github.com/lichuan0620/playground/cmd"
)

func main() {
    rootCmd := cobra.Command{
        Use: version.Name,
	}
	rootCmd.AddCommand(
		cmd.Version(),
	)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("execute: %v", err)
	}
}
