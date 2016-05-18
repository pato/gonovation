package main

import (
	"fmt"
	"github.com/pato/gonovation/launchpad"
	"time"
)

var board [3][3]int

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
	} else if boardClick(x, y) && board[x/3][y/3] == 0 {
		var r, g int
		if turn%2 == 0 {
			// player 1
			r = 0
			g = 3
			board[x/3][y/3] = 1
		} else {
			//player 2
			r = 3
			g = 3
			board[x/3][y/3] = 2
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

	/* Set up launchpad board */
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if x == 2 || x == 5 || y == 2 || y == 5 {
				launchpad.Led(x, y, 3, 0)
			}
		}
	}

	/* Set up logical board */
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			board[x][y] = 0
		}
	}
}

func checkHorizontal(line int) bool {
	return board[0][line] == board[1][line] && board[1][line] == board[2][line] && board[0][line] != 0
}

func checkVertical(line int) bool {
	return board[line][0] == board[line][1] && board[line][1] == board[line][2] && board[line][0] != 0
}

func getWinner() (bool, int) {
	for i := 0; i < 3; i++ {
		if checkHorizontal(i) {
			return true, board[0][i]
		}
		if checkVertical(i) {
			return true, board[i][0]
		}
	}
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != 0 {
		return true, board[0][0]
	} else if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[0][2] != 0 {
		return true, board[0][2]
	}
	return false, -1
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
				w, win := getWinner()
				if w {
					snake(launchpad, win)
					resetBoard(launchpad)
				}
			}
		}
	}

	defer launchpad.Close()
}

func snake(launchpad *gonovation.Launchpad, winner int) {
	var r, g int
	if winner == 1 {
		r = 0
		g = 3
	} else {
		r = 3
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
