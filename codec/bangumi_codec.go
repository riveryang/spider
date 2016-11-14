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
	o orm.Ormer
	notFound int
}

func (b *BangumiCodec) Handler(doc *goquery.Document, source string) ([]interface{}, error) {
	isLocked := doc.Find("div#wrapperNeue div#headerSubject div.tipIntro div.inner h3")
	if isLocked.Length() > 0 {
		if isLocked.Text() == "条目已锁定" {
			beego.Warn("Item is locked")
			return nil, nil
		}
	}

	bgm := new(models.Bangumi)
	var href string
	var exists bool
	override := doc.Find("div#wrapperNeue div#headerSubject div.subjectNav ul.navTabs li a.focus")
	if override.Length() > 0 {
		href, exists = override.Attr("href")
		if exists {
			bgmId, err := strconv.Atoi(href[strings.LastIndex(href, "/")+1:])
			if err != nil {
				return nil, err
			}

			beego.Debug("Subject Id:", bgmId)
			bgm.Id = uint64(bgmId)

			if b.o == nil {
				b.o = orm.NewOrm()
			}

			q := new(models.Bangumi)
			q.Id = bgm.Id
			err = b.o.Read(q, "Id")
			if err != nil && err != orm.ErrNoRows {
				return nil, err
			}
		}

		b.notFound = 0
	} else {
		if b.notFound > 10 {
			return nil, ErrNotFoundItems
		} else {
			b.notFound++
		}
	}

	title := doc.Find("div#wrapperNeue div#headerSubject h1.nameSingle a")
	if title.Length() > 0 {
		bgm.Title = title.Text()
		aliasTitle, exists := title.Attr("title")
		if exists && aliasTitle != "" {
			bgm.AliasTitle = aliasTitle
		}
	}

	subType := doc.Find("div#wrapperNeue div#headerSubject h1.nameSingle small.grey")
	if subType.Length() > 0 {
		bgm.Type = subType.Text()
	}

	img := doc.Find("div#wrapperNeue div.mainWrapper div#columnSubjectHomeA div#bangumiInfo div.infobox div[align='center'] a")
	if img.Length() > 0 {
		imgHref, exists := img.Attr("href")
		if exists {
			bgm.Image = imgHref
		}
	}

	// info
	infobox := doc.Find("div.mainWrapper div#columnSubjectHomeA div#bangumiInfo div.infobox ul#infobox li")
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
			if persons.Length() > 0 {
				texts := make([]string, persons.Length())
				persons.Each(func(personIdx int, p *goquery.Selection) {
					pName := p.Text()
					beego.Debug("Find subject info person name", pName)
					texts[personIdx] = pName
				})

				info := new(models.BangumiInfo)
				info.Tag = tipText
				info.Values = strings.Join(texts, ",")
				info.Bangumi = bgm
				infos[idx] = info
			} else {
				texts := make([]string, 1)
				if len(info.Nodes) > 0 {
					texts[0] = info.Nodes[0].LastChild.Data
					beego.Debug("Find subject info tip value", texts[0])
				}

				info := new(models.BangumiInfo)
				info.Tag = tipText
				info.Values = strings.Join(texts, ",")
				info.Bangumi = bgm
				infos[idx] = info
			}
		})

		bgm.Info = infos
	}

	// summary
	summary := doc.Find("div.mainWrapper div#columnSubjectHomeB div#columnSubjectInHomeB div#subject_detail div#subject_summary")
	if summary.Length() > 0 {
		bgm.Summary = summary.Text()
	}

	score := doc.Find("div.mainWrapper div#columnSubjectHomeB div#columnSubjectInHomeB div.rr div.SidePanel div.global_score span.number")
	if score.Length() > 0 {
		bgm.Score = score.Text()
	}

	scoreDesc := doc.Find("div.mainWrapper div#columnSubjectHomeB div#columnSubjectInHomeB div.rr div.SidePanel div.global_score span.description")
	if scoreDesc.Length() > 0 {
		bgm.ScoreDesc = scoreDesc.Text()
	}

	rank := doc.Find("div.mainWrapper div#columnSubjectHomeB div#columnSubjectInHomeB div.rr div.SidePanel div.global_score div small.alarm")
	if rank.Length() > 0 {
		bgm.Rank = strings.Replace(rank.Text(), "#", "", -1)
	}

	// characters
	characters, err := b.newDocument(source, href + "/characters", 0)
	beego.Debug("Find characters")
	if err != nil {
		return nil, err
	}

	odd := characters.Find("div.mainWrapper div#columnInSubjectA div.light_odd")
	if odd.Length() > 0 {
		crts := make([]*models.Character, odd.Length())
		odd.Each(func(oddIdx int, o *goquery.Selection) {
			crt := new(models.Character)
			cName := o.Find("div.clearit h2 a")
			if cName.Length() > 0 {
				name := cName.Text()
				beego.Debug("Find characters name:", name)
				crt.Name = name
			}

			caName := o.Find("div.clearit h2 span")
			if caName.Length() > 0 {
				aliasName := caName.Text()
				beego.Debug("Find characters alias name:", aliasName)
				crt.AliasName = aliasName
			}

			badgeJob := o.Find("div.clearit div.crt_info span.badge_job")
			if badgeJob.Length() > 0 {
				bj := badgeJob.Text()
				beego.Debug("Find characters badge job:", bj)
				crt.BadgeJob = bj
			}

			seiyuus := o.Find("div.clearit div.actorBadge.clearit p a.l")
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

	bgms := make([]interface{}, 1)
	bgms[0] = bgm
	return bgms, nil
}

func (b *BangumiCodec) newDocument(source, href string, errorCount int) (*goquery.Document, error) {
	beego.Debug("Reader document of", source + href)
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
	if topics == nil {
		return 0, nil
	}

	if b.o == nil {
		b.o = orm.NewOrm()
	}

	var size int64
	for _, item := range topics {
		topic := item.(*models.Bangumi)
		q := new(models.Bangumi)
		q.Id = topic.Id
		err := b.o.Read(q, "Id")
		if err != nil && err == orm.ErrNoRows {
			b.o.Begin()
			count, err := b.o.Insert(topic)
			if err != nil {
				b.o.Rollback()
				return 0, err
			}

			if len(topic.Info) > 0 {
				_, err = b.o.InsertMulti(100, topic.Info)
				if err != nil {
					b.o.Rollback()
					return 0, err
				}
			}

			if len(topic.Characters) > 0 {
				for _, c := range topic.Characters {
					_, err = b.o.Insert(c)
					if err != nil {
						b.o.Rollback()
						return 0, err
					}
				}

			}

			if count > 0 {
				size += 1
			}
			b.o.Commit()
		} else if err != nil {
			return 0, err
		} else {
			beego.Debug("Existing Rows of Id [", topic.Id, "]")
			continue;
		}
	}

	return size, nil
}
