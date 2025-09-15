package assets

import _ "embed"

// Numbers

//go:embed sprites/0.png
var num0 []byte

//go:embed sprites/1.png
var num1 []byte

//go:embed sprites/2.png
var num2 []byte

//go:embed sprites/3.png
var num3 []byte

//go:embed sprites/4.png
var num4 []byte

//go:embed sprites/5.png
var num5 []byte

//go:embed sprites/6.png
var num6 []byte

//go:embed sprites/7.png
var num7 []byte

//go:embed sprites/8.png
var num8 []byte

//go:embed sprites/9.png
var num9 []byte

var NumberSprites = [][]byte{
	num0,
	num1,
	num2,
	num3,
	num4,
	num5,
	num6,
	num7,
	num8,
	num9,
}

//Background Images

//go:embed sprites/background-day.png
var backgroundDay []byte

//go:embed sprites/background-night.png
var backgroundNight []byte

var BackgroundSprites = map[string][]byte{
	"day":   backgroundDay,
	"night": backgroundNight,
}

// Bird Sprites

//go:embed sprites/bluebird-upflap.png
var bluebirdUpflap []byte

//go:embed sprites/bluebird-downflap.png
var bluebirdDownflap []byte

//go:embed sprites/bluebird-midflap.png
var bluebirdMidflap []byte

//go:embed sprites/redbird-upflap.png
var redbirdUpflap []byte

//go:embed sprites/redbird-downflap.png
var redbirdDownflap []byte

//go:embed sprites/redbird-midflap.png
var redbirdMidflap []byte

//go:embed sprites/yellowbird-upflap.png
var yellowbirdUpflap []byte

//go:embed sprites/yellowbird-downflap.png
var yellowbirdDownflap []byte

//go:embed sprites/yellowbird-midflap.png
var yellowbirdMidflap []byte

var BirdSprites = map[string][][]byte{
	"blue": {
		bluebirdUpflap,
		bluebirdMidflap,
		bluebirdDownflap,
	},
	"red": {
		redbirdUpflap,
		redbirdMidflap,
		redbirdDownflap,
	},
	"yellow": {
		yellowbirdUpflap,
		yellowbirdMidflap,
		yellowbirdDownflap,
	},
}

// Pipe Sprites

//go:embed sprites/pipe-green.png
var pipeGreen []byte

//go:embed sprites/pipe-red.png
var pipeRed []byte

var PipeSprites = map[string][]byte{
	"green": pipeGreen,
	"red":   pipeRed,
}

// Ground Sprite

//go:embed sprites/ground.png
var GroundSprite []byte

// Miscellaneous Sprites

//go:embed sprites/message.png
var MessageSprite []byte

//go:embed sprites/gameover.png
var GameOverSprite []byte

// Sound Assets

//go:embed sounds/die.wav
var soundDie []byte

//go:embed sounds/hit.wav
var soundHit []byte

//go:embed sounds/point.wav
var soundPoint []byte

//go:embed sounds/swoosh.wav
var soundSwoosh []byte

//go:embed sounds/wing.wav
var soundWing []byte

var SoundAssets = map[string][]byte{
	"die":    soundDie,
	"hit":    soundHit,
	"point":  soundPoint,
	"swoosh": soundSwoosh,
	"wing":   soundWing,
}
