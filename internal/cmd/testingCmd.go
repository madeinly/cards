package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"githube.com/madeinly/cards/internal/repository"
)

func init() {
	rootCmd.AddCommand(testingCmd)
}

var testingCmd = &cobra.Command{
	Use: "testing",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := repository.InitCardsDB()

		if err != nil {
			fmt.Println(err.Error())
		}

		db, err := repository.GetCardsDB()

		if _, err := db.Exec("PRAGMA quick_check"); err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("database looks fine")

		if err != nil {
			fmt.Println(err.Error())
		}

		db.Close()

		fmt.Println("finish testing")
	},
}
