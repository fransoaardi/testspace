package mapsyncplay

import "testing"

func BenchmarkMMap_Get(b *testing.B) {

	t := NewMMap()

	for i:=0; i<b.N; i++{
		t.Get("a")
	}

}

func BenchmarkMMap_Get2(b *testing.B) {
	t := NewNoMMap()

	for i:=0; i<b.N; i++{
		t.Get("a")
	}
}

func BenchmarkMMap_Get3(b *testing.B) {
	t := NewRMMap()

	for i:=0; i<b.N; i++{
		t.Get("a")
	}
}