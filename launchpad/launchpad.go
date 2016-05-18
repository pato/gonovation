package gonovation

import (
	"github.com/rakyll/portmidi"
	"log"
	"strings"
)

type Launchpad struct {
	midiIn       portmidi.DeviceID
	midiOut      portmidi.DeviceID
	outputStream *portmidi.Stream
	inputStream  *portmidi.Stream
}

type Event struct {
	Timestamp int64
	Status    int64
	Note      int64
	Velocity  int64
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

func EventInfo(event portmidi.Event) (x, y int64, pressed bool) {
	note := event.Data1
	if event.Status == 176 {
		y = 8
		x = note - 104
	} else {
		x = note % 16
		y = 7 - (note / 16)
	}
	pressed = event.Data2 == 127
	return
}

func (launchpad *Launchpad) Close() {
	launchpad.inputStream.Close()
	launchpad.outputStream.Close()
	portmidi.Terminate()
}

func (launchpad *Launchpad) Reset() {
	launchpad.outputStream.WriteShort(0xb0, 0, 0)
}

func (launchpad *Launchpad) Events() <-chan portmidi.Event {
	return launchpad.inputStream.Listen()
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

func findLaunchPadMidis() (midiIn, midiOut portmidi.DeviceID) {
	nDevices := portmidi.CountDevices()

	for i := 0; i < nDevices; i++ {
		info := portmidi.Info(portmidi.DeviceID(i))
		if strings.HasPrefix(info.Name, "Launchpad") {
			if info.IsInputAvailable {
				midiIn = portmidi.DeviceID(i)
			} else if info.IsOutputAvailable {
				midiOut = portmidi.DeviceID(i)
			}
		}
	}
	return
}
