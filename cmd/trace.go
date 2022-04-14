package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(traceCmd)
}

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "A brief description of your command",
	Long:  `A longer description of your command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("trace called")
	},
}
