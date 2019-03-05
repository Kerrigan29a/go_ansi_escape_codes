package main

import (
	"fmt"

	. "github.com/Kerrigan29a/go_ansi_escape_codes"
	"github.com/containerd/console"
)

func main() {
	current := console.Current()
	defer current.Reset()

	fmt.Println("RBG colors")
	for r := 0; r < 256; r++ {
		for g := 0; g < 256; g++ {
			for b := 0; b < 256; b++ {
				s := RGBColor(uint8(r), uint8(g), uint8(b), true)
				fmt.Print(s.String())
				fmt.Print(" ")
			}
		}
	}
	fmt.Println(Reset.String())
}
