# Go TDD 学习笔记

[English](README.en.md) · [中文教材](https://studygolang.gitbook.io/learn-go-with-tests)

跟随 Chris James 的《Learn Go with Tests》练习测试驱动开发。每个主题使用独立的 Go 包，测试和实现放在同一目录。

## 本地练习

使用 Go 1.26，在仓库根目录运行：

```sh
go test ./...
go test -v ./lessons/hello
```

## 学习记录

| 练习 | 要点 |
| --- | --- |
| [01 问候语](lessons/hello) | 用表驱动测试描述两种行为：传入姓名时问候该人，空字符串时问候世界。 |

下一步：为只含空格的姓名编写一个失败的测试，让它通过，再在测试保护下重构。

原创练习采用 [MIT 许可证](LICENSE)。
