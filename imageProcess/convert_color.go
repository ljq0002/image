package imageProcess

import (
	"image"
	"image/draw"
)

func (i *ImageContainer) ToGray() *ImageContainer {
	grayImage := image.NewGray(i.Bounds())
	draw.Draw(grayImage, i.Bounds(), i, i.Bounds().Min, draw.Src)
	return &ImageContainer{grayImage}
}
