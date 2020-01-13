package gjson

import (
	"fmt"
	"github.com/tidwall/gjson"
)

const json = `{"name":[{"first":"Janet","last":"Prichard"}, {"abc": 1, "e" : 2}],"age":47}`

func GJson() {
	rue := gjson.Get(json, "name")
	fmt.Println(rue.Type.String(), rue.String())
}
