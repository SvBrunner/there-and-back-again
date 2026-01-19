package domain

import "time"

type Run struct {
	Distance      float64
	TimeInMinutes int32
	Timestamp     time.Time
}
