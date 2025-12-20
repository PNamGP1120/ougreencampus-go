package cloudinary

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// Upload uploads file to Cloudinary and returns secure URL
func Upload(file multipart.File, folder string) (string, error) {
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return "", err
	}

	res, err := cld.Upload.Upload(
		context.Background(),
		file,
		uploader.UploadParams{
			Folder: folder,
		},
	)
	if err != nil {
		return "", err
	}

	return res.SecureURL, nil
}
