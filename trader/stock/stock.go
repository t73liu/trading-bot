package stock

import "time"

type Info struct {
	company           string
	symbol            string
	price             float32
	outstandingShares uint32
	lastUpdated       time.Time
}
