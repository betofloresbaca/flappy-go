// Package scene provides scene management functionality for the game.
// A scene manages a collection of entities and handles their lifecycle.
package core

import (
	"sort"

	physics "flappy-go/internal/core/physics"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

// Scene manages a collection of entities and handles their updates and rendering.
// It provides efficient entity lookup and maintains the entity lifecycle.
type Scene struct {
	*BaseEntity
	*BaseUpdater
	*BaseDrawer
	entities      []Entity
	entityIndices map[uint64]int
	handlePhysics bool
	gravity       raylib.Vector2
	inTree        bool
}

// NewScene creates a new empty scene.
func NewScene(parent *Scene, name string, groups []string, zIndex int) *Scene {
	s := &Scene{
		BaseEntity:    NewBaseEntity(parent, name, groups),
		BaseUpdater:   NewBaseUpdater(),
		BaseDrawer:    NewBaseDrawer(zIndex),
		entities:      make([]Entity, 0),
		entityIndices: make(map[uint64]int),
		handlePhysics: false,
		inTree:        false,
	}
	s.BaseEntity.OnAdd = s.onAdd
	s.BaseEntity.OnRemove = s.onRemove
	s.BaseUpdater.OnPause = s.onPause
	s.BaseUpdater.OnResume = s.onResume
	return s
}

// NewPhysicsScene creates a new empty physics scene.
func NewPhysicsScene(
	parent *Scene,
	name string,
	groups []string,
	zIndex int,
	gravity raylib.Vector2,
) *Scene {
	s := &Scene{
		BaseEntity:    NewBaseEntity(parent, name, groups),
		BaseUpdater:   NewBaseUpdater(),
		BaseDrawer:    NewBaseDrawer(zIndex),
		entities:      make([]Entity, 0),
		entityIndices: make(map[uint64]int),
		handlePhysics: true,
		gravity:       gravity,
		inTree:        false,
	}
	s.BaseEntity.OnAdd = s.onPhysicsAdd
	s.BaseEntity.OnRemove = s.onPhysicsRemove
	return s
}

// Add adds an entity to the scene. If the entity is already in the scene, this is a no-op.
// The entity's OnAdd method will be called after it's successfully added.
func (s *Scene) Add(e Entity) {
	if _, exists := s.entityIndices[e.Id()]; exists {
		return
	}

	s.entities = append(s.entities, e)
	s.entityIndices[e.Id()] = len(s.entities) - 1
	if s.inTree {
		e.added()
	}
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
	e.removed()
}

// ChildById returns the entity with the given ID, or nil if not found.
// The second return value indicates whether the entity was found.
func (s *Scene) ChildById(id uint64) (Entity, bool) {
	idx, exists := s.entityIndices[id]
	if !exists {
		return nil, false
	}
	return s.entities[idx], true
}

func (s *Scene) ChildrenByGroup(group string, recursive bool) []Entity {
	var result []Entity
	for _, e := range s.entities {
		if e.IsInGroup(group) {
			result = append(result, e)
		}
		if recursive {
			if childScene, ok := e.(*Scene); ok {
				result = append(result, childScene.ChildrenByGroup(group, true)...)
			}
		}
	}
	return result
}

func (s *Scene) ChildByName(name string) Entity {
	for _, e := range s.entities {
		if e.Name() == name {
			return e
		}
	}
	return nil
}

// Root returns the root scene. If this scene has no parent, it is the root.
func (s *Scene) Root() *Scene {
	if s.parent == nil {
		return s
	}
	return s.parent.Root()
}

// Update calls the Update method on all entities in the scene.
// dt is the delta time in seconds since the last frame.
func (s *Scene) Update(dt float32) {
	// If the scene is paused, skip physics and child updates entirely
	if s.Paused() {
		return
	}
	if s.handlePhysics {
		physics.Update()
	}
	for _, e := range s.entities {
		if up, ok := e.(Updater); ok && !up.Paused() {
			up.Update(dt)
		}
	}
}

// Draw renders all drawable entities in the scene, sorted by Z-index.
// Entities with lower Z-index values are drawn first (behind entities with higher values).
func (s *Scene) Draw() {
	var drawables []Drawer
	for _, e := range s.entities {
		if d, ok := e.(Drawer); ok && d.Visible() {
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

func (s *Scene) onPhysicsAdd() {
	if s.handlePhysics {
		physics.Init()
		physics.SetGravity(s.gravity.X, s.gravity.Y)
	}
	s.onAdd()
}

func (s *Scene) onPhysicsRemove() {
	s.onRemove()
	if s.handlePhysics {
		physics.Close()
	}
}

func (s *Scene) onAdd() {
	s.inTree = true
	for _, e := range s.entities {
		e.added()
	}
}

func (s *Scene) onRemove() {
	for _, e := range s.entities {
		e.removed()
	}
	s.entities = nil
	s.entityIndices = nil
	s.inTree = false
}

func (s *Scene) onPause() {
	for _, e := range s.entities {
		if up, ok := e.(Updater); ok && !up.Paused() {
			up.Pause()
		}
	}
}

func (s *Scene) onResume() {
	for _, e := range s.entities {
		if up, ok := e.(Updater); ok && up.Paused() {
			up.Resume()
		}
	}
}
