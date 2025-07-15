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

		card, err := service.GetCardFromID("9a833fa7-5934-4c04-be42-e215a61f450e")

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
