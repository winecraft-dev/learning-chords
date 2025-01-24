package piano

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const KEYS = 88
const WHITE_KEYS = 56
const WHITE_KEYS_OCTAVE = 7
const KEY_PADDING = 2
const BLACK_SLOTS_OCTAVE = 12
const BLACK_SLOTS = WHITE_KEYS / WHITE_KEYS_OCTAVE * BLACK_SLOTS_OCTAVE
const HEIGHT_RATIO = 5 / float64(WHITE_KEYS)

type Piano struct {
}

func NewPiano() *Piano {
	return &Piano{}
}

func (p *Piano) Draw(canvas *ebiten.Image) {
	width := canvas.Bounds().Dx()
	height := canvas.Bounds().Dy()

	wkw := float64(width) / float64(WHITE_KEYS)
	wkh := float64(height)

	for i := 0; i < WHITE_KEYS; i++ {
		key := ebiten.NewImage(int(wkw-KEY_PADDING), int(wkh))
		key.Fill(color.White)

		opts := ebiten.DrawImageOptions{}
		opts.GeoM.Translate(wkw*float64(i), 0)

		canvas.DrawImage(key, &opts)
	}

	bkw := float64(width) / float64(BLACK_SLOTS)
	bkh := float64(height) / 5 * 3
	for i := 0; i < BLACK_SLOTS; i++ {
		switch i % BLACK_SLOTS_OCTAVE {
		//case 1, 3, 6, 8, 10
		//	continue
		case 0, 2, 4, 5, 7, 9, 11:
			continue
		}
		key := ebiten.NewImage(int(bkw-KEY_PADDING), int(bkh))
		key.Fill(color.Black)

		opts := ebiten.DrawImageOptions{}
		opts.GeoM.Translate(bkw*float64(i), 0)

		canvas.DrawImage(key, &opts)
	}
}

func (p *Piano) Layout(r image.Rectangle) image.Rectangle {
	return image.Rect(r.Min.X, r.Min.Y, r.Dx(), int(float64(r.Dx())*HEIGHT_RATIO))
}
