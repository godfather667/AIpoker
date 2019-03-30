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
var idx int

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
var deck = map[int]string{
	1:  "10_of_clubs.png",
	2:  "10_of_diamonds.png",
	3:  "10_of_hearts.png",
	4:  "10_of_spades.png",
	5:  "2_of_clubs.png",
	6:  "2_of_diamonds.png",
	7:  "2_of_hearts.png",
	8:  "2_of_spades.png",
	9:  "3_of_clubs.png",
	10: "3_of_diamonds.png",
	11: "3_of_hearts.png",
	12: "3_of_spades.png",
	13: "4_of_clubs.png",
	14: "4_of_diamonds.png",
	15: "4_of_hearts.png",
	16: "4_of_spades.png",
	17: "5_of_clubs.png",
	18: "5_of_diamonds.png",
	19: "5_of_hearts.png",
	20: "5_of_spades.png",
	21: "6_of_clubs.png",
	22: "6_of_diamonds.png",
	23: "6_of_hearts.png",
	24: "6_of_spades.png",
	25: "7_of_clubs.png",
	26: "7_of_diamonds.png",
	27: "7_of_hearts.png",
	28: "7_of_spades.png",
	29: "8_of_clubs.png",
	30: "8_of_diamonds.png",
	31: "8_of_hearts.png",
	32: "8_of_spades.png",
	33: "9_of_clubs.png",
	34: "9_of_diamonds.png",
	35: "9_of_hearts.png",
	36: "9_of_spades.png",
	37: "ace_of_clubs.png",
	38: "ace_of_diamonds.png",
	39: "ace_of_hearts.png",
	40: "ace_of_spades.png",
	41: "jack_of_clubs.png",
	42: "jack_of_diamonds.png",
	43: "jack_of_hearts.png",
	44: "jack_of_spades.png",
	45: "king_of_clubs.png",
	46: "king_of_diamonds.png",
	47: "king_of_hearts.png",
	48: "king_of_spades.png",
	49: "queen_of_clubs.png",
	50: "queen_of_diamonds.png",
	51: "queen_of_hearts.png",
	52: "queen_of_spades.png",
}

func shuffle() {
	s1 := time.Now().UnixNano()
	rand.Seed(s1)
	deck[52], deck[rand.Intn(52)] = deck[rand.Intn(52)], deck[52]

	rand.Shuffle(52, func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	fmt.Println(deck)
}

func update(screen *ebiten.Image) error {

	// Draw the square image
	if ebiten.IsDrawingSkipped() {
		return nil //  Insert Card Image
	}

	// Fill the Screen with the white color
	screen.Fill(color.White)

	//  Insert Table Image
	if table == nil {
		// Create an Table image
		table, _, err = ebitenutil.NewImageFromFile("images/table.png", ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
	}
	opts := &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(15, 70)
	// Draw the table image to the screen with an empty option
	screen.DrawImage(table, opts)

	// Inserting Cards Pairs and Back Pairs showing sample placement
	// This code is only for testing
	// The display code will be totally different in future versions.
	//
	//  Insert Card Image
	if card == nil {
		// Create an Table image
		card, _, err = ebitenutil.NewImageFromFile("images/king_of_clubs.png", ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
	}
	//  Insert Card2 Image
	if card2 == nil {
		card2, _, err = ebitenutil.NewImageFromFile("images/king_of_hearts.png", ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
	}

	// ---------------------
	//  left - middle
	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(0, 250)
	// Draw the card image to the screen with an empty option
	screen.DrawImage(card, opts)

	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(64, 250)
	// Draw the square image to the screen with an empty option
	screen.DrawImage(card2, opts)

	//---------------------
	// left - Top
	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(70, 80)
	// Draw the card image to the screen with an empty option
	screen.DrawImage(card, opts)

	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(134, 80)
	// Draw the square image to the screen with an empty option
	screen.DrawImage(card2, opts)

	//---------------------
	// top - left
	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(314, 20)
	// Draw the card image to the screen with an empty option
	screen.DrawImage(card, opts)

	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(378, 20)
	// Draw the square image to the screen with an empty option
	screen.DrawImage(card2, opts)

	//---------------------
	// Top - Right
	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(560, 20)
	// Draw the card image to the screen with an empty option
	screen.DrawImage(card, opts)

	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(624, 20)
	// Draw the square image to the screen with an empty option
	screen.DrawImage(card2, opts)

	//---------------------
	// right - Top
	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(800, 80)
	// Draw the card image to the screen with an empty option
	screen.DrawImage(card, opts)

	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(864, 80)
	// Draw the square image to the screen with an empty option
	screen.DrawImage(card2, opts)

	//---------------------
	// Right - Middle
	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(850, 250)
	// Draw the card image to the screen with an empty option
	screen.DrawImage(card, opts)

	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(914, 250)
	// Draw the square image to the screen with an empty option
	screen.DrawImage(card2, opts)

	//---------------------
	// Left - Bottom
	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(70, 440)
	// Draw the card image to the screen with an empty option
	screen.DrawImage(card, opts)

	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(134, 440)
	// Draw the square image to the screen with an empty option
	screen.DrawImage(card2, opts)

	//---------------------
	// Right - Bottom
	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(800, 440)
	// Draw the card image to the screen with an empty option
	screen.DrawImage(card, opts)

	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(864, 440)
	// Draw the square image to the screen with an empty option
	screen.DrawImage(card2, opts)

	//---------------------
	// bottom - Middle
	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(420, 460)
	// Draw the card image to the screen with an empty option
	screen.DrawImage(card, opts)

	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(484, 460)
	// Draw the square image to the screen with an empty option
	screen.DrawImage(card2, opts)

	// Insert Flop Cards

	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			cardf, _, err = ebitenutil.NewImageFromFile("images/ace_of_clubs.png", ebiten.FilterDefault)
			if err != nil {
				log.Fatal(err)
			}
			idx = 344
		case 1:
			cardf, _, err = ebitenutil.NewImageFromFile("images/ace_of_hearts.png", ebiten.FilterDefault)
			if err != nil {
				log.Fatal(err)
			}
			idx = 418
		case 2:
			cardf, _, err = ebitenutil.NewImageFromFile("images/ace_of_spades.png", ebiten.FilterDefault)
			if err != nil {
				log.Fatal(err)
			}
			idx = 492
		case 3:
			cardf, _, err = ebitenutil.NewImageFromFile("images/jack_of_diamonds.png", ebiten.FilterDefault)
			if err != nil {
				log.Fatal(err)
			}
			idx = 566
		case 4:
			cardf, _, err = ebitenutil.NewImageFromFile("images/jack_of_clubs.png", ebiten.FilterDefault)
			if err != nil {
				log.Fatal(err)
			}
			idx = 640
		}
		opts = &ebiten.DrawImageOptions{}
		// Add the Translate effect to the option struct.
		opts.GeoM.Translate(float64(idx)-25, 250)
		// Draw the square image to the screen with an empty option
		screen.DrawImage(cardf, opts)
	}

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
func initTable() {

	//  Insert Table Image
	if table == nil {
		// Create an Table image
		table, _, err = ebitenutil.NewImageFromFile("images/table.png", ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	instruct()  // Display current notes on project progress
	shuffle()   // Shuffle Deck
	initTable() // Initialize table display
	//
	if err := ebiten.Run(update, 1024, 768, 1, "AI POKER - You against eight AI Players!"); err != nil {
		log.Fatal(err)
	}
}
