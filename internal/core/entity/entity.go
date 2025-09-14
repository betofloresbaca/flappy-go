// Package entity provides the core entity system for the game.
// It defines the Entity and Drawable interfaces along with their base implementations.
package entity

import "sync/atomic"

// Entity represents a game object that can be updated and managed by the scene.
// All entities must have a unique ID and implement lifecycle methods.
type Entity interface {
	// Id returns the unique identifier for this entity.
	Id() uint64
	// OnAdd is called when the entity is added to a scene.
	OnAdd()
	// OnRemove is called when the entity is removed from a scene.
	OnRemove()
	// Update is called every frame with the delta time in seconds.
	Update(dt float64)
}

var nextId uint64

// BaseEntity provides a basic implementation of the Entity interface.
// It handles ID generation and provides default empty implementations of lifecycle methods.
type BaseEntity struct {
	id uint64
}

// NewBaseEntity creates a new BaseEntity with a unique ID.
func NewBaseEntity() *BaseEntity {
	return &BaseEntity{id: atomic.AddUint64(&nextId, 1)}
}

// Id returns the unique identifier for this entity.
func (e *BaseEntity) Id() uint64 {
	return e.id
}

// OnAdd provides a default empty implementation of the OnAdd lifecycle method.
func (e *BaseEntity) OnAdd() {}

// OnRemove provides a default empty implementation of the OnRemove lifecycle method.
func (e *BaseEntity) OnRemove() {}

// Update provides a default empty implementation of the Update method.
func (e *BaseEntity) Update(deltaTime float64) {}

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
func (d *BaseDrawable) ZIndex() int {
	return d.zIndex
}

// Draw provides a default empty implementation of the Draw method.
func (d *BaseDrawable) Draw() {}
