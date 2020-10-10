# raspberrypi [![GoDoc](https://godoc.org/github.com/wilhelmstoll/raspberrypi?status.svg)](https://godoc.org/github.com/wilhelmstoll/raspberrypi) [![Build Status](https://travis-ci.org/wilhelmstoll/raspberrypi.svg?branch=master)](https://travis-ci.org/wilhelmstoll/raspberrypi)

Pure go implementation of raspberry pi functions.

**Currently only the base GPIO functions are implemented!**

## Installation

```
go get -u github.com/wilhelmstoll/raspberrypi
```

## Example of usage

```go
package main

import (
	"github.com/wilhelmstoll/raspberrypi"
)

func main() {
	gpio := raspberrypi.Gpio{}

	// Initializes pin 18 (initial mode input).
	gpioPin18 := gpio.Pin("18")

	// Sets the pin to output mode (initial value low).
	gpioPin18.Output()

	// Sets the pin to high.
	gpioPin18.High()
}
```

## Reference

https://godoc.org/github.com/wilhelmstoll/raspberrypi
