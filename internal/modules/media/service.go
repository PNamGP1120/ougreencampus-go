package media

import (
	"errors"
	"mime/multipart"

	"github.com/PNamGP1120/ougreencampus-go/pkg/cloudinary"
)

type Service interface {
	Upload(file multipart.File, ownerID uint) (*Media, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Upload(file multipart.File, ownerID uint) (*Media, error) {
	if file == nil {
		return nil, errors.New("file is required")
	}

	url, err := cloudinary.Upload(file, "ougreencampus/avatar")
	if err != nil {
		return nil, err
	}

	media := &Media{
		URL:     url,
		Type:    "image",
		OwnerID: ownerID,
	}

	if err := s.repo.Create(media); err != nil {
		return nil, err
	}

	return media, nil
}
