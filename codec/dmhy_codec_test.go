package codec

import (
	"testing"
	"github.com/PuerkitoBio/goquery"
)

var source = "http://share.dmhy.org"
var topic = "topics/list/page"
var codec = "dmhy"

func TestDmhyTopicCodec_Handler(t *testing.T) {
	doc, err := goquery.NewDocument(source + "/" + topic + "/1")
	if err != nil {
		t.Fatal(err)
	}

	codec, err := WithCodec(codec)
	if err != nil {
		t.Fatal(err)
	}

	topics, err := codec.Handler(doc, source)
	if err != nil {
		t.Fatal(err)
	}

	for _, topic := range topics {
		t.Log(topic)
	}
}

func TestDmhyTopicCodec_Handler_NotFoundPage(t *testing.T) {
	doc, err := goquery.NewDocument(source + "/" + topic + "/0")
	if err != nil {
		t.Fatal(err)
	}

	codec, err := WithCodec(codec)
	if err != nil {
		t.Fatal(err)
	}

	_, err = codec.Handler(doc, source)
	if err != nil {
		if err != ErrNotFoundPage {
			t.Fatal(err)
		}
	}
}
