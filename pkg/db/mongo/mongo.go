package mongo

type MongoDB struct{}

// New ...
func New(config *Config) (*MongoDB, error) {
	return &MongoDB{}, nil
}
