package entities

import (
	"simple-go-game/internal/assets"
	"simple-go-game/internal/core"

	physics "simple-go-game/internal/core/physics"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	PipeGate_ZIndex    = -300
	PipeGate_GapYMin   = 100
	PipeGate_GapYMax   = 350
	PipeGate_GapHeight = 100
)

type PipeGate struct {
	*core.BaseEntity
	*core.BaseDrawable
	topSprite    core.Sprite
	bottomSprite core.Sprite
	gapY         float32
	speed        float32
	Running      bool
	topBody      *physics.Body
	bottomBody   *physics.Body
	initialX     float32 // Only used for initialization
}

func NewPipeGate(parent *core.Scene, x, speed float32) *PipeGate {
	topSprite := core.NewSprite(assets.PipeSprites["green"], core.PivotCenter)
	topSprite.FlipV = true
	bottomSprite := core.NewSprite(assets.PipeSprites["green"], core.PivotCenter)

	return &PipeGate{
		BaseEntity:   core.NewBaseEntity(parent, "pipe_gate"),
		BaseDrawable: core.NewBaseDrawable(PipeGate_ZIndex),
		topSprite:    *topSprite,
		bottomSprite: *bottomSprite,
		gapY:         float32(raylib.GetRandomValue(PipeGate_GapYMin, PipeGate_GapYMax)),
		speed:        speed,
		initialX:     x,
	}
}

func (pg *PipeGate) Update(dt float32) {
	if !pg.Running {
		return
	}

	// Set constant velocity for the physics bodies
	if pg.topBody != nil {
		pg.topBody.Velocity.X = -pg.speed
	}
	if pg.bottomBody != nil {
		pg.bottomBody.Velocity.X = -pg.speed
	}

	// Check if pipes have passed (using topBody position)
	if pg.topBody != nil {
		currentX := pg.topBody.Position.X
		if currentX < -float32(pg.topSprite.Texture.Width)/2 {
			pg.GetParent().Remove(pg)
		}
	}
}

func (pg *PipeGate) Draw() {
	// Draw top pipe (above the gap) - now using center pivot like the body
	if pg.topBody != nil {
		pg.topSprite.Draw(*core.NewTransform(pg.topBody.Position.X, pg.topBody.Position.Y))
	}

	// Draw bottom pipe (below the gap) - now using center pivot like the body
	if pg.bottomBody != nil {
		pg.bottomSprite.Draw(*core.NewTransform(pg.bottomBody.Position.X, pg.bottomBody.Position.Y))
	}
}

// GetX returns the current X position of the pipe gate
func (pg *PipeGate) GetX() float32 {
	if pg.topBody != nil {
		return pg.topBody.Position.X
	}
	// Fallback to initial position if body doesn't exist yet
	return pg.initialX
}

// OnAdd creates the static physics bodies for the pipes
func (pg *PipeGate) OnAdd() {
	pipeWidth := float32(pg.topSprite.Texture.Width)
	pipeHeight := float32(pg.topSprite.Texture.Height)

	// Top pipe (above the gap, pivot DownLeft)
	topCenterX := pg.initialX + pipeWidth/2
	topCenterY := (pg.gapY - float32(PipeGate_GapHeight/2)) - pipeHeight/2
	pg.topBody = physics.NewBodyRectangle(
		"PipeGate",
		raylib.Vector2{X: topCenterX, Y: topCenterY},
		pipeWidth,
		pipeHeight,
		0, // density > 0 to allow velocity
	)
	pg.topBody.UseGravity = false

	// Bottom pipe (below the gap, pivot UpLeft)
	bottomCenterX := pg.initialX + pipeWidth/2
	bottomCenterY := (pg.gapY + float32(PipeGate_GapHeight/2)) + pipeHeight/2
	pg.bottomBody = physics.NewBodyRectangle(
		"PipeGate",
		raylib.Vector2{X: bottomCenterX, Y: bottomCenterY},
		pipeWidth,
		pipeHeight,
		0, // density > 0 to allow velocity
	)
	pg.bottomBody.UseGravity = false
}

// OnRemove destroys the physics bodies for the pipes
func (pg *PipeGate) OnRemove() {
	pg.Running = false
	if pg.topBody != nil {
		pg.topBody.Destroy()
		pg.topBody = nil
	}
	if pg.bottomBody != nil {
		pg.bottomBody.Destroy()
		pg.bottomBody = nil
	}
}
