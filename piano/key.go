package piano

type Key struct {
	octave uint8
	key    uint8
	code   uint8
	color  KeyColor
	state  KeyState
}
