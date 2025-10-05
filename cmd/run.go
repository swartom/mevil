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

var beta_distro distuv.Uniform = distuv.Uniform{
	Min: 0,
	Max: 1,
	// Alpha: 2, //7.5,
	// Beta:  1, //.5,
}

type Block struct { // Largest Module in the system
	Letter uint8
	X      uint32
	Y      uint32
	// D        uint8
	// V        uint8
	Previous *Block
}

var wg sync.WaitGroup
var lim uint32 = 1
var connections = 10

// func (b *Block) RunRule() {
// 	EndBlock := b.Previous
// 	switch b.Letter {
// 	case 'A':
// 		if b.D != b.V {
// 			// Rate of decay of the meta-community
// 			var q1 float32 = .5
// 			r1 := uint8(q1*float32((b.V-b.D))) + b.D + 1
// 			// Sizes of the communities themselves
// 			var q2 float32 = .5
// 			r2 := uint32(q2*float32(b.N-uint32(b.V-r1)-(b.X+uint32(r1-b.D)))) + b.X + uint32(r1-b.D) // 3 r1s tf + 1 already

// 			a1 := new(Block)

// 			a1.Letter = 'A'
// 			a1.X = b.X
// 			a1.N = r2 - 1
// 			a1.D = b.D
// 			a1.V = r1 - 1

// 			b.X = r2
// 			b.D = r1

// 			b.Previous = a1
// 			a1.Previous = EndBlock

// 			if a1.D != a1.V {
// 				wg.Add(1)
// 				go a1.RunRule()
// 			}
// 			if b.D != b.V {
// 				b.RunRule()
// 			} else {
// 				wg.Done()
// 			}
// 		} else {
// 			wg.Done()
// 		}
// 	}
// }

func (b *Block) RunRule() {
	endblock := b.Previous
	switch b.Letter {
	case 'A':
		if b.X != b.Y {
			q := beta_distro.Rand()
			q = q / 1.01 // fpa error on machine rounds high floats to 1 -> not valid in the system
			// q := 0.5
			r := uint32((q)*float64((b.Y-b.X))) + b.X + 1

			list := make([]Block, 3)
			a2 := &list[0]
			a2.Letter = 'A'
			a2.X = b.X
			a2.Y = r - 1
			a3 := &list[1]
			a3.Letter = 'L'
			a3.X = r

			a4 := &list[2]
			a4.Letter = 'L'
			a4.X = r - 1

			// var vallist []uint32
			// vallist = append(vallist, b.Y)
			// for i := range connections {
			// 	var value uint32
			// 	var unique = false
			// 	for !unique {
			// 		unique = true
			// 		q := beta_distro.Rand()

			// 		value = uint32(int(q * float64(lim)))

			// 		for _, j := range vallist {
			// 			if value == j {
			// 				unique = false
			// 				break
			// 			}
			// 		}
			// 	}
			// 	tmp := &list[i]
			// 	tmp.Letter = 'L'
			// 	tmp.X = value
			// 	tmp.Previous = a
			// 	a = tmp
			// 	vallist = append(vallist, value)
			// }

			b.X = r
			b.Previous = a4

			a4.Previous = a2
			a2.Previous = a3

			a3.Previous = endblock
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

// func (b *Block) RunRule() {
// 	endblock := b.Previous
// 	switch b.Letter {
// 	case 'A':
// 		if b.X != b.Y {
// 			// q := beta_distro.Rand()
// 			// q = q / 1.01 // fpa error on machine
// 			q := .5
// 			r := uint32((q)*float64((b.Y-b.X))) + b.X + 1

// 			list := make([]Block, connections+1+1)
// 			a2 := &list[connections]
// 			a2.Letter = 'A'
// 			a2.X = b.X
// 			a2.Y = r - 1
// 			a3 := &list[connections+1]
// 			a3.Letter = 'L'
// 			a3.X = b.Y

// 			var a *Block = a2
// 			var vallist []uint32
// 			vallist = append(vallist, b.Y)
// 			for i := range connections {
// 				var value uint32
// 				var unique = false
// 				for !unique {
// 					unique = true
// 					q := beta_distro.Rand()

// 					value = uint32(int(q * float64(lim)))

// 					for _, j := range vallist {
// 						if value == j {
// 							unique = false
// 							break
// 						}
// 					}
// 				}
// 				tmp := &list[i]
// 				tmp.Letter = 'L'
// 				tmp.X = value
// 				tmp.Previous = a
// 				a = tmp
// 				vallist = append(vallist, value)
// 			}

// 			b.X = r
// 			b.Previous = a

// 			a2.Previous = a3
// 			a3.Previous = endblock
// 			if a2.X != r-1 {
// 				wg.Add(1)
// 				go a2.RunRule()
// 			}
// 			if r != b.Y {
// 				b.RunRule()
// 			} else {
// 				wg.Done()
// 			}
// 		} else {
// 			wg.Done()
// 		}
// 	}
// }

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
		// i3, _ := strconv.Atoi(args[2])
		i := int(math.Pow(float64(i1), float64(i2)))

		count := 100
		repeats := 50
		times := make([]int, count)
		values := make([]int, count)
		step := i / count
		for _ = range repeats {

			for i := range count {
				lim = uint32(i*step + step)

				data := Block{
					Letter: 'A',
					X:      1,
					Y:      lim,
					// D:        1,
					// V:        uint8(i3),
					Previous: nil,
				}
				start := time.Now()
				var list []*Block
				list = append(list, &data)
				wg.Add(1)
				list[0].RunRule()
				wg.Wait()
				end := time.Since(start)
				values[i] = int(lim)

				times[i] = times[i] + int(end)
				//times[i] = append(times[i], end)
				// count := PrintList(list[0])
				// log.Println(i, count)
				// list[0].debugDumpToFile()
			}
		}
		log.Println(values)
		for i := range count {
			times[i] = int(times[i] / repeats)
		}

		log.Println(times)
		writeexectofile(values, times)
	},
}

func writeexectofile(values []int, times []int) {
	fo, _ := os.Create("exec.txt")
	for i := range len(values) {
		fo.WriteString(fmt.Sprintf("%d %d\n", values[i], times[i]))
	}
	fo.Close()
}

func (b *Block) debugDumpToFile() {
	currentblock := b
	fo, err := os.Create("test_2.adjlist")
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
