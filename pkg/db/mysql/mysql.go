package mysql

type MySql struct{}

// New ...
func New(config *Config) (*MySql, error) {
	return &MySql{}, nil
}
