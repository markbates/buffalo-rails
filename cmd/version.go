package cmd

import (
	"fmt"

	"github.com/markbates/buffalo-rails/rails"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "current version of rails",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("rails", rails.Version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
