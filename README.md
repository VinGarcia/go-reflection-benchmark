# Benchmarking Reflection in Go

This repository was created compare the performance of different techniques
for serializing structs as maps in Go.

The motivation for these tests arrived when studying the interfaces of
different database tools such as GORM, Gorp, the bigquery library for Go
and others.

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
BenchmarkReflection/testing_toMapWithNoReflection-4         	34115180	       523 ns/op
BenchmarkReflection/testing_toMap-4                         	17606706	      1000 ns/op
BenchmarkReflection/testing_toMapWithCachedType-4           	19446636	       912 ns/op
BenchmarkReflection/testing_toMapUsingTag-4                 	10946109	      1496 ns/op
BenchmarkReflection/testing_toMapUsingTagWithCachedType-4   	12495646	      1408 ns/op
BenchmarkReflection/testing_toMapUsingCachedTagNames-4      	28003707	       598 ns/op
BenchmarkReflection/testing_toMapUsingMethod&DuckTyping-4   	34439281	       607 ns/op
PASS
ok  	github.com/vingarcia/go-reflection-benchmark	131.602s
```

## Conclusion

Comparing all the attempts with the first one that
uses no abstractions we can conclude that:

1. Using duck-typed methods is very fast taking only
   **16.06%** longer

2. Reflection using only the name of the struct attributes
   is **91%** slower (takes almost twice the time)

3. Reflection using the tags on the struct are surprisingly
   slower taking up to **186%** longer (almost 3 times slower)

4. It is possible to slightly improve the performance of the
   process by caching the Type of the struct instead of reading it everytime.

   Caching on the method that uses the attribute names improved the speed in **8.79%**

   Caching on the method that uses the tags improved the speed in **5.8%**

5. Since using Tags was expensive I added the function:

   - toMapUsingCachedTagNames

   This function caches only the names we need saving calls to `Field(i).Tag.Get("foo")`.
   And it is really fast faster even than the duck-typing technique with a slowdown of only
   **14.34%** which is awesome.
