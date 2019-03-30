//
// Display Table
//
package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"

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

func initCards() {

	tablePos := [][]int{
		{60, 60, 100, 60},
		{344, 20, 408, 20},
		{450, 20, 514, 20},
		{610, 20, 174, 20},
	}

	flopPos := [][]int{
		{270, 250}, // Flop
		{344, 250},
		{418, 250},
		{492, 250}, // Turn
		{566, 250}, // River
	}

	fmt.Println(tablePos)

	fmt.Println(flopPos)
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
	initCards() // init Card Display Structure

	if err := ebiten.Run(update, 1024, 768, 1, "AI POKER - You against eight AI Players!"); err != nil {
		log.Fatal(err)
	}
}
