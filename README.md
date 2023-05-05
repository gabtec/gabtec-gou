# Gabtec-Gou

![](https://img.shields.io/github/license/gabtec/gabtec-gou) ![](https://img.shields.io/github/v/release/gabtec/gabtec-gou) ![](https://img.shields.io/github/go-mod/go-version/gabtec/gabtec-gou) ![](https://img.shields.io/github/actions/workflow/status/gabtec/gabtec-gou/tests.yaml)

This is my first published Go package.

It's called gabtec-gou, because "gou" stands for "Go Utils" :smile:

# Usage

```sh
# create a new project
# after you have your go.mod file setup
# run:
$ go get -u github.com/gabtec/gabtec-gou
```

Then use it, like this:

```go
// e.g. in your main.go file

package main

import "github.com/gabtec/gabtec-gou"

func main() {
  myVar := gou.GetEnv("TheEnvVarName_aka_key", "Some_default_value_in_case_envVar_not_set_or_not_found")
  //...
}
```

# Testing

```sh
# to run tests, do:
$ go test -v ./...
```

# License

[MIT](https://opensource.org/license/mit/)
