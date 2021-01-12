package webhooks

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	url     string
	account string
	event   string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new webhook watcher",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, world!")
	},
}

func init() {
	createCmd.Flags().StringVarP(
		&url, "url", "u", "", "The callback url for your notification",
	)

	createCmd.Flags().StringVarP(
		&account, "account", "a", "",
		"The account to watch for the notification",
	)

	createCmd.Flags().StringVarP(
		&event, "event", "e", "",
		"The event type to receive notifications for "+
			"[pending_receive, receive_block, send_block]",
	)

	_ = createCmd.MarkFlagRequired("url")
	_ = createCmd.MarkFlagRequired("account")
	_ = createCmd.MarkFlagRequired("event")

	viper.BindPFlag("url", createCmd.Flags().Lookup("url"))
	viper.BindPFlag("account", createCmd.Flags().Lookup("account"))
	viper.BindPFlag("event", createCmd.Flags().Lookup("event"))
}
