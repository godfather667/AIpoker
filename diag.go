// diag.go supplies diagnostic functions for debugging and code development.
package main

import (
	"fmt"
)

func modeDump(mode Bits) {
	if dumpModes {
		if has(mode, cardDeal) {
			fmt.Println("Mode = cardDeal")
		}
		if has(mode, cardDelt) {
			fmt.Println("Mode = cardDelt")
		}
		if has(mode, cardFlop) {
			fmt.Println("Mode = cardFlop")
		}
		if has(mode, cardTurn) {
			fmt.Println("Mode = cardTurn")
		}
		if has(mode, cardRiver) {
			fmt.Println("Mode = cardRiver")
		}
		if has(mode, betValue) {
			fmt.Println("Mode = betValue")
		}
		if has(mode, betInput) {
			fmt.Println("Mode = betInput")
		}
		if has(mode, betMade) {
			fmt.Println("Mode = betMade")
		}
		if has(mode, betMade) {
			fmt.Println("Mode = betMade")
		}
		if has(mode, inputWait) {
			fmt.Println("Mode = inputWait")
		}
		if has(mode, dealWait) {
			fmt.Println("Mode = dealWait")
		}
		if has(mode, aiProcess) {
			fmt.Println("Mode = aiProcess")
		}
		if has(mode, foldMade) {
			fmt.Println("Mode = foldMode")
		}
		if has(mode, checkMade) {
			fmt.Println("Mode = checkMade")
		}
		if has(mode, callMade) {
			fmt.Println("Mode = callMade")
		}
		if has(mode, raiseMade) {
			fmt.Println("Mode = raiseMade")
		}
	}
}

func dumpPlayers(note string) {
	for i := 0; i < 18; i++ {
		j := i * 9
		fmt.Println("--", note, "--")
		fmt.Printf("Player %d: ", i)
		if players[j+7] == hide {
			fmt.Printf(" Hide ")
		} else {
			fmt.Printf("Unhide")
		}
		if players[j+8] == display {
			fmt.Println(" Display")
		} else {
			fmt.Println("Undisplay")
		}
	}
}

func displayAll(players []int) {
	for i := 0; i < 18; i++ {
		j := i * 9
		players[j+7] = unhide
	}
}

func undisplayAll(players []int) {
	for i := 0; i < 18; i++ {
		j := i * 9
		if i == 8 || i == 17 {
			players[j+7] = unhide
		} else {
			players[j+7] = hide
		}
	}
}
