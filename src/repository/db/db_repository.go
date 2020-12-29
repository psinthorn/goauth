package db

import (
	"github.com/psinthorn/goauth/src/domain/access_token"
	"github.com/psinthorn/goauth/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.ErrorRespond)
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r dbRepository) GetById(id string) (*access_token.AccessToken, *errors.ErrorRespond) {
	return nil, errors.NewInternalServerError("Please implement database connection")
}
