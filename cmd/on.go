package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	onLightId string
	onRoomId  string
)

var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Turn on a light or room",
	Long:  "Turn on a light or room",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		if onLightId == "" && onRoomId == "" {
			return fmt.Errorf("must specify either a light or room")
		}

		if onLightId != "" {
			return On(onLightId, false)
		} else {
			return On(onRoomId, true)
		}
	},
}

func init() {
	onCmd.Flags().StringVarP(&onLightId, "light", "l", "", "ID of the light")
	onCmd.Flags().StringVarP(&onRoomId, "room", "r", "", "ID of the room")
	onCmd.MarkFlagsMutuallyExclusive("light", "room")

	rootCmd.AddCommand(onCmd)
}
