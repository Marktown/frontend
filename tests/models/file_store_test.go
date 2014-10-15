package test

import (
	"github.com/Marktown/frontend/models"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// TestMain is a sample to run an endpoint test
func TestMain(t *testing.T) {
	Convey("Subject: FSFileStore\n", t, func() {
		Convey("Create", func() {
			err := models.FSFileStore().CreateFile("../assets/testfile_b.txt", "foo bar")
			So(err, ShouldBeNil)
			bytes, err := models.FSFileStore().ReadFile("../assets/testfile_b.txt")
			So(string(bytes), ShouldEqual, "foo bar")
			So(err, ShouldBeNil)
		})

		Convey("Read", func() {
			bytes, err := models.FSFileStore().ReadFile("../assets/testfile.txt")
			So(string(bytes), ShouldEqual, "this is the textfile\n")
			So(err, ShouldBeNil)
		})

		Convey("Update", func() {
			err := models.FSFileStore().UpdateFile("../assets/testfile_b.txt", "foo bar baz")
			So(err, ShouldBeNil)
			bytes, err := models.FSFileStore().ReadFile("../assets/testfile_b.txt")
			So(string(bytes), ShouldEqual, "foo bar baz")
			So(err, ShouldBeNil)
		})

	})
}
