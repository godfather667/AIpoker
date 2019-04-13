// image.go contains the primary display functions.
package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
)

// init function for setting up image Fonts.
func init() {

	// Seup a new fonts for text display.
	// Initialize Normal Fonts.

	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	smallNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    18,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	tinyNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Arcade Fonts.

	tt, err = truetype.Parse(fonts.ArcadeN_ttf)
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	smallArcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    18,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	tinyArcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

// imageDisplay inserts individual images on the specified screen.
func imageDisplay(x, y float64, cardValue, h, d int, screen *ebiten.Image) {
	// Create Card image
	if d == undisplay {
		return
	}
	if h == hide {
		card, _, err = ebitenutil.NewImageFromFile(cardBack, ebiten.FilterDefault)
		if err != nil {
			fmt.Println("CardBack Display Failed!")
			log.Fatal(cardBack, err)
		}
	} else {
		card, _, err = ebitenutil.NewImageFromFile(deck[cardValue], ebiten.FilterDefault)
		if err != nil {
			fmt.Println("Card ", cardValue, "  Filename: ", deck[cardValue],
				"  Length: ", len(deck[cardValue]), "  Display Failed!")
			log.Fatal(deck[cardValue], err)
		}
	}
	// Display Image
	opts := &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(x, y)
	// Draw the card image to the screen with an empty option
	screen.DrawImage(card, opts)
}

// imageDisplay inserts individual images on the specified screen.
func chipDisplay(x, y float64, chipName string, h, d int, screen *ebiten.Image) {
	// Create Card image
	if d == undisplay {
		return
	}
	chipImage, _, err := ebitenutil.NewImageFromFile(chipName, ebiten.FilterDefault)
	if err != nil {
		fmt.Println("Chip = ", chipName, "Display Failed!")
		log.Fatal(chipName, err)
	}
	// Display Image
	opts := &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(x, y)
	// Draw the card image to the screen with an empty option
	screen.DrawImage(chipImage, opts)
}

// messageDisplay displays a message in one of six(6) fonts/sizes
func messageDisplay(font int, msg string, x, y int, screen *ebiten.Image) {
	// Add Text
	switch font {
	case 0:
		text.Draw(screen, msg, mplusNormalFont, x, y, color.Black)
	case 1:
		text.Draw(screen, msg, smallNormalFont, x, y, color.Black)
	case 2:
		text.Draw(screen, msg, tinyNormalFont, x, y, color.Black)
	case 3:
		text.Draw(screen, msg, arcadeFont, x, y, color.Black)
	case 4:
		text.Draw(screen, msg, smallArcadeFont, x, y, color.Black)
	case 5:
		text.Draw(screen, msg, tinyArcadeFont, x, y, color.Black)
	default:
		fmt.Println("Bad Font Description!")
	}
}

//
// Error Message Functions
//   The error message is display at the bottom of the screen in red letter
//
func messageError(screen *ebiten.Image) {
	text.Draw(screen, displayErrorMessage, smallArcadeFont, 250, 760, color.NRGBA{0xff, 0x00, 0x00, 0xff}) // Color Red
}

// Set Error Message loads message buffer with message and sets download counter dec constant
func setError(msg string, screen *ebiten.Image) {
	displayError = dec // Display Error Delay Value
	displayErrorMessage = msg
}

// Clear Error Message and download counter
func clearError(screen *ebiten.Image) {
	displayError = 0
	displayErrorMessage = ""
}

func msgBlank(x, y int, screen *ebiten.Image) {
	blank, _, err := ebitenutil.NewImageFromFile(blankBack, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(blankBack, err)
	}
	// Display Image
	opts := &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(float64(x), float64(y))
	// Draw the card image to the screen with an empty option
	screen.DrawImage(blank, opts)
}

//
// Display the individual message squares (ie Large buttons)
//
func messageSquare(sx, sy int, px, py float64, colorCode color.NRGBA, screen *ebiten.Image) {

	square, _ := ebiten.NewImage(sx, sy, ebiten.FilterNearest)
	square.Fill(colorCode)

	opts := &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(px, py)
	// Draw the card image to the screen with an empty option
	screen.DrawImage(square, opts)
}

func betImage(post, bet int, screen *ebiten.Image) {
	id := post * 4
	msgBlank(betMap[id], betMap[id+1], screen)

	msg := fmt.Sprintf("$%v", bet)
	//	fmt.Println("Bet Msg: ", msg, "  Post: ", post)
	messageDisplay(aTiny, msg, betMap[id+2], betMap[id+3], screen)
}
