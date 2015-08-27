// Copyright © 2014-2015, Civis Analytics

package gelf

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

// gelfTime is a small wrapper around time.Time with methods
// for marshalling and unmarshalling JSON.
type gelfTime struct {
	Time time.Time
}

// MarshalJSON returns a slice of bytes in the form
// "seconds.milliseconds".
func (t *gelfTime) MarshalJSON() ([]byte, error) {
	asString := fmt.Sprintf("%.6f", float64(t.Time.UnixNano())/1e9)
	return []byte(asString), nil
}

// UnmarshalJSON parses a gelfTime struct from a
// "seconds.milliseconds" byte slice.
func (t *gelfTime) UnmarshalJSON(raw []byte) (err error) {
	var secRaw, nsecRaw []byte
	sep := byte('.')
	inFrac := false
	for _, b := range raw {
		if inFrac {
			nsecRaw = append(nsecRaw, b)
		} else if b == sep {
			inFrac = true
		} else {
			secRaw = append(secRaw, b)
		}
	}

	var sec, nsec int64

	err = json.Unmarshal(secRaw, &sec)
	if err != nil {
		return
	}

	err = json.Unmarshal(nsecRaw, &nsec)
	if err != nil {
		return
	}
	nsec *= 1000

	*t = gelfTime{time.Unix(sec, nsec)}
	return
}

// Message meets the Graylog2 Extended Log Format.
// http://graylog2.org/gelf#specs
type Message struct {
	Version          string                 `json:"version"`
	Host             string                 `json:"host"`
	ShortMessage     string                 `json:"short_message"`
	FullMessage      string                 `json:"full_message,omitempty"`
	Timestamp        *gelfTime              `json:"timestamp"`
	Level            Level                  `json:"level"`
	AdditionalFields string                 `json:",omitempty"`
	additional       map[string]interface{} `json:"a,omitempty"`
}

// Remote is a type for message destination configuration
type Remote int

const (
	RemoteStdout Remote = iota
	RemoteStderr
	RemoteUdp
)

var reservedFields = []string{"version", "host", "short_message", "full_message", "timestamp", "level", "_id"}

var host = ""
var remote Remote

func init() {
	remote = RemoteStdout

	var err error
	host, err = os.Hostname()
	if err != nil {
		host = "localhost"
	}
}

func SetRemote(r Remote) (err error) {
	if r == RemoteStdout {
		remote = r
	} else if r == RemoteStderr {
		remote = r
	} else if r == RemoteUdp {
		return errors.New("UDP not yet implemented")
	} else {
		return errors.New("Invalid GELF remote")
	}
	return nil
}

// NewMessage returns a new Graylog2 Extended Log Format message.
func NewMessage(l Level, short string, full string) *Message {
	a := make(map[string]interface{})

	return &Message{
		Version:      GELFVersion,
		Host:         host,
		ShortMessage: short,
		FullMessage:  full,
		Timestamp:    &gelfTime{time.Now()},
		Level:        l,
		additional:   a,
	}
}

func typeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// Add will add additional fields to a message in the form of a key and value
// pair. Values can be of JavaScript string or number type.
func (m *Message) Add(key string, value interface{}) error {
	// Verify additional fields against reserved field names.
	// If field is not reserved, add to message.
	for _, rf := range reservedFields {
		if key == rf {
			return fmt.Errorf("Invalid field[%s]", key)
		}
	}

	// Verify value is a JavaScript string or number.
	if typeOf(value) != "string" && typeOf(value) != "float64" && typeOf(value) != "int" {
		return fmt.Errorf("Invalid field type[%s]", typeOf(value))
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
	if len(m.additional) == 0 {
		baseMessageFields, _ := json.Marshal(m)
		return string(baseMessageFields)
	}

	// Maps do not marshal to JSON as top-level objects.
	// To work around we marshal the map of additional fields, modify the string
	// and append to the outbound JSON encoded struct.
	additionalFields, _ := json.Marshal(m.additional)
	filteredFields := strings.Replace(string(additionalFields[1:]), "\\\"", "\"", -1)

	baseMessageFields, _ := json.Marshal(m)
	trimBaseMessageFields := strings.TrimRight(string(baseMessageFields), "}")

	return trimBaseMessageFields + "," + filteredFields
}

// Send will currently print message's string to STDOUT
func (m *Message) Send() {
	if remote == RemoteStdout {
		fmt.Println(m.String())
	} else if remote == RemoteStderr {
		fmt.Fprintf(os.Stderr, "%s\n", m.String())
	} else if remote == RemoteUdp {
		// TODO: implement UDP
	}
}
