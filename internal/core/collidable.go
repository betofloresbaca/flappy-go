package core

// Collidable interface for entities that can participate in collision detection
type Collidable interface {
	GetCollider() *Collider
	GetTransform() *Transform
	OnCollision(other Collidable, collision CollisionInfo)
}
