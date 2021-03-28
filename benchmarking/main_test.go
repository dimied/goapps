package benchmarking

import "testing"
/*
You can redirect the results to a file and compare the results

benchcmp (deprecated)
go get golang.org/x/tools/cmd/benchcmp

Or benchstat 
go get golang.org/x/perf/cmd/benchstat 
go get rsc.io/benchstat

RUN:
go test -bench=. -count=5 > file.txt

And then:

benchcmp file1.txt file2.txt
benchstat file.txt


To generate profiles, you can use 
-cpuprofile=$FILE
-memprofile=$FILE
-blockprofile=$FILE
# -run=XXX disables tests
go test -run=XXX -bench=. -cpuprofile=myprof.p someapp
go tool pprof someapp myprof.p
*/

// Need this var to avoid optimizations
var result int

func BenchmarkAdd(b *testing.B) {
	// Heavy setup 
	// Call it after heavy setup
	b.ResetTimer()
	// Interested in allocations ?
	b.ReportAllocs()
	
	var res int

	for n:= 0; n < b.N; n++ {
		// We need to store the result 
		// to avoid optimizations (e.g. inlining)
		res = add(1000)
	}
	result =res
	// call cleanup function here or use defer its execution
}

func BenchmarkAdd2(b *testing.B) {
	for n:= 0; n < b.N; n++ {
		add2(1000)
	}
}

func add(n int) int{
	sum := 0
	for i:=0; i < n; i++ {
		sum += i
	}
	return sum
}

func add2(n int) int{
	return n*(n+1)/2
}