//
// AIpoker.go provides the primary executive functions. The main
// function shuffles the virtual deck, sets to "mode" to "cardDeal"
// which will deal the virtual cards.
//
package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
)

// uppate is called from the "run" loop. This function controls the operation
// of the program. It is called every time the "run" loop completes a cycle.
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
		createTable(screen)

		dealCards(mode, screen)     // Deal Cards
		mode = set(mode, betEnable) // Enable Message Boxes

		return nil
	}
	return nil
}

//
// main is the Primary Executive Function.
//   It performs the initial shuffle, sets the cardDeal bit in the "mode"
//   variable and then calls the "RUN" Loop -- Returns only on error or termination.
//   All further action takes place in the specified "update" Function.
//
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
