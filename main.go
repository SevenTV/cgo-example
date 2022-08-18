package main

import (
	"fmt"
	"log"
	"runtime"
	"test/point"
	"time"

	"net/http"
	_ "net/http/pprof"
)

func test(p *point.Point, q *point.Point) {
	p = point.NewPoint(p, 0.0, 0.0)
	q = point.NewPoint(q, 3.0, 4.0)

	dist := point.Distance(p, q)
	if dist != 5 {
		fmt.Println(p.X(), p.Y(), q.X(), q.Y())
		log.Fatal("Wrong distance")
	}
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	go func() {
		t := time.NewTicker(time.Second)
		for range t.C {
			PrintMemUsage()
		}
	}()

	for i := 0; i < 1e10; i++ {
		test(nil, nil)
	}
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB\tTotalAlloc = %v MiB\tSys = %v MiB\tNumGC = %v\n", bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
