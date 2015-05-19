package main

import (
	"fmt"
	"github.com/rakyll/portmidi"
	"log"
	"strings"
)

type Launchpad struct {
	midiIn       portmidi.DeviceId
	midiOut      portmidi.DeviceId
	outputStream *portmidi.Stream
	inputStream  *portmidi.Stream
}

func main() {
	Initialize()
	defer Terminate()

	launchpad := GetLaunchPad()
	launchpad.Reset()
}

func Initialize() {
	portmidi.Initialize()
}

func Terminate() {
	portmidi.Terminate()
}

func GetLaunchPad() *Launchpad {
	midiIn, midiOut := findLaunchPadMidis()
	inputStream, err := portmidi.NewInputStream(midiIn, 1024)
	handleError(err)
	outputStream, err := portmidi.NewOutputStream(midiOut, 1024, 0)
	handleError(err)
	return &Launchpad{midiIn: midiIn, midiOut: midiOut, inputStream: inputStream, outputStream: outputStream}
}

func (launchpad *Launchpad) Reset() {
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func findLaunchPadMidis() (midiIn, midiOut portmidi.DeviceId) {
	nDevices := portmidi.CountDevices()

	fmt.Printf("Devices: %d\n", nDevices)
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
