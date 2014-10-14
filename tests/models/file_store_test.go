package test

import (
	"github.com/Marktown/frontend/models"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// TestMain is a sample to run an endpoint test
func TestMain(t *testing.T) {
	Convey("Subject: FSFileStore\n", t, func() {
		Convey("Read", func() {
			bytes, err := models.FSFileStore().ReadFile("../assets/testfile.txt")
			So(string(bytes), ShouldEqual, "this is the textfile\n")
			So(err, ShouldBeNil)
		})
	})
}
