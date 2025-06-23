package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available devices",
	Long:  "List the available devices on your hue network",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		lights := List()

		fmt.Println("Available lights: ")
		for i, light := range lights {
			fmt.Printf("#%d - %s, %s, %s,\n", i, light.Name, light.Type, light.Id)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
