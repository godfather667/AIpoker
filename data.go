//
// Data Values
//
package main

import (
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/font"
)

type Bits uint8 // Bit Mask for mode

const (
	cardDeal = 1 << iota // Update Mask Encoding
	cardFlop
	cardTurn
	cardRiver
	betValue
	betInput
	betEnable
	inputWait
)

const (
	mPlus = iota // mplus fonts
	mSmall
	mTiny
	aPlus // Arcade fonts
	aSmall
	aTiny
)
const (
	unhide    = iota // Display Card Value
	hide             // Hide Card Value (Show Card Back)
	display          // Display card at position
	undisplay        // Display Nothing
)

const (
	hasCR  = 1 // Carriage Retustackrn Detected
	isNew  = 2 // Represents and New Value
	isNull = 3 // Nothing has Changed this Update

	dec = 60 // Display Error Counnt
)

const (
	screenWidth  = 1024
	screenHeight = 768

	dpi = 72
)

var displayError = 0         // Display Error Message Counter
var displayErrorMessage = "" // Error Message to Present

var mode Bits // Update Mask

var result string // Value input value

var mplusNormalFont font.Face // Font Variables
var smallNormalFont font.Face
var tinyNormalFont font.Face
var arcadeFont font.Face
var smallArcadeFont font.Face
var tinyArcadeFont font.Face

var table *ebiten.Image // Table Image
var card *ebiten.Image  // Card Image

var err error // Error

var cardTable = "images/table.png"

var cardBack = "images/playing-cards-back.png"

var counter = 0 // Blink Counter

var chip = map[int]string{
	1: "images/dealer_chip.png",
	2: "images/red_chip.png",
	3: "images/stack.png",
}

var deck = map[int]string{
	1:  "images/10_of_clubs.png",
	2:  "images/10_of_diamonds.png",
	3:  "images/10_of_hearts.png",
	4:  "images/10_of_spades.png",
	5:  "images/2_of_clubs.png",
	6:  "images/2_of_diamonds.png",
	7:  "images/2_of_hearts.png",
	8:  "images/2_of_spades.png",
	9:  "images/3_of_clubs.png",
	10: "images/3_of_diamonds.png",
	11: "images/3_of_hearts.png",
	12: "images/3_of_spades.png",
	13: "images/4_of_clubs.png",
	14: "images/4_of_diamonds.png",
	15: "images/4_of_hearts.png",
	16: "images/4_of_spades.png",
	17: "images/5_of_clubs.png",
	18: "images/5_of_diamonds.png",
	19: "images/5_of_hearts.png",
	20: "images/5_of_spades.png",
	21: "images/6_of_clubs.png",
	22: "images/6_of_diamonds.png",
	23: "images/6_of_hearts.png",
	24: "images/6_of_spades.png",
	25: "images/7_of_clubs.png",
	26: "images/7_of_diamonds.png",
	27: "images/7_of_hearts.png",
	28: "images/7_of_spades.png",
	29: "images/8_of_clubs.png",
	30: "images/8_of_diamonds.png",
	31: "images/8_of_hearts.png",

	32: "images/8_of_spades.png",
	33: "images/9_of_clubs.png",
	34: "images/9_of_diamonds.png",
	35: "images/9_of_hearts.png",
	36: "images/9_of_spades.png",
	37: "images/ace_of_clubs.png",
	38: "images/ace_of_diamonds.png",
	39: "images/ace_of_hearts.png",
	40: "images/ace_of_spades.png",
	41: "images/jack_of_clubs.png",
	42: "images/jack_of_diamonds.png",
	43: "images/jack_of_hearts.png",
	44: "images/jack_of_spades.png",
	45: "images/king_of_clubs.png",
	46: "images/king_of_diamonds.png",
	47: "images/king_of_hearts.png",
	48: "images/king_of_spades.png",
	49: "images/queen_of_clubs.png",
	50: "images/queen_of_diamonds.png",
	51: "images/queen_of_hearts.png",
	52: "images/queen_of_spades.png",
}
