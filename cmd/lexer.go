/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	// "bufio"
	// "os"
	"github.com/spf13/cobra"
)

// lexerCmd represents the lexer command
var lexerCmd = &cobra.Command{
	Use:   "lexer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lexer called")

		// scanner := bufio.NewScanner(file)

		// scanner.Split(GetTokenSkipComments)
		// for scanner.Scan() {
		// 	switch scanner.Text() {
		// 	case
		// 	}
		// }

	},
}

func init() {
	rootCmd.AddCommand(lexerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lexerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lexerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
