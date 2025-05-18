package bootstrap

import "github.com/lokesh-go/go-api-microservice/pkg/utils"

func getEnv() (env string) {
	env = utils.GetEnv(utils.EnvKey)
	if utils.IsEmpty(env) {
		env = utils.EnvDefault
	}
	return env
}
