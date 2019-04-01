//
// Display Table
//
package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var table *ebiten.Image
var card *ebiten.Image
var card2 *ebiten.Image
var cardb *ebiten.Image
var cardf *ebiten.Image
var err error

const (
	cardDeal  = iota // 0
	cardFlop         // 1
	cardTurn         // 2
	cardRiver        // 3
	cardWait         // Wait for User Action
)

var updateMode int // 0 = Deal, 1 = Flop, 2 = Turn, 3 = River
var oneShot = 0

const (
	unhide    = iota // Display Card Value
	hide             // Hide Card Value (Show Card Back)
	display          // Display card at position
	undisplay        // Display Nothing
)

var tablePos = [][]int{
	{420, 450, 484, 440}, // bottom-middle(live Player)
	{0, 250, 64, 250},    // left-middle
	{79, 80, 134, 80},    // left-Top
	{314, 20, 378, 20},   // top-left
	{560, 20, 624, 20},   // top-right
	{800, 80, 864, 80},   // right-top
	{850, 250, 914, 250}, // right-middle
	{70, 440, 134, 440},  // left-bottom
	{800, 440, 864, 440}, // right-bottom
	{270, 250, -1, -1},   // Flop
	{344, 250, -1, -1},
	{418, 250, -1. - 1},
	{492, 250, -1, -1}, // Turn
	{566, 250, -1, -1}, // River
}

var cardTable = "images/table.png"

var cardBack = "images/playing-cards-back.png"

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

func update(screen *ebiten.Image) error {

	if ebiten.IsDrawingSkipped() {
		// When the game is running slowly, the rendering result
		// will not be adopted.
		return nil
	}

	initTable(screen)
	deal(updateMode, screen) // Deal Cards

	return nil
}

func instruct() {
	fmt.Println("\n                        A I  P O K E R")
	fmt.Println("\nAIpoker is the initial concept design for an AI Based Poker Simulation")
	// Fill the Screen with the white color
	//	screen.Fill(color.White)

	fmt.Println("Currently it only displays a empty table.")
	fmt.Println("\nHowever, the 'images' directory contains all the cards and other necessary images.")
	fmt.Println("Still a huge amount of work to do.\n")
	fmt.Println("Only one 'bug' at the moment: You must put the cursor in the image box for")
	fmt.Println("the entire image to be completely displayed.")
	fmt.Println("\nThis 'Feature' also shows up in their example games. May take Big Think!")
}

// Initialize Poker Table Display
func initTable(screen *ebiten.Image) {

	// Fill the Screen with the white color
	screen.Fill(color.White)

	//  Insert Table Image
	if table == nil {
		// Fill the Screen with the white color
		screen.Fill(color.White)
		// Create an Table image
		table, _, err = ebitenutil.NewImageFromFile(cardTable, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Get Options Structure
	opts := &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(15, 70)
	// Draw the table image to the screen with an empty option
	screen.DrawImage(table, opts)
}

func shuffle() {
	s1 := time.Now().UnixNano()
	rand.Seed(s1)
	deck[52], deck[rand.Intn(52)] = deck[rand.Intn(52)], deck[52]

	rand.Shuffle(52, func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

}

func cardDisplay(x, y float64, cardValue, h, d int, screen *ebiten.Image) {

	// Create Card image
	//	fmt.Println("cval = ", cardValue, "  Card = ", deck[cardValue])
	//	fmt.Println("x = ", x, " y = ", y, " h = ", h, " d = ", d)
	if d == undisplay {
		//		fmt.Println("display = ", d)
		return
	}
	if h == hide {
		//		fmt.Println("hide = ", h)
		card, _, err = ebitenutil.NewImageFromFile(cardBack, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		//		fmt.Println("unhide = ", h)
		card, _, err = ebitenutil.NewImageFromFile(deck[cardValue], ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Display Image
	opts := &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(x, y)
	// Draw the card image to the screen with an empty option
	screen.DrawImage(card, opts)
}

func deal(cardMode int, screen *ebiten.Image) {

	if cardMode < 0 && cardMode > 3 { // Insure Valid Deal Mode
		log.Fatal(err)
	}

	if cardMode == cardDeal {

		cardDisplay(0, 250, 1, hide, display, screen)
		cardDisplay(70, 80, 2, hide, display, screen)
		cardDisplay(378, 20, 3, hide, display, screen)
		cardDisplay(560, 20, 4, hide, display, screen)
		cardDisplay(800, 80, 5, hide, display, screen)
		cardDisplay(850, 250, 6, hide, display, screen)
		cardDisplay(70, 440, 7, hide, display, screen)
		cardDisplay(800, 440, 8, hide, display, screen)
		cardDisplay(420, 460, 9, unhide, display, screen)

		cardDisplay(64, 250, 10, hide, display, screen)
		cardDisplay(134, 80, 11, hide, display, screen)
		cardDisplay(314, 20, 12, hide, display, screen)
		cardDisplay(624, 20, 13, hide, display, screen)
		cardDisplay(864, 80, 14, hide, display, screen)
		cardDisplay(914, 250, 15, hide, display, screen)
		cardDisplay(134, 440, 16, hide, display, screen)
		cardDisplay(864, 440, 17, hide, display, screen)
		cardDisplay(484, 460, 18, unhide, display, screen)
		return
	}

	if cardMode == cardFlop { // Burn Card = 20
		cardDisplay(319, 250, 21, hide, undisplay, screen)
		cardDisplay(393, 250, 22, hide, undisplay, screen)
		cardDisplay(467, 250, 23, hide, undisplay, screen)
		return
	}

	if cardMode == cardTurn { // Burn Card = 24
		cardDisplay(541, 250, 25, hide, undisplay, screen)
		return
	}

	if cardMode == cardRiver { // Burn Card = 26
		cardDisplay(615, 250, 27, hide, undisplay, screen)
		return
	}
	return
}

func main() {
	instruct()            // Display current notes on project progress
	shuffle()             // Shuffle Deck
	updateMode = cardDeal // Set Mode to Deal Cards

	//
	// Run Loop
	//
	if err := ebiten.Run(update, 1024, 768, 1, "AI POKER - You against eight AI Players!"); err != nil {
		log.Fatal(err)
	}
}
