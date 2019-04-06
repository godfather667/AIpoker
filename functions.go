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

func init() { // Initialize Normal Fonts

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

// Bit Wise Functions
//
func set(b, flag Bits) Bits { return b | flag }

func clear(b, flag Bits) Bits { return b &^ flag }

func toggle(b, flag Bits) Bits { return b ^ flag }

func has(b, flag Bits) bool { return b&flag != 0 }

func shuffle() {
	deck[52], deck[rand.Intn(52)] = deck[rand.Intn(52)], deck[52]

	rand.Shuffle(52, func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}

func cardDisplay(x, y float64, cardValue, h, d int, screen *ebiten.Image) {

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

func messageError(screen *ebiten.Image) {
	text.Draw(screen, displayErrorMessage, smallArcadeFont, 250, 760, color.NRGBA{0xff, 0x00, 0x00, 0xff}) // Color Red
}
func setError(msg string, screen *ebiten.Image) {
	displayError = dec // Display Error Delay Value
	displayErrorMessage = msg
}

func clearError(screen *ebiten.Image) {
	displayError = 0
	displayErrorMessage = ""
}

func messageSquare(sx, sy int, px, py float64, colorCode color.NRGBA, screen *ebiten.Image) {

	square, _ := ebiten.NewImage(sx, sy, ebiten.FilterNearest)
	square.Fill(colorCode)

	opts := &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(px, py)
	// Draw the card image to the screen with an empty option
	screen.DrawImage(square, opts)
}
