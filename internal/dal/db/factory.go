package db

import (
	"github.com/lokesh-go/go-api-service/internal/config"
	mongoDAL "github.com/lokesh-go/go-api-service/internal/dal/db/mongo"
	mysqlDAL "github.com/lokesh-go/go-api-service/internal/dal/db/mysql"
)

type dbType string

const (
	mongoDB   dbType = "mongoDB"
	mySQL     dbType = "mySQL"
	dbEnabled dbType = mongoDB
)

// New DB
func New(config config.Methods) (DB, error) {
	// Switch to the DB
	switch dbEnabled {
	case mySQL:
		{
			return mysqlDAL.New(config)
		}
	default:
		{
			return mongoDAL.New(config)
		}
	}
}
