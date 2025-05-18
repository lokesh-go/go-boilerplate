package dependencies

import (
	"github.com/lokesh-go/go-api-service/internal/config"
	"github.com/lokesh-go/go-api-service/internal/dal"
	"github.com/lokesh-go/go-api-service/pkg/logger"
)

// Dependencies hold all application deps
type Deps struct {
	Config config.Methods
	Logger *logger.Logger
	DAL    dal.DAL
}
