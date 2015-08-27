// Copyright © 2014-2015, Civis Analytics

package gelf

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestAdditionalFields(t *testing.T) {
	testMessage := NewMessage(LevelInfo, "This is a short test message.", "This is a long test message.\nIt includes multiple lines.")

	testMessage.Add("_StringType", "This is a string.")
	testMessage.Add("_IntegerType", 31)
	testMessage.Add("_FloatType", 1.61803398875)
	fmt.Println(testMessage.String())
}

func TestInvalidFieldNames(t *testing.T) {
	testMessage := NewMessage(LevelInfo, "This is a short test message.", "This is a long test message.\nIt includes multiple lines.")

	testMessage.Add("_id", "This is an invalid additional field.")
	testMessage.Add("host", "This is an invalid additional field.")
	testMessage.Add("valid", "This is a valid additional field without an underscore.")
	testMessage.Add("timestamp", 31)

	testMessage.Send()
}

func BenchmarkMessageCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewMessage(LevelInfo, "This is a short test message.", "This is a long test message.")
	}
}

func TestSerialization(t *testing.T) {
	msg := Message{
		Version:      "1.0",
		Host:         "snake farm",
		ShortMessage: "short snake",
		FullMessage:  "tall snake",
		Timestamp:    &gelfTime{time.Unix(1234567, 890123000)},
		Level:        6,
	}
	asJson, err := json.Marshal(msg)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	expectedJson := `{"version":"1.0","host":"snake farm","short_message":"short snake","full_message":"tall snake","timestamp":1234567.890123,"level":6}`
	if string(asJson) != expectedJson {
		t.Errorf("Unexpected marshalling. Expected %s, got %s", expectedJson, string(asJson))
	}

	var parsed Message
	err = json.Unmarshal(asJson, &parsed)

	if parsed.Version != msg.Version {
		t.Errorf("Unexpected un-marshalling. Expected %s, got %s", msg.Version, parsed.Version)
	}
	if parsed.Host != msg.Host {
		t.Errorf("Unexpected un-marshalling. Expected %s, got %s", msg.Host, parsed.Host)
	}
	if parsed.ShortMessage != msg.ShortMessage {
		t.Errorf("Unexpected un-marshalling. Expected %s, got %s", msg.ShortMessage, parsed.ShortMessage)
	}
	if parsed.FullMessage != msg.FullMessage {
		t.Errorf("Unexpected un-marshalling. Expected %s, got %s", msg.FullMessage, parsed.FullMessage)
	}
	if *parsed.Timestamp != *msg.Timestamp {
		t.Errorf("Unexpected un-marshalling. Expected %s, got %s", *msg.Timestamp, *parsed.Timestamp)
	}
	if parsed.Level != msg.Level {
		t.Errorf("Unexpected un-marshalling. Expected %s, got %s", msg.Level, parsed.Level)
	}
}
