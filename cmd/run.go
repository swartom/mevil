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
	// Alpha: .01, //7.5,
	// Beta:  1,   //.5,
}

type Block struct { // Largest Module in the system
	Letter uint8
	X      uint32
	Y      uint32
	// D        uint32
	// V        uint32
	Previous *Block
}

var wg sync.WaitGroup
var lim uint32 = 1
var connections = 4
var min uint32 = 5

// func (b *Block) RunRule() {
// 	EndBlock := b.Previous
// 	switch b.Letter {
// 	case 'L':
// 		if b.X != 1 {
// 			if beta_distro.Rand() > .8 {
// 				a1 := new(Block)
// 				a1.Letter = 'L'
// 				a1.X = b.X - 1
// 				b.Previous = a1
// 				a1.Previous = EndBlock
// 			}
// 			r := uint32(1)
// 			q := uint32(float64(b.Y) * beta_distro.Rand())
// 			if int(b.X)-int(q) > 0 {
// 				r = q
// 			}
// 			b.X = b.X - r
// 			b.RunRule()
// 		} else if b.X == 1 {
// 			b.X = b.Y
// 			wg.Done()
// 		}
// 	case 'A':
// 		if b.D != b.V && b.V != 0 {
// 			// Rate of decay of the meta-community
// 			var q1 float64 = beta_distro.Rand()
// 			r1 := uint32(q1*float64((b.V-b.D))) + b.D + 1
// 			// Sizes of the communities themselves
// 			var q2 float64 = q1
// 			r2 := uint32(q2*float64(b.Y-uint32(b.V-r1)-(b.X+uint32(r1-b.D)))) + b.X + uint32(r1-b.D) // 3 r1s tf + 1 already

// 			a1 := new(Block)

// 			a1.Letter = 'A'
// 			a1.X = b.X
// 			a1.Y = r2 - 1
// 			a1.D = b.D
// 			a1.V = r1 - 1

// 			b.X = r2
// 			b.D = r1

// 			b.Previous = a1
// 			a1.Previous = EndBlock

// 			if b.D == b.V {
// 				b.D = b.X
// 				b.V = b.X
// 			}
// 			if a1.D == a1.V {
// 				a1.D = a1.X
// 				a1.V = a1.X
// 			}

// 			wg.Add(1)
// 			go a1.RunRule()

// 			b.RunRule()

// 		} else if b.V == 0 && b.D < b.Y+1 {
// 			if beta_distro.Rand() > .25 {
// 				a1 := new(Block)
// 				a1.Letter = 'L'
// 				a1.X = b.D

// 				b.Previous = a1
// 				a1.Previous = EndBlock
// 			}
// 			b.D = b.D + 1 //  uint32(beta_distro.Rand()*float64(10))
// 			b.RunRule()
// 		} else if b.X != b.Y && b.V != 0 {
// 			var a1 *Block
// 			if b.V != 0 {
// 				a1 = new(Block)
// 				a1.Letter = 'L'
// 				a1.X = b.V
// 				a1.Y = b.X + 1
// 			}
// 			a2 := new(Block)
// 			a2.Letter = 'A'
// 			a2.X = b.X
// 			a2.Y = b.Y
// 			a2.D = b.X + 1
// 			a2.V = 0

// 			if b.V != 0 {
// 				b.Previous = a1
// 				a1.Previous = a2
// 			} else {
// 				b.Previous = a2
// 			}
// 			if b.X == b.V && b.V != 0 {
// 				a3 := new(Block)
// 				a3.Letter = 'L'
// 				a3.X = b.V
// 				a3.Y = b.X

// 				a2.Previous = a3
// 				a3.Previous = EndBlock
// 				wg.Add(1)
// 				go a3.RunRule()
// 			} else {
// 				a2.Previous = EndBlock
// 			}
// 			b.X = b.X + 1
// 			wg.Add(1)
// 			go a2.RunRule()
// 			if b.V != 0 {
// 				wg.Add(1)
// 				go a1.RunRule()
// 			}
// 			b.RunRule()
// 		} else {
// 			wg.Done()

// 		}
// 	}
// }

