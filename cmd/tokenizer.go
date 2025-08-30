/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

const (
	CONTROL_FLOW_CHARACTERS = ` (){}<>,?;:-=\"'@~¬[]`
)

/**
 * Removes comments automatically from the tokenizer before
 * the tokenization process as a preprocessing action */
func GetTokenSkipComments(data []byte, atEOF bool) (advance int, token []byte, err error) {
	length := len(data)
	if !atEOF {
		var halt bool
		var start_index = 0
		var start bool = true

		for !halt { // Continue until you find a given conditional
			// The control flow characters
			if advance < length {
				letter := rune(data[advance])
				if start && (string(letter) == ` `) {
					start_index += 1
					advance += 1
				} else {
					start = false
					for _, char := range CONTROL_FLOW_CHARACTERS {
						if rune(letter) == char {
							halt = true
							break
						}
					}
				}
				advance += 1
			} else {
				halt = true
			}
		}
		if advance != 1 {
			advance = advance - 1
		}
		token = []byte(strings.TrimSpace(string(data[start_index:advance])))
		log.Println(token)
	} // If the token is a // then remove all to next line character,

	return
}

// tokenizerCmd represents the tokenizer command
var tokenizerCmd = &cobra.Command{
	Use:   "tokenizer",
	Short: "The tokenisation transcript for a MEVIL program",
	Long:  `This is the tokeniser output to the terminal for the stack tree from a pushdown recogniser`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Tokenising: %s\n", args[0])
		if file := OpenFile(args[0]); file != nil {
			// Takes a given file, creates a reader wrapper and passes it to the scanner function
			scanner := bufio.NewScanner(file)

			scanner.Split(GetTokenSkipComments)
			for scanner.Scan() {
				fmt.Println(scanner.Text())

			}
		}
	},
}

func OpenFile(filename string) *os.File {
	if f, err := os.Open(filename); err != nil {
		log.Println(err)
	} else {
		log.Println(f)
		return f
	}
	return nil

}

// The First pass of the tokenizer takes in a file and removes comments
// And segments the file into its two constituent components Primitives and Properties.
// Primitives are some assigment or operation that needs to be taken into accoutn
// Properties are something that has to be parsed by the system.

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
