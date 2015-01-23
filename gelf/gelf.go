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
func Panic(s, l string) (*Message, error) {
	m, err := NewMessage(LevelPanic, s, l)
	if err != nil {
		return &Message{}, err
	}

	return m, nil
}

// Alert generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 1.
// Action must be taken immediately.
func Alert(s, l string) (*Message, error) {
	m, err := NewMessage(LevelAlert, s, l)
	if err != nil {
		return &Message{}, err
	}

	return m, nil
}

// Crit generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 2.
// Critical conditions.
func Crit(s, l string) (*Message, error) {
	m, err := NewMessage(LevelCritical, s, l)
	if err != nil {
		return &Message{}, err
	}

	return m, nil
}

// Error generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 3.
// Error conditions.
func Error(s, l string) (*Message, error) {
	m, err := NewMessage(LevelError, s, l)
	if err != nil {
		return &Message{}, err
	}

	return m, nil
}

// Warn generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 4.
// Warning conditions.
func Warn(s, l string) (*Message, error) {
	m, err := NewMessage(LevelWarning, s, l)
	if err != nil {
		return &Message{}, err
	}

	return m, nil
}

// Notice generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 5.
// Normal but significant condition.
func Notice(s, l string) (*Message, error) {
	m, err := NewMessage(LevelNotice, s, l)
	if err != nil {
		return &Message{}, err
	}

	return m, nil
}

// Info generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 6.
// Informational messages.
func Info(s, l string) (*Message, error) {
	m, err := NewMessage(LevelInfo, s, l)
	if err != nil {
		return &Message{}, err
	}

	return m, nil
}

// Debug generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 7.
// Developer debug messages.
func Debug(s, l string) (*Message, error) {
	m, err := NewMessage(LevelDebug, s, l)
	if err != nil {
		return &Message{}, err
	}

	return m, nil
}
