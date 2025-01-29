package piano

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type KeyState int

const KeyStateOff KeyState = 0
const KeyStateCorrect KeyState = 1
const KeyStateWrong KeyState = 2

type KeyColor int

const KeyColorWhite KeyColor = 0
const KeyColorBlack KeyColor = 1

const KEYS = 88
const WHITE_KEYS = 52
const WHITE_KEYS_OCTAVE = 7
const KEY_PADDING = 2
const BLACK_SLOTS_OCTAVE = 12
const HEIGHT_RATIO = 5 / float64(WHITE_KEYS)

func generateKeyImages(ks image.Rectangle) (w *ebiten.Image, b *ebiten.Image, ww float64, bw float64) {
	width := ks.Dx() - 10
	height := ks.Dy()

	ww = float64(width) / float64(WHITE_KEYS)
	wkh := float64(height)
	w = ebiten.NewImage(int(ww-KEY_PADDING), int(wkh))
	w.Fill(color.White)

	bw = float64(width) / (float64(KEYS) + 9.0/12.0)
	bkh := float64(height) / 5 * 3
	b = ebiten.NewImage(int(bw), int(bkh))
	b.Fill(color.Black)

	return
}
