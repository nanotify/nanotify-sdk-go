// Package cmd is responsible for handling the cli code
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nanotify/nanotify-sdk-go/cli/cmd/webhooks"
)

var rootCmd = &cobra.Command{
	Use: "nanotify-cli",
}

var key string

// Execute runs the rootCmd
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(
		&key, "key", "", "Nanotify API Key",
	)

	rootCmd.MarkPersistentFlagRequired("key")

	_ = viper.BindPFlag("apikey", rootCmd.PersistentFlags().Lookup("key"))

	rootCmd.AddCommand(webhooks.RootCmd)
}
