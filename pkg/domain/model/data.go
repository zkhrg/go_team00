package model

import "time"

// Anomaly — это доменная модель, представляющая аномалию.
type Anomaly struct {
	ID        uint
	SessionID string
	Frequency float64
	Timestamp time.Time
}
