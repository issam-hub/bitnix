package storageclient

import "github.com/cloudinary/cloudinary-go/v2"

func NewCloudinaryInstance(cloudinaryURL string) (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	return cld, err
}
