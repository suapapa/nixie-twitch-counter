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

var (
	t1, t2, t3     gpio.PinIO
	n4, n3, n2, n1 gpio.PinOut
)

func main() {
	fmt.Println("rpi-twitch-counter")

	tubeNo := atoiMust(os.Args[1])
	tubeDigit := atoiMust(os.Args[2])

	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	t1 = gpioreg.ByName(gpioT1)
	t2 = gpioreg.ByName(gpioT2)
	t3 = gpioreg.ByName(gpioT3)

	n4 = gpioreg.ByName(gpioN4)
	n3 = gpioreg.ByName(gpioN3)
	n2 = gpioreg.ByName(gpioN2)
	n1 = gpioreg.ByName(gpioN1)

	tube(tubeNo, tubeDigit)
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
