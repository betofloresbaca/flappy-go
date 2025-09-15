// Package entity provides the core entity system for the game.
// It defines the Entity and Drawable interfaces along with their base implementations.
package core

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
	Update(dt float32)
}
