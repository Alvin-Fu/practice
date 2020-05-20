package main

import (
    "github.com/bitly/go-simplejson"
    "os"
    "fmt"
    )

var jsonStr = `{
	"rc": 0,
	"result": {
        "round":[{
            "game": {
                "table": 1,
                "userCount": 6
            },
            "rank":{
                "weedOut": true
            },
            "alloc":{
                "type": 1
            }
	    },{
            "game": {
                "table": 1,
                "userCount": 6
            },
            "rank":{
                "weedOut": true
            },
            "alloc":{
                "type": 1
            }
	    }]
    }
}`

func main(){
    res, err := simplejson.NewJson([]byte(jsonStr))
    if err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
    result, ok := res.CheckGet("result")
    if !ok {
        os.Exit(-1)
    }
    levels := []string{"result"}
    for _, object := range []string{"game", "rank", "alloc"} {
        found := false
        for i := range levels {
            tl := levels[i:]
            p := append(tl, object)
            _, err := res.GetPath(p...).Map()
            if err == nil {
               found = true
               break
            }
        }
        if found {
            continue
        }
        keys := append(levels, "round")
        json := res.GetPath(keys...)
        rounds, err := json.Array()
        if err != nil {
            fmt.Println(err)
            return
        }
        rc := len(rounds) // rounds count already check above
        for i := 0; i < rc; i++ {
            fmt.Println(object, rounds[i].(map[string]interface{})[object])
            o := json.Get(object)
            if o == nil {
                return
            }
        }
    }
    fmt.Println(result)
}
