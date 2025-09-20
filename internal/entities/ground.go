package entities

import (
	"flappy-go/internal/assets"
	"flappy-go/internal/core"
	physics "flappy-go/internal/core/physics"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	Ground_ZIndex = -100
	Ground_Y      = 440
)

type Ground struct {
	*core.BaseEntity
	*core.BaseUpdater
	*core.BaseDrawer
	sprite core.Sprite
	speed  float32
	offset float32
	body   *physics.Body
}

func NewGround(parent *core.Scene, speed float32) *Ground {
	return &Ground{
		BaseEntity:  core.NewBaseEntity(parent, "ground"),
		BaseUpdater: core.NewBaseUpdater(),
		BaseDrawer:  core.NewBaseDrawer(Ground_ZIndex),
		sprite:      *core.NewSprite(assets.GroundImage, core.PivotUpLeft),
		speed:       speed,
	}
}

func (g *Ground) Update(dt float32) {
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

// OnAdd creates the static physics body for ground collision
func (g *Ground) OnAdd() {
	screenWidth := float32(raylib.GetScreenWidth())
	screenHeight := float32(raylib.GetScreenHeight())

	// Create a wide rectangle from Ground_Y to bottom of screen
	groundHeight := screenHeight - Ground_Y
	groundWidth := screenWidth

	// Position at center of the ground area
	centerX := groundWidth / 2
	centerY := Ground_Y + groundHeight/2

	// Create static body (density 0 makes it static)
	g.body = physics.NewBodyRectangle(
		"Ground",
		raylib.Vector2{X: centerX, Y: centerY},
		groundWidth,
		groundHeight,
		0, // density 0 = static body
	)
}

// OnRemove cleans up the physics body
func (g *Ground) OnRemove() {
	if g.body != nil {
		g.body.Destroy()
	}
}
