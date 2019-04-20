// ai.go is the main executive function AI Processes and Functions
package main

import (
	"fmt"
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
		// Card Values
		c1 := users[i][0]
		c2 := users[i][1]
		// Card Index (A = 0, K = 1...)
		v1 := cardReverse[deck[c1]]
		v2 := cardReverse[deck[c2]]
		// Card Suit (club = 0, diamonds = 1, hearts = 2, spades = 3)
		s1 := v1 % 4
		s2 := v2 % 4

		level := chartOne[v1][v2]
		fmt.Println(c1, "  ", v1, "  ", s1, "  ", level)
		fmt.Println(c2, "  ", v2, "  ", s2)

		betImage(dealPost, betAmount, screen)
	}
	return
}
