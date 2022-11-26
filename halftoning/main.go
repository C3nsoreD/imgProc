package main

import (
	"image/png" // used explicitly for initialization side-effects image.Decode will read png file formats.
	_ "image/jpeg"
	"log"
		
	"os"

	"github.com/C3nsoreD/imgProc/utils"
)

func main() {
	log := log.Default()

	var img, err = utils.LoadImage("apple.jpg")
	if err != nil {
		log.Fatalf("failed to load image: %v", err)
	}
	var gray = utils.RGBAToGrey(img)
	dithered := utils.GridDither(gray, 10, 8, 3)

	log.Printf("creating new image...")
	f, _ := os.Create("dithered.png")
	defer f.Close()
	png.Encode(f, dithered)
}


