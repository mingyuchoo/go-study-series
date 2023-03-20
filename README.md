<p align="center">
  <a href="https://github.com/mingyuchoo/go-study-series/blob/main/LICENSE"><img alt="license" src="https://img.shields.io/github/license/mingyuchoo/go-study-series"/></a>
  <a href="https://github.com/mingyuchoo/go-stury-series/issues"><img alt="Issues" src="https://img.shields.io/github/issues/mingyuchoo/go-stury-series?color=appveyor" /></a>
  <a href="https://github.com/mingyuchoo/go-stury-series/pulls"><img alt="GitHub pull requests" src="https://img.shields.io/github/issues-pr/mingyuchoo/go-stury-series?color=appveyor" /></a>
</p>

# go-stury-series

# How to create a project

```bash
$ mkdir <project-name>
$ cd <project-name>
$ touch main.go
```

`main.go` file

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

```bash
$ go run main.go
hello world

$ go build main.go
$ ls
main main.go

$ ./main
hello world
```
