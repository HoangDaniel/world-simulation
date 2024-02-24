package main

import (
	"fmt"
	"time"

	"hoang.sk/world/vdt"
)

func main() {
	virtualTimeTicker := time.NewTicker(1 * time.Second)

	defer virtualTimeTicker.Stop()

	vt := vdt.NewVirtualDateTime()
	fmt.Println("New Age started with date time: %v", vt.DateTime())

	for range virtualTimeTicker.C {
		fmt.Println(vt.DateTime())
		vt.ForwardDay()
	}
}
