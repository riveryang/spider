package codec

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"github.com/riveryang/spider/models"
	"github.com/pkg/errors"
	"strings"
	"strconv"
	"github.com/astaxie/beego/orm"
)

var (
	ErrNotFoundItems = errors.New("Not found items")
)

type BangumiCodec struct {

}

func (b *BangumiCodec) Handler(doc *goquery.Document, source string) ([]interface{}, error) {
	lis := doc.Find("ul#browserItemList li")
	beego.Debug("Find browserItemList li length", lis.Length())
	if lis.Length() > 0 {
		bgms := make([]interface{}, lis.Length())
		lis.Each(func(i int, li *goquery.Selection) {
			bgm := new(models.Bangumi)
			title := li.Find("div.inner h3 a.l")
			if title.Length() > 0 {
				bgm.Title = title.Text()
			}

			aliasTitle := li.Find("div.inner h3 small.grey")
			if aliasTitle.Length() > 0 {
				bgm.AliasTitle = aliasTitle.Text()
			}

			subjectCover := li.Find("a.subjectCover")
			image, exists := subjectCover.Find("span.image img.cover").Attr("src")
			beego.Debug("Find subject image exists", exists)
			if exists {
				beego.Debug("Find subject image url", image)
				bgm.Image = image
			}

			href, exists := subjectCover.Attr("href")
			if exists {
				beego.Debug("Find subject of", href)
				subject, err := b.newDocument(source, href, 0)
				bgmId, err := strconv.Atoi(href[strings.LastIndex(href, "/")+1:])
				if err != nil {
					beego.Error(err)
					return
				}

				beego.Debug("Subject Id:", bgmId)
				bgm.Id = uint64(bgmId)
				if err != nil {
					beego.Error(err)
					return
				}

				subType := subject.Find("div#wrapperNeue div#headerSubject h1.nameSingle small.grey")
				if subType.Length() > 0 {
					bgm.Type = subType.Text()
				}

				// info
				infobox := subject.Find("div.mainWrapper div#columnSubjectHomeA div#bangumiInfo div.infobox ul#infobox li")
				beego.Debug("Find subject info length", infobox.Length())
				if infobox.Length() > 0 {
					infos := make([]*models.BangumiInfo, infobox.Length())
					infobox.Each(func(idx int, info *goquery.Selection) {
						tip := info.Find("span.tip")
						var tipText string
						if tip.Length() > 0 {
							tipText = tip.Text()
							beego.Debug("Find subject info tip", tipText)
						}

						persons := info.Find("a.l")
						beego.Debug("Find subject info person length", persons.Length())
						if persons.Length() > 0 {
							texts := make([]string, persons.Length())
							persons.Each(func(personIdx int, p *goquery.Selection) {
								pName := p.Text()
								beego.Debug("Find subject info person name", pName)
								texts[personIdx] = pName
							})

							info := new(models.BangumiInfo)
							info.Tip = tipText
							info.Infos = strings.Join(texts, ",")
							info.Bangumi = bgm
							infos[idx] = info
						} else {
							texts := make([]string, 1)
							if len(info.Nodes) > 0 {
								texts[0] = info.Nodes[0].LastChild.Data
								beego.Debug("Find subject info tip value", texts[0])
							}

							info := new(models.BangumiInfo)
							info.Tip = tipText
							info.Infos = strings.Join(texts, "|")
							info.Bangumi = bgm
							infos[idx] = info
						}
					})

					bgm.Info = infos
				}

				// summary
				summary := subject.Find("div.mainWrapper div#columnSubjectHomeB div#columnSubjectInHomeB div#subject_detail div#subject_summary")
				if summary.Length() > 0 {
					bgm.Summary = summary.Text()
				}

				// characters
				characters, err := goquery.NewDocument(source + href + "/characters")
				beego.Debug("Find characters")
				if err != nil {
					beego.Error(err)
					return
				}

				odd := characters.Find("div.mainWrapper div#columnInSubjectA div.light_odd")
				beego.Debug("Find characters odd length", odd.Length())
				if odd.Length() > 0 {
					crts := make([]*models.Character, odd.Length())
					odd.Each(func(oddIdx int, o *goquery.Selection) {
						crt := new(models.Character)
						cName := o.Find("div.clearit h2 a")
						beego.Debug("Find characters name length:", cName.Length())
						if cName.Length() > 0 {
							//chref, exists := cName.Attr("href")
							//if exists {
							//	characterId, err := strconv.Atoi(chref[strings.LastIndex(chref, "/")+1:])
							//	if err != nil {
							//		beego.Error(err)
							//		return
							//	}
							//
							//	beego.Debug("Character Id:", characterId)
							//	crt.Id = uint64(characterId)
							//}

							name := cName.Text()
							beego.Debug("Find characters name:", name)
							crt.Name = name
						}

						caName := o.Find("div.clearit h2 span")
						beego.Debug("Find characters alias name length:", caName.Length())
						if caName.Length() > 0 {
							aliasName := caName.Text()
							beego.Debug("Find characters alias name:", aliasName)
							crt.AliasName = aliasName
						}

						badgeJob := o.Find("div.clearit div.crt_info span.badge_job")
						beego.Debug("Find characters badge job length:", badgeJob.Length())
						if badgeJob.Length() > 0 {
							bj := badgeJob.Text()
							beego.Debug("Find characters badge job:", bj)
							crt.BadgeJob = bj
						}

						seiyuus := o.Find("div.clearit div.actorBadge.clearit p a.l")
						beego.Debug("Find characters seiyuu length", seiyuus.Length())
						if seiyuus.Length() > 0 {
							seiyuuStrs := make([]string, seiyuus.Length())
							seiyuus.Each(func(sidx int, seiyuu *goquery.Selection) {
								seiyuuStrs[sidx] = seiyuu.Text()
							})

							beego.Debug("Find characters seiyuu", seiyuuStrs)
							crt.Seiyuu = strings.Join(seiyuuStrs, ",")
						}

						crt.Bangumi = bgm
						crts[oddIdx] = crt
					})

					bgm.Characters = crts
				}
			}

			bgms[i] = bgm
		})

		return bgms, nil
	}

	return nil, ErrNotFoundItems
}

