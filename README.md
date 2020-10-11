# raspberrypi [![GoDoc](https://godoc.org/github.com/wilhelmstoll/raspberrypi?status.svg)](https://godoc.org/github.com/wilhelmstoll/raspberrypi) [![Build Status](https://travis-ci.org/wilhelmstoll/raspberrypi.svg?branch=master)](https://travis-ci.org/wilhelmstoll/raspberrypi)

Pure go implementation of raspberry pi functions.

**Currently only the base GPIO output functions are implemented!**

**Processing time for gpio operation is set to 200ms for this module!**

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

	// Initializes pin 23 as output pin.
	gpioPin23 := gpio.Out("23")

	// Sets the pin to high.
	gpioPin23.On()
}
```

## Reference

https://godoc.org/github.com/wilhelmstoll/raspberrypi
