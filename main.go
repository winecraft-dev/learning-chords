package main

import (
	"fmt"
	"time"

	"gitlab.com/gomidi/midi/v2"

	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv" // autoregisters driver
)

func main() {
	defer midi.CloseDriver()

	fmt.Printf("%v\n", midi.GetInPorts())

	in, err := midi.FindInPort("Digital Piano")
	if err != nil {
		fmt.Println("can't find VMPK")
		return
	}

	stop, err := midi.ListenTo(in, func(msg midi.Message, timestampms int32) {
		var bt []byte
		var ch, key, vel uint8
		switch {
		case msg.GetSysEx(&bt):
			fmt.Printf("got sysex: % X\n", bt)
		case msg.GetNoteStart(&ch, &key, &vel):
			fmt.Printf("starting note %v on channel %v with velocity %v\n", key, ch, vel)
		case msg.GetNoteEnd(&ch, &key):
			fmt.Printf("ending note %v", key)
		}
	}, midi.UseSysEx())

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	time.Sleep(time.Second * 5)

	stop()
}
