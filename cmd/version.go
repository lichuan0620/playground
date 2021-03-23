package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lichuan0620/playground/pkg/version"
)

func Version() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Print(version.Message())
		},
	}
}