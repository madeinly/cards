package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"githube.com/madeinly/cards/internal/service"
)

func init() {
	rootCmd.AddCommand(initPrices)
}

var initPrices = &cobra.Command{
	Use: "initPrices",
	Run: func(cmd *cobra.Command, args []string) {
		err := service.InitCardPrices()
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("the mtg prices have been initialized")
	},
}
