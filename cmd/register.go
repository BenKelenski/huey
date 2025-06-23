package cmd

import (
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register your device with the hue bridge",
	Long:  "Registers your device with the hue bridge, provides username and key for user in other requests",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		Register()
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
