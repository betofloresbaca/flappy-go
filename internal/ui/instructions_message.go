package ui

import (
	"flappy-go/internal/assets"
	"flappy-go/internal/core"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	InstructionsMessage_Name = "instructions_message"
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
		BaseEntity: core.NewBaseEntity(parent, InstructionsMessage_Name, []string{}),
		BaseDrawer: core.NewBaseDrawer(0),
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
