// ai.go is the main executive function AI Processes and Functions
package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
)

func aiExec(users []int, betAmount, dealPost int, mode Bits, screen *ebiten.Image) {
	for i := 0; i < 9; i++ {
		betImage(dealPost, betAmount, screen)

	}
	return
}
