package getmapkeys

import (
	"math/rand"
	"strconv"
	"testing"
)

var i map[string]string
var N = 1000

func initInput() {
	if i == nil {
		i = make(map[string]string)
		for a := 0; a <= N; a++ {
			b := strconv.Itoa(rand.Intn(N))
			i[b] = b
		}
	}
}

func BenchmarkWithInit(b *testing.B) {
	initInput()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		o := WithInit(i)
		if len(o) == N {
		}
	}
}

func BenchmarkWithoutInit(b *testing.B) {
	initInput()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		o := WithoutInit(i)
		if len(o) == N {
		}
	}
}
