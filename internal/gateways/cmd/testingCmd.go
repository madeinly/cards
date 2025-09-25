package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(testingCmd)
}

var testingCmd = &cobra.Command{
	Use: "testing",
	Run: func(cmd *cobra.Command, args []string) {
		// ctx := context.Background()
		// card, err := service.GetCardFromID(ctx, "ae9a7d45-fec6-404e-965e-68e463d65fbf")

		// if err != nil {
		// 	fmt.Printf("There was an error fetching the data: %v", err)
		// 	return
		// }

		// jsonCard, err := json.MarshalIndent(card, "", " ")

		// if err != nil {
		// 	fmt.Printf("there was an error building the json: %v", err.Error())
		// 	return
		// }

		// fmt.Println(string(jsonCard))
	},
}
