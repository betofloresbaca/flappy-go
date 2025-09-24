// Package entity provides the core entity system for the game.
// It defines the Entity and Drawable interfaces along with their base implementations.
package core

import (
	"sync/atomic"
)

// Entity represents a game object that can be updated and managed by the scene.
// All entities must have a unique ID and implement lifecycle methods.
type Entity interface {
	// Id returns the unique identifier for this entity. It is unique across all entities.
	Id() uint64
	// Name returns the name of the entity.
	// It can be repeated across different entities but must be unique within the same parent.
	Name() string
	// Get groups of the entity
	Groups() []string
	// Set groups of the entity
	SetGroups(groups []string)
	// Check if the entity is in a specific group
	IsInGroup(group string) bool
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
	name     string
	groups   map[string]struct{}
	parent   *Scene
	OnAdd    func()
	OnRemove func()
}

// NewBaseEntity creates a new BaseEntity with a unique ID.
func NewBaseEntity(parent *Scene, name string, groups []string) *BaseEntity {
	be := &BaseEntity{
		id:     atomic.AddUint64(&nextId, 1),
		name:   name,
		parent: parent,
	}
	be.SetGroups(groups)
	return be
}

// Id returns the unique identifier for this entity.
func (e BaseEntity) Id() uint64 {
	return e.id
}

// Name returns the name of the entity.
func (e BaseEntity) Name() string {
	return e.name
}

// Update provides a default empty implementation of the Update method.
func (e BaseEntity) Update(deltaTime float32) {}

// Groups returns the groups of the entity. Default is empty slice.
func (e BaseEntity) Groups() []string {
	groups := make([]string, 0, len(e.groups))
	for group := range e.groups {
		groups = append(groups, group)
	}
	return groups
}

// SetGroups sets the groups of the entity. Default does nothing.
func (e *BaseEntity) SetGroups(groups []string) {
	e.groups = make(map[string]struct{})
	for _, group := range groups {
		e.groups[group] = struct{}{}
	}
}

// IsInGroup checks if the entity is in a specific group.
func (e BaseEntity) IsInGroup(group string) bool {
	_, exists := e.groups[group]
	return exists
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
