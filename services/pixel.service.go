package services

import (
	"image"
	"image/color"
	"log"
)

func PixelCaptureService(id string) *image.RGBA {
	log.Default().Println("Capture id :: " + id)
	// Create a new 1x1 pixel image
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))

	// Set the pixel to a color (e.g., white)
	img.Set(0, 0, color.RGBA{0, 0, 0, 0})
	return img
}
