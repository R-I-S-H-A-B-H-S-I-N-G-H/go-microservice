package services

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/db_utils"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/error_util"
	"image"
	"image/color"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type Pixel struct {
	ID          string `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string `bson:"name" json:"name"`
	UserId      string `bson:"user_id" json:"user_id"`
	Count       int64  `bson:"count" json:"count"`
	DateCreated int64  `bson:"date_created" json:"date_created"`
	LastUpdated int64  `bson:"last_updated" json:"last_updated"`
}

func CreateNewPixelObj(id string, name string, userId string, dateCreated string) *Pixel {
	pixel := &Pixel{
		Name:        name,
		UserId:      userId,
		Count:       0,
		LastUpdated: time.Now().Unix(),
	}

	if id != "" {
		pixel.ID = id
	} else {
		pixel.DateCreated = time.Now().Unix()
	}

	return pixel
}

func addPixelToDB(pixel *Pixel) (*Pixel, error) {
	return db_utils.CreateOne("pixels", pixel)
}

func getPixelByIdFromDB(idStr string) (*Pixel, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	error_util.Handle("Failed to parse id", err)
	filter := bson.M{"_id": id}
	return db_utils.FindOne[Pixel]("pixels", filter)
}

func updatePixelCount(idStr string, count int64) error {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}

	// Create an update document using the $set operator
	update := bson.M{"$set": bson.M{"count": count}}

	return db_utils.UpdateOne("pixels", filter, update)
}

func PixelCaptureService(pixelId string) *image.RGBA {
	pixel, err := getPixelByIdFromDB(pixelId)
	error_util.Handle("Failed to get pixel", err)
	pixel.Count = pixel.Count + 1

	updatePixelCount(pixelId, pixel.Count)

	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.RGBA{0, 0, 0, 0})
	return img
}

func PixelListService(page int, pageSize int) ([]bson.M, error) {
	filter := bson.M{}

	return db_utils.FindAllWithPagination("pixels", filter, page, pageSize, "date_created", false)
}

func PixelSaveService(pixel Pixel, userId string) (*Pixel, error) {
	new_pixel := CreateNewPixelObj("", pixel.Name, userId, "")
	return addPixelToDB(new_pixel)
}
