// data.go contains all the Constants, variables, types used in AIpoker.
package main

import (
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/font"
)

// Bit Mask type for the "mode" word.
type Bits uint16

// player postion structure (betImages)
type playStruct struct {
	active bool
	post   int
	bet    int
	screen *ebiten.Image
}

var playPost = [9]playStruct{}

// Bit Mask Defitions for the "mode" word.
// BIT MASK "MUST" MATCH VARIABLE DMAP!!!!
const (
	cardDeal = 1 << iota
	cardDelt
	cardFlop
	cardTurn
	cardRiver
	betValue
	betInput
	betEnable
	betMade
	inputWait
	dealWait
	aiProcess
	foldMade
	checkMade
	callMade
	raiseMade
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

	dec = 60 // Display Error Count
)

// Primary Display Parameters.
const (
	screenWidth  = 1024
	screenHeight = 768

	dpi = 72
)

// Common Images
const cardTable = "images/table.png"

const cardBack = "images/playing-cards-back.png"
const blankBack = "images/valueBackground.png"

// All Global Variables used by AIpoker.
var displayError = 0         // Display Error Message Counter
var displayErrorMessage = "" // Error Message to Present

var mode Bits         // Update Mask
var dumpModes = false // ModeDump Enable

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

var counter = 0 //Blink Counter

var inText = ""
var ctrl = 0
var betAmount int

var users [9][2]int                       // Users Cards
var comCards = [5]int{21, 22, 23, 25, 27} // common cards

// Card postions for each player
// The Player Position Table is duplicated so that starting at at any point
// will allow nine positions to be dealt.
var players = []int{
	70, 440, 134, 440, 140, 210, 390, hide, display,
	0, 250, 64, 250, 1, 160, 260, hide, display,
	70, 80, 134, 80, 2, 220, 170, hide, display,
	378, 20, 314, 20, 3, 390, 140, hide, display,
	560, 20, 624, 20, 4, 630, 140, hide, display,
	800, 80, 864, 80, 5, 750, 140, hide, display,
	850, 250, 914, 250, 6, 780, 260, hide, display,
	800, 440, 864, 440, 8, 730, 390, hide, display,
	420, 460, 484, 460, 9, 420, 380, unhide, display,
	70, 440, 134, 440, 140, 210, 390, hide, display,
	0, 250, 64, 250, 1, 160, 260, hide, display,
	70, 80, 134, 80, 2, 220, 170, hide, display,
	378, 20, 314, 20, 3, 390, 140, hide, display,
	560, 20, 624, 20, 4, 630, 140, hide, display,
	800, 80, 864, 80, 5, 750, 140, hide, display,
	850, 250, 914, 250, 6, 780, 260, hide, display,
	800, 440, 864, 440, 8, 730, 390, hide, display,
	420, 460, 484, 460, 9, 420, 380, unhide, display,
}

// Player Chip Postions
var chipMap = []int{
	170, 390, //  7  102, 460,
	160, 315, //  1  32,  580,
	170, 176, //  2  102, 200,
	450, 145, //  3  380, 130,
	675, 146, //  4  592, 744,
	752, 186, //  5  832, 200,
	783, 310, //  6  882, 370,
	734, 440, //  8  832, 560,
	470, 382, //  9  452, 580,
	170, 390, //  7  102, 460,
	160, 315, //  1  32,  580,
	170, 176, //  2  102, 200,
	450, 145, //  3  380, 130,
	675, 146, //  4  592, 744,
	752, 186, //  5  832, 200,
	783, 310, //  6  882, 370,
	734, 440, //  8  832, 560,
	470, 382, //  9  452, 580,
}

