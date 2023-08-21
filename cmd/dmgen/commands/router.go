package commands

import "github.com/spf13/cobra"

var Router = &cobra.Command{
	Use:   "dmgen",
	Short: "Generate data models",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
