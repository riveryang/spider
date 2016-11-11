package controllers

import (
	"testing"
	"github.com/riveryang/spider/http"
	"github.com/riveryang/spider/db"
)

func init() {
	db.InitDB()
}

func TestCodecController_exec_success(t *testing.T) {
	source := "http://share.dmhy.org"
	topic := "topics/list/page"
	codecType := "dmhy"
	page := 1
	ret := exec(source, topic, codecType, page)
	if ret.Status != http.OK {
		t.Fatal("Exec return status not OK")
	}

}

func TestCodecController_exec_NotFoundCodec(t *testing.T) {
	source := "http://share.dmhy.org"
	topic := "topics/list/page"
	codecType := "nil_codec"
	page := 1
	ret := exec(source, topic, codecType, page)
	if ret.Status != http.BAD_REQUEST {
		t.Fatal("Assert not bad_request")
	}

	if ret.Message != "Not found codec" {
		t.Fatal(`Assert message not eq "Not found codec"`)
	}

}

func TestCodecController_exec_NotFoundPage(t *testing.T) {
	source := "http://share.dmhy.org"
	topic := "topics/list/page"
	codecType := "dmhy"
	page := 0
	ret := exec(source, topic, codecType, page)
	if ret.Status != http.BAD_REQUEST {
		t.Fatal("Assert not bad_request")
	}

	if ret.Message != "Not found page" {
		t.Fatal(`Assert message not eq "Not found page"`)
	}
}