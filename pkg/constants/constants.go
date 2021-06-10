package constants

type Operation string

const (
	MaxRetries          = 20
	RetryDuration       = 10
	DefaultMongoTimeout = 20

	Add    Operation = "$push"
	Remove Operation = "$pull"
)
