/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	// "os"
	"github.com/spf13/cobra"
)

// func (p *Parser) Next() {
// 	p.scanner.Scan()
// 	p.current = p.nextToken
// 	Classify(p.scanner.Text(), &p.nextToken)
// }

const (
	IDENTIIFER int = iota // Name assigned by a programmer
	KEYWORD    int = iota
	SEPARATOR  int = iota
	OPERATOR   int = iota
	LITERAL    int = iota
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
		fmt.Printf("Lexing: %s\n", args[0])
		if file := OpenFile(args[0]); file != nil {
			// Takes a given file, creates a reader wrapper and passes it to the scanner function
			scanner := bufio.NewScanner(file)

			scanner.Split(GetTokenSkipComments)
			for scanner.Scan() {

				// Categorise token and push to a channel
				scanner.Text()

			}
		}
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
