package entities

import (
	"flappy-go/internal/assets"
	"flappy-go/internal/core"
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
	// Player-specific properties
	transform      core.Transform
	animatedSprite core.AnimatedSprite
	body           *physics.Body
	score          *ScoreDisplay
	isDead         bool
}

// NewPlayer creates a new player entity at the specified position.
func NewPlayer(parent *core.Scene, color string, score *ScoreDisplay) *Player {
	animatedSprite := core.NewAnimatedSprite()
	for _, birdColor := range []string{"blue", "red", "yellow"} {
		frames := assets.BirdImages[birdColor]
		animatedSprite.AddAnimation(birdColor, frames, Player_AnimationFrameTime, true)
	}
	animatedSprite.SetAnimation(color)
	return &Player{
		BaseEntity:     core.NewBaseEntity(parent, "player"),
		BaseUpdater:    core.NewBaseUpdater(),
		BaseDrawer:     core.NewBaseDrawer(Player_ZIndex),
		transform:      *core.NewTransform(Player_StartPositionX, Player_StartPositionY),
		animatedSprite: *animatedSprite,
		score:          score,
		isDead:         false,
	}
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

	// Clamp de velocidad vertical para controlar la sensación arcade
	if p.body != nil {
		if p.body.Velocity.Y > Player_MaxVelocityY {
			p.body.Velocity.Y = Player_MaxVelocityY
		} else if p.body.Velocity.Y < -Player_MaxVelocityY {
			p.body.Velocity.Y = -Player_MaxVelocityY
		}
	}

	// Sincroniza transform únicamente desde el cuerpo físico
	if p.body != nil {
		p.transform.Position = p.body.Position
		p.transform.Rotation = (p.body.Velocity.Y / Player_MaxVelocityY) * Player_MaxRotation
	}

	// Limita dentro de los bounds verticales de la pantalla ajustando el body
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

// Override onadd y on remove
func (p *Player) OnAdd() {
	// Densidad reducida para hacer más sensibles los impulsos
	p.body = physics.NewBodyRectangle(
		"Player",
		p.transform.Position,
		Player_Size,
		Player_Size,
		1,
	)

	// Configurar callback de colisión para logging
	p.body.OnCollision = p.OnCollision
	if p.Paused() {
		p.body.Enabled = false
	}
}

func (p *Player) OnRemove() {
	p.body.Destroy()
}

func (p *Player) OnCollision(other *physics.Body, manifold *physics.Manifold) {
	if p.Paused() || p.isDead {
		return
	}
	log.Println("Player collided with", other.Tag)
	switch other.Tag {
	case "pipe_gate_score":
		other.Destroy() // Disable score trigger after scoring
		p.score.Increment()
	case "ground", "pipe_gate":
		p.Die()
	}

}

func (p *Player) Die() {
	gates := p.Parent().GetEntitiesByGroup("pipe_gate")
	ground := p.Parent().GetEntitiesByGroup("ground")

	// Pause all gates
	for _, gate := range gates {
		if updatable, ok := gate.(core.Updater); ok {
			updatable.Pause()
		}
	}

	// Pause the ground entity (only one expected)
	if len(ground) > 0 {
		if updatable, ok := ground[0].(core.Updater); ok {
			updatable.Pause()
		}
	}
	p.isDead = true
}

func (p *Player) OnPause() {
	if p.body != nil {
		p.body.Enabled = false
	}
}

func (p *Player) OnResume() {
	if p.body != nil {
		p.body.Enabled = true
	}
}
