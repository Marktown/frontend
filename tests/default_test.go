package test

import (
	_ "github.com/Marktown/frontend/routers"

	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func request(path string) *httptest.ResponseRecorder {
	r, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestMain", "Code[%d]\n%s", w.Code, w.Body.String())
	return w
}

// TestMain is a sample to run an endpoint test
func TestMain(t *testing.T) {
	Convey("Subject: Test Station Endpoint\n", t, func() {
		w := request("/")
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})

	Convey("Subject: List files\n", t, func() {
		w := request("/files")
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Body Should Render Json", func() {
			expectedBytes, _ := ioutil.ReadFile("tests/assets/api/files.json")
			So(w.Body.String(), ShouldContainSubstring, strings.TrimSpace(string(expectedBytes)))
		})
	})
}
