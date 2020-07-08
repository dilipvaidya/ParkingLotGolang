package utils

import "time"

const (
	SPACE = " "
)

//CurrentTimeInMilli returns current system time in milli sec from epoch
func CurrentTimeInMilli() int64 {
	return time.Now().UnixNano() / 1e6
}
