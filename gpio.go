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

// ProcessingTime represents the time that the board needs to process the changes.
const ProcessingTime = 200 // ms

// Gpio represents the collection of pins of the raspberry pi.
type Gpio struct{}

// Out returns a specific pin (also initializes and configures to output the pin for the system).
func (r Gpio) Out(pin string) GpioOut {
	out := GpioOut{pin}

	filename := out.Filename()
	direction := filename + "/direction"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		writeFile("/sys/class/gpio/export", []byte(out.Pin))
		writeFile(direction, []byte("out"))
	}

	return out
}

// GpioOut represents a specific output pin of the raspberry pi.
type GpioOut struct {
	Pin string
}

// Filename returns the absolute path of the pin.
func (r GpioOut) Filename() string {
	return "/sys/class/gpio/gpio" + r.Pin
}

// Write changes to the given file.
func (r GpioOut) write(where, what string) {
	filename := r.Filename() + "/" + where
	writeFile(filename, []byte(what))
}

// On sets the value of the pin to 1.
func (r GpioOut) On() {
	r.write("value", "1")
}

// Off sets the value of the pin to 0.
func (r GpioOut) Off() {
	r.write("value", "0")
}

// Write changes to the given file and wait processing time.
func writeFile(filename string, data []byte) {
	ioutil.WriteFile(filename, data, 0666)
	time.Sleep(ProcessingTime * time.Millisecond)
}
