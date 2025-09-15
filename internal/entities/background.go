package entities

import (
	"simple-go-game/internal/assets"
	"simple-go-game/internal/components"
	"simple-go-game/internal/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Background struct {
	*core.BaseEntity
	*core.BaseDrawable
	sprite components.Sprite
}

func NewBackground() *Background {
	return &Background{
		BaseEntity:   core.NewBaseEntity(),
		BaseDrawable: core.NewBaseDrawable(-1000), // Z-index -1 for background layer
		sprite:       *components.NewSprite(assets.BackgroundSprites["day"], components.PivotUpLeft),
	}
}

func (b *Background) Update(dt float64) {
	b.BaseEntity.Update(dt)
}

func (b *Background) Draw() {
	for i := range 3 {
		b.sprite.Draw(core.Transform{
			Position: rl.Vector2{X: float32(i) * float32(b.sprite.Texture.Width), Y: 0},
			Scale:    rl.Vector2{X: 1, Y: 1},
			Rotation: 0,
		})
	}
}
