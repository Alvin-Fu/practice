package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
			_ "practice/websvr/routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"strings"
	"fmt"
)

func init() {
	path, _ := os.Getwd()
	index := strings.LastIndex(path, "/")
	path = path[:index]
	fmt.Println(index)
	beego.TestBeegoInit(path)
}


// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
	        Convey("Status Code Should Be 200", func() {
	                So(w.Code, ShouldEqual, 200)
	        })
	        Convey("The Result Should Not Be Empty", func() {
	                So(w.Body.Len(), ShouldBeGreaterThan, 0)
	        })
	})
}

