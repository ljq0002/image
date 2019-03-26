package imageProcess

import (
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func CreateImageFromFile(file string) (*ImageContainer, error) {
	fileName := file[strings.LastIndex(file, ".")+1:]
	if strings.Contains(strings.ToLower(fileName), "jpeg") || strings.Contains(strings.ToLower(fileName), "jpg") {
		return CreateImageFromJpegFile(file)
	} else if strings.Contains(strings.ToLower(fileName), "png") {
		return CreateImageFromPngFile(file)
	} else {
		return nil, unknownImageTypeError{}
	}
}

func CreateImageFromPngFile(file string) (*ImageContainer, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	pngImage, err := png.Decode(f)
	if err != nil {
		return nil, err
	}
	return &ImageContainer{pngImage}, nil
}

func CreateImageFromJpegFile(file string) (*ImageContainer, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	jpegImage, err := jpeg.Decode(f)
	if err != nil {
		return nil, err
	}
	return &ImageContainer{jpegImage}, nil

}

func (i *ImageContainer) SaveToPngFile(file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, i)
}

func (i *ImageContainer) SaveToJpegFile(file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return jpeg.Encode(f, i, nil)
}
