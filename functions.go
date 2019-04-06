//
// Utility Poker Functions
//
package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
)

//
// Initialize Normal Fonts
//
func init() {

	// Seup a new fonts for text display

	// Initialize Normal Fonts

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

	// Initialize Arcade Fonts

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

	// Insure Good Random Number - Initialize Seed with Nano-second time value!
	rand.Seed(time.Now().UnixNano())
}

//
// Bit Wise Functions - The "mode" value is used to control Deal Operationsns.
//   Set, Clear, Toggle Mode Bits
//   Toggle Bit is also available
//
func set(b, flag Bits) Bits { return b | flag }

func clear(b, flag Bits) Bits { return b &^ flag }

func toggle(b, flag Bits) Bits { return b ^ flag }

func has(b, flag Bits) bool { return b&flag != 0 }

//
// Shuffle Cards Function
//   Due to a tendency of cards at the first/end locations
//   to stay thru a number of shuffles - The cards at the locations
//   first(1) and end(52) are swapped with random middle position.
//
func shuffle() {
	// Three Shuffles for each new hand!
	deck[52], deck[rand.Intn(52)] = deck[rand.Intn(52)], deck[52]
	deck[1], deck[rand.Intn(52)] = deck[rand.Intn(52)], deck[1]

	rand.Shuffle(52, func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	// Shuffle Two
	deck[52], deck[rand.Intn(52)] = deck[rand.Intn(52)], deck[52]
	deck[1], deck[rand.Intn(52)] = deck[rand.Intn(52)], deck[1]

	rand.Shuffle(52, func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	// Shuffle Three
	deck[52], deck[rand.Intn(52)] = deck[rand.Intn(52)], deck[52]
	deck[1], deck[rand.Intn(52)] = deck[rand.Intn(52)], deck[1]

	rand.Shuffle(52, func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}

//
// Display Individual Images
//
func imageDisplay(x, y float64, cardValue, h, d int, screen *ebiten.Image) {

	// Create Card image
	if d == undisplay {
		return
	}
	if h == hide {
		card, _, err = ebitenutil.NewImageFromFile(cardBack, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
	} else {
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

//
// Display Character Message in one of six(6) fonts/sizes
//
func charDisplay(font int, msg string, x, y int, screen *ebiten.Image) {
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
