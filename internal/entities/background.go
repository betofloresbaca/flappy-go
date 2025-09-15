package entities

import (
	"simple-go-game/internal/assets"
	"simple-go-game/internal/components"
	"simple-go-game/internal/core"
)

const (
	Background_ZIndex      = -1000
	Background_SpritePivot = components.PivotUpLeft
	Background_ImageNumber = 3
)

type Background struct {
	*core.BaseEntity
	*core.BaseDrawable
	sprite components.Sprite
}

func NewBackground(style string) *Background {
	return &Background{
		BaseEntity:   core.NewBaseEntity(),
		BaseDrawable: core.NewBaseDrawable(Background_ZIndex),
		sprite:       *components.NewSprite(assets.BackgroundSprites[style], Background_SpritePivot),
	}
}

func (b *Background) Draw() {
	for i := range Background_ImageNumber {
		b.sprite.Draw(
			*core.NewTransform(
				float32(i)*float32(b.sprite.Texture.Width),
				0,
			),
		)
	}
}
