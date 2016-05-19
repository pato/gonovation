(Go)novation - Go Driver for Novation Launchpad
============

(Go)novation has three packages:
* launchpad
* tictactoe
* demo

### Launchpad

This is the go driver for communicating with the launchpad

To use first import the library

`import "github.com/pato/gonovation/launchpad`

Then open a connection

`launchpad := gonovation.GetLaunchPad()`

And defer the closing with 

`defer launchpad.Close()`

You can then reset it with `launchpad.Reset()`

Or set colors using `launchpad.Led(x,y,r,g)` where `x,y` are the coordinates and `r,g` are the red and green colors

The launchpad has three available colors red (`r=3,g=0`), yellow (`r=3,g=3`), and green (`r=0,g=3`)

Or you can poll it for events by getting the events channel with `launchpad.Events()`

For reference see `demo/demo.go`

### TicTacToe

A silly demo where you can play a game of tic tac toe on the board!


### Demo

Simple demo for listening to key presses and setting colors programmatically


## Installing

Just run

`go get github.com/pato/gonovation/launchpad`

`go get github.com/pato/gonovation/tictactoe`

`go get github.com/pato/gonovation/demo`

They will then be installed to `$GOPATH/bin`
