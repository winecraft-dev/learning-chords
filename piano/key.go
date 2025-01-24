package piano

type KeyState int

const KeyOnState KeyState = 0
const KeyOffState KeyState = 1
const KeyCorrectState KeyState = 2
const KeyWrongState KeyState = 3

type Key struct {
	state KeyState
}
