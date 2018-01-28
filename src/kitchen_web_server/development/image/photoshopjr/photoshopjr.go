// package photoshopjr
package main

// Photoshop Jr. is the code used to convert images to 500x500px images, one that retains the original color and the other turning the image into monochrome.

import (
	"image/color"
	"image/jpeg"
	"log"
	"os"

	"github.com/harrydb/go/img/grayscale"
)

type ImageSet interface {
	Set(x, y int, c color.Color)
}

func separateFilenameFromExtension(filename string) string {
	strippedfilename := filename[0 : len(filename)-4]
	return strippedfilename
}

func MakeImages(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)

	if err != nil {
		log.Fatal(os.Stderr, "%s: %v\n", filename, err)
	}

	grayImg := grayscale.Convert(img, grayscale.ToGrayLuminance)

	outFile, err := os.Create(separateFilenameFromExtension(filename) + ".mono.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	jpeg.Encode(outFile, grayImg, nil)

}

func main() {
	// separateFilenameFromExtension("filename.jpg")
	MakeImages("../files/1.jpg")

}
