// diag.go supplies diagnostic functions for debugging and code development.
package main

import (
	"fmt"
)

func modeDump(mode Bits) {
	for k, v := range dmap {
		if has(mode, k) {
			fmt.Println(v)
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
