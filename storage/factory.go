package storage

import "log"

type StorageType int

const (
	StorageTypeMemory StorageType = iota
	StorageTypeMock
	StorageTypeDB
)

func GetStorage(t StorageType) Storage {
	var s Storage
	switch t {
	case StorageTypeMemory:
		s = newMemory()
	case StorageTypeMock:
		s = mock{}
	case StorageTypeDB:
		s = newConnection()
	default:
		log.Fatal("storage type tidak dikenali")
	}
	return s
}
