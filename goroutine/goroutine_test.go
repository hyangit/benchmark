package goroutine

import (
	"fmt"
	"testing"
	"time"
)

var data chan bool
var finishCh chan bool

func f(cb func()) {
	// mem op
	var a []int
	for i := 0; i < 100; i++ {
		a = append(a, i)
	}

	// sleep
	time.Sleep(time.Duration(1) * time.Microsecond)

	// cpu op
	for i := 0; i < 1000; i++ {
	}

	cb()
}

func callback() {
	<-data
	if len(data) == 0 {
		finishCh <- true
	}
}

func run(way, chanSize, N int) {
	name := "BenchmarkGoRoutine"
	if way == GoWayChan {
		name = "BenchmarkGoChan"
	}
	data = make(chan bool, N)
	finishCh = make(chan bool, 1)
	g := New(way, chanSize)
	for n := 0; n < N; n++ {
		data <- true
	}
	tm := time.Now()
	for n := 0; n < N; n++ {
		g.Go(callback, f)
	}
	<-finishCh
	fmt.Println(name, "\t\t", N, "\t\t", time.Since(tm).Nanoseconds()/int64(N), "ns/op")
}

func TestRoutine(t *testing.T) {
	run(GoWayRoutine, 0, 300000)
}

func TestChan(t *testing.T) {
	run(GoWayChan, 1024, 1000000)
}
