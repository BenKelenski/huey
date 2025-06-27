package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var dimId string
var dimCmd = &cobra.Command{
	Use:   "dim",
	Short: "Adjust the brightness of a light",
	Long:  "Adjust the brightness of a light, takes a number between 0 and 100",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		brightness, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			return fmt.Errorf("error converting brightness to int: %s", err)
		}

		if brightness < 0 || brightness > 100 {
			return fmt.Errorf("brightness must be between 0 and 100")
		}

		fmt.Println("Dimming light with ID:", dimId, brightness)
		return Dim(dimId, brightness)
	},
}

func init() {
	dimCmd.Flags().StringVarP(&dimId, "id", "i", "", "ID of the light")
	err := dimCmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Printf("Error marking flag as required: %s\n", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(dimCmd)
}
