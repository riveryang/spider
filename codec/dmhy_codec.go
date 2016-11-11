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
	"strings"
	"github.com/pkg/errors"
	"strconv"
	"time"
	"github.com/riveryang/spider/models"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
)

var (
	ErrNotFoundPage = errors.New("Not found page")
)

type DmhyTopicCodec struct {

}

func (c *DmhyTopicCodec) Handler(doc *goquery.Document, source string) ([]interface{}, error) {
	var err error
	table := doc.Find("table#topic_list tbody tr")
	if table.Length() > 0 {
		items := doc.Find("table#topic_list tbody tr")
		topics := make([]interface{}, items.Length())
		items.Each(func(i int, s *goquery.Selection) {
			t := new(models.DmhyTopic)
			s.Find("td").Each(func(idx int, item *goquery.Selection) {
				switch idx {
				case 0:
					t.Time, err = time.Parse("2006/01/02 15:04", execText(item.Find("span").Text()))
					if err != nil {
						beego.Error(err)
						return
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
							t.Link = source + execText(href)
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

			h := md5.New()
			h.Write([]byte(t.Time.String() + t.Type + t.Tag + t.Link + t.Title + t.Magnet + t.Size))
			t.Md5 = hex.EncodeToString(h.Sum(nil))
			topics[i] = t
		})

		return topics, nil
	} else {
		return nil, ErrNotFoundPage
	}
}

func execText(text string) string {
	return strings.Replace(strings.Replace(strings.Trim(text, " "), "\t", "", -1), "\n", "", -1)
}