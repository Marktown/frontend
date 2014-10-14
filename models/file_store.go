package models

import (
  "github.com/Marktown/frontend/backends"
  "github.com/Marktown/frontend/backends/file_system"
)

func FSFileStore() backends.FileStore {
  return new(file_system.FileStore)
}
