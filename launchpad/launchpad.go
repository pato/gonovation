package main

import (
	"fmt"
	"github.com/rakyll/portmidi"
	"log"
	"strings"
	"time"
)

type Launchpad struct {
	midiIn       portmidi.DeviceId
	midiOut      portmidi.DeviceId
	outputStream *portmidi.Stream
	inputStream  *portmidi.Stream
}

func snake(launchpad *Launchpad) {
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

	launchpad := GetLaunchPad()
	launchpad.Reset()

	snake(launchpad)

	defer launchpad.Close()
}

func GetLaunchPad() *Launchpad {
	midiIn, midiOut := findLaunchPadMidis()
	inputStream, err := portmidi.NewInputStream(midiIn, 1024)
	handleError(err)
	outputStream, err := portmidi.NewOutputStream(midiOut, 1024, 0)
	handleError(err)

	portmidi.Initialize()
	return &Launchpad{midiIn: midiIn, midiOut: midiOut, inputStream: inputStream, outputStream: outputStream}
}

func (launchpad *Launchpad) Close() {
	launchpad.inputStream.Close()
	launchpad.outputStream.Close()
	portmidi.Terminate()
	fmt.Println("Closed")
}

func (launchpad *Launchpad) Reset() {
	launchpad.outputStream.WriteShort(0xb0, 0, 0)
}

func (launchpad *Launchpad) Led(x, y, r, g int) {
	vel := 16*g + r + 8 + 4
	if y == 8 && x != 8 {
		note := 104 + x
		launchpad.outputStream.WriteShort(0xb0, int64(note), int64(vel))
		return
	}
	note := x + 16*(7-y)
	launchpad.outputStream.WriteShort(0x90, int64(note), int64(vel))
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func findLaunchPadMidis() (midiIn, midiOut portmidi.DeviceId) {
	nDevices := portmidi.CountDevices()

	for i := 0; i < nDevices; i++ {
		info := portmidi.GetDeviceInfo(portmidi.DeviceId(i))
		if strings.HasPrefix(info.Name, "Launchpad") {
			if info.IsInputAvailable {
				midiIn = portmidi.DeviceId(i)
			} else if info.IsOutputAvailable {
				midiOut = portmidi.DeviceId(i)
			}
		}
	}
	return
}
