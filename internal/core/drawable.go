package core

// Drawable represents an entity that can be rendered to the screen.
// Drawables are sorted by ZIndex before rendering.
type Drawable interface {
	// ZIndex returns the rendering layer index. Lower values are drawn first.
	ZIndex() int
	// Draw renders the drawable to the screen.
	Draw()
}

// BaseDrawable provides a basic implementation of the Drawable interface.
type BaseDrawable struct {
	zIndex int
}

// NewBaseDrawable creates a new BaseDrawable with the specified Z-index.
func NewBaseDrawable(zIndex int) *BaseDrawable {
	return &BaseDrawable{zIndex: zIndex}
}

// ZIndex returns the rendering layer index for this drawable.
func (d BaseDrawable) ZIndex() int {
	return d.zIndex
}

// Draw provides a default empty implementation of the Draw method.
func (d BaseDrawable) Draw() {}
