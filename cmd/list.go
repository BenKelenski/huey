package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sort"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available devices",
	Long:  "List the available devices on your hue network",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		reset := "\033[0m"
		red := "\033[31m"
		green := "\033[32m"
		cyan := "\033[36m"

		lights := List()

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

func init() {
	rootCmd.AddCommand(listCmd)
}