// func (b *Block) RunRule() {
// 	endblock := b.Previous
// 	switch b.Letter {
// 	case 'A':
// 		if b.X != b.Y {
// 			q := beta_distro.Rand()
// 			// q = q / 1.01 // fpa error on machine rounds high floats to 1 -> not valid in the system
// 			// q := 0.5
// 			r := uint32((q)*float64((b.Y-b.X))) + b.X + 1

// 			list := make([]Block, 3)
// 			a2 := &list[0]
// 			a2.Letter = 'A'
// 			a2.X = b.X
// 			a2.Y = r - 1
// 			a3 := &list[1]
// 			a3.Letter = 'L'
// 			a3.X = r

// 			a4 := &list[2]
// 			a4.Letter = 'L'
// 			a4.X = r - 1

// 			// var vallist []uint32
// 			// vallist = append(vallist, b.Y)
// 			// for i := range connections {
// 			// 	var value uint32
// 			// 	var unique = false
// 			// 	for !unique {
// 			// 		unique = true
// 			// 		q := beta_distro.Rand()

// 			// 		value = uint32(int(q * float64(lim)))

// 			// 		for _, j := range vallist {
// 			// 			if value == j {
// 			// 				unique = false
// 			// 				break
// 			// 			}
// 			// 		}
// 			// 	}
// 			// 	tmp := &list[i]
// 			// 	tmp.Letter = 'L'
// 			// 	tmp.X = value
// 			// 	tmp.Previous = a
// 			// 	a = tmp
// 			// 	vallist = append(vallist, value)
// 			// }

// 			b.X = r
// 			b.Previous = a4

// 			a4.Previous = a2
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

func (b *Block) RunRule() {
	endblock := b.Previous
	switch b.Letter {
	case 'A':
		if b.X != b.Y {
			// q := beta_distro.rand()
			// q = q / 1.01 // fpa error on machine
			q := .5
			r := uint32((q)*float64((b.Y-b.X))) + b.X + 1

			list := make([]Block, connections+1)
			a2 := &list[connections]
			a2.Letter = 'A'
			a2.X = b.X
			a2.Y = r - 1

			var a *Block = a2
			var vallist []uint32

			connhere := uint32(connections)
			if r < connhere {
				for i := range b.X - 1 {
					tmp := &list[i]
					tmp.Letter = 'L'
					tmp.X = b.X - i - 1
					tmp.Previous = a
					a = tmp
				}
			} else {
				//TODO fix duplicate edegs issue (its a code problem)
				for i := range connhere {
					var value uint32

					q := beta_distro.Rand()

					value = uint32(int(q*float64(r-uint32(len(vallist)))) + len(vallist))

					for _, j := range vallist {
						if value == j {
							if value-1 > 0 {
								value = value - 1
							}
						}
					}

					tmp := &list[i]
					tmp.Letter = 'L'
					tmp.X = value
					tmp.Previous = a
					a = tmp
					vallist = append(vallist, value)
				}
			}

			b.X = r
			b.Previous = a

			a2.Previous = endblock

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
	log.Printf("%+v", block)
	current := block.Previous
	var count int
	for current != nil {
		count = count + 1
		log.Printf("%+v", current)
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
		// r := rand.New(rand.NewSource(5))
		// beta_distro.Src = r

		i1, _ := strconv.Atoi(args[0])
		i2, _ := strconv.Atoi(args[1])
		// i3, _ := strconv.Atoi(args[2])
		i := int(math.Pow(float64(i1), float64(i2)))

		count := 1
		repeats := 100
		times := make([]int, count)
		values := make([]int, count)
		step := i / count
		time.Sleep(1000)
		for c := range repeats {
			log.Println("Repeat No. : ", c)

			for i := range count {
				lim = uint32(i*step + step)

				data := Block{
					Letter: 'A',
					X:      1,
					Y:      lim,
					// D:        1,
					// V:        uint32(i3),
					Previous: nil,
				}
				start := time.Now()
				var list []*Block
				list = append(list, &data)
				wg.Add(1)
				start = time.Now()
				list[0].RunRule()
				wg.Wait()
				end := time.Since(start)
				values[i] = int(lim)
				log.Println(end)
				times[i] = times[i] + int(end)
				//times[i] = append(times[i], end)
				// PrintList(list[0])
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
