package main

import "periph.io/x/periph/conn/gpio"

func tube(idx, num int) {
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
