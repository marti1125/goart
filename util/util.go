package util

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func LoadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return img, nil
}

func SaveOutPut(img image.Image, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
