package engine

type Scene interface {
	// Scene
	Init(*Engine, string) error
	SetActive(bool)
	GetSceneID() string

	// Widgets
	InsertWidget(Widget) error
	DeleteWidget(string) error
	GetWidget(string) Widget
	HasWidgetID(string) bool
	GetWidgetIDs() []string
}
