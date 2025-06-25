package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var colorId string
var redValue float64
var greenValue float64
var blueValue float64
var colorCmd = &cobra.Command{
	Use:   "color",
	Short: "Set the color of a light",
	Long:  "Set the color of a light",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Changing light to color: ", colorId)
	},
}

func init() {
	colorCmd.Flags().StringVarP(&colorId, "id", "i", "", "ID of the light")
	err := colorCmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Printf("Error marking flag as required: %s\n", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(colorCmd)
}
