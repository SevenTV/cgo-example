package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"text/tabwriter"

	"test/point"

	go_point "test/go_point"
	"time"

	"net/http"
	_ "net/http/pprof"
)

func randfloat() float64 {
	return rand.Float64() * 100
}

func test(p *point.Point, q *point.Point) {
	p = point.NewPoint(p, randfloat()+1, randfloat()+2)
	q = point.NewPoint(q, randfloat()+3, randfloat()+4)

	point.Distance(p, q)

	point.DeletePoint(p)
	point.DeletePoint(q)
}

func testGo(p *go_point.Point, q *go_point.Point) {
	p = go_point.NewPoint(p, randfloat()+1, randfloat()+2)
	q = go_point.NewPoint(q, randfloat()+3, randfloat()+4)

	go_point.Distance(p, q)

	go_point.DeletePoint(p)
	go_point.DeletePoint(q)
}

const TEST_SIZE = 10_000_000

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	go func() {
		w := tabwriter.NewWriter(os.Stdout, 10, 1, 1, ' ', 0)
		t := time.NewTicker(time.Second)
		for range t.C {
			PrintMemUsage(w)
		}
	}()

	w := tabwriter.NewWriter(os.Stderr, 10, 1, 1, ' ', 0)
	{
		start := time.Now()
		rand.Seed(0)
		for i := 0; i < TEST_SIZE; i++ {
			testGo(nil, nil)
		}

		fmt.Fprintf(w, "GO-OPS\t%d\t%s\n", int64(float64(time.Second)/float64(time.Since(start))*TEST_SIZE), time.Since(start))
		w.Flush()
	}

	{
		start := time.Now()
		rand.Seed(1)
		for i := 0; i < TEST_SIZE; i++ {
			test(nil, nil)
		}
		fmt.Fprintf(w, "C-OPS\t%d\t%s\n", int64(float64(time.Second)/float64(time.Since(start))*TEST_SIZE), time.Since(start))
		w.Flush()
	}

	select {}
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage(w *tabwriter.Writer) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Fprintf(w, "Alloc = %v MiB\tTotalAlloc = %v MiB\tSys = %v MiB\tNumGC = %v\n", bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC)
	w.Flush()
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
