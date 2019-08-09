# Errlog: reduce debugging time while programming [![Go Report Card](https://goreportcard.com/badge/github.com/snwfdhmp/errlog)](https://goreportcard.com/report/github.com/snwfdhmp/errlog) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [![Documentation](https://godoc.org/github.com/snwfdhmp/errlog?status.svg)](http://godoc.org/github.com/snwfdhmp/errlog) [![GitHub issues](https://img.shields.io/github/issues/snwfdhmp/errlog.svg)](https://github.com/snwfdhmp/errlog/issues) [![license](https://img.shields.io/github/license/snwfdhmp/errlog.svg?maxAge=6000)](https://github.com/snwfdhmp/errlog/LICENSE)

![Example](https://i.imgur.com/Ulf1RGw.png)

## Introduction

Use errlog to improve error logging and speed up your debugging while programming :

- Highlight source code
- Detect and point out which func call is causing the fail
- Pretty stack trace
- No-op mode for production
- Easy implementation, adaptable logger
- Plug to any existing project without changing your/your teammates habits
- Plug to your existing logging system

## Get started

### Install

```shell
go get github.com/snwfdhmp/errlog
```

### Usage

Replace your `if err != nil` with `if errlog.Debug(err)` to add debugging informations.

```golang
func someFunc() {
    //...
    if errlog.Debug(err) { // will debug & pass if err != nil, will ignore if err == nil
        return
    }
}
```

In production, call `errlog.Disable(true)` to enable no-op (equivalent to `if err != nil`)

## Tweak as you need

You can configure your own logger with the following options :

```golang
type Config struct {
    PrintFunc          func(format string, data ...interface{}) //Printer func (eg: fmt.Printf)
    LinesBefore        int  //How many lines to print *before* the error line when printing source code
    LinesAfter         int  //How many lines to print *after* the error line when printing source code
    PrintStack         bool //Shall we print stack trace ? yes/no
    PrintSource        bool //Shall we print source code along ? yes/no
    PrintError         bool //Shall we print the error of Debug(err) ? yes/no
    ExitOnDebugSuccess bool //Shall we os.Exit(1) after Debug has finished logging everything ? (doesn't happen when err is nil). Will soon be replaced by ExitFunc to enable panic-ing the current goroutine. (if you need this quick, please open an issue)
}
```

> As we don't yet update automatically this README immediately when we add new features, this definition may be outdated. (Last update: 2019/08/07)
> [See the struct definition in godoc.org](https://godoc.org/github.com/snwfdhmp/errlog#Config) for the up to date definition


## Example

### Try yourself

| Name and link | Description |
| --- | --- |
| [Basic](examples/basic/basic.go) | standard usage, quick setup
| [Custom](examples/custom/custom.go) | guided configuration for fulfilling your needs |
| [Disabled](examples/disabled/disabled.go) | how to disable the logging & debugging (eg: for production use) |
| [Failing line far away](examples/failingLineFar/failingLineFar.go) | example of finding the func call that caused the error while it is lines away from the errlog.Debug call |
| [Pretty stack trace](examples/stackTrace/stackTrace.go) | pretty stack trace printing instead of debugging. |

### Just read

#### Basic example

We're going to use this sample program :

```golang
package main

import (
    "errors"
    "fmt"

    "github.com/snwfdhmp/errlog"
)

func init() {
    errlog.DefaultLogger.Disable(true)
}

func main() {
    fmt.Println("Example start")

    wrapingFunc()

    fmt.Println("Example end")
}

func wrapingFunc() {
    someBigFunction()
}

func someBigFunction() {
    someDumbFunction()

    someSmallFunction()

    someDumbFunction()

    if err := someNastyFunction(); errlog.Debug(err) {
        return
    }

    someSmallFunction()

    someDumbFunction()
}

func someSmallFunction() {
    _ = fmt.Sprintf("I do things !")
}

func someNastyFunction() error {
    return errors.New("I'm failing for some reason")
}

func someDumbFunction() bool {
    return false
}

```


#### Output

![Console Output examples/basic.go](https://i.imgur.com/tOkDgwP.png)


We are able to detect and point out which line is causing the error.

### Custom Configuration Example

Now let's see what we can do with a custom configuration.

```golang
debug := errlog.NewLogger(&errlog.Config{
    // PrintFunc is of type `func (format string, data ...interface{})`
    // so you can easily implement your own logger func.
    // In this example, logrus is used, but any other logger can be used.
    // Beware that you should add '\n' at the end of format string when printing.
    PrintFunc:          logrus.Printf,
    PrintSource:        true, //Print the failing source code
    LinesBefore:        2, //Print 2 lines before failing line
    LinesAfter:         1, //Print 1 line after failing line
    PrintError:         true, //Print the error
    PrintStack:         false, //Don't print the stack trace
    ExitOnDebugSuccess: true, //Exit if err
})
```

> As we don't yet update automatically this README immediately when we add new features, this definition may be outdated. (Last update: 2019/08/07)
> [See the struct definition in godoc.org](https://godoc.org/github.com/snwfdhmp/errlog#Config) for the up to date definition

#### Output

![Console Output examples/custom.go](https://i.imgur.com/vh2iEnS.png)


### Example

Errlog finds the exact line where the error is defined.

### Output

![Source Example: error earlier in the code](https://i.imgur.com/wPBrYqs.png)

## Documentation

Documentation can be found here : [![Documentation](https://godoc.org/github.com/snwfdhmp/errlog?status.svg)](http://godoc.org/github.com/snwfdhmp/errlog)

## Feedback

Feel free to open an issue for any feedback or suggestion.

I fix process issues quickly.

## Contributions

PR are accepted as soon as they follow Golang common standards.
For more information: https://golang.org/doc/effective_go.html

## License information

[![license](https://img.shields.io/github/license/snwfdhmp/errlog.svg?maxAge=60000)](https://github.com/snwfdhmp/errlog/LICENSE)

## Contributors

### Project contribution

- [snwfdhmp](https://github.com/snwfdhmp): Original author
- [chemidy](https://github.com/chemidy): Add badges

### Minor fixes

- [orisano](https://github.com/orisano)
- [programmingman](https://github.com/programmingman)
