package codec

import (
	"testing"
	"github.com/Unknwon/com"
)

func TestWithCodec_NotFoundCodec(t *testing.T) {
	_, err := WithCodec("nil")
	if err == nil {
		t.Fatal("Codec Finded")
	}

	if err != ErrNotFoundCodec {
		t.Fatal(err)
	}
}

func TestURLEncode(t *testing.T) {
	t.Log(com.UrlEncode("anime/browser?sort=date&page="))
}