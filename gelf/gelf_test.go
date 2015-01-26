// Copyright © 2014-2015, Civis Analytics

package gelf

import (
	"fmt"
	"testing"
)

func TestPanicMessage(t *testing.T) {
	testMessage := Panic("This is a panic test message.", "This is a long panic test message.\nIt includes multiple lines.")

	fmt.Println(testMessage.String())
}
