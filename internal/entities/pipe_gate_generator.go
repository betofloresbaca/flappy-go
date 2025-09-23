package entities

import (
	"flappy-go/internal/core"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	PipeGateGenerator_HSpacing         = 150
	PipeGateGenerator_XStart           = 400
	PipeGateGenerator_PreloadZoneWidth = 100
)

type PipeGateGenerator struct {
	*core.BaseEntity
	*core.BaseUpdater
	speed        float32
	lastPipeGate *PipeGate
	pipeIndex    int
	Running      bool
}

func NewPipeGateGenerator(parent *core.Scene, speed float32) *PipeGateGenerator {
	return &PipeGateGenerator{
		BaseEntity:  core.NewBaseEntity(parent, "pipe_gate_generator", ""),
		BaseUpdater: core.NewBaseUpdater(),
		speed:       speed,
	}
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

func (pgg *PipeGateGenerator) addPipe(x float32) {
	newPipe := NewPipeGate(pgg.BaseEntity.Parent(), pgg.pipeIndex, x, pgg.speed)
	newPipe.Running = true
	pgg.lastPipeGate = newPipe
	pgg.Parent().Add(newPipe)
	pgg.pipeIndex++
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
