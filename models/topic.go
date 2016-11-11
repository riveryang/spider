// Copyright © 2015-2016 River Yang <comicme_yanghe@nanoframework.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"time"
)

type Topic struct {
	Id uint64 `json:"id,omitempty" orm:"auto"`
	Md5 string `json:"md5,omitempty" orm:"size(32);unique"`
}

type DmhyTopic struct {
	Topic
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

