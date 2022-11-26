package utils

import (
	"image"
	"image/color"
	"math/rand"
	"time"
	"math"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func ThresholdDither(gray *image.Gray, threshold uint8) *image.Gray {
	bounds := gray.Bounds()
	dithered := image.NewGray(bounds)
	width := bounds.Dx()
	height := bounds.Dy()
	for x := 0; x <= width; x++ {
		for y := 0; y <= height; y++ {
			var c = blackOrWhite(gray.GrayAt(x, y), threshold)
			dithered.SetGray(x, y, c)
		}
	}
	return dithered
}

// find the avg intensity of a grayscale image
func avgIntensity(gray *image.Gray) float64 {
	var sum float64
	for _, pix := range gray.Pix {
		sum += float64(pix)
	}
	return sum / float64(len(gray.Pix)*256)
}

// nk is the number of pixels in a cell, gamma is the number of points to sample.
func GridDither(gray *image.Gray, nK int, gamma, alpha float64) *image.Gray {
	bounds := gray.Bounds()

	dithered := image.NewGray(bounds) // black background

	// we iterate thought each cell size.
	for x := bounds.Min.X; x < bounds.Max.X; x += nK {
		for y := bounds.Min.Y; y < bounds.Max.Y; y += nK {
			cell := RGBAToGrey(gray.SubImage(image.Rect(x, y, x+nK, y+nK)))
			mu := avgIntensity(cell)
			n := math.Pow((1-mu)*gamma, 2) / 3 // the number of black pixels in a cell based off the density of "darkness"
			if n < alpha {
				n = 0
			}
			for k := 0; k < int(n); k++ {	
				var i = randInt(x, min(x+nK, bounds.Max.X), rng)
				var j = randInt(y, min(y+nK, bounds.Max.Y), rng)

				dithered.SetGray(i, j, color.Gray{255}) // white pixels. 
			}
		}
	}
	return dithered
}
// generates a randInt between [min, max)
func randInt(min, max int, rng *rand.Rand) int {
	return rng.Intn(max-min) + min
}

// returns the smallest fo two ints
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}