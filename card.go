// card.go supplies functions that manipulate virtual cards.
package main

import (
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// init function for card manipulation that insures a good Random Number.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// createTable sets up basic table images.
func createTable(screen *ebiten.Image) {

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

	// If the user has a bet/fold/check decision - Setup up the Buttons.

	if has(mode, betEnable) {
		messageSquare(150, 90, 20, 600, color.NRGBA{0xff, 0xff, 0x00, 0xff}, screen)
		messageDisplay(aSmall, "CHECK", 40, 650, screen)

		messageSquare(150, 90, 170, 600, color.NRGBA{0xff, 0x00, 0x00, 0xff}, screen)
		messageDisplay(aSmall, "FOLD", 210, 650, screen)

		messageSquare(150, 90, 320, 600, color.NRGBA{0x00, 0xff, 0x00, 0xff}, screen)
		messageDisplay(aSmall, " BET", 360, 650, screen)

		t := ""
		if has(mode, betValue) {
			// Blink the cursor.
			if counter%60 < 30 {
				t = "_"
			}
			// Error Message Control
			if displayError > 0 {
				messageError(screen)
				displayError--
			} else {
				clearError(screen)
			}
			counter++
			messageDisplay(aTiny, "Value:"+result+t, 500, 650, screen)
			mode = set(mode, betInput)
		}
	}
}

// dealCards place cards images on the displayed table.
//   The actual cards dealt depend on the mode bits
func dealCards(mode Bits, screen *ebiten.Image) {

	if has(mode, cardDeal) {
		imageDisplay(0, 250, 1, hide, display, screen)
		imageDisplay(70, 80, 2, hide, display, screen)
		imageDisplay(378, 20, 3, hide, display, screen)
		imageDisplay(560, 20, 4, hide, display, screen)
		imageDisplay(800, 80, 5, hide, display, screen)
		imageDisplay(850, 250, 6, hide, display, screen)
		imageDisplay(70, 440, 7, hide, display, screen)
		imageDisplay(800, 440, 8, hide, display, screen)
		imageDisplay(420, 460, 9, unhide, display, screen)

		imageDisplay(64, 250, 10, hide, display, screen)
		imageDisplay(134, 80, 11, hide, display, screen)
		imageDisplay(314, 20, 12, hide, display, screen)
		imageDisplay(624, 20, 13, hide, display, screen)
		imageDisplay(864, 80, 14, hide, display, screen)
		imageDisplay(914, 250, 15, hide, display, screen)
		imageDisplay(134, 440, 16, hide, display, screen)
		imageDisplay(864, 440, 17, hide, display, screen)
		imageDisplay(484, 460, 18, unhide, display, screen)
		return
	}

	if (mode & cardFlop) != 0 { // Burn Card = 20
		imageDisplay(319, 250, 21, unhide, display, screen)
		imageDisplay(393, 250, 22, unhide, display, screen)
		imageDisplay(467, 250, 23, unhide, display, screen)
		return
	}

	if (mode & cardTurn) != 0 { // Burn Card = 24
		imageDisplay(541, 250, 25, unhide, display, screen)
		return
	}

	if (mode & cardRiver) != 0 { // Burn Card = 26
		imageDisplay(615, 250, 27, unhide, display, screen)
		return
	}

	return
}

// Mode Bit Functions - The "mode" value is used to control Deal Operationns,
//   Set, Clear, Toggle Mode Bits modify the "mode" word.
//   Has tests the state a particular bit in the "mode" word.
func set(b, flag Bits) Bits { return b | flag }

func clear(b, flag Bits) Bits { return b &^ flag }

func toggle(b, flag Bits) Bits { return b ^ flag }

func has(b, flag Bits) bool { return b&flag != 0 }

// shuffle randomly rearranges the cards in the virtual deck.
//   Due to a tendency of cards at the first and last locations
//   to stay thru a number of shuffles.  The cards at the locations
//   first(1) and last(52) are swapped with random middle positions.
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
