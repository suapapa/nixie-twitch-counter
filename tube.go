package main

import (
	"sync"
	"time"

	"periph.io/x/periph/conn/gpio"
)

type tube struct {
	no100st, no10st, no1st int
	sync.Mutex
}

func (t *tube) Start() {
	sleepDuration := time.Duration(sleepBetweenEachTube) * time.Millisecond
	for {
		t.Lock()
		tubeOne(1, t.no1st)
		time.Sleep(sleepDuration)
		tubeOne(2, t.no10st)
		time.Sleep(sleepDuration)
		tubeOne(3, t.no100st)
		t.Unlock()
		time.Sleep(sleepDuration)
	}
}

func (t *tube) Set(num int) {
	t.Lock()
	defer t.Unlock()
	t.no100st, t.no10st, t.no1st = getStDigits(num)
}

func tubeOne(idx, num int) {
	chk(t1.Out(gpio.Low))
	chk(t2.Out(gpio.Low))
	chk(t3.Out(gpio.Low))
	switch idx {
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
	if num&0b1000 != 0 {
		chk(n4.Out(gpio.High)) // 8
	}
	if num&0b0100 != 0 {
		chk(n3.Out(gpio.High)) // 4
	}
	if num&0b0010 != 0 {
		chk(n2.Out(gpio.High)) // 2
	}
	if num&0b0001 != 0 {
		chk(n1.Out(gpio.High)) // 1
	}
}

func getStDigits(num int) (int, int, int) {
	no1st := num % 10
	num = num / 10
	no10st := num % 10
	num = num / 10
	no100st := num % 10

	return no100st, no10st, no1st
}
