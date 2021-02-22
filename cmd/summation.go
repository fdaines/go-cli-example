package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var summationCmd = &cobra.Command{
	Use:   "summation n",
	Short: "Returns the summation of desired number 'n'",
	Long: `Returns the summation of desired number 'n'.
Formula:
  Sum of 'i', from i='from' to 'n'`,
	Args: cobra.ExactArgs(1),
	Run: summation,
}

var from int

func init() {
	rootCmd.AddCommand(summationCmd)
	summationCmd.Flags().IntVarP(&from, "from", "f", 1, "Define the init value for the summation.")
}

func summation(cmd *cobra.Command, args []string) {
	number, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Cannot use '%v' as number.\n", args[0])
		return
	}

	result := 0
	for i := number; i >= from; i-- {
		result += i
	}
	fmt.Printf("Summation of %d is: %d\n", number, result)
}