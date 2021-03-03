package access_token

import (
	"strings"

	"github.com/psinthorn/gostack_goauth-api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.ErrorRespond)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.ErrorRespond)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(id string) (*AccessToken, *errors.ErrorRespond) {
	id = strings.TrimSpace("id")
	if len(id) == 0 {
		return nil, errors.NewBadRequestError("Invalid Access Token ID")
	}
	accessToken, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
