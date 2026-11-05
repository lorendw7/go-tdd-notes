# Go TDD study notes

[中文](README.md) · [English book](https://quii.gitbook.io/learn-go-with-tests)

Personal exercises following Chris James's *Learn Go with Tests*. Each topic has its own Go package with tests alongside the implementation.

**Current progress: Mocking.** The earlier dependency injection exercise is complete; the countdown exercise now uses test doubles to verify output content and operation order.

## Prerequisites

The module requires Go 1.25 or newer and is currently practiced with Go 1.26. Confirm the installation first:

```powershell
go version
```

The root [`go.mod`](go.mod) defines the project, so the repository does not need to live inside `GOPATH`. Go discovers the module automatically when commands run from its directory.

## Quick start

Run every test from the repository root:

```powershell
go test ./...
```

Run the `hello` exercise with individual test results:

```powershell
go test -v ./lessons/hello
```

Run the iteration benchmark:

```powershell
go test -bench "." ./lessons/iteration
```

Run the arrays and slices exercise:

```powershell
go test -v ./lessons/arrays
```

Run the structs and interfaces exercise:

```powershell
go test -v ./lessons/structs
```

Run the maps dictionary exercise:

```powershell
go test -v ./lessons/maps
```

Run the dependency injection exercise:

```powershell
go test -v ./lessons/di
```

Run the mocking countdown tests:

```powershell
go test -v ./lessons/mocking
```

Useful variations:

| Command | Purpose |
| --- | --- |
| `go test` | Test the package in the current directory |
| `go test ./...` | Test the current module and all subdirectories |
| `go test -v ./...` | Show each test result |
| `go test -count=1 ./...` | Bypass the test cache and run again |
| `go test -cover ./...` | Report test coverage |
| `go test -run '^TestHello$' ./lessons/hello` | Run one named test |
| `go test -v -run 'TestHello/in Spanish' ./lessons/hello` | Run one named subtest |

Test files end in `_test.go`; test functions start with `Test` and accept `t *testing.T`. The basic TDD loop is to write a failing test, add the smallest implementation that passes, and then refactor under the protection of the tests.

## Run a Go program

Most chapters in this repository are packages exercised with `go test`. An executable program must declare `package main` and provide `func main()` as its entry point. The dependency injection chapter includes a runnable HTTP server:

```powershell
go run ./lessons/di
```

Open <http://localhost:5000> after it starts to see `Hello, world`, and press `Ctrl+C` in the terminal to stop it. `go run` compiles and runs a temporary executable; to create a Windows executable instead, use:

```powershell
go build -o greeter.exe ./lessons/di
```

Run the actual countdown:

```powershell
go run ./lessons/mocking
```

The program waits one second before each output: `3`, `2`, `1`, and `Go!`, for about four seconds in total. Tests replace real sleeping with a spy and do not wait.

## Packages and directories

The current exercise has this structure:

```text
lessons/
└── hello/
    ├── hello.go
    └── hello_test.go
```

Both files declare `package hello`, so they belong to the same Go package. Ordinary Go files in one directory must use the same package name; the test can therefore call `Hello` directly.

## Read Go documentation

Go includes a command-line documentation tool:

```powershell
go doc ./lessons/hello
go doc ./lessons/hello.Hello
go doc testing.T.Errorf
```

The book's `godoc` web server is not bundled with Go and must be installed separately:

```powershell
go install golang.org/x/tools/cmd/godoc@latest
godoc -http=:8000
```

Open <http://localhost:8000> after it starts and press `Ctrl+C` to stop it. Modern Go projects can also use [`pkgsite`](https://pkg.go.dev/golang.org/x/pkgsite/cmd/pkgsite) to browse local documentation.

## Learning log

| Exercise | Observation |
| --- | --- |
| [01 Greetings](lessons/hello) | `Hello(name, language)` supports English and Spanish, defaulting an empty name to `World`. |
| [02 Integers](lessons/integers) | Use `Add` to learn integer return values, the `ExampleAdd` example test, and documentation comments for exported functions. |
| [03 Iteration](lessons/iteration) | Use `for` to repeat a string and `BenchmarkRepeat` to measure its average execution time. |
| [04 Arrays and slices](lessons/arrays) | Use `Sum`, `SumAll`, and `SumAllTails` to learn slices, variadic arguments, `append`, slice tails, and slice comparison. |
| [05 Structs, methods, and interfaces](lessons/structs) | Implement `Area` methods for rectangles, circles, and triangles, then verify their shared behavior through the `Shape` interface and table-driven subtests. |
| [06 Maps](lessons/maps) | Use `Dictionary` to search, add, update, and delete words, with custom errors distinguishing failure cases. |
| [07 Dependency injection](lessons/di) | Pass an `io.Writer` to `Greet`, writing to an in-memory buffer in tests and an HTTP response in the server. |
| [08 Mocking](lessons/mocking) | Inject `io.Writer` and `Sleeper` into a countdown, testing output content separately from the order of sleeping and printing. |

The greetings exercise uses `t.Run` to organize scenarios as subtests, `t.Helper()` to mark the shared assertion helper, and a `switch` to choose the language-specific greeting prefix. Failures can therefore identify both the scenario and the assertion call site.

Benchmark functions must live in a `_test.go` file, start with `Benchmark`, and accept `*testing.B`. Run one with `go test -bench "." ./lessons/iteration`; the reported `ns/op` is the average number of nanoseconds per operation. Plain `go test` runs functional tests but does not execute benchmarks.

The arrays exercise replaces a fixed-length `[5]int` with an arbitrary-length `[]int`. `SumAll(numbersToSum ...[]int)` accepts any number of slices and returns one total for each; `SumAllTails` uses `numbers[1:]` to omit the first element while safely handling empty slices. Tests use `reflect.DeepEqual` to compare slice contents.

The structs exercise describes `Rectangle`, `Circle`, and `Triangle` with fields and gives each type an `Area() float64` method. Any type implementing that method satisfies `Shape` automatically. A table of names, shapes, and expected areas lets `t.Run` apply the same assertion to every shape.

The maps exercise defines `map[string]string` as the named `Dictionary` type so search, add, update, and delete behavior can be expressed as methods. `DictionaryErr` satisfies `error` by implementing `Error() string`; the stable `ErrNotFound`, `ErrWordExists`, and `ErrWordDoesNotExist` values let tests compare exact failure cases. Shared assertions call `t.Helper()` so failures point to the scenario that invoked them.

The dependency injection exercise makes `Greet` depend on the abstract `io.Writer` instead of calling `fmt.Println` directly. Both `bytes.Buffer` and `http.ResponseWriter` implement that interface, so the same greeting logic is easy to test and reuse in an HTTP response. Although the code is an executable `package main`, tests call `Greet` directly and never need to start the server.

### Mocking: recording operation order

`Countdown(out io.Writer, sleeper Sleeper)` accepts an output destination and a sleeping dependency. The executable uses `os.Stdout` and a `ConfigurableSleeper` that calls `time.Sleep`; the test's `CountdownOperationsSpy` records calls without actually sleeping.

Two subtests verify different behaviors:

- **Output content**: `Countdown(buffer, &CountdownOperationsSpy{})` captures output in a `bytes.Buffer` and checks for `3\n2\n1\nGo!`. The spy only handles sleeping, and this test does not inspect its records.
- **Operation order**: `Countdown(spy, spy)` passes the same spy as both dependencies. Its `Write` and `Sleep` methods satisfy both interfaces and append to the same `Calls` slice, preserving their relative order. `reflect.DeepEqual` compares the recorded sequence with the expected sequence.

Checking output and sleep counts alone cannot detect an implementation that sleeps repeatedly before printing all the numbers together. The order test requires a sleep before every output, including `Go!`:

```text
Expected:      sleep write sleep write sleep write sleep write
Missing sleep: sleep write sleep write sleep write write
```

The earlier failure occurred because the loop handled only `3`, `2`, and `1`, then printed `Go!` immediately after leaving the loop. Adding `sleeper.Sleep()` before `fmt.Fprint(out, finalWord)` restores the wait between the final two outputs. The existing subtest name, `sleep after every print`, follows the book, but its assertion actually requires sleeping before each print; follow the sequence in `want`.

The earlier `SpySleeper` only counts sleep calls. It remains in the code, but the tests now use `CountdownOperationsSpy`. The order assertion also verifies four sleep calls, so a separate count assertion is no longer needed.

Original exercises are available under the [MIT license](LICENSE).
