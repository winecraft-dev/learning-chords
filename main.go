package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"winecraft.dev/learning-chords/mididriver"
	"winecraft.dev/learning-chords/piano"
)

func main() {
	//driver := mididriver.NewMIDIDriver()
	piano := piano.NewPiano()

	learner := Learner{
		driver: nil,
		piano:  piano,
	}

	ebiten.SetWindowSize(1200, 300)
	ebiten.SetWindowTitle("Learning Chords")
	if err := ebiten.RunGame(&learner); err != nil {
		log.Fatalln(err)
	}

}

type Learner struct {
	driver *mididriver.MIDIDriver
	piano  *piano.Piano
}

func (l *Learner) Update() error {
	return nil
}

func (l *Learner) Draw(screen *ebiten.Image) {
	piano := l.piano

	image := screen.SubImage(piano.Layout(screen.Bounds())).(*ebiten.Image)
	piano.Draw(image)
}

func (l *Learner) Layout(ow, oh int) (int, int) {
	return ow, oh
}
