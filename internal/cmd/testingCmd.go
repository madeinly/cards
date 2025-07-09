package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(testingCmd)
}

var testingCmd = &cobra.Command{
	Use: "testing",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("yo aqui aprendi")

	},
}
