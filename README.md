# raspberrypi [![GoDoc](https://godoc.org/github.com/wilhelmstoll/raspberrypi?status.svg)](https://godoc.org/github.com/wilhelmstoll/raspberrypi) [![Build Status](https://travis-ci.org/wilhelmstoll/raspberrypi.svg?branch=master)](https://travis-ci.org/wilhelmstoll/raspberrypi)

Pure go implementation of raspberry pi functions.

**Currently only the base GPIO output functions are implemented!**

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

	// Initializes pin 18 as output pin.
	gpioPin18 := gpio.Out("18")

	// Sets the pin to high.
	gpioPin18.On()
}
```

## Reference

https://godoc.org/github.com/wilhelmstoll/raspberrypi
