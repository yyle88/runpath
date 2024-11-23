[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/runpath/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/runpath/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/runpath)](https://pkg.go.dev/github.com/yyle88/runpath)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/runpath/master.svg)](https://coveralls.io/github/yyle88/runpath?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/runpath.svg)](https://github.com/yyle88/runpath/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/runpath)](https://goreportcard.com/report/github.com/yyle88/runpath)

# runpath

The `runpath` package provides functionality to get the execution location of Go code, including the absolute path of the source file and the line number.

## README

[中文说明](README.zh.md)

The package name is `runpath`, which utilizes the Go `runtime` package. means `executionlocation`.

You can also use the built-in approach:
```go
filepath.Abs(".") // Get the current directory
```
However, this doesn’t always provide the expected result in certain situations.

### Usage

```shell
go get github.com/yyle88/runpath
```

#### Example usage:

```go
path := runpath.Path()
```
This will return the absolute path of the current source file where the code is running.

#### Reading Configuration Files

You can also use the package to easily build paths to configuration files:

```go
path := runpath.DIR.Join("config.json")
```

```go
path := runpath.PARENT.Join("config.json")
```
This is especially useful in test cases, where different configurations may be loaded depending on the test.

#### Locating Source Code in Test Cases

If you need to generate source code in your tests and reference the source file path, you can use the following approach:

```go
func TestSrcPath(t *testing.T) {
    path := runpath.SrcPath(t)
    t.Log(path)
    require.True(t, strings.HasSuffix(path, "runpath/runtestpath/utils_runtestpath.go"))
}
```
This helps when generating source code that needs to be placed alongside the original files.

#### Modifying File Extensions

You can change the file extension based on the test context:

```go
func TestSrcPathChangeExtension(t *testing.T) {
    path := runpath.SrcPathChangeExtension(t, ".json")
    t.Log(path)
    require.True(t, strings.HasSuffix(path, "runpath/runtestpath/utils_runtestpath.json"))
}
```
This is particularly useful for loading different file types, such as configuration files.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Give Stars

Feel free to contribute or improve the package! Your stars and pull requests are welcome.

## Thank You

If you find this package valuable, please give it a star on GitHub! Thank you!!!

---

### Explanation of Functions:

- `Path()`: Returns the absolute path of the source file at the point of execution.
- `Current()`, `CurrentPath()`, `CurrentName()`, `Name()`: Variations of `Path()` to fetch the file path or name based on the current execution context.
- `Skip(int)`: Allows you to get the path from a specified call frame (useful for getting caller locations).
- `GetPathChangeExtension()`: Returns the path of the current source file with a new extension, e.g., changing `.go` to `.json`.
- `GetPathRemoveExtension()`: Returns the path of the current source file without the `.go` extension.
- `Join()`: Joins the current directory with additional path components, useful for building paths dynamically.
- `Up()`, `UpTo()`: Navigate up the directory structure a specified number of levels.

### For Test File Specific Operations:
- `SrcPath(t *testing.T)`: Gets the source path of the file being tested.
- `SrcName(t *testing.T)`: Gets the name of the source file being tested.
- `SrcPathChangeExtension(t *testing.T, ext string)`: Changes the file extension of the test file path (e.g., from `.go` to `.json`).
- `SrcSkipRemoveExtension(t *testing.T)`: Removes the `.go` extension from the test file path.

This package is intended for use in test files where you need to reference source code paths or configuration files based on the location of your test file.

--- 

Give me stars! Thank you!!!
