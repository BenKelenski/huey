package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	offLightId string
	offRoomId  string
)

var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn off a light or room",
	Long:  "Turn off a light or room",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {

		if offLightId == "" && offRoomId == "" {
			return fmt.Errorf("must specify either a light or room")
		}

		if offLightId != "" {
			return Off(offLightId, false)
		} else {
			return Off(offRoomId, true)
		}

	},
}

func init() {
	offCmd.Flags().StringVarP(&offLightId, "light", "l", "", "ID of the light")
	offCmd.Flags().StringVarP(&offRoomId, "room", "r", "", "ID of the room")
	offCmd.MarkFlagsMutuallyExclusive("light", "room")

	rootCmd.AddCommand(offCmd)
}
