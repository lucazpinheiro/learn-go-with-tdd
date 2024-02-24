## Learn GO with tests


### useful commands

* `go test` 

run all tests in the project
```
go test ./... 
```

run tests for an specific package
```
go test ./path/to/sub/package
```

run tests with benchmarks
```
go test -bench=. ./path/to/sub/package
```

* `godoc`

generate docs, access on: http://localhost:6060/pkg/

```
godoc -http=:6060
```