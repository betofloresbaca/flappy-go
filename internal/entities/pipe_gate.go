package entities

import (
	"flappy-go/internal/assets"
	"flappy-go/internal/core"
	"fmt"

	physics "flappy-go/internal/core/physics"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	PipeGate_Group           = "pipe_gate"
	PipeGate_ScoreTriggerTag = "pipe_gate_score"
	PipeGate_PipeBodyTag     = "pipe_gate_body"
	PipeGate_ZIndex          = -300
	PipeGate_GapYMin         = 100
	PipeGate_GapYMax         = 350
	PipeGate_GapHeight       = 100
)

type PipeGate struct {
	*core.BaseEntity
	*core.BaseUpdater
	*core.BaseDrawer
	topSprite    *core.Sprite
	bottomSprite *core.Sprite
	topBody      *physics.Body
	bottomBody   *physics.Body
	scoreBody    *physics.Body
	gapY         float32
	speed        float32
	Running      bool
	initialX     float32 // Only used for initialization
}

func NewPipeGate(parent *core.Scene, index int, x, speed float32) *PipeGate {
	topSprite := core.NewSprite(assets.PipeSprites["green"], core.PivotCenter)
	topSprite.FlipV = true
	bottomSprite := core.NewSprite(assets.PipeSprites["green"], core.PivotCenter)

	pg := &PipeGate{
		BaseEntity: core.NewBaseEntity(
			parent,
			fmt.Sprintf("pipe_gate_%d", index),
			[]string{PipeGate_Group},
		),
		BaseUpdater:  core.NewBaseUpdater(),
		BaseDrawer:   core.NewBaseDrawer(PipeGate_ZIndex),
		topSprite:    topSprite,
		bottomSprite: bottomSprite,
		gapY:         float32(raylib.GetRandomValue(PipeGate_GapYMin, PipeGate_GapYMax)),
		speed:        speed,
		initialX:     x,
	}
	pg.BaseEntity.OnAdd = pg.onAdd
	pg.BaseEntity.OnRemove = pg.onRemove
	pg.BaseUpdater.OnPause = pg.onPause
	pg.BaseUpdater.OnResume = pg.onResume
	return pg
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
	if pg.scoreBody != nil {
		pg.scoreBody.Velocity.X = -pg.speed
	}

	// Check if pipes have passed (using topBody position)
	if pg.topBody != nil {
		currentX := pg.topBody.Position.X
		if currentX < -float32(pg.topSprite.Texture.Width)/2 {
			pg.Parent().Remove(pg)
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
	if pg.bottomBody != nil {
		return pg.bottomBody.Position.X
	}
	// Fallback to initial position if body doesn't exist yet
	return pg.initialX
}

// onAdd creates the static physics bodies for the pipes
func (pg *PipeGate) onAdd() {
	pipeWidth := float32(pg.topSprite.Texture.Width)
	pipeHeight := float32(pg.topSprite.Texture.Height)

	// Top pipe (above the gap, pivot DownLeft)
	topCenterX := pg.initialX + pipeWidth/2
	topCenterY := (pg.gapY - float32(PipeGate_GapHeight/2)) - pipeHeight/2
	pg.topBody = physics.NewBodyRectangle(
		PipeGate_PipeBodyTag,
		raylib.Vector2{X: topCenterX, Y: topCenterY},
		pipeWidth,
		pipeHeight,
		0,
	)
	if pg.topBody != nil {
		pg.topBody.UseGravity = false
	}

	// Score body (covers the gap between pipes)
	scoreCenterX := pg.initialX + pipeWidth/2
	scoreCenterY := pg.gapY
	scoreWidth := pipeWidth / 4
	scoreHeight := float32(PipeGate_GapHeight)
	pg.scoreBody = physics.NewTriggerRectangle(
		PipeGate_ScoreTriggerTag,
		raylib.Vector2{X: scoreCenterX, Y: scoreCenterY},
		scoreWidth,
		scoreHeight,
	)

	// Bottom pipe (below the gap, pivot UpLeft)
	bottomCenterX := pg.initialX + pipeWidth/2
	bottomCenterY := (pg.gapY + float32(PipeGate_GapHeight/2)) + pipeHeight/2
	pg.bottomBody = physics.NewBodyRectangle(
		PipeGate_PipeBodyTag,
		raylib.Vector2{X: bottomCenterX, Y: bottomCenterY},
		pipeWidth,
		pipeHeight,
		0,
	)
	if pg.bottomBody != nil {
		pg.bottomBody.UseGravity = false
	}

	if pg.Paused() {
		pg.onPause()
	}
}

// onRemove destroys the physics bodies for the pipes
func (pg *PipeGate) onRemove() {
	pg.Running = false
	if pg.topBody != nil {
		pg.topBody.Destroy()
		pg.topBody = nil
	}
	if pg.scoreBody != nil {
		pg.scoreBody.Destroy()
		pg.scoreBody = nil
	}
	if pg.bottomBody != nil {
		pg.bottomBody.Destroy()
		pg.bottomBody = nil
	}
}

func (pg *PipeGate) onPause() {
	if pg.topBody != nil {
		pg.topBody.Paused = true
	}
	if pg.scoreBody != nil {
		pg.scoreBody.Paused = true
	}
	if pg.bottomBody != nil {
		pg.bottomBody.Paused = true
	}
}

func (pg *PipeGate) onResume() {
	if pg.topBody != nil {
		pg.topBody.Paused = false
	}
	if pg.scoreBody != nil {
		pg.scoreBody.Paused = false
	}
	if pg.bottomBody != nil {
		pg.bottomBody.Paused = false
	}
}
