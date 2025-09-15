package assets

import _ "embed"

// Numbers

//go:embed images/0.png
var num0 []byte

//go:embed images/1.png
var num1 []byte

//go:embed images/2.png
var num2 []byte

//go:embed images/3.png
var num3 []byte

//go:embed images/4.png
var num4 []byte

//go:embed images/5.png
var num5 []byte

//go:embed images/6.png
var num6 []byte

//go:embed images/7.png
var num7 []byte

//go:embed images/8.png
var num8 []byte

//go:embed images/9.png
var num9 []byte

var NumberImages = [][]byte{
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

//go:embed images/background-day.png
var backgroundDay []byte

//go:embed images/background-night.png
var backgroundNight []byte

var BackgroundImages = map[string][]byte{
	"day":   backgroundDay,
	"night": backgroundNight,
}

// Bird Sprites

//go:embed images/bluebird-upflap.png
var bluebirdUpflap []byte

//go:embed images/bluebird-downflap.png
var bluebirdDownflap []byte

//go:embed images/bluebird-midflap.png
var bluebirdMidflap []byte

//go:embed images/redbird-upflap.png
var redbirdUpflap []byte

//go:embed images/redbird-downflap.png
var redbirdDownflap []byte

//go:embed images/redbird-midflap.png
var redbirdMidflap []byte

//go:embed images/yellowbird-upflap.png
var yellowbirdUpflap []byte

//go:embed images/yellowbird-downflap.png
var yellowbirdDownflap []byte

//go:embed images/yellowbird-midflap.png
var yellowbirdMidflap []byte

var BirdImages = map[string][][]byte{
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

//go:embed images/pipe-green.png
var pipeGreen []byte

//go:embed images/pipe-red.png
var pipeRed []byte

var PipeSprites = map[string][]byte{
	"green": pipeGreen,
	"red":   pipeRed,
}

// Ground Sprite

//go:embed images/ground.png
var GroundImage []byte

// Miscellaneous Sprites

//go:embed images/message.png
var MessageImage []byte

//go:embed images/gameover.png
var GameOverImage []byte

// Sound Assets

//go:embed sounds/die.wav
var dieSound []byte

//go:embed sounds/hit.wav
var hitSound []byte

//go:embed sounds/point.wav
var pointSound []byte

//go:embed sounds/swoosh.wav
var swooshSound []byte

//go:embed sounds/wing.wav
var wingSound []byte

var Sounds = map[string][]byte{
	"die":    dieSound,
	"hit":    hitSound,
	"point":  pointSound,
	"swoosh": swooshSound,
	"wing":   wingSound,
}
