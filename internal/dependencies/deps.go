package dependencies

import (
	"github.com/lokesh-go/go-api-microservice/internal/config"
	"github.com/lokesh-go/go-api-microservice/internal/dal"
	"github.com/lokesh-go/go-api-microservice/pkg/logger"
)

// Dependencies hold all application deps
type Deps struct {
	Config config.Methods
	Logger *logger.Logger
	DAL    dal.DAL
}
