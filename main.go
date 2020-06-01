package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

const (
	gpioT1 = "GPIO5"
	gpioT2 = "GPIO6"
	gpioT3 = "GPIO13"

	gpioN4 = "GPIO18"
	gpioN3 = "GPIO23"
	gpioN2 = "GPIO24"
	gpioN1 = "GPIO25"
)

func main() {
	fmt.Println("rpi-twitch-counter")

	tubeNo := atoiMust(os.Args[1])
	tubeDigit := atoiMust(os.Args[2])

	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	t1 := gpioreg.ByName(gpioT1)
	t2 := gpioreg.ByName(gpioT2)
	t3 := gpioreg.ByName(gpioT3)

	n4 := gpioreg.ByName(gpioN4)
	n3 := gpioreg.ByName(gpioN3)
	n2 := gpioreg.ByName(gpioN2)
	n1 := gpioreg.ByName(gpioN1)

	chk(t1.Out(gpio.Low))
	chk(t2.Out(gpio.Low))
	chk(t3.Out(gpio.Low))
	switch tubeNo {
	case 1:
		chk(t1.Out(gpio.High))
	case 2:
		chk(t2.Out(gpio.High))
	case 3:
		chk(t3.Out(gpio.High))
	}

	chk(n4.Out(gpio.Low)) // 8
	chk(n3.Out(gpio.Low)) // 4
	chk(n2.Out(gpio.Low)) // 2
	chk(n1.Out(gpio.Low)) // 1
	if tubeDigit&0b1000 != 0 {
		chk(n4.Out(gpio.High)) // 8
	}
	if tubeDigit&0b0100 != 0 {
		chk(n3.Out(gpio.High)) // 4
	}
	if tubeDigit&0b0010 != 0 {
		chk(n2.Out(gpio.High)) // 2
	}
	if tubeDigit&0b0001 != 0 {
		chk(n1.Out(gpio.High)) // 1
	}

}

func atoiMust(str string) int {
	v, err := strconv.Atoi(str)
	chk(err)
	return v
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
