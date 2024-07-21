package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type BorderMode int

const (
	NO_BORDER = iota
	BORDER_IN
	BORDER_MID
	BORDER_OUT
)

type Rect struct {
	// Rect
	rect         sdl.Rect
	fillColor    sdl.Color
	hoverColor   sdl.Color
	pressedColor sdl.Color
	borderColor  sdl.Color
	mode         BorderMode
	borderWeight int
	draggable    bool

	// Widget
	isActive bool
	widgetID string

	// Mouse Events
	entered   bool
	pressed   bool
	currColor sdl.Color
}

// Rect
func (r *Rect) Init(
	id string, rect sdl.Rect, mode BorderMode, weight int, draggable bool, colors ...sdl.Color,
) error {
	r.widgetID = id
	r.rect = rect
	r.mode = mode
	r.borderWeight = weight

	if len(colors) == 0 {
		r.fillColor = sdl.Color{R: 0x7F, G: 0x7F, B: 0x7F, A: 0xFF}
		r.borderColor = r.fillColor
		r.hoverColor = r.fillColor
		r.pressedColor = r.fillColor
		r.currColor = r.fillColor
	}

	for idx, color := range colors {
		switch idx {
		case 0:
			r.fillColor = color
			r.borderColor = r.fillColor
			r.hoverColor = r.fillColor
			r.pressedColor = r.fillColor
			r.currColor = r.fillColor
		case 1:
			r.borderColor = color
		case 2:
			r.hoverColor = color
		case 3:
			r.pressedColor = color
		}
	}

	return nil
}

// Widget
func (r *Rect) SetActive(active bool) {
	r.isActive = active
}

func (r *Rect) GetWidgetID() string {
	return r.widgetID
}

func (r *Rect) Render(e *Engine) error {
	if r.isActive {
		oldR, oldG, oldB, oldA, err := e.Renderer.GetDrawColor()

		if err != nil {
			return err
		}

		// Draw the fill
		if err = e.Renderer.SetDrawColor(
			r.currColor.R, r.currColor.G, r.currColor.B, r.currColor.A,
		); err != nil {
			return err
		}

		if err = e.Renderer.FillRect(&r.rect); err != nil {
			return err
		}

		// Draw the border
		if err = e.Renderer.SetDrawColor(
			r.borderColor.R, r.borderColor.G, r.borderColor.B, r.borderColor.A,
		); err != nil {
			return err
		}

		var offs int32

		switch r.mode {
		case BORDER_IN:
			offs = 0
		case BORDER_OUT:
			offs = int32(r.borderWeight)
		case BORDER_MID:
			offs = int32(r.borderWeight / 2)
		}

		for i := int32(0); i < int32(r.borderWeight); i++ {
			borderRect := sdl.Rect{
				X: r.rect.X - offs + i, Y: r.rect.Y - offs + i,
				W: r.rect.W + 2*(offs-i), H: r.rect.H + 2*(offs-i),
			}

			if err = e.Renderer.DrawRect(&borderRect); err != nil {
				return err
			}
		}

		if err = e.Renderer.SetDrawColor(oldR, oldG, oldB, oldA); err != nil {
			return err
		}
	}

	return nil
}

func (r *Rect) Free() error {
	return nil
}

// Mouse Events
func (r *Rect) OnMouseDown(e *Engine, pos sdl.Point) error {
	if r.entered {
		r.pressed = true
		r.currColor = r.pressedColor
	}

	return nil
}

func (r *Rect) OnMouseUp(e *Engine, pos sdl.Point) error {
	r.pressed = false

	if r.entered {
		r.currColor = r.hoverColor
	} else {
		r.currColor = r.fillColor
	}

	return nil
}

func (r *Rect) OnMouseEnter(e *Engine, pos sdl.Point) error {
	if !r.entered && pos.InRect(&r.rect) {
		r.entered = true
		r.pressed = false
		r.currColor = r.hoverColor
	}

	return nil
}

func (r *Rect) OnMouseLeave(e *Engine, pos sdl.Point) error {
	if r.entered && !pos.InRect(&r.rect) {
		r.entered = false
		r.pressed = false
		r.currColor = r.fillColor
	}

	return nil
}

func (r *Rect) OnMouseDrag(e *Engine, pos1 sdl.Point, pos2 sdl.Point) error {
	if r.draggable {
		xOffs := pos2.X - pos1.X
		yOffs := pos2.Y - pos1.Y

		r.rect.X += xOffs
		r.rect.Y += yOffs
	}

	return nil
}
