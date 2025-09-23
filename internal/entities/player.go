package entities

import (
	"flappy-go/internal/assets"
	"flappy-go/internal/core"
	"flappy-go/internal/ui"
	"log"

	physics "flappy-go/internal/core/physics"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	Player_ZIndex             = 0
	Player_Size               = 20
	Player_StartPositionX     = 150
	Player_StartPositionY     = 100
	Player_MaxVelocityY       = 500.0
	Player_JumpForce          = 300.0
	Player_DeathForce         = 800.0
	Player_MaxRotation        = 75.0
	Player_AnimationFrameTime = 0.2
)

// Player represents the main player character in the game.
// It embeds BaseEntity and BaseDrawable to inherit core functionality.
type Player struct {
	*core.BaseEntity
	*core.BaseUpdater
	*core.BaseDrawer
	animatedSprite *core.AnimatedSprite
	body           *physics.Body
	scoreDisplay   *ui.ScoreDisplay
	transform      core.Transform
	isDead         bool
}

// NewPlayer creates a new player entity at the specified position.
func NewPlayer(parent *core.Scene, color string) *Player {
	animatedSprite := core.NewAnimatedSprite()
	for _, birdColor := range []string{"blue", "red", "yellow"} {
		frames := assets.BirdImages[birdColor]
		animatedSprite.AddAnimation(birdColor, frames, Player_AnimationFrameTime, true)
	}
	animatedSprite.SetAnimation(color)
	p := &Player{
		BaseEntity:     core.NewBaseEntity(parent, "player", ""),
		BaseUpdater:    core.NewBaseUpdater(),
		BaseDrawer:     core.NewBaseDrawer(Player_ZIndex),
		animatedSprite: animatedSprite,
		transform:      *core.NewTransform(Player_StartPositionX, Player_StartPositionY),
		isDead:         false,
	}
	p.BaseUpdater.OnPause = p.onPause
	p.BaseUpdater.OnResume = p.onResume
	p.BaseEntity.OnAdd = p.onAdd
	p.BaseEntity.OnRemove = p.onRemove
	return p
}

// Update handles player input and movement.
func (p *Player) Update(dt float32) {
	if p.isDead {
		if p.body != nil {
			p.body.Velocity.X = 0
		}
	} else {
		p.animatedSprite.Update(dt)
		// Input: jump
		if raylib.IsKeyPressed(raylib.KeySpace) ||
			raylib.IsMouseButtonPressed(raylib.MouseLeftButton) {
			if p.body != nil {
				p.body.Velocity.Y = -Player_JumpForce
			}
		}
	}

	// Clamp vertical velocity to control arcade feel
	if p.body != nil {
		if p.body.Velocity.Y > Player_MaxVelocityY {
			p.body.Velocity.Y = Player_MaxVelocityY
		} else if p.body.Velocity.Y < -Player_MaxVelocityY {
			p.body.Velocity.Y = -Player_MaxVelocityY
		}
	}

	// Synchronize transform only from the physics body
	if p.body != nil {
		p.transform.Position = p.body.Position
		p.transform.Rotation = (p.body.Velocity.Y / Player_MaxVelocityY) * Player_MaxRotation
	}

	// Limit within the vertical bounds of the screen by adjusting the body
	screenHeight := float32(raylib.GetScreenHeight())
	if p.body != nil {
		if p.body.Position.Y < 0 {
			p.body.Position.Y = 0
			if p.body.Velocity.Y < 0 {
				p.body.Velocity.Y = 0
			}
		} else if p.body.Position.Y > screenHeight-Player_Size {
			p.body.Position.Y = screenHeight - Player_Size
			if p.body.Velocity.Y > 0 {
				p.body.Velocity.Y = 0
			}
		}
		p.transform.Position = p.body.Position
	}
}

// Draw renders the player to the screen.
func (p *Player) Draw() {
	p.animatedSprite.Draw(p.transform)
}

// Override onAdd and OnRemove
func (p *Player) onAdd() {
	// Reduced density to make impulses more sensitive
	p.body = physics.NewBodyRectangle(
		"Player",
		p.transform.Position,
		Player_Size,
		Player_Size,
		1,
	)

	// Set collision callback for logging
	p.body.OnCollision = p.onCollision
	if p.Paused() {
		p.body.Enabled = false
	}
}

func (p *Player) onRemove() {
	p.body.Destroy()
}

func (p *Player) onPause() {
	if p.body != nil {
		p.body.Enabled = false
	}
}

func (p *Player) onResume() {
	if p.body != nil {
		p.body.Enabled = true
	}
}

func (p *Player) onCollision(other *physics.Body, manifold *physics.Manifold) {
	if p.Paused() || p.isDead {
		return
	}
	log.Println("Player collided with", other.Tag)
	switch other.Tag {
	case "pipe_gate_score":
		other.Destroy() // Disable score trigger after scoring
		p.searchScoreDisplay().Increment()
	case "ground", "pipe_gate":
		p.die()
	}
}

func (p *Player) searchScoreDisplay() *ui.ScoreDisplay {
	if p.scoreDisplay == nil {
		p.scoreDisplay = p.Root().ChildByName("ui").(*core.Scene).ChildByName("score_display").(*ui.ScoreDisplay)
		if p.scoreDisplay == nil {
			log.Println("Warning: Player could not find ScoreDisplay entity in scene")
		}
	}
	return p.scoreDisplay
}

func (p *Player) die() {
	pipeGates := p.Parent().ChildrenByGroup("pipe_gate", false)
	ground := p.Parent().ChildByName("ground").(*Ground)
	// Pause all gates
	for _, gate := range pipeGates {
		if updater, ok := gate.(*PipeGate); ok {
			updater.Pause()
		}
	}

	// Pause the ground entity (only one expected)
	ground.Pause()
	p.isDead = true
}
