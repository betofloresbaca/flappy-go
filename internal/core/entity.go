// Package entity provides the core entity system for the game.
// It defines the Entity and Drawable interfaces along with their base implementations.
package core

import (
	"sync/atomic"
)

// Entity represents a game object that can be updated and managed by the scene.
// All entities must have a unique ID and implement lifecycle methods.
type Entity interface {
	// Id returns the unique identifier for this entity.
	Id() uint64
	// Get group of the entity
	Group() string
	// Set group of the entity
	SetGroup(group string)
	// Parent entity
	Parent() *Scene
	// The root entity
	Root() *Scene
	// Call when entity is added to the scene
	added()
	// Call when entity is removed from the scene
	removed()
}

var nextId uint64

// BaseEntity provides a basic implementation of the Entity interface.
// It handles ID generation and provides default empty implementations of lifecycle methods.
type BaseEntity struct {
	id       uint64
	group    string
	parent   *Scene
	OnAdd    func()
	OnRemove func()
}

// NewBaseEntity creates a new BaseEntity with a unique ID.
func NewBaseEntity(parent *Scene, group string) *BaseEntity {
	return &BaseEntity{
		id:     atomic.AddUint64(&nextId, 1),
		group:  group,
		parent: parent,
	}
}

// Id returns the unique identifier for this entity.
func (e BaseEntity) Id() uint64 {
	return e.id
}

// Update provides a default empty implementation of the Update method.
func (e BaseEntity) Update(deltaTime float32) {}

// Group returns the group of the entity. Default is empty string.
func (e BaseEntity) Group() string {
	return e.group
}

// SetGroup sets the group of the entity. Default does nothing.
func (e *BaseEntity) SetGroup(group string) {
	e.group = group
}

// Parent returns the parent entity of the entity. Default is nil.
func (e BaseEntity) Parent() *Scene {
	return e.parent
}

// Root returns the root entity of the entity. Default is nil.
func (e *BaseEntity) Root() *Scene {
	if e.parent == nil {
		return nil
	}
	return e.parent.Root()
}

func (e BaseEntity) added() {
	if e.OnAdd != nil {
		e.OnAdd()
	}
}

func (e BaseEntity) removed() {
	if e.OnRemove != nil {
		e.OnRemove()
	}
}
