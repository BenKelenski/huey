package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var onId string
var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Turn on a light",
	Long:  "Turn on a light",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}

		fmt.Printf("Turning ON light with ID: %s\n", id)

		return On(id)
	},
}

func init() {
	onCmd.Flags().StringVarP(&onId, "id", "i", "", "ID of the light")
	err := onCmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Printf("Error marking flag as required: %s\n", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(onCmd)
}
