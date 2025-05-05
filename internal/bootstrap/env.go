package bootstrap

import "github.com/lokesh-go/go-boilerplate/pkg/utils"

func getEnv() (env string) {
	env = utils.GetEnv(utils.EnvKey)
	if utils.IsEmpty(env) {
		env = utils.EnvDefault
	}
	return env
}
