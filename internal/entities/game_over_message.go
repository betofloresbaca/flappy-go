package entities

import (
	"flappy-go/internal/assets"
	"flappy-go/internal/core"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type GameOverMessage struct {
	*core.BaseEntity
	*core.BaseDrawer
	sprite    *core.Sprite
	transform core.Transform
}

func NewGameOverMessage(
	parent *core.Scene,
) *GameOverMessage {
	return &GameOverMessage{
		BaseEntity: core.NewBaseEntity(parent, "game_over_message"),
		BaseDrawer: core.NewBaseDrawer(2000),
		sprite:     core.NewSprite(assets.GameOverImage, core.PivotCenter),
		transform: core.Transform{
			Position: raylib.Vector2{
				X: float32(raylib.GetScreenWidth()) / 2,
				Y: float32(raylib.GetScreenHeight()) / 2,
			},
			Scale:    raylib.Vector2{X: 2, Y: 2},
			Rotation: 0,
		},
	}
}

func (sm *GameOverMessage) Draw() {
	sm.sprite.Draw(sm.transform)
}
