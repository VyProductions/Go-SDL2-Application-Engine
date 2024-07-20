package engine

import "github.com/veandco/go-sdl2/sdl"

type ImageButton struct {
	// Image Button
	background Rect

	// Widget
	isActive bool
	widgetID string

	// Mouse Events
	isPressed bool
	mousePos  sdl.Point
}

func (i *ImageButton) SetActive(active bool) { i.isActive = active }
func (i *ImageButton) GetWidgetID() string   { return i.widgetID }

func (i *ImageButton) OnMouseDown(*Engine, sdl.Point)            {}
func (i *ImageButton) OnMouseUp(*Engine, sdl.Point)              {}
func (i *ImageButton) OnMouseEnter(*Engine, sdl.Point)           {}
func (i *ImageButton) OnMouseLeave(*Engine, sdl.Point)           {}
func (i *ImageButton) OnMouseDrag(*Engine, sdl.Point, sdl.Point) {}
