package core

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// ColliderType defines the type of collider
type ColliderType int

const (
	ColliderRectangle ColliderType = iota
	ColliderCircle
)

// Collider represents a collision shape that can be attached to entities
type Collider struct {
	Type   ColliderType
	Bounds rl.Rectangle // For rectangles
	Radius float32      // For circles
	Offset rl.Vector2   // Offset from the transform position
	Layer  string       // The layer this collider belongs to
	Mask   []string     // Which layers this collider can collide with
}

// NewRectangleCollider creates a new rectangular collider
func NewRectangleCollider(width, height float32, layer string, mask []string) *Collider {
	return &Collider{
		Type:   ColliderRectangle,
		Bounds: rl.NewRectangle(0, 0, width, height),
		Layer:  layer,
		Mask:   mask,
	}
}

// NewCircleCollider creates a new circular collider
func NewCircleCollider(radius float32, layer string, mask []string) *Collider {
	return &Collider{
		Type:   ColliderCircle,
		Radius: radius,
		Layer:  layer,
		Mask:   mask,
	}
}

// GetWorldBounds returns the collider bounds in world coordinates
func (c Collider) GetWorldBounds(transform Transform) rl.Rectangle {
	if c.Type == ColliderRectangle {
		return rl.NewRectangle(
			transform.Position.X+c.Offset.X,
			transform.Position.Y+c.Offset.Y,
			c.Bounds.Width*transform.Scale.X,
			c.Bounds.Height*transform.Scale.Y,
		)
	}
	// For circles, return a bounding rectangle
	size := c.Radius * 2 * transform.Scale.X
	return rl.NewRectangle(
		transform.Position.X+c.Offset.X-c.Radius*transform.Scale.X,
		transform.Position.Y+c.Offset.Y-c.Radius*transform.Scale.Y,
		size,
		size,
	)
}

// GetWorldCenter returns the center point of the collider in world coordinates
func (c Collider) GetWorldCenter(transform Transform) rl.Vector2 {
	return rl.Vector2{
		X: transform.Position.X + c.Offset.X,
		Y: transform.Position.Y + c.Offset.Y,
	}
}

// GetWorldRadius returns the radius of a circular collider scaled by transform
func (c Collider) GetWorldRadius(transform Transform) float32 {
	if c.Type == ColliderCircle {
		return c.Radius * transform.Scale.X
	}
	return 0
}

// ShouldCollideWith checks if this collider should collide with another collider based on layers and masks
func (c Collider) ShouldCollideWith(other *Collider) bool {
	// Check if this collider's mask includes the other's layer OR vice versa
	return c.HasLayerInMask(other.Layer) || other.HasLayerInMask(c.Layer)
}

// HasLayerInMask checks if a specific layer is in this collider's mask
func (c Collider) HasLayerInMask(layer string) bool {
	for _, maskLayer := range c.Mask {
		if maskLayer == layer {
			return true
		}
	}
	return false
}