// Chip Images
var chip = map[int]string{
	0: "images/red_chip.png",
	1: "images/stack.png",
	2: "images/table.png",
	3: "images/white_deal_chip.png",
}
var cardValue = map[int]string{
	0:  "A",
	1:  "K",
	2:  "Q",
	3:  "J",
	4:  "10",
	5:  "9",
	6:  "8",
	7:  "7",
	8:  "6",
	9:  "5",
	10: "4",
	11: "3",
	12: "2",
}
var deckInitial = map[int]string{
	0:  "images/ace_of_clubs.png",
	1:  "images/ace_of_diamonds.png",
	2:  "images/ace_of_hearts.png",
	3:  "images/ace_of_spades.png",
	4:  "images/king_of_clubs.png",
	5:  "images/king_of_diamonds.png",
	6:  "images/king_of_hearts.png",
	7:  "images/king_of_spades.png",
	8:  "images/queen_of_clubs.png",
	9:  "images/queen_of_diamonds.png",
	10: "images/queen_of_hearts.png",
	11: "images/queen_of_spades.png",
	12: "images/jack_of_clubs.png",
	13: "images/jack_of_diamonds.png",
	14: "images/jack_of_hearts.png",
	15: "images/jack_of_spades.png",
	16: "images/10_of_clubs.png",
	17: "images/10_of_diamonds.png",
	18: "images/10_of_hearts.png",
	19: "images/10_of_spades.png",
	20: "images/9_of_clubs.png",
	21: "images/9_of_diamonds.png",
	22: "images/9_of_hearts.png",
	23: "images/9_of_spades.png",
	24: "images/8_of_clubs.png",
	25: "images/8_of_diamonds.png",
	26: "images/8_of_hearts.png",
	27: "images/8_of_spades.png",
	28: "images/7_of_clubs.png",
	29: "images/7_of_diamonds.png",
	30: "images/7_of_hearts.png",
	31: "images/7_of_spades.png",
	32: "images/6_of_clubs.png",
	33: "images/6_of_diamonds.png",
	34: "images/6_of_hearts.png",
	35: "images/6_of_spades.png",
	36: "images/5_of_clubs.png",
	37: "images/5_of_diamonds.png",
	38: "images/5_of_hearts.png",
	39: "images/5_of_spades.png",
	40: "images/4_of_clubs.png",
	41: "images/4_of_diamonds.png",
	42: "images/4_of_hearts.png",
	43: "images/4_of_spades.png",
	44: "images/3_of_clubs.png",
	45: "images/3_of_diamonds.png",
	46: "images/3_of_hearts.png",
	47: "images/3_of_spades.png",
	48: "images/2_of_clubs.png",
	49: "images/2_of_diamonds.png",
	50: "images/2_of_hearts.png",
	51: "images/2_of_spades.png",
}

var deck = map[int]string{
	0:  "images/ace_of_clubs.png",
	1:  "images/ace_of_diamonds.png",
	2:  "images/ace_of_hearts.png",
	3:  "images/ace_of_spades.png",
	4:  "images/king_of_clubs.png",
	5:  "images/king_of_diamonds.png",
	6:  "images/king_of_hearts.png",
	7:  "images/king_of_spades.png",
	8:  "images/queen_of_clubs.png",
	9:  "images/queen_of_diamonds.png",
	10: "images/queen_of_hearts.png",
	11: "images/queen_of_spades.png",
	12: "images/jack_of_clubs.png",
	13: "images/jack_of_diamonds.png",
	14: "images/jack_of_hearts.png",
	15: "images/jack_of_spades.png",
	16: "images/10_of_clubs.png",
	17: "images/10_of_diamonds.png",
	18: "images/10_of_hearts.png",
	19: "images/10_of_spades.png",
	20: "images/9_of_clubs.png",
	21: "images/9_of_diamonds.png",
	22: "images/9_of_hearts.png",
	23: "images/9_of_spades.png",
	24: "images/8_of_clubs.png",
	25: "images/8_of_diamonds.png",
	26: "images/8_of_hearts.png",
	27: "images/8_of_spades.png",
	28: "images/7_of_clubs.png",
	29: "images/7_of_diamonds.png",
	30: "images/7_of_hearts.png",
	31: "images/7_of_spades.png",
	32: "images/6_of_clubs.png",
	33: "images/6_of_diamonds.png",
	34: "images/6_of_hearts.png",
	35: "images/6_of_spades.png",
	36: "images/5_of_clubs.png",
	37: "images/5_of_diamonds.png",
	38: "images/5_of_hearts.png",
	39: "images/5_of_spades.png",
	40: "images/4_of_clubs.png",
	41: "images/4_of_diamonds.png",
	42: "images/4_of_hearts.png",
	43: "images/4_of_spades.png",
	44: "images/3_of_clubs.png",
	45: "images/3_of_diamonds.png",
	46: "images/3_of_hearts.png",
	47: "images/3_of_spades.png",
	48: "images/2_of_clubs.png",
	49: "images/2_of_diamonds.png",
	50: "images/2_of_hearts.png",
	51: "images/2_of_spades.png",
}

