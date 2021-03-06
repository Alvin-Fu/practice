// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en":      &dictionary{index: enIndex, data: enData},
		"zh_Hans": &dictionary{index: zh_HansIndex, data: zh_HansData},
		"zh_Hant": &dictionary{index: zh_HantIndex, data: zh_HantData},
	}
	fallback := language.MustParse("en")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"client request error": 1,
	"login multiple times": 0,
}

var enIndex = []uint32{ // 3 elements
	0x00000000, 0x00000015, 0x0000002a,
} // Size: 36 bytes

const enData string = "" + // Size: 42 bytes
	"\x02login multiple times\x02client request error"

var zh_HansIndex = []uint32{ // 3 elements
	0x00000000, 0x0000000d, 0x00000023,
} // Size: 36 bytes

const zh_HansData string = "\x02重复登录\x02客户端请求错误"

var zh_HantIndex = []uint32{ // 3 elements
	0x00000000, 0x0000000d, 0x00000023,
} // Size: 36 bytes

const zh_HantData string = "\x02重復登錄\x02客戶端請求錯誤"

// Total table size 220 bytes (0KiB); checksum: CCF91E74
