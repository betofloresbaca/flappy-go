package entities

import (
	"simple-go-game/internal/assets"
	"simple-go-game/internal/components"
	"simple-go-game/internal/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Player_ZIndex         = 0
	Player_Size           = 20
	Player_StartPositionX = 150
	Player_StartPositionY = 100
	Player_Gravity        = 900.0
	Player_MaxVelocityY   = 500.0
	Player_JumpForce      = 300.0
	Player_MaxRotation    = 75.0
)

// Player represents the main player character in the game.
// It embeds BaseEntity and BaseDrawable to inherit core functionality.
type Player struct {
	*core.BaseEntity
	*core.BaseDrawable

	// Player-specific properties
	transform      core.Transform
	velocity       rl.Vector2
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
	p.animatedSprite.Update(dt)
	p.applyGravity(dt)
	p.applyInputforce()
	p.setFacingAngle()
	p.updatePosition(dt)

}

func (p *Player) applyGravity(dt float32) {
	p.velocity.Y += Player_Gravity * dt
	p.velocity.Y = rl.Clamp(p.velocity.Y, -Player_MaxVelocityY, Player_MaxVelocityY)
}

func (p *Player) applyInputforce() {
	if rl.IsKeyPressed(rl.KeySpace) || rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		p.velocity.Y = -Player_JumpForce
	}
}

func (p *Player) setFacingAngle() {
	p.transform.Rotation = (p.velocity.Y / Player_MaxVelocityY) * Player_MaxRotation
}

func (p *Player) updatePosition(dt float32) {
	// Update position based on velocity
	p.transform.Position.X += p.velocity.X * dt
	p.transform.Position.Y += p.velocity.Y * dt

	// Keep player within screen bounds using raylib's Clamp function
	screenHeight := float32(rl.GetScreenHeight())
	p.transform.Position.Y = rl.Clamp(p.transform.Position.Y, 0, screenHeight-Player_Size)
}

// Draw renders the player to the screen.
func (p *Player) Draw() {
	p.animatedSprite.Draw(p.transform)
}
