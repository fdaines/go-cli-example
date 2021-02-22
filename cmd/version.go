package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows CLI version",
	Long: "Shows CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version: 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}