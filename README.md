# JLUtils

### A utility library useful mostly to myself, because I hate rewriting code with a passion

## Install

```shell
go get github.com/jacoblockett/jlutils
```

## Usage

Import the package:

```go
import "github.com/jacoblockett/jlutils"
```

## API

Slice Helpers:

- [ContainsString](#containsstring)

Parsing:

- [ParseURL](#parseurl)

---

### Slice Helpers

#### ContainsString

iterates over a string slice/array and returns a boolean if a given string argument exists

```go
func ContainsString(collection []string, s string) bool
```

```go
s := []string{"a", "b", "c"}
b := jlutils.ContainsString(s, "b")

fmt.Printf("Result: %t", b)

// Result: true
```

### Parsing

#### ParseURL

Parses a given url and returns a struct with all the parsed bits

```go
type URL struct {
	Protocol  string
	Hostname  string
	Subdomain string
	Domain    string
	Port      int
	Path      string
	Query     string
	Fragment  string
}

func ParseURL(url string) URL{}
```

```go
u := jlutils.ParseURL("https://test.abc.com/somepath?lang=go")

fmt.Printf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s", u.Protocol, u.Hostname, u.Subdomain, u.Domain, u.Port, u.Path, u.Query, u.Fragment)

/*
https
test.abc.com
test
abc.com
0
somepath
lang=go

*/
```
