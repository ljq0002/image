package imageProcess

import (
	"image"
	"image/color"
	"image/draw"
)

func (i *ImageContainer) ToGray() *ImageContainer {
	grayImage := image.NewGray(i.Bounds())
	draw.Draw(grayImage, i.Bounds(), i, i.Bounds().Min, draw.Src)
	return &ImageContainer{grayImage}
}

func (i *ImageContainer) ToBinary(threshold int) *ImageContainer {
	res := image.NewGray(i.Bounds())
	grayImage := i.ToGray()

	width := i.Bounds().Dx()
	height := i.Bounds().Dy()

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			Y, _, _, _ := grayImage.At(j, i).RGBA()
			y := uint8(Y)
			if int(y) > threshold {
				y = 255
			} else {
				y = 0
			}
			res.Set(j, i, color.Gray{Y: y})
		}
	}
	return &ImageContainer{res}
}
