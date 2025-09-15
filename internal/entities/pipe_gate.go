package entities

import (
	"simple-go-game/internal/assets"
	"simple-go-game/internal/components"
	"simple-go-game/internal/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PipeGate struct {
	*core.BaseEntity
	*core.BaseDrawable
	topSprite    components.Sprite
	bottomSprite components.Sprite
	x            float32
	yOffset      float32
	gapHeight    float32
	speed        float32
	Passed       bool
	Running      bool
}

func NewPipeGate(x, yOffset, gapHeight, speed float32) *PipeGate {
	topSprite := components.NewSprite(assets.PipeSprites["green"], components.PivotDownLeft)
	topSprite.FlipV = true
	bottomSprite := components.NewSprite(assets.PipeSprites["green"], components.PivotUpLeft)
	return &PipeGate{
		BaseEntity:   core.NewBaseEntity(),
		BaseDrawable: core.NewBaseDrawable(-200),
		topSprite:    *topSprite,
		bottomSprite: *bottomSprite,
		x:            x,
		yOffset:      yOffset,
		gapHeight:    gapHeight,
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
	topY := p.yOffset
	p.topSprite.Draw(core.Transform{
		Position: rl.Vector2{X: p.x, Y: topY},
		Scale:    rl.Vector2{X: 1, Y: 1},
		Rotation: 0,
	})

	// Draw bottom pipe (below the gap)
	bottomY := p.yOffset + p.gapHeight
	p.bottomSprite.Draw(core.Transform{
		Position: rl.Vector2{X: p.x, Y: bottomY},
		Scale:    rl.Vector2{X: 1, Y: 1},
		Rotation: 0,
	})
}
