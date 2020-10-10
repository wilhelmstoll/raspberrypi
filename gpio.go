// Copyright 2020 Wilhelm Stoll. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Pure go implementation of raspberry pi functions.

package raspberrypi

import (
	"io/ioutil"
	"os"
	"time"
)

// Gpio represents the collection of pins of the raspberry pi.
type Gpio struct{}

// Pin returns a specific pin (also initialize the pin for the system).
func (r Gpio) Pin(name string) GpioPin {
	pin := GpioPin{name}
	filename := pin.Filename()
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		ioutil.WriteFile("/sys/class/gpio/export", []byte(pin.Name), 0666)
		waitProcessingTime()
	}

	return pin
}

// GpioPin represents a specific pin of the raspberry pi.
type GpioPin struct {
	Name string
}

// Filename returns the absolute path of the pin.
func (r GpioPin) Filename() string {
	return "/sys/class/gpio/gpio" + r.Name
}

// Write changes the transferred file with the transferred value.
func (r GpioPin) write(where, what string) {
	filename := r.Filename() + "/" + where
	ioutil.WriteFile(filename, []byte(what), 0666)
	waitProcessingTime()
}

// Output sets the pin mode to output.
func (r GpioPin) Output() {
	r.write("direction", "out")
}

// High sets the value to 1.
func (r GpioPin) High() {
	r.write("value", "1")
}

// Low sets the value to 0.
func (r GpioPin) Low() {
	r.write("value", "0")
}

// WaitProcessingTime waits a defined timespan.
func waitProcessingTime() {
	time.Sleep(200 * time.Millisecond)
}
