package entities

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"simple-go-game/internal/assets"
	"simple-go-game/internal/components"
	"simple-go-game/internal/core"
)

const (
	Player_ZIndex         = 0
	Player_Size           = 20
	Player_StartPositionX = 100
	Player_StartPositionY = 100
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
func NewPlayer(color string) *Player {
	animatedSprite := components.NewAnimatedSprite()
	for _, birdColor := range []string{"blue", "red", "yellow"} {
		frames := assets.BirdSprites[birdColor]
		animatedSprite.AddAnimation(birdColor, frames, 0.2, true)
	}
	animatedSprite.SetAnimation(color)
	return &Player{
		BaseEntity:     core.NewBaseEntity(),
		BaseDrawable:   core.NewBaseDrawable(Player_ZIndex),
		transform:      *core.NewTransform(Player_StartPositionX, Player_StartPositionY),
		speed:          100.0, // pixels per second
		animatedSprite: *animatedSprite,
	}
}

// Update handles player input and movement.
func (p *Player) Update(dt float32) {
	// Update animated sprite
	p.animatedSprite.Update(dt)

	// Handle player movement input
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		p.transform.Position.X += p.speed * dt
	}
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		p.transform.Position.X -= p.speed * dt
	}
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		p.transform.Position.Y -= p.speed * dt
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		p.transform.Position.Y += p.speed * dt
	}

	// Keep player within screen bounds using raylib's Clamp function
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	p.transform.Position.X = rl.Clamp(p.transform.Position.X, 0, screenWidth-Player_Size)
	p.transform.Position.Y = rl.Clamp(p.transform.Position.Y, 0, screenHeight-Player_Size)
}

// Draw renders the player to the screen.
func (p *Player) Draw() {
	p.animatedSprite.Draw(p.transform)
}
