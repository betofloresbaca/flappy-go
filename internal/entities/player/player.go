// Package player provides the player entity implementation.
package player

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"simple-go-game/internal/core/entity"
)

const (
	// PlayerSize is the width and height of the player rectangle
	PlayerSize = 20
)

// Player represents the main player character in the game.
// It embeds BaseEntity and BaseDrawable to inherit core functionality.
type Player struct {
	*entity.BaseEntity
	*entity.BaseDrawable

	// Player-specific properties
	position rl.Vector2
	speed    float32
	color    rl.Color
}

// NewPlayer creates a new player entity at the specified position.
func NewPlayer(x, y float32) *Player {
	return &Player{
		BaseEntity:   entity.NewBaseEntity(),
		BaseDrawable: entity.NewBaseDrawable(1), // Z-index 1 for player layer
		position:     rl.Vector2{X: x, Y: y},
		speed:        100.0, // pixels per second
		color:        rl.Blue,
	}
}

// Position returns the current position of the player.
func (p *Player) Position() rl.Vector2 {
	return p.position
}

// SetPosition sets the player's position.
func (p *Player) SetPosition(x, y float32) {
	p.position = rl.Vector2{X: x, Y: y}
}

// Update handles player input and movement.
func (p *Player) Update(dt float64) {
	// Call base implementation first
	p.BaseEntity.Update(dt)

	// Handle player movement input
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		p.position.X += p.speed * float32(dt)
	}
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		p.position.X -= p.speed * float32(dt)
	}
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		p.position.Y -= p.speed * float32(dt)
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		p.position.Y += p.speed * float32(dt)
	}

	// Keep player within screen bounds using raylib's Clamp function
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	p.position.X = rl.Clamp(p.position.X, 0, screenWidth-PlayerSize)
	p.position.Y = rl.Clamp(p.position.Y, 0, screenHeight-PlayerSize)
}

// Draw renders the player to the screen.
func (p *Player) Draw() {
	// Draw player as a simple rectangle for now
	rl.DrawRectangle(int32(p.position.X), int32(p.position.Y), PlayerSize, PlayerSize, p.color)

	// Draw a simple border
	rl.DrawRectangleLines(int32(p.position.X), int32(p.position.Y), PlayerSize, PlayerSize, rl.DarkBlue)
}
