package logger

var callerSkipNo int = 3

// Config ...
type Config struct {
	Env          string
	AppName      string
	AppVersion   string
	DebugEnabled bool
	CallerSkipNo int
}
