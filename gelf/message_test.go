// Copyright © 2014-2015, Civis Analytics

package gelf

import (
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

func TestGelfTime(t *testing.T) {
	gt := gelfTime{time.Unix(1234567, 890123000)}
	expected := "1234567.890123"
	if gt.String() != expected {
		t.Errorf("Error formatting string. Expected %s, got %s", expected, gt.String())
	}
}

func BenchmarkMessageCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewMessage(LevelInfo, "This is a short test message.", "This is a long test message.")
	}
}
