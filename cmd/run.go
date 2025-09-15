/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"time"
)

type Block struct {
	Letter   rune
	X        int
	Y        int
	Previous *Block
}

func (b *Block) RunRule() (list []*Block) {
	EndBlock := b.Previous

	switch b.Letter {
	case 'A':
		if b.X != b.Y {
			c := make([]Block, 1)
			c[0].Previous = EndBlock
			c[0].X = b.X * 2
			b.Previous = &c[0]
			list = append(list, b)
			list = append(list, &c[0])

		}
	}
	return list
}

type BlockList = *[]Block

func PrintList(block *Block) {
	log.Printf("%+v", block)
	current := block.Previous
	for current != nil {
		log.Printf("%+v", current)
		current = current.Previous
	}
}

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		data := Block{
			Letter:   'A',
			X:        1,
			Y:        10,
			Previous: nil,
		}
		start := time.Now()
		var list []*Block
		list = append(list, &data)
		for _ = range 1_000_000_000 {
			list = list[0].RunRule()
		}
		log.Println(time.Since(start))
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
