// Copyright © 2014-2015, Civis Analytics

package gelf

// Level is a generic type for severity.
type Level int

// Levels matching RFC5424 severity.
const (
	LevelPanic Level = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInfo
	LevelDebug
)

// Panic generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 0.
// System is unusable.
func Panic(s, l string) *Message {
	return NewMessage(LevelPanic, s, l)
}

// Alert generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 1.
// Action must be taken immediately.
func Alert(s, l string) *Message {
	return NewMessage(LevelAlert, s, l)
}

// Crit generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 2.
// Critical conditions.
func Crit(s, l string) *Message {
	return NewMessage(LevelCritical, s, l)
}

// Error generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 3.
// Error conditions.
func Error(s, l string) *Message {
	return NewMessage(LevelError, s, l)
}

// Warn generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 4.
// Warning conditions.
func Warn(s, l string) *Message {
	return NewMessage(LevelWarning, s, l)
}

// Notice generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 5.
// Normal but significant condition.
func Notice(s, l string) *Message {
	return NewMessage(LevelNotice, s, l)
}

// Info generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 6.
// Informational messages.
func Info(s, l string) *Message {
	return NewMessage(LevelInfo, s, l)
}

// Debug generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 7.
// Developer debug messages.
func Debug(s, l string) *Message {
	return NewMessage(LevelDebug, s, l)
}
