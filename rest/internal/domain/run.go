package domain

import "time"

type Run struct {
	ID            string
	DistanceInKm  float64
	TimeInMinutes int32
	Timestamp     time.Time
}
