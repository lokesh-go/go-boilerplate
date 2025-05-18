package user

import (
	"github.com/lokesh-go/go-api-microservice/internal/dal/cache"
	"github.com/lokesh-go/go-api-microservice/internal/dal/db"
)

// User data access layer
type UserDAL struct {
	cache cache.Cache
	db    db.DB
}

// New creates a new user data access layer
func New(cache cache.Cache, db db.DB) *UserDAL {
	return &UserDAL{
		cache: cache,
		db:    db,
	}
}

// GetUserByID fetch a user data by ID
func (u *UserDAL) GetUserByID(id string) (data *UserDocument, err error) {
	return nil, nil
}
