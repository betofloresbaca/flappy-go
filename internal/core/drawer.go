package core

// Drawer represents an entity that can be rendered to the screen.
// Only visible drawers are rendered.
// Drawers are sorted by ZIndex before rendering.
type Drawer interface {
	// Visible returns whether the drawer is currently visible.
	Visible() bool
	// Shows the drawer.
	Show()
	// Hides the drawer.
	Hide()
	// ZIndex returns the rendering layer index. Lower values are drawn first.
	ZIndex() int
	// Draw renders the drawer to the screen.
	Draw()
}

// BaseDrawer provides a basic implementation of the Drawable interface.
type BaseDrawer struct {
	visible bool
	zIndex  int
}

// NewBaseDrawer creates a new BaseDrawer with the specified Z-index.
func NewBaseDrawer(zIndex int) *BaseDrawer {
	return &BaseDrawer{zIndex: zIndex, visible: true}
}

// ZIndex returns the rendering layer index for this drawable.
func (bd BaseDrawer) ZIndex() int {
	return bd.zIndex
}

// Draw provides a default empty implementation of the Draw method.
func (bd BaseDrawer) Draw() {}

// Visible returns whether the drawable is currently visible.
func (bd BaseDrawer) Visible() bool {
	return bd.visible
}

// Show makes the drawable visible.
func (bd *BaseDrawer) Show() {
	bd.visible = true
}

// Hide makes the drawable invisible.
func (bd *BaseDrawer) Hide() {
	bd.visible = false
}
