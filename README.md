# benchmark
### getmapkeys
benchmark `append with init []` and `append without init []`

```
BenchmarkWithInit-4      	 5000000	       297 ns/op
BenchmarkWithoutInit-4   	 2000000	       724 ns/op
```

### goroutine
benchmark `go func` and `chan <- work`
```
BenchmarkGoRoutine 		 300000 		 6012 ns/op
BenchmarkGoChan 		 1000000 		 2249 ns/op
```