func (b *BangumiCodec) newDocument(source, href string, errorCount int) (*goquery.Document, error) {
	doc, err := goquery.NewDocument(source + href)
	if err != nil {
		return nil, err
	}

	if doc == nil {
		if errorCount < 3 {
			errorCount++
			return b.newDocument(source, href, errorCount)
		} else {
			return nil, errors.New("Reader document timeout of " + source + href)
		}
	}

	return doc, nil
}

func (b *BangumiCodec) Save(topics []interface{}) (int64, error) {
	o := orm.NewOrm()
	var size int64
	for _, item := range topics {
		topic := item.(*models.Bangumi)
		q := new(models.Bangumi)
		q.Id = topic.Id
		err := o.Read(q, "Id")
		if err != nil && err == orm.ErrNoRows {
			o.Begin()
			count, err := o.Insert(topic)
			if err != nil {
				o.Rollback()
				return 0, err
			}

			if len(topic.Info) > 0 {
				_, err = o.InsertMulti(100, topic.Info)
				if err != nil {
					o.Rollback()
					return 0, err
				}
			}

			if len(topic.Characters) > 0 {
				for _, c := range topic.Characters {
					_, err = o.Insert(c)
					if err != nil {
						o.Rollback()
						return 0, err
					}
				}

			}

			if count > 0 {
				size += 1
			}
			o.Commit()
		} else if err != nil {
			return 0, err
		} else {
			beego.Debug("Existing Rows of Id [", topic.Id, "]")
			continue;
		}
	}

	return size, nil
}
