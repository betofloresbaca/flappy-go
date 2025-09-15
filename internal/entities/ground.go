package entities

import (
	"simple-go-game/internal/assets"
	"simple-go-game/internal/components"
	"simple-go-game/internal/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const groundY = 440

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
		BaseDrawable: core.NewBaseDrawable(-100),
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
		g.sprite.Draw(core.Transform{
			Position: rl.Vector2{X: float32(i)*float32(g.sprite.Texture.Width) + g.offset, Y: groundY},
			Scale:    rl.Vector2{X: 1, Y: 1},
			Rotation: 0,
		})
	}
}
