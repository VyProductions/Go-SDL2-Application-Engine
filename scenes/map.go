package scenes

import (
	"fmt"
	"main/engine"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	NUM_ROWS int32 = 32
	NUM_COLS int32 = 56
)

type MapDisplay struct {
	screen    *engine.Screen
	pathSheet *sdl.Texture
}

func generate(m *MapDisplay) error {
	blocked := map[int32]bool{}
	marked := map[int32]bool{}

	for r := int32(0); r < NUM_ROWS; r++ {
		for c := int32(0); c < NUM_COLS; c++ {
			imgID := fmt.Sprintf("Image_%d_%d", r, c)

			if m.screen.HasWidgetID(imgID) {
				if err := m.screen.DeleteWidget(imgID); err != nil {
					return err
				}
			}

			randChoice := int32(rand.Intn(9))
			blocked[r*NUM_COLS+c] = randChoice < 4
			marked[r*NUM_COLS+c] = false
		}
	}

	yOffs := [8]int32{-1, -1, -1, 0, 0, 1, 1, 1}
	xOffs := [8]int32{-1, 0, 1, -1, 1, -1, 0, 1}

	for r := int32(0); r < NUM_ROWS; r++ {
		for c := int32(0); c < NUM_COLS; c++ {
			currBlocked := blocked[r*NUM_COLS+c]
			oppositeCount := 0

			for i := 0; i < 8; i++ {
				x, y := xOffs[i], yOffs[i]

				if r+y < 0 || r+y >= NUM_ROWS || c+x < 0 || c+x >= NUM_COLS {
					oppositeCount++
					continue
				}

				otherBlocked := blocked[(r+y)*NUM_COLS+c+x]

				if (currBlocked || otherBlocked) && (!currBlocked || !otherBlocked) {
					oppositeCount++
				}
			}

			if oppositeCount >= 5 {
				marked[r*NUM_COLS+c] = true
			}
		}
	}

	for idx, mark := range marked {
		r, c := idx/NUM_COLS, idx%NUM_COLS
		imgID := fmt.Sprintf("Image_%d_%d", r, c)

		blocked[idx] = (blocked[idx] || mark) && (!blocked[idx] || !mark)

		var offs int32

		if blocked[idx] {
			offs = 1
		} else {
			offs = 0
		}

		img := new(engine.Image)
		if err := img.Init(
			imgID, m.pathSheet,
			sdl.Rect{X: offs * 32, Y: 0, W: 32, H: 32},
			sdl.Rect{X: 64 + 32*c, Y: 16 + 32*r, W: 32, H: 32},
			engine.BORDER_IN, 0, engine.NOT_DRAGGABLE,
			[]sdl.Color{
				{R: 0x7F, G: 0x7F, B: 0x7F, A: 0xFF},
				{R: 0xBF, G: 0xBF, B: 0xBF, A: 0xFF},
			}...,
		); err != nil {
			return err
		}

		img.SetActive(true)

		if err := m.screen.InsertWidget(img); err != nil {
			return err
		}
	}

	return nil
}

func (m *MapDisplay) Init(e *engine.Engine) error {
	m.screen = new(engine.Screen)

	pathSurface, err := sdl.LoadBMP("images/PathSpriteSheet.bmp")

	if err != nil {
		return err
	}

	if m.pathSheet, err = e.Renderer.CreateTextureFromSurface(pathSurface); err != nil {
		return err
	}

	pathSurface.Free()

	if err := m.screen.Init(e, "MapScene"); err != nil {
		return err
	}

	if err := generate(m); err != nil {
		return err
	}

	regenerateButton := new(engine.Button)
	regenerateButton.Init(
		"RegenerateButton", sdl.Rect{X: 832, Y: 1048, W: 256, H: 32}, "Regenerate", engine.NOT_DRAGGABLE,
		func() error {
			return generate(m)
		},
		engine.BORDER_IN, 4,
		[]sdl.Color{
			{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
			{R: 0x4F, G: 0x4F, B: 0x4F, A: 0xFF},
			{R: 0x7F, G: 0x7F, B: 0x7F, A: 0xFF},
			{R: 0x5F, G: 0x5F, B: 0x5F, A: 0xFF},
			{R: 0x3F, G: 0x3F, B: 0x3F, A: 0xFF},
		}...,
	)
	regenerateButton.SetActive(true)

	return m.screen.InsertWidget(regenerateButton)
}

func (m *MapDisplay) Free() error {
	return m.pathSheet.Destroy()
}
