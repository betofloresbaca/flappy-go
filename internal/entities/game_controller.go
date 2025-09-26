package entities

import (
	"flappy-go/internal/core"
	"flappy-go/internal/ui"
	"flappy-go/internal/utils"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	GameController_Name = "game_controller"
)

// Make an enum indicating the status Start, Playing,GameOver
type GameStatus int

const (
	Initial GameStatus = iota
	Start
	Playing
	GameOver
)

type GameController struct {
	*core.BaseEntity
	*core.BaseUpdater
	status          GameStatus
	createGameBoard func() *core.Scene
	gameBoard       *core.Scene
	player          *Player
	scoreDisplay    *utils.Lazy[*ui.ScoreDisplay]
	gameOverMessage *utils.Lazy[*ui.GameOverMessage]
	startMessage    *utils.Lazy[*ui.StartMessage]
}

func NewGameController(parent *core.Scene, createGameBoard func() *core.Scene) *GameController {
	lc := &GameController{
		BaseEntity:  core.NewBaseEntity(parent, GameController_Name, []string{}),
		BaseUpdater: core.NewBaseUpdater(),
		status:      Initial,
	}
	lc.createGameBoard = createGameBoard
	lc.gameOverMessage = utils.NewLazy(lc.findGameOverMessage)
	lc.startMessage = utils.NewLazy(lc.findStartMessage)
	lc.scoreDisplay = utils.NewLazy(lc.findScoreDisplay)
	return lc
}

func (gc *GameController) Update(dt float32) {
	// If space or up key is pressed, change status to Playing
	switch gc.status {
	case Initial:
		gc.transitToStart()
	case Start:
		if isInputPressed() {
			gc.transitToPlaying()
		}
	case Playing:
		if gc.player.IsDead() {
			gc.transitToGameOver()
		}
	case GameOver:
		if isInputPressed() {
			gc.transitToStart()
		}
	}
}

func isInputPressed() bool {
	return raylib.IsKeyPressed(raylib.KeySpace) ||
		raylib.IsMouseButtonPressed(raylib.MouseLeftButton)
}

func (gc *GameController) transitToStart() {
	gc.status = Start
	// If there is an existing game board, remove it from the tree to cleanup physics/world
	if gc.gameBoard != nil {
		gc.Root().Remove(gc.gameBoard)
		gc.gameBoard = nil
		gc.player = nil
	}
	gc.gameBoard = gc.createGameBoard()
	gc.Root().Add(gc.gameBoard)
	gc.gameBoard.Pause()
	gc.player = gc.findPlayer()
	gc.startMessage.Value().Show()
	gc.scoreDisplay.Value().Hide()
	gc.gameOverMessage.Value().Hide()
}

func (gc *GameController) transitToPlaying() {
	gc.status = Playing
	gc.gameBoard.Resume()
	gc.startMessage.Value().Hide()
	gc.scoreDisplay.Value().Reset()
	gc.scoreDisplay.Value().Show()
	gc.gameOverMessage.Value().Hide()
}

func (gc *GameController) transitToGameOver() {
	gc.status = GameOver
	gc.gameBoard.Pause()
	// Show the game over message
	gc.gameOverMessage.Value().Show()
}

func (gc *GameController) findScoreDisplay() *ui.ScoreDisplay {
	return gc.Root().ChildByName("ui").(*core.Scene).
		ChildByName(ui.ScoreDisplay_Name).(*ui.ScoreDisplay)
}

func (gc *GameController) findGameOverMessage() *ui.GameOverMessage {
	return gc.Root().ChildByName("ui").(*core.Scene).
		ChildByName(ui.GameOverMessage_Name).(*ui.GameOverMessage)
}

func (gc *GameController) findStartMessage() *ui.StartMessage {
	return gc.Root().ChildByName("ui").(*core.Scene).
		ChildByName(ui.StartMessage_Name).(*ui.StartMessage)
}
func (gc *GameController) findPlayer() *Player {
	return gc.Root().ChildByName("game_board").(*core.Scene).
		ChildByName(Player_Name).(*Player)
}
