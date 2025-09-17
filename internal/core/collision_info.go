package core

import rl "github.com/gen2brain/raylib-go/raylib"

// CollisionInfo contains information about a collision
type CollisionInfo struct {
	Point      rl.Vector2
	Normal     rl.Vector2
	Overlap    float32
	OtherLayer string
}
