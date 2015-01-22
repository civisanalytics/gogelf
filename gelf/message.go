// Copyright © 2014-2015, Civis Analytics

package gelf

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

// Message meets the Graylog2 Extended Log Format.
// http://graylog2.org/gelf#specs
type Message struct {
	Version          string                 `json:"version"`
	Host             string                 `json:"host"`
	ShortMessage     string                 `json:"short_message"`
	FullMessage      string                 `json:"full_message,omitempty"`
	Timestamp        int64                  `json:"timestamp"`
	Level            Level                  `json:"level"`
	AdditionalFields string                 `json:",omitempty"`
	additional       map[string]interface{} `json:"a,omitempty"`
}

var reservedFields = []string{"version", "host", "short_message", "full_message", "timestamp", "level", "_id"}

// NewMessage returns a new Graylog2 Extended Log Format message.
func NewMessage(l Level, short string, full string) (*Message, error) {
	a := make(map[string]interface{})

	host, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return &Message{
		Version:      GELFVersion,
		Host:         host,
		ShortMessage: short,
		FullMessage:  full,
		Timestamp:    time.Now().UnixNano(),
		Level:        l,
		additional:   a,
	}, nil
}

// Add will add additional fields to a message in the form of a key and value
// pair. Values can be of string or int type.
func (m *Message) Add(key string, value interface{}) error {
	// Verify additional fields against reserved field names.
	// If field is not reserved, add to message.
	for _, rf := range reservedFields {
		if key == rf {
			return fmt.Errorf("Invalid field[%s]", key)
		}
	}

	// Verify underscore prefix
	r, _ := utf8.DecodeRuneInString(key)
	if string(r) == "_" {
		m.additional[key] = value
	} else {
		m.additional["_"+key] = value
	}

	return nil
}

// String is a convience method that meets the fmt.String interface providing an
// easy way to print the string JSON representation of a message.
func (m *Message) String() string {
	// Maps do not marshal to JSON as top-level objects.
	// To work around we marshal the map of additional fields, modify the string
	// and append to the outbound JSON encoded struct.
	additionalFields, _ := json.Marshal(m.additional)
	filteredFields := strings.Replace(string(additionalFields[1:]), "\\\"", "\"", -1)

	baseMessageFields, _ := json.Marshal(m)
	trimBaseMessageFields := strings.TrimRight(string(baseMessageFields), "}")

	return trimBaseMessageFields + "," + filteredFields
}
