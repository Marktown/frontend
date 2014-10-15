package test

import (
	"github.com/Marktown/frontend/models"

	"io"
	"strings"
	"testing"
	"text/scanner"

	. "github.com/smartystreets/goconvey/convey"
)

func readAll(reader io.Reader) (data string) {
	data = ""
	var s scanner.Scanner
	s.Init(reader)
	s.Whitespace = 1
	tok := s.Scan()
	for tok != scanner.EOF {
		data += s.TokenText()
		tok = s.Scan()
	}
	return
}

// TestMain is a sample to run an endpoint test
func TestMain(t *testing.T) {
	Convey("Subject: FSFileStore\n", t, func() {
		Convey("Create", func() {
			err := models.FSFileStore().CreateFile("../assets/testfile_b.txt", strings.NewReader("foo bar"))
			So(err, ShouldBeNil)
			reader, err := models.FSFileStore().ReadFile("../assets/testfile_b.txt")
			So(readAll(reader), ShouldEqual, "foo bar")
			So(err, ShouldBeNil)
		})

		Convey("Read", func() {
			reader, err := models.FSFileStore().ReadFile("../assets/testfile.txt")
			So(readAll(reader), ShouldEqual, "this is the textfile\n")
			So(err, ShouldBeNil)
		})

		Convey("Update", func() {
			err := models.FSFileStore().UpdateFile("../assets/testfile_b.txt", strings.NewReader("foo bar baz"))
			So(err, ShouldBeNil)
			reader, err := models.FSFileStore().ReadFile("../assets/testfile_b.txt")
			So(readAll(reader), ShouldEqual, "foo bar baz")
			So(err, ShouldBeNil)
		})

	})
}
