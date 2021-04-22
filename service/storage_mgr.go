package service

import (
	"github/qm012/nacos-adress/util"
	"strings"
)

type StorageMgr struct {
	Storage
}

func NewStorageMgr() *StorageMgr {
	model := util.GetStorageModel()
	var storage Storage
	switch model {
	case util.StorageModelRedis:
		storage = newCluster()
	default:
		storage = newStandalone()
	}

	return &StorageMgr{
		Storage: storage,
	}
}

func (s *StorageMgr) Get() (ipStr string, err error) {
	ips, err := s.Storage.get()
	if err != nil {
		return
	}

	var builder strings.Builder
	builder.Grow(15 * len(ips))
	for _, v := range ips {
		builder.WriteString(v)
		builder.WriteString("\n")
	}
	return builder.String(), nil
}

func (s *StorageMgr) Add(ips []string) (err error) {
	err = s.Storage.add(ips)
	return
}
func (s *StorageMgr) Delete(ips []string) (err error) {
	err = s.Storage.delete(ips)
	return
}

func (s *StorageMgr) DeleteAll() (err error) {
	err = s.Storage.deleteAll()
	return
}
