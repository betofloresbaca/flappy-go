package core

import (
	"sync/atomic"
)

var nextId uint64

// BaseEntity provides a basic implementation of the Entity interface.
// It handles ID generation and provides default empty implementations of lifecycle methods.
type BaseEntity struct {
	id     uint64
	group  string
	parent *Scene
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

// OnAdd provides a default empty implementation of the OnAdd lifecycle method.
func (e BaseEntity) OnAdd() {}

// OnRemove provides a default empty implementation of the OnRemove lifecycle method.
func (e BaseEntity) OnRemove() {}

// Update provides a default empty implementation of the Update method.
func (e BaseEntity) Update(deltaTime float32) {}

// GetGroup returns the group of the entity. Default is empty string.
func (e BaseEntity) GetGroup() string {
	return e.group
}

// SetGroup sets the group of the entity. Default does nothing.
func (e *BaseEntity) SetGroup(group string) {
	e.group = group
}

// GetParent returns the parent scene of the entity. Default is nil.
func (e BaseEntity) GetParent() *Scene {
	return e.parent
}
