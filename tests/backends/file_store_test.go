package test

import (
	"github.com/Marktown/frontend/backends/file_system"

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
	Convey("Subject: FileStore\n", t, func() {
		fs := file_system.NewFileStore()
		fs.RootPath = "../../tmp/tests/fs_file_store/"
		Convey("Create", func() {
			So(fs.RootPath, ShouldEqual, "../../tmp/tests/fs_file_store/")
			err := fs.CreateFile("testfile_b.txt", strings.NewReader("foo bar"))
			So(err, ShouldBeNil)
			reader, err := fs.ReadFile("testfile_b.txt")
			So(readAll(reader), ShouldEqual, "foo bar")
			So(err, ShouldBeNil)
		})

		Convey("Read", func() {
			fs.RootPath = "../../tests/assets/"
			reader, err := fs.ReadFile("testfile.txt")
			So(readAll(reader), ShouldEqual, "this is the textfile\n")
			So(err, ShouldBeNil)
			fs.RootPath = "../../tmp/tests/fs_file_store/"
		})

		Convey("Update", func() {
			err := fs.UpdateFile("testfile_b.txt", strings.NewReader("foo bar baz"))
			So(err, ShouldBeNil)
			reader, err := fs.ReadFile("testfile_b.txt")
			So(readAll(reader), ShouldEqual, "foo bar baz")
			So(err, ShouldBeNil)
		})

	})
}
