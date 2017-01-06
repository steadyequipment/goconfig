package goutility

import (
	"time"
)

// NowFileString the current date and time as a file name safe string
func NowFileString() string {
	now := time.Now()
	return now.Format("2006-01-02_15-04-05-000Z")
}
