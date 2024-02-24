package main

import (
	"fmt"
	"time"

	"hoang.sk/world/vdt"
)

func main() {

	ticker := time.NewTicker(1 * time.Nanosecond)
	vt := vdt.NewVirtualDateTime()

	fmt.Print("New Age started with date time: %v", vt.DateTime())
	defer ticker.Stop()

	for range ticker.C {
		vt.IncreaseSecond()
		fmt.Println(vt.DateTime())
	}
}
