package main

import (
	"art/util"
	"image"
	"image/color"
)

type GrayscaleFilter struct {
	image.Image
}

func (f *GrayscaleFilter) At(x, y int) color.Color {
	r, g, b, a := f.Image.At(x, y).RGBA()
	grey := uint16(float64(r)*0.21 + float64(g)*0.72 + float64(b)*0.07)
	return color.RGBA64{
		R: grey,
		G: grey,
		B: grey,
		A: uint16(a),
	}
}

func main() {

	img, err := util.LoadImage("C:\\Users\\Willy\\Pictures\\[GR]_concept_01.png")
	if err != nil {
		panic(err.Error())
	}
	util.SaveOutPut(&GrayscaleFilter{img}, "C:\\Users\\Willy\\Pictures\\sample2.png")

}
