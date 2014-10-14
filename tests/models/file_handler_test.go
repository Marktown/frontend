package test

import (
	"github.com/Marktown/frontend/models"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// TestMain is a sample to run an endpoint test
func TestMain(t *testing.T) {
	Convey("Subject: FSFileHandler\n", t, func() {
		Convey("Read", func() {
			content, err := models.FileHandler().Read(&models.File{"../assets/testfile.txt", ""})
			So(content, ShouldEqual, "this is the textfile\n")
			So(err, ShouldBeNil)
		})
	})
}
