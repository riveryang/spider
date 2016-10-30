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

package codec

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/riveryang/spider/topic"
	"log"
	"strings"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

type DmhyTopicCodec struct {

}

func (c *DmhyTopicCodec) Decode(doc *goquery.Document, baseLink string) ([]interface{}, error) {
	var err error
	table := doc.Find("table#topic_list tbody tr")
	if table.Length() > 0 {
		items := doc.Find("table#topic_list tbody tr")
		topics := make([]interface{}, items.Length())
		items.Each(func(i int, s *goquery.Selection) {
			t := topic.Topic{}
			s.Find("td").Each(func(idx int, item *goquery.Selection) {
				switch idx {
				case 0:
					t.Time, err = time.Parse("2006/01/02 15:04", execText(item.Find("span").Text()))
					if err != nil {
						log.Fatal(err)
					}
				case 1:
					t.Type = execText(item.Find("a font").Text())
				case 2:
					tag := item.Find("span.tag")
					if tag.Length() > 0 {
						t.Tag = execText(tag.Find("a").Text())
					}

					title := item.Find("a[target='_blank']")
					if title.Length() > 0 {
						href, exists := title.Attr("href")
						if exists {
							t.Link = baseLink + execText(href)
						}

						t.Title = execText(title.Text())
					}
				case 3:
					magnet := item.Find("a")
					href, exists := magnet.Attr("href")
					if exists {
						t.Magnet = execText(href)
					}
				case 4:
					t.Size = execText(item.Text())
				case 5:
					t.Seed, err = strconv.Atoi(execText(item.Find("span").Text()))
					if err != nil {
						t.Seed = 0
					}
				case 6:
					t.Downloads, err = strconv.Atoi(execText(item.Find("span").Text()))
					if err != nil {
						t.Downloads = 0
					}
				case 7:
					t.Complete, err = strconv.Atoi(execText(item.Text()))
					if err != nil {
						t.Complete = 0
					}
				default:
					break
				}
			});

			topics[i] = t
		})

		return topics, nil
	} else {
		return nil, errors.New("Not found page")
	}
}

func execText(text string) string {
	return strings.Replace(strings.Replace(strings.Trim(text, " "), "\t", "", -1), "\n", "", -1)
}