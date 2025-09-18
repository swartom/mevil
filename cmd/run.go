/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
	"time"
)

var beta_distro distuv.Beta = distuv.Beta{
	// Min: 0,
	// Max: 1,
	Alpha: 40, //7.5,
	Beta:  .5, //.5,
}

type Block struct { // Largest Module in the system
	Letter   uint8
	X        uint32
	Y        uint32
	Previous *Block
}

var wg sync.WaitGroup
var lim uint32 = 1

func (b *Block) RunRule() {
	EndBlock := b.Previous

	switch b.Letter {
	case 'A':
		if b.X != b.Y {
			q := beta_distro.Rand()
			r := uint32(int(q*float64(b.Y-b.X))) + b.X + 1

			list := make([]Block, 3)
			a := &list[0]
			a.Letter = 'L'
			a.X = r

			a2 := &list[1]
			a2.Letter = 'A'
			a2.X = b.X
			a2.Y = r - 1

			a3 := &list[2]
			a3.Letter = 'L'
			a3.X = r - 1

			b.X = r

			b.Previous = a3
			a3.Previous = a2
			a2.Previous = a
			a.Previous = EndBlock

			if a2.X != r-1 {
				wg.Add(1)
				go a2.RunRule()
			}
			if r != b.Y {
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
	// log.Printf("%+v", block)
	current := block.Previous
	var count int
	for current != nil {
		count = count + 1
		// log.Printf("%+v", current)
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
		i1, _ := strconv.Atoi(args[0])
		i2, _ := strconv.Atoi(args[1])
		i := int(math.Pow(float64(i1), float64(i2)))
		{
			lim = uint32(i)

			data := Block{
				Letter:   'A',
				X:        1,
				Y:        uint32(i),
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
			list[0].debugDumpToFile()
		}
	},
}

func (b *Block) debugDumpToFile() {
	currentblock := b
	fo, err := os.Create("test.adjlist")
	if err != nil {
		panic(err)
	}
	separator := " "
	for currentblock != nil {
		switch currentblock.Letter {
		case 'A':
			separator = "\n"
		case 'L':
			separator = " "
		}
		if currentblock != b {
			fo.WriteString(fmt.Sprintf("%s%d", separator, currentblock.X-1))
		} else {
			fo.WriteString(fmt.Sprintf("%d", currentblock.X-1))
		}
		currentblock = currentblock.Previous
	}
	fo.Close()
	return
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
