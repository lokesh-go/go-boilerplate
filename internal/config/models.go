package config

type config struct {
	App    appConfig    `yaml:"app"`
	Server serverConfig `yaml:"server`
	Logger loggerConfig `yaml:"logger"`
	DAL    dalConfig    `yaml:"dal"`
}

type appConfig struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
	Author      string `yaml:"author"`
	URL         string `yaml:"url"`
}

type serverConfig struct {
	HTTP httpServerConfig `yaml:"http"`
	GRPC grpcServerConfig `yaml:"grpc"`
}

type httpServerConfig struct {
	PublicAddr               string `yaml:"publicAddr"`
	InternalAddr             string `yaml:"internalAddr"`
	ReadTimeoutInSeconds     int64  `yaml:"readTimeoutInSeconds"`
	WriteTimeoutInSeconds    int64  `yaml:"WriteTimeoutInSeconds"`
	ShutdownTimeoutInSeconds int64  `yaml:"shutdownTimeoutInSeconds"`
}

type grpcServerConfig struct{}

type loggerConfig struct {
	Debug        bool `yaml:"debug"`
	CallerSkipNo int  `yaml:"callerSkipNo"`
}

type dalConfig struct {
	Cache cacheConfig `yaml:"cache"`
	DB    dbConfig    `yaml:"db"`
}

type cacheConfig struct {
	AppCache appCacheConfig   `yaml:"appCache"`
	Redis    redisCacheConfig `yaml:"redis"`
}

type appCacheConfig struct {
	Enabled                    bool  `yaml:"enabled"`
	DefaultExpirationInSeconds int64 `yaml:"defaultExpirationInSeconds"`
	CleanupIntervalInMinutes   int64 `yaml:"cleanupIntervalInMinutes"`
}

type redisCacheConfig struct {
	Enabled bool `yaml:"enabled"`
}

type dbConfig struct{}
