package refrigerator

import "time"

type Beer struct {
	ID        int64
	Title     string
	ABV       string
	ExpiresAt time.Time
}
