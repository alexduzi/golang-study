package main

import "testing"

// execute using -> go test -bench=.
// go test -bench=. -count=10
// go test -bench=. -cpuprofile cpu.prof -count=3
// go test -bench=. -memprofile mem.prof -count=5
// go tool pprof cpu.prof -> profiling tool
// inside pprof tool execute those commands:
//   top -> diagnosis
//   list NameOfFunction -> will print func diagnosis
//   web -> will open svg on browser with all tracing calls

func BenchmarkHeavyComputation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HeavyComputation()
	}
}

func BenchmarkMemoryIntensiveOperation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MemoryIntensiveOperation()
	}
}
