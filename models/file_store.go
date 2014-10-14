package models

func FileStore() FileStoreInterface {
  return new(FSFileStore)
}
