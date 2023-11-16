package logstash

type Level int

const (
	DebugLevel Level = 1
	InfoLevel  Level = 2
	WarnLevel  Level = 3
	ErrorLevel Level = 4
)
