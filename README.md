# Benchmarking Reflection in Go

This repository was created to test how much slower is the code when we use reflections.

## Results

You can run the benchmark yourself using the command:

```bash
make bench
```

The result on my machine was as follows:

```
go test -bench=. -benchtime=15s
goos: linux
goarch: amd64
pkg: github.com/vingarcia/go-reflection-benchmark
BenchmarkReflection/testing_toMapWithNoReflection-4         	25710033	       630 ns/op
BenchmarkReflection/testing_toMap-4                         	14361529	      1204 ns/op
BenchmarkReflection/testing_toMapWithCachedType-4           	15633762	      1043 ns/op
BenchmarkReflection/testing_toMapUsingTag-4                 	 9483343	      1744 ns/op
BenchmarkReflection/testing_toMapUsingTagWithCachedType-4   	10080223	      1575 ns/op
BenchmarkReflection/testing_toMapUsingMethod&DuckTyping-4   	26961108	       644 ns/op
PASS
ok  	github.com/vingarcia/go-reflection-benchmark	107.135s
```

## Conclusion

Comparing all the attempts with the first one that
uses no abstractions we can conclude that:

1. Using duck-typed methods is very fast taking only
   **2.22%** longer

2. Reflection using only the name of the struct attributes
   is **91%** slower (takes almost twice the time)

3. Reflection using the tags on the struct are surprisingly
   slower taking up to **176%** longer (almost 3 times slower)

4. It is possible to slightly improve the performance of the
   process by caching the Type of the struct instead of reading it everytime.

   Caching on the method that uses the attribute names improved the speed in **13.37%**

   Caching on the method that uses the tags improved the speed in **9.69%**

