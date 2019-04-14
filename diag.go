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
		if has(mode, betEnable) {
			fmt.Println("Mode = betEnable")
		}
		if has(mode, inputWait) {
			fmt.Println("Mode = inputWait")
		}
		if has(mode, waitDeal) {
			fmt.Println("Mode = waitDeal")
		}
		if has(mode, aiProcess) {
			fmt.Println("Mode = aiProcess")
		}
		if has(mode, betMade) {
			fmt.Println("Mode = betMade")
		}
		if has(mode, foldMade) {
			fmt.Println("Mode = foldMode")
		}
		if has(mode, checkMade) {
			fmt.Println("Mode = checkMade")
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

func setOneShot() {
	shotGate = 0
	return
}

func oneShot(msg string) {
	if shotGate == 0 {
		fmt.Println(msg)
	}
	shotGate = 1
	return
}
