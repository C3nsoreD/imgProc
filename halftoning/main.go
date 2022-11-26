package main

import (
	"image/png" // used explicitly for initialization side-effects image.Decode will read png file formats.
	_ "image/jpeg"
	"log"
	
)

func main() {
	var img, err = utils.LoadImage("apple.jpg")
	if err != nil {
		log.Fatalf("failed to load image: %v", err)
	}
	var gray = utils.RGBAToGrey(img)

	f, _ := os.Create("gray_apple.png")
	defer f.Close()
	png.Encode(f, gray)
}


