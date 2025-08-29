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
			split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
				advance, token, err = bufio.ScanWords(data, atEOF)
				// check for comment
				if string(token) == `//` {
					comment := []byte(strings.Split(string(data[1:]), "\n")[0])
					commentlength := len(comment)
					advance = commentlength + 1
					// If we have a comment read the next token
					var nextTokenAdvance int
					nextTokenAdvance, token, err = bufio.ScanWords(data[advance:], atEOF)
					advance = advance + nextTokenAdvance
				}

				switch string(token) {
				case PRIMITIVE_TEXT:
					log.Println("primtive text")
				default:
					log.Println("non primitive")
				}

				return
			}
			scanner.Split(split)

			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}
	},
}

// Take the scanner and create a whole tokenizer for the language,
// - Convert reserved words to identifier objects

type ReservedLookup struct {
	Marker int64
}

type ReservedWord interface {
	int64
}

type Module struct {
	Letter     uint8
	Parameters []uint64
}

type Conditional struct {
	LHS     string
	operand uint8
	RHS     string
}

type TokenTypes interface {
	int64 |
		string |
		float64 |
		Module
}

// List of tokens, wrapped in a monadic structure to parse as a single item in the stream channel
type TokenList[T TokenTypes] interface {
	[]T
}

// Type uint8 is a overloaded abstract monad classifier,
// It effectively acts as a switch flow controller because
// go won't let me implement this under its rather restrictive type system.

// *
// * Type
// |8|7|6|5|4|3|2|1|
// |---------------|
// |0|1|0|1|0|1|0|1|
// LE - as Little Endian
// 1. isList (Important for if the parser needs to loop on all the next states)
// 2-3. abstract primitive type class where,
//
//	00: Module
//	01: int64
//	10: string
//	11: float64
//
// 4-5: prediction of future classes, i.e. what does the parser think this?
//
//	00: primitive list
//	01: rule
//	10: GENERATOR
//	11: validation and setup
//
// 6-7:
//
//	reserved
//
// */
type Token struct {
	Type uint8
	V    string
}

/*
Evaluates the string against the reserved words dictionary

	Computes in O(1) due to preloading the dictionary into memory and lookup

	*
*/
func EvaluateForReservedWords(word string) {

}

func (T *Token) classify() {
	EvaluateForReservedWords(T.V)
}

func (T *Token) ClassifyTokenType() uint8 {
	T.classify()
	return 0
}

// func (T *Token) WrapInMonad() AbstractMonadicToken {

// }

// All the parseable types of tokens in the token stream

// Object for the tokeniser,
type Tokenizer struct {
	file       *os.File    // The file being read from
	Primitives chan string // Assigned Variable Namespaces
	Properties chan string // Properties
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
