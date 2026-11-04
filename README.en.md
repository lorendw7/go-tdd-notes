# Go TDD study notes

[中文](README.md) · [English book](https://quii.gitbook.io/learn-go-with-tests)

Personal exercises following Chris James's *Learn Go with Tests*. Each topic has its own Go package with tests alongside the implementation.

## Run locally

With Go 1.26 installed, run these commands from the repository root:

```sh
go test ./...
go test -v ./lessons/hello
```

## Learning log

| Exercise | Observation |
| --- | --- |
| [01 Greetings](lessons/hello) | A table-driven test specifies two behaviors: greet the named person or greet the world when the name is empty. |

Next exercise: write a failing test for a whitespace-only name, make it pass, then refactor with the tests running.

Original exercises are available under the [MIT license](LICENSE).
