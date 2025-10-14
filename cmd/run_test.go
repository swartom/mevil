package cmd

import (
	"math"
	"runtime"
	"testing"
)

func Benchmark_MillonVertex(b *testing.B) {
	lim = uint32(math.Pow(float64(10), float64(8)))

	for b.Loop() {
		data := Block{
			Letter: 'A',
			X:      1,
			Y:      lim,
			// D:        1,
			// V:        uint32(i3),
			Previous: nil,
		}
		var list []*Block
		list = append(list, &data)
		wg.Add(1)
		b.ResetTimer()
		list[0].RunRule()
		wg.Wait()
	}
	b.ReportMetric(float64(int(lim)*connections)/float64(b.Elapsed().Seconds())/float64(runtime.NumCPU())/math.Pow(float64(10), float64(6)), "mE/PE/s")
	b.ReportMetric(float64(runtime.NumCPU()), "CPUs")
	b.ReportMetric(float64(b.Elapsed().Seconds()), "s")
}
