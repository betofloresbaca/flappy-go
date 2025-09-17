package entities

import (
	"simple-go-game/internal/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	PipeGateGenerator_HSpacing         = 150
	PipeGateGenerator_XStart           = 400
	PipeGateGenerator_PreloadZoneWidth = 100
)

type PipeGateGenerator struct {
	*core.BaseEntity
	speed     float32
	pipeGates []*PipeGate
	Running   bool
	physics   *core.PhysicsSystem
}

func NewPipeGateGenerator(speed float32) *PipeGateGenerator {
	return &PipeGateGenerator{
		BaseEntity: core.NewBaseEntity(),
		speed:      speed,
		pipeGates:  make([]*PipeGate, 0),
	}
}

func (pg *PipeGateGenerator) addPipe(x float32) {
	newPipe := NewPipeGate(x, pg.speed)
	newPipe.Running = true
	pg.pipeGates = append(pg.pipeGates, newPipe)
}

func (pg *PipeGateGenerator) Update(dt float32) {
	if !pg.Running {
		return
	}
	// Update existing pipes
	activePipes := make([]*PipeGate, 0, len(pg.pipeGates))
	for i := range pg.pipeGates {
		pg.pipeGates[i].Update(dt)
		if !pg.pipeGates[i].Passed {
			activePipes = append(activePipes, pg.pipeGates[i])
		} else {
			// Unregister passed pipes from physics system
			if pg.physics != nil {
				pg.physics.Unregister(pg.pipeGates[i])
			}
		}
	}
	pg.pipeGates = activePipes

	// Generate new pipes if needed
	if len(pg.pipeGates) == 0 {
		pg.addPipe(PipeGateGenerator_XStart)
	}
	nextXStart := pg.getNextXStart()
	for nextXStart < float32(rl.GetScreenWidth()+PipeGateGenerator_PreloadZoneWidth) {
		pg.addPipe(nextXStart)
		nextXStart = pg.getNextXStart()
	}

}

func (pg *PipeGateGenerator) getNextXStart() float32 {
	lastPipeX := pg.pipeGates[len(pg.pipeGates)-1].x
	return lastPipeX +
		PipeGateGenerator_HSpacing +
		float32(pg.pipeGates[len(pg.pipeGates)-1].topSprite.Texture.Width)
}
