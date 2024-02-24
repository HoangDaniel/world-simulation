package main

import (
	"fmt"
	"time"

	"hoang.sk/world/vdt"
)

func main() {
	virtualTimeTicker := time.NewTicker(250 * time.Millisecond)
	realSecondTicker := time.NewTicker(1 * time.Second)

	defer virtualTimeTicker.Stop()
	defer realSecondTicker.Stop()

	vt := vdt.NewVirtualDateTime()
	fmt.Println("New Age started with date time: %v", vt.DateTime())

	for {
		select {
		case <-virtualTimeTicker.C:
			vt.ForwardMonth()
		case <-realSecondTicker.C:
			fmt.Println(vt.DateTime())
		}
	}
}
