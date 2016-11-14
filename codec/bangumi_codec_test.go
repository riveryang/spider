package codec

import (
	"testing"
	"github.com/PuerkitoBio/goquery"
	"github.com/docker/go/canonical/json"
	"github.com/astaxie/beego"
	"github.com/riveryang/spider/models"
	"sync"
	"github.com/riveryang/spider/db"
)

var bangumiSource = "http://bangumi.tv"
var bangumiTopic = "subject/"

var once sync.Once

func TestBangumiCodec_Handler(t *testing.T) {
	//beego.SetLevel(beego.LevelInformational)
	once.Do(db.InitDB)

	url := bangumiSource + "/" + bangumiTopic + "1"
	beego.Debug("URL:", url)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		t.Fatal(err)
	}

	c, err := WithCodec(models.BANGUMI_CODEC)
	if err != nil {
		t.Fatal(err)
	}

	topics, err := c.Handler(doc, bangumiSource)
	if err != nil {
		t.Fatal(err)
	}

	for _, topic := range topics {
		j, err := json.Marshal(topic)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(string(j))
	}
}