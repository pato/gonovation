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

	fmt.Println("Initalized")

	launchpad := GetLaunchPad()
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
