package cmd

import (
	"encoding/json"
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

		card, err := service.GetCardFromID("653cc07b-0f53-4b5b-9c5f-885b8b4a6e5f")

		if err != nil {
			fmt.Printf("There was an error fetching the data: %v", err)
			return
		}

		jsonCard, err := json.MarshalIndent(card, "", " ")

		if err != nil {
			fmt.Printf("there was an error building the json: %v", err.Error())
			return
		}

		fmt.Println(string(jsonCard))
	},
}
