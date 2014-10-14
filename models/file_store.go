package models

import (
  "github.com/Marktown/frontend/backends"
  "github.com/Marktown/frontend/backends/file_system"
)

func FSFileStore() backends.FileStore {
  fs := new(file_system.FileStore)
  fs.FilePath = "tmp/file_system/"
  return fs
}
