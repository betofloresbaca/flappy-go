package entities

import (
	"flappy-go/internal/assets"
	"flappy-go/internal/core"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type InstructionsMessage struct {
	*core.BaseEntity
	*core.BaseDrawer
	sprite    *core.Sprite
	transform core.Transform
}

func NewInstructionsMessage(
	parent *core.Scene,
) *InstructionsMessage {
	return &InstructionsMessage{
		BaseEntity: core.NewBaseEntity(parent, "instructions_message"),
		BaseDrawer: core.NewBaseDrawer(2000),
		sprite:     core.NewSprite(assets.MessageImage, core.PivotCenter),
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

func (sm *InstructionsMessage) Draw() {
	sm.sprite.Draw(sm.transform)
}
