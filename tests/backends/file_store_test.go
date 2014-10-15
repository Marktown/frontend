package test

import (
	"github.com/Marktown/frontend/backends/file_system"

	"io"
	"os"
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
	fs := file_system.NewFileStore()
	fs.RootPath = "../../tmp/tests/fs_file_store/"

	_ = os.Remove(fs.RootPath + "testfile_b.txt")
	_ = os.Remove(fs.RootPath + "testfile_b_new.txt")
	_ = os.Remove(fs.RootPath + "dir_a")

	Convey("Subject: FileStore\n", t, func() {
		Convey("Create", func() {
			fileErr := fs.CreateFile("file_a.txt", strings.NewReader("content_a"))
			So(fileErr, ShouldBeNil)
			fileReader, fileErr := fs.ReadFile("file_a.txt")
			So(readAll(fileReader), ShouldEqual, "content_a")
			So(fileErr, ShouldBeNil)
		})

		Convey("Read", func() {
			fs.RootPath = "../../tests/assets/"
			fileReader, fileErr := fs.ReadFile("testfile.txt")
			So(readAll(fileReader), ShouldEqual, "this is the textfile\n")
			So(fileErr, ShouldBeNil)

			// TODO dir
			fs.RootPath = "../../tmp/tests/fs_file_store/"
		})

		Convey("Update", func() {
			fileErr := fs.UpdateFile("file_a.txt", strings.NewReader("content_b"))
			So(fileErr, ShouldBeNil)
			fileReader, fileErr := fs.ReadFile("file_a.txt")
			So(readAll(fileReader), ShouldEqual, "content_b")
			So(fileErr, ShouldBeNil)
		})

		Convey("CreateDir", func() {
			dirErr := fs.CreateDir("dir_a")
			So(dirErr, ShouldBeNil)
			_, dirErr = os.Stat(fs.RootPath + "dir_a")
			So(os.IsExist(dirErr), ShouldBeTrue)
		})

		Convey("Move", func() {
			fileErr := fs.Move("file_a.txt", "file_b.txt")
			So(fileErr, ShouldBeNil)
			_, fileErrNot := os.Stat(fs.RootPath + "file_a.txt")
			So(os.IsNotExist(fileErrNot), ShouldBeTrue)
			_, fileErrExist := os.Stat(fs.RootPath + "file_b.txt")
			So(os.IsExist(fileErrExist), ShouldBeTrue)
		})

		Convey("Delete", func() {
			fileErr := fs.Delete("file_b.txt")
			So(fileErr, ShouldBeNil)
			_, fileErr = os.Stat(fs.RootPath + "file_b.txt")
			So(os.IsNotExist(fileErr), ShouldBeTrue)

			dirErr := fs.Delete("dir_a")
			So(dirErr, ShouldBeNil)
			_, dirErr = os.Stat(fs.RootPath + "dir_a")
			So(os.IsNotExist(dirErr), ShouldBeTrue)
		})
	})
}
