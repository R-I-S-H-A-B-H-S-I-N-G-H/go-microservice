package services

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/db_utils"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/error_util"
	"image"
	"image/color"
)

type Pixel struct {
	ID     string `bson:"_id,omitempty" json:"id,omitempty"`
	Name   string `bson:"name" json:"name"`
	UserId string `bson:"user_id" json:"user_id"`
}

func addPixelToDB(pixel *Pixel) error {
	return db_utils.CreateOne("pixels", pixel)
}

func PixelCaptureService(userId string) *image.RGBA {
	// Create a new 1x1 pixel image
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))

	// Set the pixel to a color (e.g., white)
	img.Set(0, 0, color.RGBA{0, 0, 0, 0})

	pixel := Pixel{
		Name:   "new Pixel",
		UserId: userId,
	}
	err := addPixelToDB(&pixel)

	error_util.Handle("Failed to create pixel", err)

	return img
}
