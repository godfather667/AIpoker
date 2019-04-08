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
