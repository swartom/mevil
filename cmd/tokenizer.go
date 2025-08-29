/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tokenizerCmd represents the tokenizer command
var tokenizerCmd = &cobra.Command{
	Use:   "tokenizer",
	Short: "The tokenisation transcript for a MEVIL program",
	Long:  `This is the tokeniser output to the terminal for the stack tree from a pushdown recogniser`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Tokenising %s", args[0])
	},
}


func OpenFile() {

}




func init() {
	rootCmd.AddCommand(tokenizerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tokenizerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tokenizerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
