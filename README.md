# Benchmark Go Combinatorics

## How to run the benchmark

### Local

```shell
go test -bench=. -benchmem ./...
```

### Docker

```shell
docker run -i --rm -v $(pwd):/go/src/github.com/ikngtty/benchmark-go-combinatorics golang go test -bench=. -benchmem github.com/ikngtty/benchmark-go-combinatorics/combinatorics
```

## Result

```
goos: darwin
goarch: amd64
pkg: github.com/ikngtty/benchmark-go-combinatorics/combinatorics
BenchmarkCombinations/Recursive0-4         	       1	5270823668 ns/op	5060925400 B/op	70747903 allocs/op
BenchmarkCombinations/Recursive1-4         	      20	  55564969 ns/op	      96 B/op	       1 allocs/op
BenchmarkCombinations/Recursive2-4         	      19	  58998589 ns/op	     112 B/op	       1 allocs/op
BenchmarkCombinations/WithStack0-4         	       4	 313034402 ns/op	166409778 B/op	 5200303 allocs/op
BenchmarkCombinations/WithSlice0-4         	      26	  44367916 ns/op	    8256 B/op	       9 allocs/op
BenchmarkCombinations/WithCarrying0-4      	      31	  36928434 ns/op	      96 B/op	       1 allocs/op
BenchmarkCombinations/WithCarrying1-4      	      21	  49481598 ns/op	      96 B/op	       1 allocs/op
BenchmarkDupCombinations/Recursive0-4      	       1	4282572542 ns/op	4187600032 B/op	60615826 allocs/op
BenchmarkDupCombinations/Recursive1-4      	      24	  47771265 ns/op	      80 B/op	       1 allocs/op
BenchmarkDupCombinations/Recursive2-4      	      21	  52088408 ns/op	      80 B/op	       1 allocs/op
BenchmarkDupCombinations/WithStack0-4      	       4	 276439302 ns/op	149978604 B/op	 4686827 allocs/op
BenchmarkDupCombinations/WithSlice0-4      	      26	  42335223 ns/op	    8240 B/op	       9 allocs/op
BenchmarkDupCombinations/WithCarrying0-4   	      33	  35562403 ns/op	      80 B/op	       1 allocs/op
BenchmarkDupCombinations/WithCarrying1-4   	      22	  48869797 ns/op	      80 B/op	       1 allocs/op
BenchmarkDupPermutations/Recursive0-4      	       1	2145949059 ns/op	2020474264 B/op	30689223 allocs/op
BenchmarkDupPermutations/Recursive1-4      	      49	  23847135 ns/op	      64 B/op	       1 allocs/op
BenchmarkDupPermutations/WithStack0-4      	       8	 136736002 ns/op	76696018 B/op	 2396747 allocs/op
BenchmarkDupPermutations/WithSlice0-4      	      46	  23043210 ns/op	    2080 B/op	       7 allocs/op
BenchmarkDupPermutations/WithCarrying0-4   	      74	  15678811 ns/op	      64 B/op	       1 allocs/op
BenchmarkDupPermutations/WithCarrying1-4   	      51	  22888059 ns/op	      64 B/op	       1 allocs/op
BenchmarkPermutations/Recursive0-4         	       1	6338217782 ns/op	5168828928 B/op	89707742 allocs/op
BenchmarkPermutations/Recursive1-4         	       1	4109722365 ns/op	3087847792 B/op	85046601 allocs/op
BenchmarkPermutations/Recursive2-4         	       1	3807560486 ns/op	2581434768 B/op	91281958 allocs/op
BenchmarkPermutations/Recursive3-4         	       1	3787541809 ns/op	2339606912 B/op	85046640 allocs/op
BenchmarkPermutations/Recursive4-4         	       1	5024220142 ns/op	2513788976 B/op	95933042 allocs/op
BenchmarkPermutations/Recursive5-4         	       1	1198652505 ns/op	655839184 B/op	19728209 allocs/op
BenchmarkPermutations/Recursive6-4         	       6	 191816791 ns/op	      96 B/op	       2 allocs/op
BenchmarkPermutations/Recursive7-4         	       2	 741371156 ns/op	      80 B/op	       1 allocs/op
BenchmarkPermutations/WithStack0-4         	       1	1004648553 ns/op	315651624 B/op	19728208 allocs/op
BenchmarkPermutations/WithStack1-4         	       2	 572861294 ns/op	157825764 B/op	 9864104 allocs/op
BenchmarkPermutations/WithStack2-4         	       2	 843198178 ns/op	315651816 B/op	19728210 allocs/op
BenchmarkPermutations/WithStack3-4         	       2	 648371094 ns/op	315651576 B/op	 9864106 allocs/op
BenchmarkPermutations/WithStack4-4         	       2	 663686560 ns/op	315651548 B/op	 9864106 allocs/op
BenchmarkPermutations/WithStack5-4         	       2	 969391196 ns/op	557476872 B/op	16099412 allocs/op
BenchmarkPermutations/WithStack6-4         	       1	1170586330 ns/op	315651632 B/op	 9864109 allocs/op
BenchmarkPermutations/WithStack7-4         	       1	1766143475 ns/op	946954664 B/op	29592321 allocs/op
BenchmarkPermutations/WithStack8-4         	       1	3088975270 ns/op	1420431936 B/op	59184625 allocs/op
BenchmarkPermutations/WithSlice0-4         	       5	 234075281 ns/op	     576 B/op	       6 allocs/op
BenchmarkPermutations/WithSlice1-4         	       6	 184274404 ns/op	     336 B/op	       6 allocs/op
BenchmarkPermutations/WithSlice2-4         	       6	 186323035 ns/op	    2112 B/op	       8 allocs/op
BenchmarkPermutations/WithSlice4-4         	       6	 183515108 ns/op	    2496 B/op	       6 allocs/op
BenchmarkPermutations/WithSlice5-4         	       3	 360493641 ns/op	84004413 B/op	 6235310 allocs/op
BenchmarkPermutations/WithSlice6-4         	       2	 657643570 ns/op	    2096 B/op	       7 allocs/op
BenchmarkPermutations/WithSlice7-4         	       6	 183175502 ns/op	    9504 B/op	       9 allocs/op
BenchmarkPermutations/WithSlice8-4         	       1	2023963760 ns/op	946958520 B/op	29592322 allocs/op
BenchmarkPermutations/WithCarrying0-4      	       2	 602414196 ns/op	      80 B/op	       1 allocs/op
BenchmarkPermutations/WithCarrying1-4      	       7	 149944171 ns/op	      96 B/op	       2 allocs/op
BenchmarkPermutations/WithCarrying2-4      	       6	 170103707 ns/op	      96 B/op	       2 allocs/op
```
