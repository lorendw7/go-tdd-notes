# Go TDD 学习笔记

[English](README.en.md) · [中文教材](https://studygolang.gitbook.io/learn-go-with-tests)

跟随 Chris James 的《Learn Go with Tests》练习测试驱动开发。每个主题使用独立的 Go 包，测试和实现放在同一目录。

**当前进度：Select 章节。** 已加入并发网站检查和网站竞速练习，正在学习 goroutine、channel、`select` 与超时处理。

## 环境准备

项目要求 Go 1.25 或更高版本，目前使用 Go 1.26 练习。安装后先确认命令可用：

```powershell
go version
```

项目由根目录的 [`go.mod`](go.mod) 定义，因此不需要放进 `GOPATH`。进入仓库根目录后，Go 会自动识别模块。

## 快速开始

在仓库根目录运行全部测试：

```powershell
go test ./...
```

运行 `hello` 练习并显示每个测试：

```powershell
go test -v ./lessons/hello
```

运行迭代练习的基准测试：

```powershell
go test -bench "." ./lessons/iteration
```

运行数组与切片练习：

```powershell
go test -v ./lessons/arrays
```

运行结构体与接口练习：

```powershell
go test -v ./lessons/structs
```

运行 Maps 字典练习：

```powershell
go test -v ./lessons/maps
```

运行依赖注入练习：

```powershell
go test -v ./lessons/di
```

运行 Mocking 倒计时练习的测试：

```powershell
go test -v ./lessons/mocking
```

运行并发与 Select 练习：

```powershell
go test -v ./lessons/concurrency ./lessons/select
go test -race ./lessons/concurrency ./lessons/select
go test -bench . -run '^$' ./lessons/concurrency
```

`-race` 用于检测运行过程中出现的数据竞争，需要支持竞态检测的 Go 环境；基准测试使用模拟检查函数，不访问外部网站。

常用选项：

| 命令 | 用途 |
| --- | --- |
| `go test` | 测试当前目录中的包 |
| `go test ./...` | 测试当前模块及所有子目录 |
| `go test -v ./...` | 显示每个测试的运行结果 |
| `go test -count=1 ./...` | 忽略测试缓存，强制重新运行 |
| `go test -cover ./...` | 显示测试覆盖率 |
| `go test -run '^TestHello$' ./lessons/hello` | 只运行指定测试 |
| `go test -v -run 'TestHello/in Spanish' ./lessons/hello` | 只运行指定子测试 |

测试文件以 `_test.go` 结尾，测试函数以 `Test` 开头，并接收 `t *testing.T`。TDD 的基本循环是：先写一个失败的测试，写最少的代码让它通过，再在测试保护下重构。

## 运行 Go 程序

仓库中的大多数章节是供 `go test` 使用的包。可执行程序必须声明 `package main`，并提供 `func main()` 作为入口。依赖注入章节包含一个可运行的 HTTP 服务器入口：

```powershell
go run ./lessons/di
```

启动后访问 <http://localhost:5000>，页面会显示 `Hello, world`。在终端按 `Ctrl+C` 停止服务器。`go run` 会临时编译并运行程序；如需生成 Windows 可执行文件，可以使用：

```powershell
go build -o greeter.exe ./lessons/di
```

运行实际倒计时程序：

```powershell
go run ./lessons/mocking
```

程序每次等待一秒后输出 `3`、`2`、`1` 和 `Go!`，共等待约四秒。测试使用 spy 替代真实休眠，无需等待。

## 包与目录

当前练习的结构如下：

```text
lessons/
└── hello/
    ├── hello.go
    └── hello_test.go
```

两个文件都声明 `package hello`，表示它们属于同一个 Go 包。一个目录中的普通 Go 文件必须使用相同的包名；测试因此可以直接调用同包中的 `Hello` 函数。

## 查看 Go 文档

Go 自带命令行文档工具：

```powershell
go doc ./lessons/hello
go doc ./lessons/hello.Hello
go doc testing.T.Errorf
```

教材中的 `godoc` 网页服务器不是 Go 安装包自带的命令，需要另外安装：

```powershell
go install golang.org/x/tools/cmd/godoc@latest
godoc -http=:8000
```

启动后访问 <http://localhost:8000>，按 `Ctrl+C` 停止服务器。现代 Go 项目也可以使用 [`pkgsite`](https://pkg.go.dev/golang.org/x/pkgsite/cmd/pkgsite) 浏览本地文档。

## 学习记录

| 练习 | 要点 |
| --- | --- |
| [01 问候语](lessons/hello) | `Hello(name, language)` 支持英文和西班牙语，并在名字为空时默认问候 `World`。 |
| [02 整数](lessons/integers) | 用 `Add` 学习整数返回值、`ExampleAdd` 示例测试和可导出函数的文档注释。 |
| [03 迭代](lessons/iteration) | 用 `for` 重复字符串，并通过 `BenchmarkRepeat` 测量平均执行时间。 |
| [04 数组与切片](lessons/arrays) | 用 `Sum`、`SumAll` 和 `SumAllTails` 学习切片、可变参数、`append`、切片尾部和切片比较。 |
| [05 结构体、方法与接口](lessons/structs) | 为矩形、圆形和三角形实现 `Area` 方法，通过 `Shape` 接口和表驱动子测试验证共同的面积行为。 |
| [06 Maps](lessons/maps) | 用 `Dictionary` 实现单词的查找、添加、更新和删除，并用自定义错误区分失败原因。 |
| [07 依赖注入](lessons/di) | 让 `Greet` 接收 `io.Writer`，在测试中写入内存缓冲区，在 HTTP 服务中写入响应。 |
| [08 Mocking](lessons/mocking) | 为倒计时注入 `io.Writer` 和 `Sleeper`，分别验证输出内容与休眠、打印的调用顺序。 |
| [09 并发](lessons/concurrency) | 每个网址由独立 goroutine 检查，通过 channel 汇总结果，由接收方统一写入 map。 |
| [10 Select](lessons/select) | 同时等待两个网站的完成通知和超时信号，用本地测试服务器验证竞速与超时行为。 |

问候语练习用 `t.Run` 将不同场景组织为子测试，用 `t.Helper()` 标记共享的断言辅助函数，并通过 `switch` 选择不同语言的问候前缀。这样测试失败时会指出具体场景和调用测试辅助函数的位置。

基准测试函数必须放在 `_test.go` 文件中，以 `Benchmark` 开头并接收 `*testing.B`。用 `go test -bench "." ./lessons/iteration` 运行；输出中的 `ns/op` 表示每次操作平均耗费的纳秒数。普通 `go test` 只运行功能测试，不会执行基准测试。

数组练习将固定长度的 `[5]int` 改成任意长度的 `[]int`。`SumAll(numbersToSum ...[]int)` 接收任意数量的切片，并为每个切片返回一个总和；`SumAllTails` 使用 `numbers[1:]` 忽略第一个元素，并安全处理空切片。测试用 `reflect.DeepEqual` 比较两个切片的内容。

结构体练习用字段描述 `Rectangle`、`Circle` 和 `Triangle`，并为每种类型定义返回 `float64` 的 `Area()` 方法。只要类型实现这个方法，就会自动满足 `Shape` 接口。测试把名称、形状和期望面积放入表格，通过 `t.Run` 对所有形状运行相同的断言。

Maps 练习把 `map[string]string` 定义为命名类型 `Dictionary`，让查找、添加、更新和删除可以写成它的方法。`DictionaryErr` 实现 `Error() string` 后满足 `error` 接口；`ErrNotFound`、`ErrWordExists` 和 `ErrWordDoesNotExist` 作为稳定的错误值，让测试能够直接比较具体失败原因。共享断言使用 `t.Helper()`，使失败位置指向实际测试场景。

依赖注入练习让 `Greet` 依赖抽象的 `io.Writer`，而不是直接调用 `fmt.Println`。`bytes.Buffer` 和 `http.ResponseWriter` 都实现了这个接口，因此相同的问候逻辑既能被快速测试，也能用于 HTTP 响应。虽然代码属于可执行的 `package main`，测试会直接调用 `Greet`，不需要先启动服务器。

### Mocking：记录调用顺序

`Countdown(out io.Writer, sleeper Sleeper)` 接收输出目标和休眠工具。实际运行时使用 `os.Stdout` 和调用 `time.Sleep` 的 `ConfigurableSleeper`；测试中的 `CountdownOperationsSpy` 只记录调用，不进行真实休眠。

两个子测试分别验证不同的行为：

- **输出内容**：`Countdown(buffer, &CountdownOperationsSpy{})` 用 `bytes.Buffer` 保存输出，检查内容是否为 `3\n2\n1\nGo!`。spy 只承担休眠职责，这个测试不读取它的记录。
- **操作顺序**：`Countdown(spy, spy)` 把同一个 spy 同时作为输出目标和休眠工具。它拥有 `Write` 和 `Sleep` 方法，因此同时满足两个接口；两种操作会写入同一个 `Calls` 切片，保留相对顺序。`reflect.DeepEqual` 检查实际记录是否与期望一致。

只检查输出和休眠次数，无法发现“先连续休眠，再集中打印”的错误。顺序测试要求每次打印前都休眠，包括最后的 `Go!`：

```text
期望：sleep write sleep write sleep write sleep write
少一次休眠：sleep write sleep write sleep write write
```

曾经失败的原因是循环只处理 `3`、`2`、`1`，循环外直接打印了 `Go!`。在 `fmt.Fprint(out, finalWord)` 前补上 `sleeper.Sleep()` 后，最后两次输出之间才有等待。现有子测试名 `sleep after every print` 沿用教材，实际断言的是每次打印前休眠，应以 `want` 中的顺序为准。

早期的 `SpySleeper` 只统计休眠次数，目前仍保留在代码中，但测试已改用 `CountdownOperationsSpy`。顺序断言也覆盖了休眠四次的要求，因此不再需要单独的计数断言。

### 并发：goroutine 与 channel

`checkWebsite(url)` 是普通函数调用，当前代码等它返回才继续；`go checkWebsite(url)` 启动一个 goroutine 执行函数，当前代码继续向下运行。执行顺序不确定，`main` 返回时也不会自动等待其他 goroutine。

channel 是 goroutine 之间传递数据的通道。下面三行分别表示创建、发送和接收，发送与接收需要在能相互配合的 goroutine 中执行：

```go
ch := make(chan string) // 创建只能传递 string 的无缓冲通道
ch <- "完成"            // 发送；等待接收方准备好
message := <-ch        // 接收；等待发送方准备好
```

可以运行的最小配合方式是：

```go
ch := make(chan string)
go func() {
    ch <- "完成"
}()
message := <-ch
fmt.Println(message)
```

`go func(u string) { ... }(url)` 定义并调用匿名函数，把 `url` 作为参数 `u`，并让这次调用在 goroutine 中执行。

在 `CheckWebsites` 中，`WebsiteChecker` 是 `func(string) bool` 的命名函数类型。每个 goroutine 执行 `wc(u)`，将网址和结果打包成 `result{u, wc(u)}`，发送到 `resultChannel`。`result` 的 `string` 和 `bool` 是嵌入字段，因此可以用 `.string` 和 `.bool` 访问。

接收方循环 `len(urls)` 次，每次接收一个结果并写入 map。结果到达顺序不固定，但每个结果都带着网址。只有接收方修改 map，因此避免了多个检查任务同时写这个 map；函数收齐结果后才返回。测试使用替身函数，基准测试为每次检查模拟 20 毫秒延迟。

### Select：等待多个通道

`select` 的 `case` 是通道的发送或接收操作。一个分支就绪就可以执行；多个分支同时就绪时随机选一个；没有就绪分支时等待。如果有 `default`，没有就绪分支时立即执行它。一次 `select` 只执行一个分支，不会自动循环。

教材中的竞速逻辑可以展开为下面的示意代码：

```go
channelA := ping(a)
channelB := ping(b)

select {
case <-channelA:
    return a, nil
case <-channelB:
    return b, nil
case <-time.After(timeout):
    return "", fmt.Errorf("等待超时")
}
```

`ping(url string) chan bool` 接收网址并返回一个布尔通道。函数创建通道，启动 goroutine 执行 `http.Get(url)`，然后立即返回通道。请求调用返回后，goroutine 执行 `ch <- true` 发送通知。`true` 只是完成信号，不代表 HTTP 请求成功。

`case <-ping(a):` 表示调用 `ping(a)` 取得通道，然后等待接收通知，不保存收到的值。进入 `select` 时两个 `ping` 都会调用，不会等第一个请求结束才启动第二个。两个通道都就绪时随机选一个，因此这里不是精确的响应时间测量。

`time.After(timeout)` 返回一个到时会发送通知的通道。如果两个请求都未能在这之前提供可接收的通知，超时分支可以返回错误。`Racer` 使用默认十秒超时，`ConfigurableRacer` 允许测试传入更短的时间。`(winner string, error error)` 是命名返回值列表；`return a, nil` 返回网址和“没有错误”。

测试使用 `httptest.NewServer` 创建本地 HTTP 服务，通过不同延迟控制响应速度。`defer server.Close()` 将关闭服务器的操作安排在当前测试函数返回前执行。这样无需依赖真实网站；超时测试使用毫秒级配置，无需等十秒。

当前保留的是教材的中间实现，后续需要完善：

- `ping` 使用无缓冲通道发送 `true`，落败或超时后的发送没有接收方，会使 goroutine 一直阻塞。可以改为关闭通道发出完成通知，例如使用 `chan struct{}` 配合 `close(ch)`。
- HTTP 错误被忽略，成功响应的 `Body` 也尚未关闭；`measureResponseTime` 是保留的早期实现，当前竞速流程不使用它。
- `select` 超时只停止等待，并不会取消 HTTP 请求。取消请求需要另行传递取消信号，例如使用带超时的 context。

原创练习采用 [MIT 许可证](LICENSE)。
