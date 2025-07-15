package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"githube.com/madeinly/cards/internal/service"
)

func init() {
	rootCmd.AddCommand(updateCardsDB)
}

var updateCardsDB = &cobra.Command{
	Use: "updateDB",
	Run: func(cmd *cobra.Command, args []string) {
		err := service.UpdateCardsDB()
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("the mtg cards database has been updated")
	},
}
