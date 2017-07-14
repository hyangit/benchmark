package distinct

import (
	"math/rand"
	"strconv"
	"testing"
)

var i []string
var N = 1000
var M = 100

func initInput() {
	if i == nil {
		i = make([]string, 0, N)
		for a := 0; a <= N; a++ {
			b := strconv.Itoa(rand.Intn(M))
			i = append(i, b)
		}
	}
}

func BenchmarkByMap(b *testing.B) {
	initInput()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		o := ByMap(i)
		if len(o) != N {
		}
	}
}

func BenchmarkBySlice(b *testing.B) {
	initInput()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		o := BySlice(i)
		if len(o) != N {
		}
	}
}
