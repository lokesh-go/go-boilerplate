package dependencies

import (
	"github.com/lokesh-go/go-boilerplate/internal/config"
	"github.com/lokesh-go/go-boilerplate/internal/dal"
	"github.com/lokesh-go/go-boilerplate/pkg/logger"
)

// Dependencies hold all application deps
type Deps struct {
	Config config.Methods
	Logger *logger.Logger
	DAL    dal.DAL
}
