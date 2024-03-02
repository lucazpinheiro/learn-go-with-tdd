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

run tests and get coverage report
```
go test -cover ./path/to/sub/package

# or

go test -cover ./...
```

* `godoc`

generate docs, access on: http://localhost:6060/pkg/

```
godoc -http=:6060
```

### how to know what to test (extracted from 'mocking chapter')

* The definition of refactoring is that the code changes but the behaviour stays the same. If you have decided to do some refactoring in theory you should be able to make the commit without any test changes. So when writing a test ask yourself
    * Am I testing the behaviour I want, or the implementation details?
    * If I were to refactor this code, would I have to make lots of changes to the tests?
* Although Go lets you test private functions, I would avoid it as private functions are implementation detail to support public behaviour. Test the public behaviour. Sandi Metz describes private functions as being "less stable" and you don't want to couple your tests to them.
* I feel like if a test is working with more than 3 mocks then it is a red flag - time for a rethink on the design
* Use spies with caution. Spies let you see the insides of the algorithm you are writing which can be very useful but that means a tighter coupling between your test code and the implementation. Be sure you actually care about these details if you're going to spy on them