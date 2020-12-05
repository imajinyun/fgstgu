package main

import (
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

const (
	row = 10000
	col = 10000
)

func fillMatrix(m *[row][col]int) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m[i][j] = s.Intn(999999)
		}
	}
}

func calculate(m *[row][col]int) {
	for i := 0; i < row; i++ {
		tmp := 0
		for j := 0; j < col; j++ {
			tmp += m[i][j]
		}
	}
}

/**
 * -> cd src/d/d02
 * -> go build main.go
 * -> ./main
 * -> go tool pprof main cpu.prof
 */
func main() {
	cpu, err := os.Create("cpu.prof")
	if err != nil { log.Fatal("Could not create CPU profile:", err) }
	if err := pprof.StartCPUProfile(cpu); err != nil {
		log.Fatal("Could not start CPU profile", err)
	}
	defer pprof.StopCPUProfile()

	x := [row][col]int{}
	fillMatrix(&x)
	calculate(&x)

	mem, err := os.Create("mem.prof")
	if err != nil { log.Fatal("Could not create memory profile", err) }
	runtime.GC()
	if err := pprof.WriteHeapProfile(mem); err != nil {
		log.Fatal("Could not create memory profile", err)
	}
	mem.Close()

	grt, err := os.Create("grt.prof")
	if err != nil { log.Fatal("Could not create groutine") }
	if gprof := pprof.Lookup("goroutine"); gprof == nil {
		log.Fatal("Cound not write goroutine profile")
	} else {
		gprof.WriteTo(grt, 0)
	}
	grt.Close()
}
