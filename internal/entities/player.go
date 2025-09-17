package entities

import (
	"simple-go-game/internal/assets"
	"simple-go-game/internal/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Player_ZIndex             = 0
	Player_Size               = 20
	Player_StartPositionX     = 150
	Player_StartPositionY     = 100
	Player_Gravity            = 900.0
	Player_MaxVelocityY       = 500.0
	Player_JumpForce          = 300.0
	Player_MaxRotation        = 75.0
	Player_AnimationFrameTime = 0.2
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
	animatedSprite core.AnimatedSprite
	score          *ScoreDisplay
	collider       core.Collider
}

// NewPlayer creates a new player entity at the specified position.
func NewPlayer(color string, score *ScoreDisplay) *Player {
	animatedSprite := core.NewAnimatedSprite()
	for _, birdColor := range []string{"blue", "red", "yellow"} {
		frames := assets.BirdImages[birdColor]
		animatedSprite.AddAnimation(birdColor, frames, Player_AnimationFrameTime, true)
	}
	animatedSprite.SetAnimation(color)
	return &Player{
		BaseEntity:     core.NewBaseEntity(),
		BaseDrawable:   core.NewBaseDrawable(Player_ZIndex),
		transform:      *core.NewTransform(Player_StartPositionX, Player_StartPositionY),
		speed:          100.0, // pixels per second
		animatedSprite: *animatedSprite,
		score:          score,
		collider:       *core.NewCircleCollider(Player_Size/2, "player", []string{"pipe", "ground"}),
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
		p.score.Increment()
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

// GetCollider returns the player's collider
func (p *Player) GetCollider() interface{} {
	return p.collider
}

// GetTransform returns the player's transform
func (p *Player) GetTransform() *core.Transform {
	return &p.transform
}

// OnPipeCollision handles collision with pipes
func (p *Player) OnPipeCollision() {
	// Reset player position on pipe collision
	p.transform.Position.X = Player_StartPositionX
	p.transform.Position.Y = Player_StartPositionY
	p.velocity = rl.Vector2{X: 0, Y: 0}
	// Could also trigger game over state here
}

// OnGroundCollision handles collision with ground
func (p *Player) OnGroundCollision() {
	// Stop downward velocity when hitting ground
	if p.velocity.Y > 0 {
		p.velocity.Y = 0
	}
	// Keep player above ground
	screenHeight := float32(rl.GetScreenHeight())
	groundHeight := float32(112) // Assuming ground sprite height
	p.transform.Position.Y = screenHeight - groundHeight - Player_Size
}

// OnCollision handles generic collision events using the new layer system
func (p *Player) OnCollision(other core.Collidable, collision core.CollisionInfo) {
	switch collision.OtherLayer {
	case "pipe":
		p.OnPipeCollision()
	case "ground":
		p.OnGroundCollision()
	}
}
