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
			topic := topic.Topic{}
			s.Find("td").Each(func(idx int, item *goquery.Selection) {
				switch idx {
				case 0:
					topic.Time, err = time.Parse("2006/01/02 15:04", execText(item.Find("span").Text()))
					if err != nil {
						log.Fatal(err)
					}
				case 1:
					topic.Type = execText(item.Find("a font").Text())
				case 2:
					tag := item.Find("span.tag")
					if tag.Length() > 0 {
						topic.Tag = execText(tag.Find("a").Text())
					}

					title := item.Find("a[target='_blank']")
					if title.Length() > 0 {
						href, exists := title.Attr("href")
						if exists {
							topic.Link = baseLink + execText(href)
						}

						topic.Title = execText(title.Text())
					}
				case 3:
					magnet := item.Find("a")
					href, exists := magnet.Attr("href")
					if exists {
						topic.Magnet = execText(href)
					}
				case 4:
					topic.Size = execText(item.Text())
				case 5:
					topic.Seed, err = strconv.Atoi(execText(item.Find("span").Text()))
					if err != nil {
						topic.Seed = 0
					}
				case 6:
					topic.Downloads, err = strconv.Atoi(execText(item.Find("span").Text()))
					if err != nil {
						topic.Downloads = 0
					}
				case 7:
					topic.Complete, err = strconv.Atoi(execText(item.Text()))
					if err != nil {
						topic.Complete = 0
					}
				default:
					break
				}
			});

			topics[i] = topic
		})

		return topics, nil
	} else {
		return nil, errors.New("Not found page")
	}
}

func execText(text string) string {
	return strings.Replace(strings.Replace(strings.Trim(text, " "), "\t", "", -1), "\n", "", -1)
}