package codec

import "testing"

func TestWithCodec_NotFoundCodec(t *testing.T) {
	_, err := WithCodec("nil")
	if err == nil {
		t.Fatal("Codec Finded")
	}

	if err != ErrNotFoundCodec {
		t.Fatal(err)
	}
}
