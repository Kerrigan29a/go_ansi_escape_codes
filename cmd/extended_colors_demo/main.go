package main

import (
	"fmt"

	. "github.com/Kerrigan29a/go_ansi_escape_codes"
	"github.com/containerd/console"
)

func main() {
	current := console.Current()
	defer current.Reset()

	fmt.Println("Standard colors")
	for i := 0; i < 8; i++ {
		code := uint8(i)
		s := ExtendedColor(code, true)
		fmt.Print(s.String())
		fmt.Printf(" %03d ", code)
	}
	fmt.Println(Reset.String())

	fmt.Println("Bright colors")
	for i := 8; i < 16; i++ {
		code := uint8(i)
		s := ExtendedColor(code, true)
		fmt.Print(s.String())
		fmt.Printf(" %03d ", code)
	}
	fmt.Println(Reset.String())

	fmt.Println("Extended colors")
	for r := 0; r < 6; r++ {
		for g := 0; g < 6; g++ {
			for b := 0; b < 6; b++ {
				code := uint8(16 + 36*r + 6*g + b)
				s := ExtendedColor(code, true)
				fmt.Print(s.String())
				fmt.Printf(" %03d ", code)
			}
		}
		fmt.Println(Reset.String())
	}
	fmt.Println("Grayscale colors")
	for i := 232; i < 256; i++ {
		code := uint8(i)
		s := ExtendedColor(code, true)
		fmt.Print(s.String())
		fmt.Printf(" %03d ", code)
	}
	fmt.Println(Reset.String())
}
