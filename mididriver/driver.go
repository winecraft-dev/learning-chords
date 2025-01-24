package mididriver

import (
	"errors"
	"fmt"
	"log"

	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/drivers"

	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv" // autoregisters driver
)

const DEFAULT_IN_PORT = "Digital Piano"

var ErrStartListen = errors.New("problem starting listening")
var ErrNotListening = errors.New("not listening")

type MIDIDriver struct {
	in       drivers.In
	messages chan midi.Message
	stop     func()
}

func NewMIDIDriver() *MIDIDriver {
	driver := &MIDIDriver{
		messages: make(chan midi.Message),
		stop:     nil,
	}

	in, err := midi.FindInPort(DEFAULT_IN_PORT)
	if err != nil {
		log.Fatalf("Can't find %s\n", DEFAULT_IN_PORT)
		return nil
	}
	driver.in = in

	return driver
}

func (d *MIDIDriver) StartListening() error {
	stop, err := midi.ListenTo(d.in, func(msg midi.Message, ms int32) {
		var ch, k, v uint8
		switch {
		case msg.GetNoteStart(&ch, &k, &v):
			fmt.Printf("Note Start on %v %v @ %v\n", ch, k, v)
		case msg.GetNoteEnd(&ch, &k):
			fmt.Printf("Note end on %v: %v\n", ch, k)
		default:
		}
	}, midi.UseSysEx())

	if err != nil {
		return fmt.Errorf("%w: %w", ErrStartListen, err)
	}

	d.stop = stop
	return nil
}

func (d *MIDIDriver) StopListening() error {
	if d.stop == nil {
		return ErrNotListening
	}

	d.stop()
	return nil
}
