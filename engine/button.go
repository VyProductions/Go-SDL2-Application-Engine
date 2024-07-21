package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Button struct {
	// Button
	background *Rect
	text       string
	draggable  bool

	// Widget
	isActive bool
	widgetID string

	// Mouse Events
	entered   bool
	pressed   bool
	onClicked func() error
}

// Button
func (b *Button) Init(
	id string, rect sdl.Rect, text string, draggable bool,
	onClicked func() error,
	mode BorderMode, weight int, colors ...sdl.Color,
) error {
	b.widgetID = id
	b.background = new(Rect)
	b.text = text
	b.draggable = draggable

	b.onClicked = onClicked

	if len(colors) > 1 {
		b.background.Init(id, rect, mode, weight, draggable, colors[1:]...)
	} else {
		b.background.Init(id, rect, mode, weight, draggable, []sdl.Color{}...)
	}

	return nil
}

// Widget
func (b *Button) SetActive(active bool) {
	b.isActive = active
	b.background.SetActive(active)
}

func (b *Button) GetWidgetID() string {
	return b.widgetID
}

func (b *Button) Render(e *Engine) error {
	if b.isActive {
		if err := b.background.Render(e); err != nil {
			return err
		}
	}

	return nil
}

func (b *Button) Free() error {
	return nil
}

// Mouse Events
func (b *Button) OnMouseDown(e *Engine, pos sdl.Point) error {
	if b.entered {
		b.pressed = true
		return b.background.OnMouseDown(e, pos)
	}

	return nil
}

func (b *Button) OnMouseUp(e *Engine, pos sdl.Point) error {
	if err := b.background.OnMouseUp(e, pos); err != nil {
		return err
	}

	if b.pressed && pos.InRect(&b.background.rect) {
		return b.onClicked()
	}

	return nil
}

func (b *Button) OnMouseEnter(e *Engine, pos sdl.Point) error {
	if !b.entered && pos.InRect(&b.background.rect) {
		b.entered = true
		b.pressed = false
		return b.background.OnMouseEnter(e, pos)
	}

	return nil
}

func (b *Button) OnMouseLeave(e *Engine, pos sdl.Point) error {
	if b.entered && !pos.InRect(&b.background.rect) {
		b.entered = false
		b.pressed = false
		return b.background.OnMouseLeave(e, pos)
	}

	return nil
}

func (b *Button) OnMouseDrag(e *Engine, pos1 sdl.Point, pos2 sdl.Point) error {
	if b.draggable {
		return b.background.OnMouseDrag(e, pos1, pos2)
	}

	return nil
}
