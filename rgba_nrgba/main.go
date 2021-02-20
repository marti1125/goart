package main

import (
	"fmt"
	"image/color"
)

func main()  {
	var alpha uint8 = 42
	rgba := color.RGBA{
		R: 42,
		G: 42,
		B: 42,
		A: alpha,
	}
	r1, g1, b1, _ := rgba.RGBA()
	fmt.Println("RGBA: ", r1, g1, b1)

	nrgba := color.NRGBA{
		R: 42,
		G: 42,
		B: 42,
		A: alpha,
	}
	r2, g2, b2, _ := nrgba.RGBA()
	fmt.Println("RGBA: ", r2, g2, b2)

}
