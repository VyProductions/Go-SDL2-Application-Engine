package engine

import "github.com/veandco/go-sdl2/sdl"

type TextBox struct {
	// Text Box
	text string

	// Widget
	isActive bool
	widgetID string

	// Mouse Events
}

func (t *TextBox) SetActive(active bool) { t.isActive = active }
func (t *TextBox) GetWidgetID() string   { return t.widgetID }

func (t *TextBox) OnMouseDown(*Engine, sdl.Point)            {}
func (t *TextBox) OnMouseUp(*Engine, sdl.Point)              {}
func (t *TextBox) OnMouseEnter(*Engine, sdl.Point)           {}
func (t *TextBox) OnMouseLeave(*Engine, sdl.Point)           {}
func (t *TextBox) OnMouseDrag(*Engine, sdl.Point, sdl.Point) {}
