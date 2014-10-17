package base

import (
	"os"
	"io/ioutil"
)

type FileStore struct{}

func (this *FileStore) ReadDirTree(path string, depth int, tmpList []os.FileInfo) (list []os.FileInfo, err error) {
	// use func of file_system.FileStore
	infoList, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	for _, info := range infoList {
		tmpList = append(tmpList, []os.FileInfo{info}...)
		if info.IsDir() {
			depthCopy := depth - 1
			if depthCopy > 0 {
				tmpList, err = this.ReadDirTree(path+"/"+info.Name(), depthCopy, tmpList)
			}
		}
	}
	list = tmpList
	return
}
