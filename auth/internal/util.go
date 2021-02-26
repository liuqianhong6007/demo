package internal

import "time"

func NowUnix() int64 {
	return time.Now().Unix()
}
