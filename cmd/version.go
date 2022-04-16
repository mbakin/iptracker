package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "get the version of the application",
	Long:  `get the version of the application`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)
			}
		} else {
			fmt.Println("v1.0.1")
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
