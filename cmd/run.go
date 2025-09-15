/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"math"
	"strconv"
	"sync"
	"time"
)

type Block struct {
	Letter   uint8
	X        uint16
	Y        uint16
	Previous *Block
}

var wg sync.WaitGroup
var lim uint16 = 1

func (b *Block) RunRule() {
	EndBlock := b.Previous

	switch b.Letter {
	case 'A':
		if b.Y < lim {
			c := new(Block)
			c.Letter = 'A'
			c.Previous = EndBlock
			c.X = uint16(math.Pow(2, float64(b.Y))) + b.X
			c.Y = b.Y + 1
			b.Y = b.Y + 1
			b.Previous = c

			if b.Y < lim {
				wg.Add(1)
				go c.RunRule()
				b.RunRule()
			} else {
				wg.Done()
			}
		} else {
			wg.Done()
		}
	}
}

func PrintList(block *Block) int {
	log.Printf("%+v", block)
	current := block.Previous
	var count int
	for current != nil {
		count = count + 1

		current = current.Previous
	}
	return count
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
		i, _ := strconv.Atoi(args[0])
		{
			lim = uint16(i)

			data := Block{
				Letter:   'A',
				X:        1,
				Y:        0,
				Previous: nil,
			}
			start := time.Now()
			var list []*Block
			list = append(list, &data)
			wg.Add(1)
			list[0].RunRule()
			wg.Wait()
			log.Println(time.Since(start))
			count := PrintList(list[0])
			log.Println(i, count)
		}
	},
}

func LinearLookupParser(*Block) {

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
