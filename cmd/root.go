package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sort"
)

var rootCmd = &cobra.Command{
	Use:   "Huey",
	Short: "Get all Hue devices",
	Long:  "Get all Hue devices connected to the Hue bridge",
	Run: func(cmd *cobra.Command, args []string) {
		reset := "\033[0m"
		red := "\033[31m"
		green := "\033[32m"
		cyan := "\033[36m"

		lights := Devices()

		if len(lights) == 0 {
			fmt.Printf(red + "❗ NO LIGHTS FOUND ❗\n" + reset)
			return
		}

		sort.Slice(lights, func(i, j int) bool {
			return lights[i].Name < lights[j].Name
		})

		for i, light := range lights {
			fmt.Printf("%s #%d %s- %s, %s, %s %s %s\n", green, i, reset, light.Name, light.Type, cyan, light.Id, reset)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Oops. An error while executing Huey '%s'\n", err)
		os.Exit(1)
	}
}
