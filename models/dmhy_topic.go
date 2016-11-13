package models

import "time"

type DmhyTopic struct {
	Topic
	Md5 string `json:"md5,omitempty" orm:"size(32);unique"`
	Time time.Time `json:"time,omitempty" orm:"type(datetime)"`
	Type string `json:"type,omitempty" orm:"size(64);null"`
	Tag string `json:"tag,omitempty" orm:"size(64);null"`
	Link string `json:"link,omitempty" orm:"size(256)"`
	Title string `json:"title,omitempty" orm:"size(512)"`
	Magnet string `json:"magnet,omitempty" orm:"type(text)"`
	Size string `json:"size,omitempty" orm:"size(32)"`
	Seed int `json:"seed,omitempty" orm:"digits(6)"`
	Downloads int `json:"downloads,omitempty" orm:"digits(6)"`
	Complete int `json:"complete,omitempty" orm:"digits(6)"`

}

