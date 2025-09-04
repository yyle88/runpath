[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/runpath/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/runpath/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/runpath)](https://pkg.go.dev/github.com/yyle88/runpath)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/runpath/master.svg)](https://coveralls.io/github/yyle88/runpath?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/runpath.svg)](https://github.com/yyle88/runpath/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/runpath)](https://goreportcard.com/report/github.com/yyle88/runpath)

# runpath

`runpath` package provides func to get the execution location of Go code, including the absolute path of the source file.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Installation

Install the package with:

```shell
go get github.com/yyle88/runpath
```

## Key Features

`runpath` is a runtime path utils package that obtains the absolute path of Go code execution locations.

**Core Capabilities:**

1. **Runtime Path Retrieval** - Get code execution location's absolute path via `runtime.Caller()`
2. **Parent DIR Operations** - Parent DIR related operations through `PARENT`/`DIR` namespace
3. **Testing Support** - `runtestpath` sub-package specifically supports getting source file paths from test files
4. **Path Extension Handling** - Support for changing and removing file extensions

**runpath Advantages:**

Unlike the built-in `filepath.Abs(".")` which doesn't always provide the expected result in certain situations, `runpath` uses Go's `runtime` package to offer precise location tracking, making it especially useful when the exact execution path is needed.

## Core API Overview

**Path Operations:**
- `Path()`, `Current()`, `CurrentPath()` - Get current source file absolute path
- `Name()`, `CurrentName()` - Get current source file name  
- `Skip(skip int)` - Skip specified levels to get caller path
- `GetPathChangeExtension()`, `GetRex()` - Change file extension
- `GetPathRemoveExtension()`, `GetNox()` - Remove .go extension

**DIR Operations:**
- `PARENT.Path()`, `DIR.Path()` - Get parent DIR path
- `PARENT.Join()`, `DIR.Join()` - Join paths
- `PARENT.Up()`, `DIR.UpTo()` - Navigate up DIR structure

**Test Utilities (runtestpath):**
- `SrcPath(t)` - Get tested source file path (from _test.go to corresponding .go file)
- `SrcName(t)` - Get tested source file name
- `SrcPathChangeExtension(t, ext)` - Change tested file extension

## Common Use Cases

- **Dynamic Config File Paths** - Read config.json from config.go dynamically
- **Test Code Generation** - Locate source files when generating code in tests  
- **Accurate Execution Paths** - More reliable than `filepath.Abs(".")` for execution location

## Example Usage

### Get the absolute path of the current source file (execution location abs-path):

```go
path := runpath.Path()
```

This method returns the absolute path of the source file at the point of execution, which is the **execution location** of the current code.

### Get the absolute path of the DIR containing the current source file (execution DIR abs-path):

```go
path := runpath.PARENT.Path()
```

This method returns the absolute path of the DIR where the current source file resides, which can be useful for identifying the **execution DIR abs-path**.

### Get the source file path of the test being executed in testing:

```go
path := runtestpath.SrcPath(t)
```

This method returns the source path of the test file that is currently being run. It is particularly useful when testing and needing to know the exact location of the test file.

---

## Working with Configuration Files

You can use `runpath` to easily build paths to configuration files based on the **execution location** of the code. This is especially useful in tests where different configurations are loaded depending on where the test is being executed.

```go
path := runpath.DIR.Join("config.json")
```

This dynamically constructs the path to `config.json` relative to the **execution DIR abs-path**.

You can also use `PARENT` for similar func:

```go
path := runpath.PARENT.Join("config.json")
```

If you need to navigate up multiple DIR levels from the execution location, use:

```go
path := runpath.PARENT.UpTo(3, "config.json")
```

This will return the path to `config.json` located three levels up from the **execution DIR abs-path**.

---

## Locating Source Code in Test Cases

When running tests, you may need to generate code or reference the source file path dynamically. Here's how you can locate the test file’s source path:

```go
func TestSrcPath(t *testing.T) {
    path := runpath.SrcPath(t)
    t.Log(path)
    require.True(t, strings.HasSuffix(path, "runpath/runtestpath/utils_runtestpath.go"))
}
```

This approach helps when generating code that needs to be placed alongside the original source files based on the **execution location**.

---

## Changing File Extensions Based on Context

You can also change the file extension depending on the test context (e.g., from `.go` to `.json`):

```go
func TestSrcPathChangeExtension(t *testing.T) {
    path := runpath.SrcPathChangeExtension(t, ".json")
    t.Log(path)
    require.True(t, strings.HasSuffix(path, "runpath/runtestpath/utils_runtestpath.json"))
}

func TestSrcRex(t *testing.T) {
    path := SrcRex(t, ".json")
    t.Log(path)
    require.True(t, strings.HasSuffix(path, "runpath/runtestpath/runtestpath.json"))
}
```

This allows you to load different types of files (e.g., configuration files) based on the **execution location** and test requirements.

---

## Function Overview

- `Path()`: Returns the absolute path of the source file where the code is executed, representing the **execution location**.
- `Current()`, `CurrentPath()`, `CurrentName()`, `Name()`: Variants of `Path()` that retrieve the file path or name based on the current execution context.
- `Skip(int)`: Retrieves the path from a specified call frame, useful for getting the execution location of the caller.
- `GetPathChangeExtension()`: Returns the current source file path with a new extension (e.g., changing `.go` to `.json`).
- `GetPathRemoveExtension()`: Returns the current source file path without the `.go` extension.
- `Join()`: Joins the current DIR abs-path with additional path components, dynamically constructing paths based on the **execution location**.
- `Up()`, `UpTo()`: Navigates up the DIR structure a specified number of levels from the **execution location**.

---

## Test-Specific Operations

- `SrcPath(t *testing.T)`: Retrieves the source path of the file being tested.
- `SrcName(t *testing.T)`: Retrieves the name of the source file being tested.
- `SrcPathChangeExtension(t *testing.T, ext string)`: Changes the extension of the test file path (e.g., from `.go` to `.json`).
- `SrcSkipRemoveExtension(t *testing.T)`: Removes the `.go` extension from the test file path.

This package is particularly useful in test files, where you need to reference source code paths or configuration files based on the **execution location** of the test file.

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-08-28 08:33:43.829511 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a bug?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share your use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize by reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo for new releases and features
- 🌟 **Success stories?** Share how this package improved your workflow
- 💬 **General feedback?** All suggestions and comments are welcome

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage interface).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement your changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation for user-facing changes and use meaningful commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a pull request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Happy Coding with this package!** 🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![starring](https://starchart.cc/yyle88/runpath.svg?variant=adaptive)](https://starchart.cc/yyle88/runpath)

Give me stars! Thank you!!!

---
