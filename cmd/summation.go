package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var summationCmd = &cobra.Command{
	Use:   "summation number",
	Short: "Returns the summation of desired number",
	Long: `Returns the summation of desired number
Usage:
        go-cli-example summation 45`,
	Args: cobra.ExactArgs(1),
	Run: summation,
}

func init() {
	rootCmd.AddCommand(summationCmd)
}

func summation(cmd *cobra.Command, args []string) {
	number, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Cannot use '%v' as number.\n", args[0])
		return
	}

	result := 0
	for i := number; i > 0; i-- {
		result += i
	}
	fmt.Printf("Summation of %d is: %d\n", number, result)
}