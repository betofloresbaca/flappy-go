// Package scene provides scene management functionality for the game.
// A scene manages a collection of entities and handles their lifecycle.
package core

import (
	"sort"
)

// Scene manages a collection of entities and handles their updates and rendering.
// It provides efficient entity lookup and maintains the entity lifecycle.
type Scene struct {
	entities      []Entity
	entityIndices map[uint64]int
	PhysicsSystem PhysicsSystem
}

// NewScene creates a new empty scene.
func NewScene() *Scene {
	return &Scene{
		entities:      make([]Entity, 0),
		entityIndices: make(map[uint64]int),
		PhysicsSystem: *NewPhysicsSystem(),
	}
}

// Add adds an entity to the scene. If the entity is already in the scene, this is a no-op.
// The entity's OnAdd method will be called after it's successfully added.
func (s *Scene) Add(e Entity) {
	if _, exists := s.entityIndices[e.Id()]; exists {
		return
	}

	s.entities = append(s.entities, e)
	s.entityIndices[e.Id()] = len(s.entities) - 1
	e.OnAdd(s)
}

// Remove removes an entity from the scene. If the entity is not in the scene, this is a no-op.
// The entity's OnRemove method will be called after it's successfully removed.
// Uses swap-and-pop for O(1) removal.
func (s *Scene) Remove(e Entity) {
	idToRemove := e.Id()
	idxToRemove, exists := s.entityIndices[idToRemove]
	if !exists {
		return
	}
	lastIndex := len(s.entities) - 1
	lastEntity := s.entities[lastIndex]
	s.entities[idxToRemove] = lastEntity
	s.entityIndices[lastEntity.Id()] = idxToRemove
	s.entities = s.entities[:lastIndex]
	delete(s.entityIndices, idToRemove)
	e.OnRemove(s)
}

// Update calls the Update method on all entities in the scene.
// dt is the delta time in seconds since the last frame.
func (s *Scene) Update(dt float32) {
	for _, e := range s.entities {
		e.Update(dt)
	}
}

// Draw renders all drawable entities in the scene, sorted by Z-index.
// Entities with lower Z-index values are drawn first (behind entities with higher values).
func (s *Scene) Draw() {
	var drawables []Drawable
	for _, e := range s.entities {
		if d, ok := e.(Drawable); ok {
			drawables = append(drawables, d)
		}
	}

	sort.Slice(drawables, func(i, j int) bool {
		return drawables[i].ZIndex() < drawables[j].ZIndex()
	})

	for _, d := range drawables {
		d.Draw()
	}
}

// EntityById returns the entity with the given ID, or nil if not found.
// The second return value indicates whether the entity was found.
func (s *Scene) EntityById(id uint64) (Entity, bool) {
	idx, exists := s.entityIndices[id]
	if !exists {
		return nil, false
	}
	return s.entities[idx], true
}
