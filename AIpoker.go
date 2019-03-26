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

var img *ebiten.Image
var table *ebiten.Image
var card *ebiten.Image

var err error

/* func init() {
	var err error
	table, _, err = ebitenutil.NewImageFromFile("images/table.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}
*/

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

	// Draw the square image to the screen with an empty option
	screen.DrawImage(table, opts)

	// Draw Initial Image
	//	screen.DrawImage(img, nil)
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
