package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func ExecuteContext(ctx context.Context) error {
	return rootCmd.ExecuteContext(ctx)
}
