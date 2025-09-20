package entities

import (
	"flappy-go/internal/assets"
	"flappy-go/internal/core"
)

const (
	Background_ZIndex       = -1000
	Background_SpritePivot  = core.PivotUpLeft
	Background_ImagesNumber = 3
)

type Background struct {
	*core.BaseEntity
	*core.BaseDrawer
	sprite core.Sprite
}

func NewBackground(parent *core.Scene, style string) *Background {
	return &Background{
		BaseEntity: core.NewBaseEntity(parent, "background"),
		BaseDrawer: core.NewBaseDrawer(Background_ZIndex),
		sprite:     *core.NewSprite(assets.BackgroundImages[style], Background_SpritePivot),
	}
}

func (b *Background) Draw() {
	for i := range Background_ImagesNumber {
		b.sprite.Draw(
			*core.NewTransform(
				float32(i)*float32(b.sprite.Texture.Width),
				0,
			),
		)
	}
}
