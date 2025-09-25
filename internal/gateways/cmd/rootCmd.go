package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "use users feature",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("prefix with cards for a simple interface with this feature")

	},
}
