package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "Huey",
	Short: "Huey is a cli tool for controlling Philips Hue lights",
	Long:  "Huey is a cli tool for controlling Philips Hue lights - listing available devices, switching on and off, and controlling color.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Oops. An error while executing Huey '%s'\n", err)
		os.Exit(1)
	}
}
