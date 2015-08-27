// Copyright © 2014-2015, Civis Analytics

package gelf

import (
	"fmt"
	"math"
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

func TestTimestamp(t *testing.T) {
	now := float64(time.Now().Unix())
	testMessage := NewMessage(LevelInfo, "This is a short test message.", "This is a long test message.")
	if math.Abs(testMessage.Timestamp-now) > 1 {
		t.Errorf("Wanted a timestamp in seconds since epoch, got %f")
	}
}

func BenchmarkMessageCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewMessage(LevelInfo, "This is a short test message.", "This is a long test message.")
	}
}
