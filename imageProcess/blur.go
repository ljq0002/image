package imageProcess

import (
	"image"
	"image/color"
	"image/draw"
	"math"
	"sort"
)

func (ic *ImageContainer) GaussianBlur(px int, rho float64) *ImageContainer {
	res := image.NewRGBA(ic.Bounds())

	rgbImage := image.NewRGBA(ic.Bounds())
	draw.Draw(rgbImage, ic.Bounds(), ic, ic.Bounds().Min, draw.Src)

	dis := getGaussianDistribution(px, rho)
	width := ic.Bounds().Dx()
	height := ic.Bounds().Dy()
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			c := getGaussianWeight(rgbImage, dis, j, i)
			res.Set(j, i, c)
		}
	}
	return &ImageContainer{Image: res}
}

func getGaussianDistributionFactor(x, y, rho float64) float64 {
	return 1.0 / math.Sqrt(2.0*math.Pi*rho*rho) * math.Exp(-(x*x + y*y)/(2*rho*rho))
}

func getGaussianDistribution(length int, rho float64) *[][]float64 {
	sum := 0.0
	res := make([][]float64, 2*length+1, 2*length+1)
	for i := 0; i < 2*length+1; i++ {
		res[i] = make([]float64, 2*length+1, 2*length+1)
	}
	for i := 0; i < length; i++ {
		weight := getGaussianDistributionFactor(float64(i-length), 0, rho)
		res[i][length] = weight
		sum += 4 * weight
		for j := 0; j < length; j++ {
			weight := getGaussianDistributionFactor(float64(i-length), float64(j-length), rho)
			res[i][j] = weight
			sum += 4 * weight
		}
	}
	weight := getGaussianDistributionFactor(0, 0, rho)
	res[length][length] = weight
	sum += weight

	for i := 0; i < length; i++ {
		res[i][length] /= sum
		res[2*length-i][length], res[length][i], res[length][2*length-i] = res[i][length], res[i][length], res[i][length]
		for j := 0; j < length; j++ {
			res[i][j] /= sum
			res[i][2*length-j], res[2*length-i][j], res[2*length-i][2*length-j] = res[i][j], res[i][j], res[i][j]
		}
	}
	res[length][length] /= sum
	return &res
}

func getGaussianWeight(im image.Image, dis *[][]float64, x, y int) color.RGBA {
	R, G, B := 0.0, 0.0, 0.0
	length := len(*dis)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			xx := x - length/2 + j
			yy := y - length/2 + i
			if xx < 0 {
				xx = 0
			} else if xx > im.Bounds().Max.X-1 {
				xx = length - 1
			}
			if yy < 0 {
				yy = 0
			} else if yy > im.Bounds().Max.Y-1 {
				yy = length - 1
			}

			r, g, b, _ := im.At(xx, yy).RGBA()
			r8 := uint8(r)
			g8 := uint8(g)
			b8 := uint8(b)
			R += float64(r8) * (*dis)[i][j]
			G += float64(g8) * (*dis)[i][j]
			B += float64(b8) * (*dis)[i][j]
		}
	}
	_, _, _, a := im.At(x, y).RGBA()
	return color.RGBA{R: uint8(R), G: uint8(G), B: uint8(B), A: uint8(a)}
}

func (ic *ImageContainer) MedianBlur(size int) *ImageContainer {
	res := image.NewRGBA(ic.Bounds())
	draw.Draw(res, ic.Bounds(), ic, ic.Bounds().Min, draw.Src)

	rgbImage := image.NewRGBA(ic.Bounds())
	draw.Draw(rgbImage, ic.Bounds(), ic, ic.Bounds().Min, draw.Src)

	width := ic.Bounds().Dx()
	height := ic.Bounds().Dy()
	start := size / 2
	for i := start; i < height-start; i++ {
		for j := start; j < width-start; j++ {
			r := make([]int, 0, size*size)
			g := make([]int, 0, size*size)
			b := make([]int, 0, size*size)
			for m := i - start; m < i+start+1; m++ {
				for n := j - start; n < j+start+1; n++ {
					R, G, B, _ := rgbImage.At(n, m).RGBA()
					r = append(r, int(R))
					g = append(g, int(G))
					b = append(b, int(B))
				}
			}
			sort.Ints(r)
			sort.Ints(g)
			sort.Ints(b)
			_, _, _, A := rgbImage.At(j, i).RGBA()
			res.Set(j, i, color.RGBA{
				R: uint8(r[size*size/2]),
				G: uint8(g[size*size/2]),
				B: uint8(b[size*size/2]),
				A: uint8(A),
			})
		}
	}
	return &ImageContainer{Image: res}
}

func (ic *ImageContainer) MeanBlur(size int) *ImageContainer {
	res := image.NewRGBA(ic.Bounds())
	draw.Draw(res, ic.Bounds(), ic, ic.Bounds().Min, draw.Src)

	rgbImage := image.NewRGBA(ic.Bounds())
	draw.Draw(rgbImage, ic.Bounds(), ic, ic.Bounds().Min, draw.Src)

	width := ic.Bounds().Dx()
	height := ic.Bounds().Dy()

	start := size / 2
	px := uint32(size * size)
	for i := start; i < height-start; i++ {
		for j := start; j < width-start; j++ {
			r, g, b := uint32(0), uint32(0), uint32(0)
			for m := i - start; m < i+start+1; m++ {
				for n := j - start; n < j+start+1; n++ {
					R, G, B, _ := rgbImage.At(n, m).RGBA()
					r += uint32(uint8(R))
					g += uint32(uint8(G))
					b += uint32(uint8(B))
				}
			}
			_, _, _, A := rgbImage.At(j, i).RGBA()
			r = r / px
			g = g / px
			b = b / px
			res.Set(j, i, color.RGBA{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
				A: uint8(A),
			})
		}
	}
	return &ImageContainer{Image: res}
}
