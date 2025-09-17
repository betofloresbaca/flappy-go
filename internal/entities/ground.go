package entities

import (
	"simple-go-game/internal/assets"
	"simple-go-game/internal/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Ground_ZIndex = -100
	Ground_Y      = 440
)

type Ground struct {
	*core.BaseEntity
	*core.BaseDrawable
	sprite   core.Sprite
	speed    float32
	offset   float32
	Running  bool
	collider *core.Collider
}

func NewGround(speed float32) *Ground {
	sprite := core.NewSprite(assets.GroundImage, core.PivotUpLeft)
	groundWidth := float32(sprite.Texture.Width) * 4 // Account for multiple tiles
	groundHeight := float32(sprite.Texture.Height)

	return &Ground{
		BaseEntity:   core.NewBaseEntity(),
		BaseDrawable: core.NewBaseDrawable(Ground_ZIndex),
		sprite:       *sprite,
		speed:        speed,
		collider:     core.NewRectangleCollider(groundWidth, groundHeight, "ground", []string{"player"}),
	}
}

func (g *Ground) Update(dt float32) {
	if !g.Running {
		return
	}
	// Move the ground to the left
	g.offset -= g.speed * dt
	if g.offset <= -float32(g.sprite.Texture.Width) {
		g.offset += float32(g.sprite.Texture.Width)
	}
}

func (g *Ground) Draw() {
	for i := range 4 {
		g.sprite.Draw(
			*core.NewTransform(float32(i)*float32(g.sprite.Texture.Width)+g.offset, Ground_Y),
		)
	}
}

// GetBounds returns the collision bounds for the ground
func (g *Ground) GetBounds() rl.Rectangle {
	transform := core.NewTransform(g.offset, Ground_Y)
	return g.collider.GetWorldBounds(*transform)
}

// GetCollider returns the ground's collider
func (g *Ground) GetCollider() interface{} {
	return g.collider
}

// GetTransform returns the ground's transform
func (g *Ground) GetTransform() *core.Transform {
	return core.NewTransform(g.offset, Ground_Y)
}

// OnCollision handles collision events (ground doesn't need to react to collisions)
func (g *Ground) OnCollision(other core.Collidable, collision core.CollisionInfo) {
	// Ground doesn't need to react to collisions, the other entity will handle it
}
