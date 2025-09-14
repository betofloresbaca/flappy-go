// Package player provides the player entity implementation.
package player

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"simple-go-game/internal/assets"
	"simple-go-game/internal/components"
	"simple-go-game/internal/core"
)

const (
	// PlayerSize is the width and height of the player rectangle
	PlayerSize = 20
)

// Player represents the main player character in the game.
// It embeds BaseEntity and BaseDrawable to inherit core functionality.
type Player struct {
	*core.BaseEntity
	*core.BaseDrawable

	// Player-specific properties
	transform      core.Transform
	speed          float32
	animatedSprite components.AnimatedSprite
}

// NewPlayer creates a new player entity at the specified position.
func NewPlayer(x, y float32, color string) *Player {
	animatedSprite := components.NewAnimatedSprite()
	for _, birdColor := range []string{"blue", "red", "yellow"} {
		frames := assets.BirdSprites[birdColor]
		animatedSprite.AddAnimation(birdColor, frames, 0.2, true)
	}
	animatedSprite.SetAnimation(color)
	return &Player{
		BaseEntity:     core.NewBaseEntity(),
		BaseDrawable:   core.NewBaseDrawable(1), // Z-index 1 for player layer
		transform:      core.Transform{Position: rl.Vector2{X: x, Y: y}, Scale: rl.Vector2{X: 1, Y: 1}, Rotation: 0},
		speed:          100.0, // pixels per second
		animatedSprite: *animatedSprite,
	}
}

// Update handles player input and movement.
func (p *Player) Update(dt float64) {
	// Call base implementation first
	p.BaseEntity.Update(dt)
	// Update animated sprite
	p.animatedSprite.Update(dt)

	// Handle player movement input
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		p.transform.Position.X += p.speed * float32(dt)
	}
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		p.transform.Position.X -= p.speed * float32(dt)
	}
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		p.transform.Position.Y -= p.speed * float32(dt)
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		p.transform.Position.Y += p.speed * float32(dt)
	}

	// Keep player within screen bounds using raylib's Clamp function
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	p.transform.Position.X = rl.Clamp(p.transform.Position.X, 0, screenWidth-PlayerSize)
	p.transform.Position.Y = rl.Clamp(p.transform.Position.Y, 0, screenHeight-PlayerSize)
}

// Draw renders the player to the screen.
func (p *Player) Draw() {
	p.animatedSprite.Draw(p.transform)
}
