package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"githube.com/madeinly/cards/internal/card"
)

func init() {
	rootCmd.AddCommand(testingCmd)
}

var testingCmd = &cobra.Command{
	Use: "testing",
	Run: func(cmd *cobra.Command, args []string) {
		err := card.SetupPriceTable()

		if err != nil {
			println(err.Error())
		}

		fmt.Println("finish parsing")
	},
}
