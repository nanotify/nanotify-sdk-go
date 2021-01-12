// Package webhooks is responsible for the webhook cmds
package webhooks

import "github.com/spf13/cobra"

// RootCmd is the root command for the webhooks
var RootCmd = &cobra.Command{
	Use: "webhooks",
}

func init() {
	RootCmd.AddCommand(createCmd)
}
