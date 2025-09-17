package core

import (
	"sync/atomic"
)

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
func (e BaseEntity) Id() uint64 {
	return e.id
}

// OnAdd provides a default empty implementation of the OnAdd lifecycle method.
func (e BaseEntity) OnAdd() {}

// OnRemove provides a default empty implementation of the OnRemove lifecycle method.
func (e BaseEntity) OnRemove() {}

// Update provides a default empty implementation of the Update method.
func (e BaseEntity) Update(deltaTime float32) {}
