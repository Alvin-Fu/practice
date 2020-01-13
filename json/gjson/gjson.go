package gjson

import (
	"fmt"
	"github.com/tidwall/gjson"
)

const json = `{"name":[{"first":"Janet","last":"Prichard"}, {"abc": 1, "e" : 2}],"age":47}`

func GJson() {
	rue := gjson.Get(json, "name").Get("abc")
	fmt.Println(rue.Type.String(), rue.String())
}
func selector(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Println(v)
	case int64:
	default:

	}
}
