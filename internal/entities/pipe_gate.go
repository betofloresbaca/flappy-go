package entities

import (
	"simple-go-game/internal/assets"
	"simple-go-game/internal/core"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	PipeGate_GapYMin   = 100
	PipeGate_GapYMax   = 350
	PipeGate_GapHeight = 100
)

type PipeGate struct {
	*core.BaseEntity
	*core.BaseDrawable
	topSprite    core.Sprite
	bottomSprite core.Sprite
	x            float32
	gapY         float32
	speed        float32
	Passed       bool
	Running      bool
}

func NewPipeGate(x, speed float32) *PipeGate {
	topSprite := core.NewSprite(assets.PipeSprites["green"], core.PivotDownLeft)
	topSprite.FlipV = true
	bottomSprite := core.NewSprite(assets.PipeSprites["green"], core.PivotUpLeft)

	return &PipeGate{
		BaseEntity:   core.NewBaseEntity(),
		BaseDrawable: core.NewBaseDrawable(0),
		topSprite:    *topSprite,
		bottomSprite: *bottomSprite,
		x:            x,
		gapY:         float32(raylib.GetRandomValue(PipeGate_GapYMin, PipeGate_GapYMax)),
		speed:        speed,
	}
}

func (p *PipeGate) Update(dt float32) {
	if !p.Running {
		return
	}
	// Move the pipes to the left
	p.x -= p.speed * float32(dt)
	if p.x < -float32(p.topSprite.Texture.Width) {
		p.Passed = true
	}
}

func (p *PipeGate) Draw() {
	// Draw top pipe (above the gap)
	topY := p.gapY - float32(PipeGate_GapHeight/2)
	p.topSprite.Draw(*core.NewTransform(p.x, topY))

	// Draw bottom pipe (below the gap)
	bottomY := p.gapY + float32(PipeGate_GapHeight/2)
	p.bottomSprite.Draw(*core.NewTransform(p.x, bottomY))
}
