// ai.go is the main executive function AI Processes and Functions
package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
)

/* -- Mode Switches List
if has(mode, cardDelt) {
	mode = clear(mode, cardDelt)
	mode = set(mode, cardFlop)
} else if has(mode, cardFlop) {
	mode = clear(mode, cardFlop)
	mode = set(mode, cardTurn)
} else if has(mode, cardTurn) {
	mode = clear(mode, cardTurn)
	mode = set(mode, cardRiver)
} else if has(mode, cardRiver) {
	mode = clear(mode, cardRiver)
	mode = set(mode, dealWait)
}
*/

func aiExec(users [9][2]int, betAmount, dealPost int, mode Bits, screen *ebiten.Image) {
	//di := dealPost
	for i := 1; i < 9; i++ { // AI Cards at postions 1 - 8  (Zero is Person)

		betImage(dealPost, betAmount, screen)
	}
	return
}
