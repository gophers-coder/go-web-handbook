package log

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

func TestLog(test *testing.T) {
	Printf("Hello World, %s", "Golang")

}

func TestLogWithNew(test *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, "My Project ", log.LstdFlags)
	logger.Printf("Hello World, %s", "Golang")
	fmt.Println(&buf)

}
