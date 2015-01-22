// Copyright © 2014-2015, Civis Analytics

package gelf

import (
	"fmt"
	"testing"
)

func TestPanicMessage(t *testing.T) {
	testMessage, err := Panic("This is a panic test message.", "This is a long panic test message.\nIt includes multiple lines.")
	if err != nil {
		t.Error("Unable to create Graylog message.")
	}

	fmt.Println(testMessage.String())
}
