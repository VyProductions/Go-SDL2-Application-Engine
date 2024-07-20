package engine

import "github.com/veandco/go-sdl2/sdl"

type Button struct {
	// Button
	background Rect

	// Widget
	isActive bool
	widgetID string

	// Mouse Events
	isPressed bool
	mousePos  sdl.Point
}

func (b *Button) SetActive(active bool) { b.isActive = active }
func (b *Button) GetWidgetID() string   { return b.widgetID }

func (b *Button) OnMouseDown(e *Engine, pos sdl.Point)                  {}
func (b *Button) OnMouseUp(e *Engine, pos sdl.Point)                    {}
func (b *Button) OnMouseEnter(e *Engine, pos sdl.Point)                 {}
func (b *Button) OnMouseLeave(e *Engine, pos sdl.Point)                 {}
func (b *Button) OnMouseDrag(e *Engine, pos1 sdl.Point, pos2 sdl.Point) {}
