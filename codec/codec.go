package codec

import "github.com/PuerkitoBio/goquery"

type Decoder interface {
	Decode(doc *goquery.Document, baseLink string) ([]interface{}, error)
}