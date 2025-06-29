package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sort"
)

var roomsFlag bool
var rootCmd = &cobra.Command{
	Use:   "Huey",
	Short: "Get all Hue devices",
	Long:  "Get all Hue devices connected to the Hue bridge",
	RunE: func(cmd *cobra.Command, args []string) error {
		reset := "\033[0m"
		red := "\033[31m"
		green := "\033[32m"
		cyan := "\033[36m"

		devices, err := Devices(roomsFlag)

		if err != nil {
			return err
		}

		if len(devices) == 0 {
			fmt.Printf(red + "❗ NO LIGHTS FOUND ❗\n" + reset)
			return nil
		}

		sort.Slice(devices, func(i, j int) bool {
			return devices[i].Name < devices[j].Name
		})

		for i, light := range devices {
			fmt.Printf("%s #%d %s- %s, %s, %s %s %s\n", green, i, reset, light.Name, light.Type, cyan, light.Id, reset)
		}

		return nil
	},
}

func Execute() {
	rootCmd.Flags().BoolVarP(&roomsFlag, "rooms", "r", false, "Get all rooms known to the bridge")
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Oops. An error while executing Huey '%s'\n", err)
		os.Exit(1)
	}
}
