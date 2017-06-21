package getmapkeys

import (
	"testing"
)

var i = map[string]string{
	"a": "aa",
	"b": "bb",
	"c": "cc",
	"d": "dd",
	"e": "ee",
	"f": "ff",
	"g": "gg",
	"h": "hh",
	"i": "ii",
	"j": "jj",
}

func BenchmarkWithInit(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var o []string
		GetMapKeys(i, GetWayWithInit, o)
	}
}

func BenchmarkWithoutInit(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var o []string
		GetMapKeys(i, GetWayWithoutInit, o)
	}
}
