package core

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// PhysicsSystem manages collision detection for all collidable entities using a layer-based system
type PhysicsSystem struct {
	collidables []Collidable
}

// NewPhysicsSystem creates a new physics system
func NewPhysicsSystem() *PhysicsSystem {
	return &PhysicsSystem{
		collidables: make([]Collidable, 0),
	}
}

// Register adds a collidable entity to the physics system
func (ps *PhysicsSystem) Register(collidable Collidable) {
	ps.collidables = append(ps.collidables, collidable)
}

// Unregister removes a collidable entity from the physics system
func (ps *PhysicsSystem) Unregister(collidable Collidable) {
	for i, c := range ps.collidables {
		if c == collidable {
			ps.collidables = append(ps.collidables[:i], ps.collidables[i+1:]...)
			break
		}
	}
}

// Update checks for collisions between all registered collidables based on their layers and masks
func (ps *PhysicsSystem) Update(dt float32) {
	for i := 0; i < len(ps.collidables); i++ {
		for j := i + 1; j < len(ps.collidables); j++ {
			ps.checkCollision(ps.collidables[i], ps.collidables[j])
		}
	}
}

// checkCollision tests collision between two collidable entities using their layer/mask configuration
func (ps *PhysicsSystem) checkCollision(a, b Collidable) {
	colliderA := a.GetCollider()
	colliderB := b.GetCollider()

	// Check if these colliders should collide based on their layers and masks
	if !colliderA.ShouldCollideWith(colliderB) {
		return
	}

	transformA := a.GetTransform()
	transformB := b.GetTransform()

	var collision bool
	var infoA, infoB CollisionInfo

	if colliderA.Type == ColliderRectangle && colliderB.Type == ColliderRectangle {
		// Rectangle-Rectangle collision
		recA := colliderA.GetWorldBounds(*transformA)
		recB := colliderB.GetWorldBounds(*transformB)
		collision = rl.CheckCollisionRecs(recA, recB)
		if collision {
			overlapRec := rl.GetCollisionRec(recA, recB)
			point := rl.Vector2{
				X: overlapRec.X + overlapRec.Width/2,
				Y: overlapRec.Y + overlapRec.Height/2,
			}
			infoA = CollisionInfo{
				Point:      point,
				Overlap:    overlapRec.Width * overlapRec.Height,
				OtherLayer: colliderB.Layer,
			}
			infoB = CollisionInfo{
				Point:      point,
				Overlap:    overlapRec.Width * overlapRec.Height,
				OtherLayer: colliderA.Layer,
			}
		}
	} else if colliderA.Type == ColliderCircle && colliderB.Type == ColliderCircle {
		// Circle-Circle collision
		centerA := colliderA.GetWorldCenter(*transformA)
		centerB := colliderB.GetWorldCenter(*transformB)
		radiusA := colliderA.GetWorldRadius(*transformA)
		radiusB := colliderB.GetWorldRadius(*transformB)
		collision = rl.CheckCollisionCircles(centerA, radiusA, centerB, radiusB)
		if collision {
			point := rl.Vector2{
				X: (centerA.X + centerB.X) / 2,
				Y: (centerA.Y + centerB.Y) / 2,
			}
			infoA = CollisionInfo{
				Point:      point,
				OtherLayer: colliderB.Layer,
			}
			infoB = CollisionInfo{
				Point:      point,
				OtherLayer: colliderA.Layer,
			}
		}
	} else {
		// Mixed Circle-Rectangle collision
		var center rl.Vector2
		var radius float32
		var rec rl.Rectangle

		if colliderA.Type == ColliderCircle {
			center = colliderA.GetWorldCenter(*transformA)
			radius = colliderA.GetWorldRadius(*transformA)
			rec = colliderB.GetWorldBounds(*transformB)
		} else {
			center = colliderB.GetWorldCenter(*transformB)
			radius = colliderB.GetWorldRadius(*transformB)
			rec = colliderA.GetWorldBounds(*transformA)
		}

		collision = rl.CheckCollisionCircleRec(center, radius, rec)
		if collision {
			infoA = CollisionInfo{
				Point:      center,
				OtherLayer: colliderB.Layer,
			}
			infoB = CollisionInfo{
				Point:      center,
				OtherLayer: colliderA.Layer,
			}
		}
	}

	if collision {
		a.OnCollision(b, infoA)
		b.OnCollision(a, infoB)
	}
}
