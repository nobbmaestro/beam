package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "beam",
	Long:  "Command-line utility for generating and broadcasting DMX over sACN",
	RunE:  runRoot,
}

func runRoot(cmd *cobra.Command, args []string) error {
	return nil
}

func SetVersion(version, commit, date string) {
	rootCmd.Version = version
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
}
