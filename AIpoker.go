//
// Main Poker Functions
//
package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

//
// Update called each Run Cycle
//
func update(screen *ebiten.Image) error {

	// Logical Operations to Setup Rendering

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && has(mode, betEnable) {
		x, y := ebiten.CursorPosition()
		if x > 320 && x < 470 && y > 600 && y < 690 {
			if !has(mode, betValue) {
				result = ""
			}
			mode = set(mode, betValue) // Process various Bet Options
			mode = set(mode, betInput)
		}
	}
	if has(mode, betInput) {
		inText, inputMode := inputUpdate(screen)
		if inputMode != isNull {
			result = inText //Update Output Image
		}

		if inputMode == hasCR {
			// Finialize Result Number and Clear Input Flags
			result = inText
			mode = clear(mode, betValue)
			mode = clear(mode, betInput)
		}
	}

	if has(mode, cardDeal) {

		if ebiten.IsDrawingSkipped() {
			// When the game is running slowly, the rendering data
			// will not be adopted.
			return nil
		}
		initTable(screen)

		deal(mode, screen)          // Deal Cards
		mode = set(mode, betEnable) // Enable Message Boxes

		return nil
	}
	return nil
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

	if has(mode, betEnable) {
		messageSquare(150, 90, 20, 600, color.NRGBA{0xff, 0xff, 0x00, 0xff}, screen)
		charDisplay(aSmall, "CHECK", 40, 650, screen)

		messageSquare(150, 90, 170, 600, color.NRGBA{0xff, 0x00, 0x00, 0xff}, screen)
		charDisplay(aSmall, "FOLD", 210, 650, screen)

		messageSquare(150, 90, 320, 600, color.NRGBA{0x00, 0xff, 0x00, 0xff}, screen)
		charDisplay(aSmall, " BET", 360, 650, screen)

		t := ""
		if has(mode, betValue) {
			// Blink the cursor.
			if counter%60 < 30 {
				t = "_"
			}
			if displayError > 0 {
				messageError(screen)
				displayError--
			} else {
				clearError(screen)
			}
			counter++
			charDisplay(aTiny, "Value:"+result+t, 500, 650, screen)
			mode = set(mode, betInput)
		}
	}
}

func deal(mode Bits, screen *ebiten.Image) {

	if has(mode, cardDeal) {
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

	if (mode & cardFlop) != 0 { // Burn Card = 20
		cardDisplay(319, 250, 21, hide, undisplay, screen)
		cardDisplay(393, 250, 22, hide, undisplay, screen)
		cardDisplay(467, 250, 23, hide, undisplay, screen)
		return
	}

	if (mode & cardTurn) != 0 { // Burn Card = 24
		cardDisplay(541, 250, 25, hide, undisplay, screen)
		return
	}

	if (mode & cardRiver) != 0 { // Burn Card = 26
		cardDisplay(615, 250, 27, hide, undisplay, screen)
		return
	}

	return
}

func main() {

	shuffle()       // Shuffle Deck
	mode = cardDeal // Clear Deal Mode

	//
	// Run Loop
	//
	if err := ebiten.Run(update, 1024, 768, 1, "AI POKER - You against eight AI Players!"); err != nil {
		log.Fatal(err)
	}
}
