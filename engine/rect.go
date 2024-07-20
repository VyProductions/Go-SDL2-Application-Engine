package engine

import "github.com/veandco/go-sdl2/sdl"

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
	borderColor  sdl.Color
	mode         BorderMode
	borderWeight int

	// Widget
	isActive bool
	widgetID string

	// Mouse Events
}

func (r *Rect) SetActive(active bool) { r.isActive = active }
func (r *Rect) GetWidgetID() string   { return r.widgetID }

func (r *Rect) OnMouseDown(*Engine, sdl.Point)            {}
func (r *Rect) OnMouseUp(*Engine, sdl.Point)              {}
func (r *Rect) OnMouseEnter(*Engine, sdl.Point)           {}
func (r *Rect) OnMouseLeave(*Engine, sdl.Point)           {}
func (r *Rect) OnMouseDrag(*Engine, sdl.Point, sdl.Point) {}
