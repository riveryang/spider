package utils

import (
	"testing"
)

func TestSimilarity(t *testing.T) {
	str1 := "[110727] TVアニメ 神様ドォルズ OP＆ED「不完全燃焼／スイッチが入ったら」／石川智晶 (320K+BK)"
	arr := []string{
		"[傲嬌字幕組][TNDR][神様ドォルズ][01][BIG5][1080p] - 徵才內詳",
		"[110727] TVアニメ 神様ドォルズ OP＆ED「不完全燃焼／スイッチが入ったら」／石川智晶 (320K+BK)",
		"[烏賊發佈][EAC]TVアニメ「神様ドォルズ」OPEDテーマ 「不完全燃焼／スイッチが入ったら」／石川智晶 (ape+cue+png)",
		"[烏賊發佈]TVアニメ「神様ドォルズ」OPEDテーマ 「不完全燃焼／スイッチが入ったら」／石川智晶 (320K+BK)",
		"[R8-Audio] TVアニメ 神様ドォルズ OP&ED 「不完全燃焼／スイッチが入ったら」／石川智晶 (320kbps+Scans).zip",
		"[HKG字幕組][10月新番][Toaru_Kagaku_no_Railgun 科學超電磁炮][18][BIG5][848x480][RMVB]",
		"(81) 幪面超人電王 Kamen Rider Den-O 第11話 AVI 粵語配音 日語OP 外掛繁體字幕 [嚴禁轉載]",
		"【伊妹儿·异域字幕组】★[火影忍者364][Naruto364][big5][rmvb]",
	}

	for _, str2 := range arr {
		diff, err := Similarity(str1, str2)
		if err != nil {
			t.Error(err)
		} else {
			t.Log("Compare rate:", diff)
		}
	}
}
