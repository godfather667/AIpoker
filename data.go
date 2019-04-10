// data.go contains all the Constants, variables, types used in AIpoker.
package main

import (
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/font"
)

// Bit Mask type for the "mode" word.
type Bits uint16

// Bit Mask Defitions for the "mode" word.
const (
	cardDeal = 1 << iota
	cardFlop
	cardTurn
	cardRiver
	betValue
	betInput
	betEnable
	inputWait
	waitDeal
	aiProcess
	betMade
	foldMade
	checkMade
)

// Define the font selector values.
const (
	mPlus = iota // mplus fonts
	mSmall
	mTiny
	aPlus // Arcade fonts
	aSmall
	aTiny
)

// The card display control bits.
const (
	unhide    = iota // Display Card Value
	hide             // Hide Card Value (Show Card Back)
	display          // Display card at position
	undisplay        // Display Nothing
)

// Input control flags for the text_input functions.
const (
	hasCR  = 1 // Carriage Return Detected
	isNew  = 2 // Represents and New Value
	isNull = 3 // Nothing has Changed this Update

	dec = 60 // Display Error Counnt
)

// Primary Display Parameters.
const (
	screenWidth  = 1024
	screenHeight = 768

	dpi = 72
)

// All Global Variables used bu AIpoker.
var displayError = 0         // Display Error Message Counter
var displayErrorMessage = "" // Error Message to Present

var mode Bits // Update Mask

var dealPost = 0 // Deal Position

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

var (
	inText    = ""
	ctrl      = 0
	betAmount int
)

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
	53: "images/red_chip.png",
	54: "images/stack.png",
	55: "images/table.png",
	56: "images/white_deal_chip.png",
}

// Card postions for each player
// The Player Position Table is duplicated so that starting at at any point
// will allow nine positions to be dealt.
var players = []int{
	70, 440, 134, 440, 7, 210, 390, hide, display,
	0, 250, 64, 250, 1, 160, 260, hide, display,
	70, 80, 134, 80, 2, 210, 160, hide, display,
	378, 20, 314, 20, 3, 360, 160, hide, display,
	560, 20, 624, 20, 4, 590, 160, hide, display,
	800, 80, 864, 80, 5, 750, 180, hide, display,
	850, 250, 914, 250, 6, 750, 270, hide, display,
	800, 440, 864, 440, 8, 730, 390, hide, display,
	420, 460, 484, 460, 9, 460, 400, unhide, display,
	70, 440, 134, 440, 7, 210, 420, hide, display,
	0, 250, 64, 250, 1, 210, 420, hide, display,
	70, 80, 134, 80, 2, 210, 420, hide, display,
	378, 20, 314, 20, 3, 210, 420, hide, display,
	560, 20, 624, 20, 4, 210, 420, hide, display,
	800, 80, 864, 80, 5, 210, 420, hide, display,
	850, 250, 914, 250, 6, 210, 420, hide, display,
	800, 440, 864, 440, 8, 210, 420, hide, display,
	420, 460, 484, 460, 9, 210, 420, unhide, display,
}

/*
var players = []int{
	70, 440, 134, 440, 7, 16, // AI Player 1
	0, 250, 64, 250, 1, 10,
	70, 80, 134, 80, 2, 11,
	378, 20, 314, 20, 3, 12,
	560, 20, 624, 20, 4, 13,
	800, 80, 864, 80, 5, 14,
	850, 250, 914, 250, 6, 15,
	800, 440, 864, 440, 8, 17, // AI Player 8
	420, 460, 484, 460, 9, 18, // Human Player
	70, 440, 134, 440, 7, 16, // AI Player 1
	0, 250, 64, 250, 1, 10,
	70, 80, 134, 80, 2, 11,
	378, 20, 314, 20, 3, 12,
	560, 20, 624, 20, 4, 13,
	800, 80, 864, 80, 5, 14,
	850, 250, 914, 250, 6, 15,
	800, 440, 864, 440, 8, 17, // AI Player 8
	420, 460, 484, 460, 9, 18, // Human Player
}
*/

// Bits to String Conversion Map
var dmap = map[Bits]string{
	cardDeal:  "cardDeal",
	cardFlop:  "cardFlop",
	cardTurn:  "cardTurn",
	cardRiver: "cardRive",
	betValue:  "betValue",
	betInput:  "betInput",
	inputWait: "inputWait",
	waitDeal:  "waitDeal",
	aiProcess: "aiProcess",
	betMade:   "betMade",
	foldMade:  "foldMade",
	checkMade: "checkMade",
}
