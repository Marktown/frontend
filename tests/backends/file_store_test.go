package test

import (
	"github.com/Marktown/frontend/backends/file_system"

	"io"
	"os"
	"testing"
	"text/scanner"

	"bytes"
	"crypto/sha1"
	"strings"

	"fmt"

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

	_ = os.MkdirAll(fs.RootPath, os.ModePerm)
	_ = os.Remove(fs.RootPath + "testfile_b.txt")
	_ = os.Remove(fs.RootPath + "dir_b")

	Convey("Subject: FileStore\n", t, func() {
		Convey("Create file", func() {
			fileErr := fs.WriteFile("file_a.txt", strings.NewReader("content_a"))
			So(fileErr, ShouldBeNil)
			fileReader, fileErr := fs.ReadFile("file_a.txt")
			So(readAll(fileReader), ShouldEqual, "content_a")
			So(fileErr, ShouldBeNil)
		})

		Convey("Read file and dir", func() {
			// test root path
			fs.RootPath = "../../tests/assets/"
			fileReader, fileErr := fs.ReadFile("testfile.txt")
			So(readAll(fileReader), ShouldEqual, "this is the textfile\n")
			So(fileErr, ShouldBeNil)

			// TODO dir
			// check if each counter is correct
			list, err := fs.ReadDir("testfolder")
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 6)

			// check for top-level directory if file- and dirnames are correct
			var buffer bytes.Buffer
			for _, file := range list {
				buffer.WriteString(file.Name() + ";")
			}
			So(buffer.String(), ShouldEqual, "dir_a;dir_b;dir_c;file_a.md;file_b.md;file_c.md;")

			list, err = fs.ReadDir("testfolder/dir_a")
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 3)

			list, err = fs.ReadDir("testfolder/dir_b")
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 1)

			list, err = fs.ReadDir("testfolder/dir_c")
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 4)

			list, err = fs.ReadDir("testfolder/dir_c/dir_c_a/")
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 3)

			list, err = fs.ReadDir("testfolder/dir_c/dir_c_b/")
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 1)

			list, err = fs.ReadDir("testfolder/dir_c/dir_c_c/")
			So(err, ShouldBeNil)
			So(len(list), ShouldEqual, 1)

			// default root path
			fs.RootPath = "../../tmp/tests/fs_file_store/"
		})

		Convey("Create dir", func() {
			dirErr := fs.CreateDir("dir_a")
			So(dirErr, ShouldBeNil)
			_, dirErr = os.Stat(fs.RootPath + "dir_a")
			So(dirErr, ShouldBeNil)
		})

		Convey("Move file and dir", func() {
			fileErr := fs.Move("file_a.txt", "file_b.txt")
			So(fileErr, ShouldBeNil)
			_, fileErrNot := os.Stat(fs.RootPath + "file_a.txt")
			So(os.IsNotExist(fileErrNot), ShouldBeTrue)
			_, fileErrExist := os.Stat(fs.RootPath + "file_b.txt")
			So(fileErrExist, ShouldBeNil)

			dirErr := fs.Move("dir_a", "dir_b")
			So(dirErr, ShouldBeNil)
			_, dirErrNot := os.Stat(fs.RootPath + "dir_a")
			So(os.IsNotExist(dirErrNot), ShouldBeTrue)
			_, dirErrExist := os.Stat(fs.RootPath + "dir_b")
			So(dirErrExist, ShouldBeNil)

		})

		Convey("Delete", func() {
			fileErr := fs.Delete("file_b.txt")
			So(fileErr, ShouldBeNil)
			_, fileErr = os.Stat(fs.RootPath + "file_b.txt")
			So(os.IsNotExist(fileErr), ShouldBeTrue)

			dirErr := fs.Delete("dir_b")
			So(dirErr, ShouldBeNil)
			_, dirErr = os.Stat(fs.RootPath + "dir_b")
			So(os.IsNotExist(dirErr), ShouldBeTrue)
		})

		Convey("Checksum", func() {
			fs.RootPath = "../../tests/assets/"
			sum, err := fs.Checksum("testfile.txt")
			So(err, ShouldBeNil)
			stringCorrect := fmt.Sprintf("%x", sha1.Sum([]byte("this is the textfile\n")))
			stringTest := fmt.Sprintf("%x", sum)
			So(stringCorrect, ShouldEqual, stringTest)
		})

		Convey("ReadDirTree", func() {
			fs.RootPath = "../../tests/assets/testfolder/"
			list, err := fs.ReadDirTree("", 2)
			So(err, ShouldBeNil)
			So(14, ShouldEqual, len(list))
		})
	})
}
