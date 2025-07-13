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
	Use: "update-db",
	Run: func(cmd *cobra.Command, args []string) {
		err := service.FetchCardsDB()
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}
