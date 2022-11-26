package utils

import (
	"fmt"
	"image"
	_ "image/jpeg" // used explicitly for initialization side-effects image.Decode will read jpg file formats.
	"os"
	"image/color"
)

// load an image from disk
func LoadImage(src string) (image.Image, error) {
	in, err := os.Open(src)
	if err != nil {
		return nil, fmt.Errorf("failed to load image %v", err)
	}
	defer in.Close()
	img, _, err := image.Decode(in)
	if err != nil {
		return nil, fmt.Errorf("failed to decode %v", err)
	}
	return img, nil
}

// convert a rgba image into a grayscale image
func RGBAToGrey(img image.Image) *image.Gray {
	bounds := img.Bounds() // image boundaries 
	gray := image.NewGray(bounds) // empty grayscale image
	for x := 0; x <= bounds.Max.X; x++ {
		for y := 0; y <= bounds.Max.Y; y++ {
			var rgba = img.At(x, y) //return color value at x,y
			gray.Set(x, y, rgba)
		}
	}
	return gray
}

// give graysacele color a threshold 
func blackOrWhite(g color.Gray, thr uint8) color.Gray {
	if g.Y < thr {
		return color.Gray{0} // black
	}
	return color.Gray{255} // white
}

func newWhite(bounds image.Rectangle) *image.Gray {
	var white = image.NewGray(bounds)
	for i := range white.Pix {
		white.Pix[i] = 255 // make every pixel white
	}
	return white
}