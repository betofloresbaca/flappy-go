package entities

import (
	"simple-go-game/internal/assets"
	"simple-go-game/internal/components"
	"simple-go-game/internal/core"
)

const (
	Ground_ZIndex = -100
	Ground_Y      = 440
)

type Ground struct {
	*core.BaseEntity
	*core.BaseDrawable
	sprite  components.Sprite
	speed   float32
	offset  float32
	Running bool
}

func NewGround(speed float32) *Ground {
	return &Ground{
		BaseEntity:   core.NewBaseEntity(),
		BaseDrawable: core.NewBaseDrawable(Ground_ZIndex),
		sprite:       *components.NewSprite(assets.GroundSprite, components.PivotUpLeft),
		speed:        speed,
	}
}

func (g *Ground) Update(dt float32) {
	if !g.Running {
		return
	}
	// Move the ground to the left
	g.offset -= g.speed * dt
	if g.offset <= -float32(g.sprite.Texture.Width) {
		g.offset += float32(g.sprite.Texture.Width)
	}
}

func (g *Ground) Draw() {
	for i := range 4 {
		g.sprite.Draw(
			*core.NewTransform(float32(i)*float32(g.sprite.Texture.Width)+g.offset, Ground_Y),
		)
	}
}
