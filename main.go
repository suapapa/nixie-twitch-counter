package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

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

	sleepBetweenEachTube int
	sleepBetweenEachNum  int
	tubeNum, tubeDigit   int
)

func init() {
	flag.IntVar(&sleepBetweenEachTube, "b", 10, "sleep millisecons between tubes")
	flag.IntVar(&sleepBetweenEachNum, "bn", 30*1000, "sleep millisecons between num")
	flag.IntVar(&tubeNum, "n", 0, "tube num")
	flag.IntVar(&tubeDigit, "d", 0, "tube digit")

	rand.Seed(time.Now().UnixNano())
}

func main() {
	flag.Parse()
	fmt.Println("rpi-twitch-counter")

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

	if tubeNum != 0 {
		tubeOne(tubeNum, tubeDigit)
		return
	}

	tubeOne(1, 10)
	tubeOne(2, 10)
	tubeOne(3, 10)
	time.Sleep(500 * time.Millisecond)

	t := &tube{}
	go t.Start()
	sleepDuration := time.Duration(sleepBetweenEachNum) * time.Millisecond
	for {
		n := rand.Intn(1000)
		log.Printf("set tube to %d", n)
		t.Set(n)
		time.Sleep(sleepDuration)
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
