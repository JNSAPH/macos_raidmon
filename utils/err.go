package utils

import (
	"fmt"
	"time"
)

// return an error with a long message
func TestError() error {
	currentDate := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Errorf("this is a test error. It's a long message that should be wrapped to multiple lines. This is a test error. Today is %s", currentDate)
}
