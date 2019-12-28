package crypto

import "time"

type Info struct {
	name        string
	symbol      string
	price       float32
	lastUpdated time.Time
}
