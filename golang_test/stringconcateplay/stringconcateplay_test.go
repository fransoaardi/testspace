package main

import (
	"testing"
	"bytes"
)

// Result of  `go test -bench . `
//
//	goos: darwin
//	goarch: amd64
//	BenchmarkAppendString-8      	 10000000	       178 ns/op
//	BenchmarkBuffWrite-8         	200000000	      8.34 ns/op
//	BenchmarkBuffWriteString-8   	200000000	      7.18 ns/op
//	BenchmarkConcatString-8      	  1000000	     44952 ns/op
//

// Using append to slice
func BenchmarkAppendString(b *testing.B){
	var str []string
	for i:=0; i<b.N; i++{
		str = append(str, "a")
	}
}

// Using bytes.Buffer and write
func BenchmarkBuffWrite(b *testing.B){
	var buf bytes.Buffer
	var a = []byte("a")
	for i:=0; i<b.N; i++{
		buf.Write(a)
	}
}

// Using bytes.Buffer WriteString
func BenchmarkBuffWriteString(b *testing.B){
	var buf bytes.Buffer
	for i:=0; i<b.N; i++{
		buf.WriteString("a")
	}
}

// Using concatenate and very slow
func BenchmarkConcatString(b *testing.B){
	var str string
	for i:=0; i<b.N; i++{
		str+="a"
	}
}

