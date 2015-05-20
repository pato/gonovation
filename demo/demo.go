package main

import (
	"fmt"
	"github.com/pato/gonovation/launchpad"
	"time"
)

func snake(launchpad *gonovation.Launchpad) {
	var r, g int
	for on := 0; ; on++ {
		if on%2 == 0 {
			r = 3
			g = 0
		} else {
			r = 0
			g = 3
		}
		for x := 0; x <= 8; x++ {
			for y := 0; y <= 8; y++ {
				if x%2 == 0 {
					launchpad.Led(x, y, r, g)
				} else {
					launchpad.Led(x, 8-y, r, g)
				}
				time.Sleep(30 * time.Millisecond)
			}
		}
	}
}

func main() {
	fmt.Println("Initalized")

	launchpad := gonovation.GetLaunchPad()
	launchpad.Reset()

	//snake(launchpad)
	events := launchpad.Events()
	for {
		event := <-events
		x, y, pressed := gonovation.EventInfo(event)
		//launchpad.outputStream.WriteShort(0x90, event.Data1, event.Data2)
		if pressed {
			launchpad.Led(int(x), int(y), 3, 0)
		} else {
			launchpad.Led(int(x), int(y), 0, 3)
		}
		fmt.Printf("(%d,%d): %t\n", x, y, pressed)
	}

	defer launchpad.Close()
}
