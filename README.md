# Go TDD 学习笔记

[English](README.en.md) · [中文教材](https://studygolang.gitbook.io/learn-go-with-tests)

跟随 Chris James 的《Learn Go with Tests》练习测试驱动开发。每个主题使用独立的 Go 包，测试和实现放在同一目录。

**当前进度：Maps 章节。** 已完成问候语、整数、迭代、数组以及结构体与接口练习，正在学习用 map 构建可增删改查的字典。

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

问候语练习用 `t.Run` 将不同场景组织为子测试，用 `t.Helper()` 标记共享的断言辅助函数，并通过 `switch` 选择不同语言的问候前缀。这样测试失败时会指出具体场景和调用测试辅助函数的位置。

基准测试函数必须放在 `_test.go` 文件中，以 `Benchmark` 开头并接收 `*testing.B`。用 `go test -bench "." ./lessons/iteration` 运行；输出中的 `ns/op` 表示每次操作平均耗费的纳秒数。普通 `go test` 只运行功能测试，不会执行基准测试。

数组练习将固定长度的 `[5]int` 改成任意长度的 `[]int`。`SumAll(numbersToSum ...[]int)` 接收任意数量的切片，并为每个切片返回一个总和；`SumAllTails` 使用 `numbers[1:]` 忽略第一个元素，并安全处理空切片。测试用 `reflect.DeepEqual` 比较两个切片的内容。

结构体练习用字段描述 `Rectangle`、`Circle` 和 `Triangle`，并为每种类型定义返回 `float64` 的 `Area()` 方法。只要类型实现这个方法，就会自动满足 `Shape` 接口。测试把名称、形状和期望面积放入表格，通过 `t.Run` 对所有形状运行相同的断言。

Maps 练习把 `map[string]string` 定义为命名类型 `Dictionary`，让查找、添加、更新和删除可以写成它的方法。`DictionaryErr` 实现 `Error() string` 后满足 `error` 接口；`ErrNotFound`、`ErrWordExists` 和 `ErrWordDoesNotExist` 作为稳定的错误值，让测试能够直接比较具体失败原因。共享断言使用 `t.Helper()`，使失败位置指向实际测试场景。

原创练习采用 [MIT 许可证](LICENSE)。
