package main

import (
	"art/util"
	"image"
	"image/color"
)

type CircularMask struct {
	source image.Image
	center image.Point
	radius int
}

func (c *CircularMask) ColorModel() color.Model  {
	return c.source.ColorModel()
}

func (c *CircularMask) Bounds() image.Rectangle {
	return image.Rect(c.center.X-c.radius, c.center.Y-c.radius, c.center.X+c.radius, c.center.Y+c.radius)
}

func (c *CircularMask) At(x, y int) color.Color {
	xx :=float64(x-c.center.X)
	yy :=float64(y-c.center.Y)
	rr := float64(c.radius)
	if xx*xx+yy*yy < rr*rr {
		return c.source.At(x, y)
	}
	return color.Alpha{}
}

func main()  {
	img, err := util.LoadImage("C:\\Users\\Willy\\Pictures\\[GR]_concept_01.png")
	if err != nil {
		panic("error")
	}
	c := &CircularMask{
		source: img,
		center: image.Point{
			X: 500,
			Y: 500,
		},
		radius: 300,
	}
	util.SaveOutPut(c, "C:\\Users\\Willy\\Pictures\\sample.png")
}
