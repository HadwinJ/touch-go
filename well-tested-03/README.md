# Benchmark tests
* Writing Benchmark Tests
package main
func BenchmarkSHA1(b *testing.B) {
	data := []byte("Mary had a little lamb")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha1.Sum(data)
	}
}
* Key testing.B Members
b.N
b.StartTimer, b.StopTimer, b.ResetTimer
b.RunParallel

* Running
go test -bench .
go test -bench . -benchtime 10s

# Profiling tests
* go test -benchmem  // Report memory allocation statistics for benchmarks
* go test -trace {trace.out}    // Record execution trace to {trace.out} for analysis
* go test -{type}profile {file} // Generate profile of requested type: block, cover, cpu, mem, mutex
* Examples
go test -bench 512
go test -bench 512 -benchmem
go test -bench Alloc -memprofile profile.out
go tool pprof profile.out   // svg