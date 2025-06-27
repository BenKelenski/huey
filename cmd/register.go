package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register your device with the hue bridge",
	Long:  "Registers your device with the hue bridge, provides username and key for user in other requests",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := Register()

		if err != nil {
			return err
		}

		fmt.Printf("client: response body - %s", res)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