var deckReverse = map[string]int{
	"images/ace_of_clubs.png":      0,
	"images/ace_of_diamonds.png":   1,
	"images/ace_of_hearts.png":     2,
	"images/ace_of_spades.png":     3,
	"images/king_of_clubs.png":     4,
	"images/king_of_diamonds.png":  5,
	"images/king_of_hearts.png":    6,
	"images/king_of_spades.png":    7,
	"images/queen_of_clubs.png":    8,
	"images/queen_of_diamonds.png": 9,
	"images/queen_of_hearts.png":   10,
	"images/queen_of_spades.png":   11,
	"images/jack_of_clubs.png":     12,
	"images/jack_of_diamonds.png":  13,
	"images/jack_of_hearts.png":    14,
	"images/jack_of_spades.png":    15,
	"images/10_of_clubs.png":       16,
	"images/10_of_diamonds.png":    17,
	"images/10_of_hearts.png":      18,
	"images/10_of_spades.png":      19,
	"images/9_of_clubs.png":        20,
	"images/9_of_diamonds.png":     21,
	"images/9_of_hearts.png":       22,
	"images/9_of_spades.png":       23,
	"images/8_of_clubs.png":        24,
	"images/8_of_diamonds.png":     25,
	"images/8_of_hearts.png":       26,
	"images/8_of_spades.png":       27,
	"images/7_of_clubs.png":        28,
	"images/7_of_diamonds.png":     29,
	"images/7_of_hearts.png":       30,
	"images/7_of_spades.png":       31,
	"images/6_of_clubs.png":        32,
	"images/6_of_diamonds.png":     33,
	"images/6_of_hearts.png":       34,
	"images/6_of_spades.png":       35,
	"images/5_of_clubs.png":        36,
	"images/5_of_diamonds.png":     37,
	"images/5_of_hearts.png":       38,
	"images/5_of_spades.png":       39,
	"images/4_of_clubs.png":        40,
	"images/4_of_diamonds.png":     41,
	"images/4_of_hearts.png":       42,
	"images/4_of_spades.png":       43,
	"images/3_of_clubs.png":        44,
	"images/3_of_diamonds.png":     45,
	"images/3_of_hearts.png":       46,
	"images/3_of_spades.png":       47,
	"images/2_of_clubs.png":        48,
	"images/2_of_diamonds.png":     49,
	"images/2_of_hearts.png":       50,
	"images/2_of_spades.png":       51,
}

// CardReverse give Card Value ( A = 0, K = 1, Q = 2 ...)
// DeckReverse above allows computing Suits as follows:
//    Suits are "deckReverse" modulo 4 - (0 = club, 1 = diamond, 2 = hearts, 3 = spades)

