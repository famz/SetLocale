SetLocale for Golang
====================

A wrapper for setlocale(3).

Usage
-----

Just import "github.com/famz/SetLocale", then call SetLocale:

```go
SetLocale.SetLocale(SetLocale.LC_*, "<your-locale>")
```

Install
-------

Simply run `go get github.com/famz/SetLocale`

Example
-------

```go
package main
import (
    "github.com/famz/SetLocale"
    "fmt"
)

func main() {
    SetLocale.SetLocale(SetLocale.LC_ALL, "")
}
```
