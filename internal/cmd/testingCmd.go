package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"githube.com/madeinly/cards/internal/service"
)

func init() {
	rootCmd.AddCommand(testingCmd)
}

var testingCmd = &cobra.Command{
	Use: "testing",
	Run: func(cmd *cobra.Command, args []string) {
		err := service.SetPricesDB()

		if err != nil {
			println(err.Error())
		}

		fmt.Println("mtg database set up")
	},
}
