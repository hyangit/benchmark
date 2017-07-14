# benchmark
### getmapkeys
benchmark `append with init []` and `append without init []`

```
BenchmarkWithInit-4             100000          16551 ns/op
BenchmarkWithoutInit-4          50000           25053 ns/op
```

### goroutine
benchmark `go func` and `chan <- work`
```
BenchmarkGoRoutine              300000          8475 ns/op
BenchmarkGoChan                 1000000         2678 ns/op
```

### distinct
benchmark `by map` and `by slice`
```
BenchmarkByMap-4                20000           72772 ns/op
BenchmarkBySlice-4              5000            263969 ns/op
```