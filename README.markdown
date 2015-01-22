# Go GELF

A Graylog2 Extended Message Format library for Go.

This library aims to be orthogonal to the Go standard library `log`. Future versions will include the ability to compress message and transmit via UDP to a Graylog2 node.

## Usage

```
package main

import (
	"fmt"
	"github.com/civisanalytics/gogelf/gelf"
)

func main() {
	gelfMessage, err := gelf.Info("This is a short message.", "This is a long message.\nIt includes multiple lines.")
	if err != nil {
		fmt.Error("Unable to create Graylog message.")
	}

	// Add additional fields to the returned message object.
	gelfMessage.Add("_HTTPMethod", "GET")
	gelfMessage.Add("_ResponseCode", 301)

	// Print the message in JSON format to stdout.
	fmt.Println(testMessage.String())
}
```
