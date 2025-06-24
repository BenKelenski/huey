package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var offId string
var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn off a light",
	Long:  "Turn off a light",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			panic(err)
		}

		fmt.Printf("Turning OFF light with ID: %s\n", id)

		Off(id)
	},
}

func init() {
	offCmd.Flags().StringVarP(&offId, "id", "i", "", "ID of the light")
	rootCmd.AddCommand(offCmd)
}
