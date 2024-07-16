package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version v0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
