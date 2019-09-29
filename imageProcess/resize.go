package imageProcess

import "image"
import "github.com/nfnt/resize"

func (i *ImageContainer) Sub(rectangle image.Rectangle) *ImageContainer {
	var res image.Image
	switch i.Image.(type) {
	case *image.RGBA:
		temp := i.Image.(*image.RGBA)
		res = temp.SubImage(rectangle).(*image.RGBA)
		return &ImageContainer{res}
	case *image.YCbCr:
		temp := i.Image.(*image.YCbCr)
		res = temp.SubImage(rectangle).(*image.YCbCr)
		return &ImageContainer{res}
	case *image.NRGBA:
		temp := i.Image.(*image.NRGBA)
		res = temp.SubImage(rectangle).(*image.NRGBA)
		return &ImageContainer{res}
	case *image.Paletted:
		temp := i.Image.(*image.Paletted)
		res = temp.SubImage(rectangle).(*image.Paletted)
		return &ImageContainer{res}
	default:
		return nil
	}
}

func (i *ImageContainer) Resize(width, height int) *ImageContainer {
	res := resize.Resize(uint(width), uint(height), i.Image, resize.Lanczos3)
	return &ImageContainer{res}
}
