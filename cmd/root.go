/*
Copyright © 2022 James Daniels <james.daniels@hotmail.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "cGem",
	Short: "Use cGem to buy and sell crypto",
	Long: "Use cGem to quickly buy and sell cryto on the Gemini Exchange",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}