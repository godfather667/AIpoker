// ai.go is the main executive function AI Processes and Functions
package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
)

func aiExec(betAmount, dealPost int, mode Bits, screen *ebiten.Image) {
	betImage(dealPost, betAmount, screen)
	return
}
