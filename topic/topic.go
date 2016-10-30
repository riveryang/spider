package topic

import "time"

type Topic struct {
	Time time.Time
	Type string
	Tag string
	Link string
	Title string
	Magnet string
	Size string
	Seed int
	Downloads int
	Complete int

}

