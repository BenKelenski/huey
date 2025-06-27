package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var offId string
var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn off a light",
	Long:  "Turn off a light",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}

		fmt.Printf("Turning OFF light with ID: %s\n", id)

		return Off(id)
	},
}

func init() {
	offCmd.Flags().StringVarP(&offId, "id", "i", "", "ID of the light")
	err := offCmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Printf("Error marking flag as required: %s\n", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(offCmd)
}
