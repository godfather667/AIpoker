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
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Fill the Screen with the white color
	screen.Fill(color.White)

	//  Insert Table Image
	if table == nil {
		// Create an Table image
		//		table, _ = ebiten.NewImage(950, 475, ebiten.FilterNearest)
		table, _, err = ebitenutil.NewImageFromFile("images/table.png", ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
	}

	opts := &ebiten.DrawImageOptions{}

	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(0, 37)

	// Draw the table image to the screen with an empty option
	screen.DrawImage(table, opts)

	// Inserting Cards Pairs and Back Pairs showing sample placement
	// This code is only for testing
	// The display code will be totally different in future versions.
	//ls
	//  Insert Card Image
	if card == nil {
		// Create an Table image
		//		table, _ = ebiten.NewImage(950, 475, ebiten.FilterNearest)
		card, _, err = ebitenutil.NewImageFromFile("images/king_of_clubs.png", ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
	}
	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(100, 37)
	// Draw the card image to the screen with an empty option
	screen.DrawImage(card, opts)

	//  Insert Card2 Image
	if card2 == nil {
		// Create an Table image
		//		table, _ = ebiten.NewImage(950, 475, ebiten.FilterNearest)
		card2, _, err = ebitenutil.NewImageFromFile("images/king_of_hearts.png", ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
	}
	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(165, 37)
	// Draw the square image to the screen with an empty option
	screen.DrawImage(card2, opts)

	//  Insert Cardb Image
	if cardb == nil {
		cardb, _, err = ebitenutil.NewImageFromFile("images/playing-cards-back.png", ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
	}
	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(10, 220)
	// Draw the square image to the screen with an empty option
	screen.DrawImage(cardb, opts)

	// Show second card back
	opts = &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(74, 220)
	// Draw the square image to the screen with an empty option
	screen.DrawImage(cardb, opts)

	// Insert Flop Cards

	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			cardf, _, err = ebitenutil.NewImageFromFile("images/ace_of_clubs.png", ebiten.FilterDefault)
			if err != nil {
				log.Fatal(err)
			}
			idx = 250
		case 1:
			cardf, _, err = ebitenutil.NewImageFromFile("images/ace_of_hearts.png", ebiten.FilterDefault)
			if err != nil {
				log.Fatal(err)
			}
			idx = 324
		case 2:
			cardf, _, err = ebitenutil.NewImageFromFile("images/ace_of_spades.png", ebiten.FilterDefault)
			if err != nil {
				log.Fatal(err)
			}
			idx = 398
		case 3:
			cardf, _, err = ebitenutil.NewImageFromFile("images/jack_of_diamonds.png", ebiten.FilterDefault)
			if err != nil {
				log.Fatal(err)
			}
			idx = 472
		case 4:
			cardf, _, err = ebitenutil.NewImageFromFile("images/jack_of_clubs.png", ebiten.FilterDefault)
			if err != nil {
				log.Fatal(err)
			}
			idx = 546
		}
		opts = &ebiten.DrawImageOptions{}
		// Add the Translate effect to the option struct.
		opts.GeoM.Translate(float64(idx)+20, 230)
		// Draw the square image to the screen with an empty option
		screen.DrawImage(cardf, opts)
	}

	return nil
}

func instruct() {
	fmt.Println("\n                        A I  P O K E R")
	fmt.Println("\nAIpoker is the initial concept design for an AI Based Poker Simulation")
	fmt.Println("Currently it only displays a empty table.")
	fmt.Println("\nHowever, the 'images' directory contains all the cards and other necessary images.")
	fmt.Println("Still a huge amount of work to do.\n")
	fmt.Println("Only one 'bug' at the moment: You must put the cursor in the image box for")
	fmt.Println("the entire image to be completely displayed.")
	fmt.Println("\nThis 'Feature' also shows up in their example games. May take Big Think!")
}

func main() {
	instruct() // Display current notes on project progress
	if err := ebiten.Run(update, 1024, 768, 1, "AI POKER - You against eight AI Players!"); err != nil {
		log.Fatal(err)
	}
}
