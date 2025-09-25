package cmd

import (
	"fmt"

	"github.com/madeinly/cards/internal/flows"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCardsDB)
}

var updateCardsDB = &cobra.Command{
	Use: "updateDB",
	Run: func(cmd *cobra.Command, args []string) {
		err := flows.UpdateCardsDB()
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("the mtg cards database has been updated")
	},
}