var cardReverse = map[string]int{
	"images/ace_of_clubs.png":      0,
	"images/ace_of_diamonds.png":   0,
	"images/ace_of_hearts.png":     0,
	"images/ace_of_spades.png":     0,
	"images/king_of_clubs.png":     1,
	"images/king_of_diamonds.png":  1,
	"images/king_of_hearts.png":    1,
	"images/king_of_spades.png":    1,
	"images/queen_of_clubs.png":    2,
	"images/queen_of_diamonds.png": 2,
	"images/queen_of_hearts.png":   2,
	"images/queen_of_spades.png":   2,
	"images/jack_of_clubs.png":     3,
	"images/jack_of_diamonds.png":  3,
	"images/jack_of_hearts.png":    3,
	"images/jack_of_spades.png":    3,
	"images/10_of_clubs.png":       4,
	"images/10_of_diamonds.png":    4,
	"images/10_of_hearts.png":      4,
	"images/10_of_spades.png":      4,
	"images/9_of_clubs.png":        5,
	"images/9_of_diamonds.png":     5,
	"images/9_of_hearts.png":       5,
	"images/9_of_spades.png":       5,
	"images/8_of_clubs.png":        6,
	"images/8_of_diamonds.png":     6,
	"images/8_of_hearts.png":       6,
	"images/8_of_spades.png":       6,
	"images/7_of_clubs.png":        7,
	"images/7_of_diamonds.png":     7,
	"images/7_of_hearts.png":       7,
	"images/7_of_spades.png":       7,
	"images/6_of_clubs.png":        8,
	"images/6_of_diamonds.png":     8,
	"images/6_of_hearts.png":       8,
	"images/6_of_spades.png":       8,
	"images/5_of_clubs.png":        9,
	"images/5_of_diamonds.png":     9,
	"images/5_of_hearts.png":       9,
	"images/5_of_spades.png":       9,
	"images/4_of_clubs.png":        10,
	"images/4_of_diamonds.png":     10,
	"images/4_of_hearts.png":       10,
	"images/4_of_spades.png":       10,
	"images/3_of_clubs.png":        11,
	"images/3_of_diamonds.png":     11,
	"images/3_of_hearts.png":       11,
	"images/3_of_spades.png":       11,
	"images/2_of_clubs.png":        12,
	"images/2_of_diamonds.png":     12,
	"images/2_of_hearts.png":       12,
	"images/2_of_spades.png":       12,
}

var betMap = []int{
	230, 440, 238, 456, // 7    32, 580,
	130, 360, 130, 375, // 1   102, 200,
	74, 180, 74, 196, // 2   410, 140,
	300, 143, 300, 158, // 3   592, 744,
	540, 146, 550, 164, // 4   832, 200,
	800, 190, 806, 206, // 5   882, 370,
	842, 352, 842, 370, // 6   832, 560,
	800, 400, 810, 420, // 8   452, 580,
	420, 430, 424, 450, // 9   102, 460,
	230, 440, 238, 456, // 7    32, 580,
	130, 360, 130, 375, // 1   102, 200,
	74, 180, 74, 196, // 2   410, 140,
	300, 143, 300, 158, // 3   592, 744,
	540, 146, 550, 164, // 4   832, 200,
	800, 190, 806, 206, // 5   882, 370,
	842, 352, 842, 370, // 6   832, 560,
	800, 400, 810, 420, // 8   452, 580,
	420, 430, 424, 450, // 9   102, 460,
}
var chartOne = [13][13]int{
	//A K  Q  J  10 9  8  7  6  5  4  3  2
	{9, 9, 8, 8, 8, 5, 3, 3, 3, 3, 3, 3, 3}, // A
	{9, 9, 3, 3, 8, 5, 7, 7, 7, 7, 3, 3, 3}, // K
	{9, 2, 9, 8, 8, 5, 5, 3, 3, 3, 3, 3, 3}, // Q
	{2, 2, 2, 9, 8, 5, 5, 3, 3, 3, 3, 3, 3}, // J
	{9, 2, 2, 2, 9, 9, 5, 5, 3, 3, 3, 3, 3}, // 10
	{1, 1, 1, 1, 1, 9, 9, 5, 5, 3, 3, 3, 3}, // 9
	{1, 1, 1, 1, 1, 1, 9, 9, 5, 5, 3, 3, 3}, // 8
	{1, 1, 1, 1, 1, 1, 3, 9, 9, 5, 5, 3, 3}, // 7
	{1, 1, 1, 1, 1, 1, 0, 3, 9, 5, 5, 5, 0}, // 6
	{0, 0, 0, 0, 0, 0, 0, 0, 3, 8, 5, 5, 5}, // 5
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 8, 5, 5}, // 4
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 8, 5}, // 3
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 9, 5}, // 2
}

// Bits to String Conversion Map
var dmap = map[Bits]string{
	cardDeal:  "cardDeal",
	cardDelt:  "cardDelt",
	cardFlop:  "cardFlop",
	cardTurn:  "cardTurn",
	cardRiver: "cardRiver",
	betValue:  "betValue",
	betInput:  "betInput",
	betMade:   "betMade",
	inputWait: "inputWait",
	dealWait:  "dealWait",
	aiProcess: "aiProcess",
	foldMade:  "foldMade",
	checkMade: "checkMade",
	callMade:  "callMade",
	raiseMade: "raiseMade",
}
