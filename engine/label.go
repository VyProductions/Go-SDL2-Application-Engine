package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Label struct {
	// Label
	background Rect
	text       string

	// Widget
	isActive bool
	widgetID string

	// Mouse Events
	/* Pass Events Down? */
	/* - Z-index system for engine */
}

func (l *Label) SetActive(active bool) { l.isActive = active }
func (l *Label) GetWidgetID() string   { return l.widgetID }

func (l *Label) OnMouseDown(e Engine, pos sdl.Point)                  {}
func (l *Label) OnMouseUp(e Engine, pos sdl.Point)                    {}
func (l *Label) OnMouseEnter(e Engine, pos sdl.Point)                 {}
func (l *Label) OnMouseLeave(e Engine, pos sdl.Point)                 {}
func (l *Label) OnMouseDrag(e Engine, pos1 sdl.Point, pos2 sdl.Point) {}
