package entities

import (
	"simple-go-game/internal/core"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	PipeGateGenerator_HSpacing         = 150
	PipeGateGenerator_XStart           = 400
	PipeGateGenerator_PreloadZoneWidth = 100
)

type PipeGateGenerator struct {
	*core.BaseEntity
	speed        float32
	lastPipeGate *PipeGate
	Running      bool
}

func NewPipeGateGenerator(parent *core.Scene, speed float32) *PipeGateGenerator {
	return &PipeGateGenerator{
		BaseEntity: core.NewBaseEntity(parent),
		speed:      speed,
	}
}

func (pgg *PipeGateGenerator) addPipe(x float32) {
	newPipe := NewPipeGate(pgg.BaseEntity.Parent, x, pgg.speed)
	newPipe.Running = true
	pgg.lastPipeGate = newPipe
	pgg.BaseEntity.Parent.Add(newPipe)
}

func (pgg *PipeGateGenerator) Update(dt float32) {
	if !pgg.Running {
		return
	}
	// Generate new pipes if needed
	if pgg.lastPipeGate == nil {
		pgg.addPipe(PipeGateGenerator_XStart)
	}
	nextXStart := pgg.getNextXStart()
	for nextXStart < float32(raylib.GetScreenWidth()+PipeGateGenerator_PreloadZoneWidth) {
		pgg.addPipe(nextXStart)
		nextXStart = pgg.getNextXStart()
	}

}

func (pgg *PipeGateGenerator) getNextXStart() float32 {
	if pgg.lastPipeGate == nil {
		return PipeGateGenerator_XStart
	}
	lastPipeX := pgg.lastPipeGate.GetX()
	return lastPipeX +
		PipeGateGenerator_HSpacing +
		float32(pgg.lastPipeGate.topSprite.Texture.Width)
}
