package piano

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Piano struct {
	keys map[uint8]Key

	whiteKeys []uint8
	blackKeys []uint8
}

func NewPiano() *Piano {
	keys := make(map[uint8]Key)
	whiteKeys := make([]uint8, WHITE_KEYS)
	blackKeys := make([]uint8, KEYS-WHITE_KEYS)

	for i := 0; i < KEYS; i++ {
		var code = uint8(i + 9)
		var color KeyColor
		switch code % BLACK_SLOTS_OCTAVE {
		case 0, 2, 3, 5, 7, 8, 10:
			color = KeyColorWhite
			whiteKeys = append(whiteKeys, code)
		case 1, 4, 6, 9, 11:
			color = KeyColorBlack
			blackKeys = append(blackKeys, code)
		}

		keys[code] = Key{
			octave: uint8(code / BLACK_SLOTS_OCTAVE),
			key:    uint8(code % BLACK_SLOTS_OCTAVE),
			code:   code,
			color:  color,
		}
	}

	return &Piano{
		keys:      keys,
		whiteKeys: whiteKeys,
		blackKeys: blackKeys,
	}
}

func (p *Piano) Draw(canvas *ebiten.Image) {
	white, black, ww, bw := generateKeyImages(canvas.Bounds())
	fmt.Println(ww, bw)

	for i := 0; i < WHITE_KEYS; i++ {
		key := p.keys[p.whiteKeys[i]]

		switch key.state {

		}

		opts := ebiten.DrawImageOptions{}
		opts.GeoM.Translate(ww*float64(i), 0)

		canvas.DrawImage(white, &opts)
	}

	for i, b := 0, 0; b < KEYS-WHITE_KEYS; i++ {
		switch i % BLACK_SLOTS_OCTAVE {
		case 0, 2, 3, 5, 7, 8, 10:
			continue
		}

		opts := ebiten.DrawImageOptions{}
		opts.GeoM.Translate(bw*float64(i)+bw*4.0/12.0, 0)

		canvas.DrawImage(black, &opts)

		b++
	}

}

func (p *Piano) Layout(r image.Rectangle) image.Rectangle {
	return image.Rect(r.Min.X, r.Min.Y, r.Dx(), int(float64(r.Dx())*HEIGHT_RATIO))
}
