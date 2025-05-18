package address

import (
	"github.com/lokesh-go/go-boilerplate/internal/dal/cache"
	"github.com/lokesh-go/go-boilerplate/internal/dal/db"
)

// Address data access layer
type AddressDAL struct {
	cache cache.Cache
	db    db.DB
}

// New creates a new address data access layer
func New(cache cache.Cache, db db.DB) *AddressDAL {
	return &AddressDAL{
		cache: cache,
		db:    db,
	}
}

// GetAddressByUserID retrieves an address by user ID
func (a *AddressDAL) GetAddressByUserID(userID string) (data *AddressDocument, err error) {
	return nil, nil
}
