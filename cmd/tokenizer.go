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
	PRIMITIVE_TEXT = "prim"
)

/**
 * Removes comments automatically from the tokenizer before
 * the tokenization process as a preprocessing action */
func GetTokenSkipComments(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if advance, token, err = bufio.ScanWords(data, atEOF); err == nil {
		// check for comment
		for string(token) == `//` {
			comment := []byte(strings.Split(string(data[advance:]), "\n")[0])
			commentlength := len(comment)
			advance = advance + commentlength
			// If we have a comment read the next token
			var nextTokenAdvance int
			nextTokenAdvance, token, err = bufio.ScanWords(data[advance:], atEOF)
			advance = advance + nextTokenAdvance
			log.Println("comment", string(comment))
		}
	}
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
