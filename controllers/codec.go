package controllers

import (
	"github.com/astaxie/beego"
	"github.com/riveryang/spider/codec"
	"github.com/PuerkitoBio/goquery"
	"github.com/riveryang/spider/result"
	"github.com/riveryang/spider/http"
	"strconv"
	"time"
	"math/rand"
)

var stop = make(map[string] chan bool)

type CodecController struct {
	beego.Controller
}

// @Description Test: http://localhost:8080/v1/topic/dmhy/1?source=http://share.dmhy.org&topic=topics/list/page
// @Title Get
// @Param codec path string true "The codec type"
// @Param page path int true "The source page index"
// @Param source query string true "The source site"
// @Param topic query string true "The source site topic token"
// @Success 200 {object} result.Result
// @router /:codec/:page [get]
func (c *CodecController) Get() {
	defer c.ServeJSON()
	codecType := c.GetString(":codec")
	page, err := c.GetInt(":page")
	if err != nil {
		c.Data["json"] = &result.Result{Status: http.BAD_REQUEST, Message: err.Error()}
		return
	} else if page < 1 {
		c.Data["json"] = &result.Result{Status: http.BAD_REQUEST, Message: "Unknown page index"}
		return
	}

	source := c.GetString("source")
	topic := c.GetString("topic")
	if codecType != "" && source != "" {
		c.Data["json"] = exec(source, topic, codecType, page)
	}
}

// @Description Test: http://localhost:8080/v1/topic/dmhy/auto/1?source=http://share.dmhy.org&topic=topics/list/page
// @Title Auto
// @Param codec path string true "The codec type"
// @Param source query string true "The source site"
// @Param topic query string true "The source site topic token"
// @Success 200 {object} result.Result
// @router /:codec/auto/:page [get]
func (c *CodecController) Auto() {
	defer c.ServeJSON()
	codecType := c.GetString(":codec")
	source := c.GetString("source")
	topic := c.GetString("topic")
	page, err := c.GetInt(":page")
	if err != nil || page < 1 {
		page = 1
	}

	if codecType != "" && source != "" {
		stop[codecType] = make(chan bool)
		go func () {
			nowPage := page
			for {
				select {
				case <- stop[codecType]:
					beego.Info("Last reader page of", nowPage)
					delete(stop, codecType)
					return
				default:
					ret := exec(source, topic, codecType, nowPage)
					if ret.Status != http.OK {
						beego.Warn(ret.Message + ", Stoped auto reader")
						delete(stop, codecType)
						return
					}

					dur := time.Second + time.Millisecond * time.Duration(rand.Int31n(2000))
					beego.Info("Now reader page of href [", source + "/" + topic + strconv.Itoa(nowPage), "]. next reader of", dur)
					nowPage++
					time.Sleep(dur)
				}
			}
		}()
	}

	c.Data["json"] = &result.Result{Status: http.OK}
}

// @Description Test: http://localhost:8080/v1/topic/dmhy/stop
// @Title Stop
// @Param codec path string true "The codec type"
// @Success 200 {object} result.Result
// @router /:codec/stop [get]
func (c *CodecController) Stop() {
	defer c.ServeJSON()
	codecType := c.GetString(":codec")
	stopChan := stop[codecType]
	if stopChan != nil {
		stopChan <- true
		c.Data["json"] = &result.Result{Status: http.OK, Message: codecType +  " runner to be stop"}
	} else {
		c.Data["json"] = &result.Result{Status: http.FORBIDDEN, Message: "Not found stop chan"}
	}
}

func exec(source, topic, codecType string, page int) *result.Result {
	c, err := codec.WithCodec(codecType)
	if err != nil {
		return &result.Result{Status: http.BAD_REQUEST, Message: err.Error()}
	} else {
		doc, err := goquery.NewDocument(source + "/" + topic + strconv.Itoa(page))
		if err != nil {
			return &result.Result{Status: http.BAD_REQUEST, Message: err.Error()}
		}

		topics, err := c.Handler(doc, source)
		if err != nil {
			return &result.Result{Status: http.BAD_REQUEST, Message: err.Error()}
		}

		size, err := c.Save(topics)
		if err != nil {
			return &result.Result{Status: http.BAD_REQUEST, Message: err.Error()}
		}

		beego.Info("Insert size: ", size)
		return &result.Result{Status: http.OK, Data: size}
	}
}
