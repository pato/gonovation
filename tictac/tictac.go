package main

import (
	"fmt"
	"github.com/pato/gonovation/launchpad"
)

func boardOn(launchpad *gonovation.Launchpad, x, y, r, g int) {
	if !(x < 0 || y < 0 || x == 2 || x == 5 || x == 8 || y == 2 || y == 5 || y == 8) {
		launchpad.Led(x, y, r, g)
	}
}

func boardClick(x, y int64) bool {
	return !(x == 2 || x == 5 || x == 8 || y == 2 || y == 5 || y == 8)
}

func processClick(launchpad *gonovation.Launchpad, x, y int64, turn int) {
	if x == 8 {
		// right buttons
	} else if y == 8 {
		// top buttons
		resetBoard(launchpad)
	} else if boardClick(x, y) {
		var r, g int
		if turn%2 == 0 {
			// player 1
			r = 0
			g = 3
		} else {
			//player 2
			r = 3
			g = 3
		}
		boardOn(launchpad, int(x), int(y), r, g)
		boardOn(launchpad, int(x+1), int(y), r, g)
		boardOn(launchpad, int(x-1), int(y), r, g)
		boardOn(launchpad, int(x+1), int(y-1), r, g)
		boardOn(launchpad, int(x+1), int(y+1), r, g)
		boardOn(launchpad, int(x-1), int(y-1), r, g)
		boardOn(launchpad, int(x-1), int(y+1), r, g)
		boardOn(launchpad, int(x-1), int(y-1), r, g)
		boardOn(launchpad, int(x), int(y+1), r, g)
		boardOn(launchpad, int(x), int(y-1), r, g)
	}
}

func resetBoard(launchpad *gonovation.Launchpad) {
	launchpad.Reset()

	/* Set up board */
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if x == 2 || x == 5 || y == 2 || y == 5 {
				launchpad.Led(x, y, 3, 0)
			}
		}
	}
}

func main() {
	fmt.Println("Initalized")

	launchpad := gonovation.GetLaunchPad()

	resetBoard(launchpad)

	events := launchpad.Events()
	turn := 0
	for {
		event := <-events
		x, y, pressed := gonovation.EventInfo(event)

		if pressed {
			processClick(launchpad, x, y, turn)
			if boardClick(x, y) {
				turn++
			}
		}
	}

	defer launchpad.Close()
}
