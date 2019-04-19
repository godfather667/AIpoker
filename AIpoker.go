//
// AIpoker.go provides the primary executive functions. The main
// function shuffles the virtual deck, sets to "mode" to "cardDeal"
// which will deal the virtual cards.
//
package main

import (
	_ "image/png"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten"
)

// uppate is called from the "run" loop. This function controls the operation
// of the program. It is called every time the "run" loop completes a cycle.
func update(screen *ebiten.Image) error {
	dumpModes = true // Dump Mode Settings

	createTable(screen)

	if has(mode, cardDeal) {
		dealCards(mode, screen) // Deal Cards
		modeDump(mode)
	}

	// Logical Operations to Setup Rendering

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && has(mode, waitDeal) {
		x, y := ebiten.CursorPosition()
		if x > 10 && x < 160 && y > 600 && y < 690 {
			mode = set(mode, cardDeal)   // Process various Bet Options
			mode = clear(mode, waitDeal) // Deal and remove Deal Button
			modeDump(mode)
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) &&
		has(mode, betEnable) && !has(mode, betMade) {
		x, y := ebiten.CursorPosition()
		if x > 450 && x < 600 && y > 600 && y < 690 {
			if !has(mode, betValue) {
				result = ""
			}
			mode = set(mode, betValue) // Process various Bet Options
			mode = set(mode, betInput)
			mode = clear(mode, betEnable)
			modeDump(mode)
		}
	}
	if has(mode, betInput) {
		inText, inputMode := inputUpdate(screen)
		if inputMode != isNull {
			result = inText //Update Output Image
			modeDump(mode)
		}

		if inputMode == hasCR {
			// Finialize Result Number and Clear Input Flags
			result = inText
			if betAmount, err = strconv.Atoi(result); err != nil {
				setError("Numeric Error - Bad Number", screen)
				inText = ""
				return nil
			}

			mode = clear(mode, betValue)
			mode = clear(mode, betInput)
			mode = clear(mode, betEnable)
			mode = set(mode, betMade)
			mode = set(mode, aiProcess)
			modeDump(mode)
		}
	}

	// After User Action - Run the AI Processing
	if has(mode, aiProcess) {
		aiExec(betAmount, dealPost, mode, screen)
		modeDump(mode)
	}

	if has(mode, cardDeal) {

		if ebiten.IsDrawingSkipped() {
			// When the game is running slowly, the rendering data
			// will not be adopted.
			return nil
		}
		mode = set(mode, betEnable) // Enable Message Boxes
		modeDump(mode)

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
	mode = waitDeal // Wait for user to select Deal!
	dealPost = 0    // Initialize Deal Position
	//
	// Run Loop
	//
	if err := ebiten.Run(update, 1024, 768, 1, "AI POKER - You against eight AI Players!"); err != nil {
		log.Fatal(err)
	}
}
