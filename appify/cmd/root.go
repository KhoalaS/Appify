package cmd

import (
	"fmt"
	"os"

	"github.com/KhoalaS/Appify/appify/cmd/generate"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "appify",
	Short: "Appify is a generator for webview android apps",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(generate.GenerateCmd)
}
