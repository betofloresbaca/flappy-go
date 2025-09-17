package entities

import (
	"simple-go-game/internal/assets"
	"simple-go-game/internal/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	PipeGate_GapYMin   = 100
	PipeGate_GapYMax   = 350
	PipeGate_GapHeight = 100
)

type PipeGate struct {
	*core.BaseEntity
	*core.BaseDrawable
	topSprite      core.Sprite
	bottomSprite   core.Sprite
	x              float32
	gapY           float32
	speed          float32
	Passed         bool
	Running        bool
	topCollider    core.Collider
	bottomCollider core.Collider
	physicsSystem  *core.PhysicsSystem
}

func NewPipeGate(x, speed float32) *PipeGate {
	topSprite := core.NewSprite(assets.PipeSprites["green"], core.PivotDownLeft)
	topSprite.FlipV = true
	bottomSprite := core.NewSprite(assets.PipeSprites["green"], core.PivotUpLeft)

	pipeWidth := float32(topSprite.Texture.Width)
	pipeHeight := float32(topSprite.Texture.Height)

	return &PipeGate{
		BaseEntity:     core.NewBaseEntity(),
		BaseDrawable:   core.NewBaseDrawable(0),
		topSprite:      *topSprite,
		bottomSprite:   *bottomSprite,
		x:              x,
		gapY:           float32(rl.GetRandomValue(PipeGate_GapYMin, PipeGate_GapYMax)),
		speed:          speed,
		topCollider:    *core.NewRectangleCollider(pipeWidth, pipeHeight, "pipe", []string{"player"}),
		bottomCollider: *core.NewRectangleCollider(pipeWidth, pipeHeight, "pipe", []string{"player"}),
	}
}

func (pg PipeGate) OnAdd(scene *core.Scene) {
	pg.physicsSystem = &scene.PhysicsSystem
	if pg.physicsSystem != nil {
		pg.physicsSystem.Register(pg)
	}
}

func (pg PipeGate) OnRemove(scene *core.Scene) {
	if pg.physicsSystem != nil {
		pg.physicsSystem.Unregister(pg)
	}
}

func (pg PipeGate) Update(dt float32) {
	if !pg.Running {
		return
	}
	// Move the pipes to the left
	pg.x -= pg.speed * float32(dt)
	if pg.x < -float32(pg.topSprite.Texture.Width) {
		pg.Passed = true
	}
}

func (pg PipeGate) Draw() {
	// Draw top pipe (above the gap)
	topY := pg.gapY - float32(PipeGate_GapHeight/2)
	pg.topSprite.Draw(*core.NewTransform(pg.x, topY))

	// Draw bottom pipe (below the gap)
	bottomY := pg.gapY + float32(PipeGate_GapHeight/2)
	pg.bottomSprite.Draw(*core.NewTransform(pg.x, bottomY))
}

// GetCollider implements core.Collidable.
func (pg PipeGate) GetCollider() *core.Collider {
	return &pg.topCollider
}

// GetTransform implements core.Collidable.
func (pg PipeGate) GetTransform() *core.Transform {
	topY := pg.gapY - float32(PipeGate_GapHeight/2)
	return core.NewTransform(pg.x, topY)
}

// OnCollision handles collision events (pipes don't need to react to collisions)
func (pg PipeGate) OnCollision(other core.Collidable, collision core.CollisionInfo) {
	// Pipes don't need to react to collisions, the other entity will handle it
}
