package mysql

import (
	"github.com/lokesh-go/go-api-microservice/internal/config"
	mysqlPkg "github.com/lokesh-go/go-api-microservice/pkg/db/mysql"
)

// MySql connection
type mySql struct {
	conn *mysqlPkg.MySql
}

// New mysql connection
func New(config config.Methods) (*mySql, error) {
	// Forms mysql config
	mysqlConfig := &mysqlPkg.Config{}

	// Creates mysql connection
	conn, err := mysqlPkg.New(mysqlConfig)
	if err != nil {
		return nil, err
	}

	// Returns
	return &mySql{
		conn: conn,
	}, nil
}
