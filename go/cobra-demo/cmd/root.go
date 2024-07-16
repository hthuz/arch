package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Verbose bool

// Root command is my app
var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a static site generator",
	Long:  "Hugo is a fantastic and fast static site generator",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running hugo")
		fmt.Println("your args ", args)
		fmt.Printf("Verbose: %v\n", Verbose)
	},
	TraverseChildren: true,
}

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.Execute()
}
