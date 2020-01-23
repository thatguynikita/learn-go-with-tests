package racer

import (
	"net/http"
	"time"
)

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

// Racer return the fastest of 2 URLs
func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration > bDuration {
		return b
	}
	return a
}
