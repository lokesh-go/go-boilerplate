package dal

import (
	"github.com/lokesh-go/go-boilerplate/internal/config"
	"github.com/lokesh-go/go-boilerplate/internal/dal/cache"
	"github.com/lokesh-go/go-boilerplate/internal/dal/db"
	addressEntity "github.com/lokesh-go/go-boilerplate/internal/dal/entities/address"
	userEntity "github.com/lokesh-go/go-boilerplate/internal/dal/entities/user"
)

// Data Access Layer
type DAL interface {
	User() userDAO
	Address() addressDAO
}

type userDAO interface {
	GetUserByID(id string) (data *userEntity.UserDocument, err error)
}

type addressDAO interface {
	GetAddressByUserID(userID string) (data *addressEntity.AddressDocument, err error)
}

type dal struct {
	user    *userEntity.UserDAL
	address *addressEntity.AddressDAL
}

// Initialize
func Initialize(config config.Methods) (DAL, error) {
	// Connects cache
	cache, err := cache.New(config)
	if err != nil {
		return nil, err
	}

	// Connects DB
	db, err := db.New(config)
	if err != nil {
		return nil, err
	}

	// Returns
	return &dal{
		user:    userEntity.New(cache, db),
		address: addressEntity.New(cache, db),
	}, nil
}

// User
func (d *dal) User() userDAO {
	return d.user
}

// Address
func (d *dal) Address() addressDAO {
	return d.address
}
