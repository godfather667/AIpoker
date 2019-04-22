// ai.go is the main executive function AI Processes and Functions
package main

import (
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
)

func aiExec(users [9][2]int, betAmount, dealPost int, mode Bits, screen *ebiten.Image) {

	if has(mode, cardDelt) {
		// Process opening
		for bi := 0; bi < 9; bi++ { // AI Cards at postions 1 - 8  (Zero is Person)
			c1 := users[bi][0]
			c2 := users[bi][1]
			// Card Index (A = 0, K = 1...)
			v1 := cardReverse[deck[c1]]
			v2 := cardReverse[deck[c2]]
			// Card Suit (club = 0, diamonds = 1, hearts = 2, spades = 3)
			s1 := v1 % 4
			s2 := v2 % 4
			// Compute Opening Values
			level := chartOne[v1][v2]
			fmt.Println(c1, "  ", v1, "  ", s1, "  ", level)
			fmt.Println(c2, "  ", v2, "  ", s2)
			mode = clear(mode, cardDelt)
			mode = set(mode, cardFlop)
			if level == 9 {
				fmt.Println("Raise  ", level)
				betAmount = betAmount * 2
				mode = set(mode, raiseMade)
				storeImage(bi, betAmount, screen)
			}
			if level == 8 {
				if has(mode, raiseMade) {
					fmt.Println("Call ")
					mode = set(mode, callMade)
				} else {
					fmt.Println("Raise  ", level)
					betAmount = betAmount * 2
					mode = set(mode, raiseMade)
				}
				storeImage(bi, betAmount, screen)
			}
			if level >= 7 {
				if has(mode, raiseMade) {
					fmt.Println("Call ")
					mode = set(mode, callMade)
				} else {
					fmt.Println("Raise  ", level)
					betAmount = betAmount * 2
					mode = set(mode, raiseMade)
				}
				storeImage(bi, betAmount, screen)
			}
			if level >= 2 && !has(mode, raiseMade) {
				fmt.Println("call  ", level)
				mode = set(mode, callMade)
				storeImage(bi, betAmount, screen)
			}
			if level < 2 && has(mode, raiseMade) {
				fmt.Println("fold  ", level)
				mode = set(mode, foldMade)
				changeCard(bi, hide, undisplay, screen)
				fmt.Println("clear ", bi)
				clearImage(bi)
			}
		}

	}
	return
}
