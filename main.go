package main

import "github.com/ljq0002/image/imageProcess"

// sample

func main() {
	image, err := imageProcess.CreateImageFromFile("resource/lena.jpeg")
	if err != nil {
		panic(err)
	}
	image.GaussianBlur(5, 1.5).SaveToJpegFile("blur.jpeg")
}
