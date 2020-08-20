# Benchmarking Generic For Each Loop

Until actual generics are available in Go,
to implement a generic loop we need to use reflection.

This benchmark compares how much slower this generic
implementation is in comparison with a function that
only works with a single type of slice or a slice of interfaces.

Since the size of the elements of the list affect the results
greatly we had to split the benchmark in two groups:

The first group test lists of pointers and has 3 methods:

- The unaltered generic method receiving a list of pointers to structs
- An static method that receives a list of pointers to structs
- An static method that receives a list of interfaces
  (which are internally implemented as pointers)

The second group compares the methods that uses copy by value,
this excludes the version using a list of interfaces because these
are implemented internally using pointers, so we have:

- The unaltered generic method
- An static method that receives a list of structs

## Results

You can run the benchmark yourself using the command:

```
  make for-each TIME=15s
```

The result on my machine was as follows:

```
go test -bench=. ./genericforeach/... -benchtime=5s
goos: linux
goarch: amd64
pkg: github.com/vingarcia/go-reflection-benchmark/genericforeach
BenchmarkReflectionWithPtr/testing_GenericForEach-4         	 2777156	      2146 ns/op
BenchmarkReflectionWithPtr/testing_StaticForEachWithPtr-4   	15592579	       372 ns/op
BenchmarkReflectionWithPtr/testing_ForEachWithInterface-4   	15041979	       401 ns/op
BenchmarkReflection/testing_GenericForEach-4                	  797018	      7606 ns/op
BenchmarkReflection/testing_StaticForEach-4                 	 1000000	      5865 ns/op
PASS
ok  	github.com/vingarcia/go-reflection-benchmark/genericforeach	32.851s
```

## Conclusion

It is easy to notice that using lists of pointers,
which also includes lists of interfaces greatly optimizes
the speed of the operations since the first 3 are clearly
faster.

On the first group, using pointers, the generic method
is **576.88%** slower than the static method and **535.16%**
slower than the method with interfaces.

On the second group the difference is less extreme:
The generic version is only **129.68%**.

These results show that if the operation inside the loop is too small
the reflection slowness shows more than otherwise.
Not a surprise, but it interesting to see that even copying such a small struct
the cost changes so much.

This suggests that in many cases using bigger structs would be acceptable.
