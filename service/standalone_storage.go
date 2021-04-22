package service

import (
	"errors"
	"github/qm012/nacos-adress/global"
	"github/qm012/nacos-adress/util"
	"go.uber.org/zap"
	"sync"
	"time"
)

type standalone struct {
	maps   map[string]interface{}
	rwLock sync.RWMutex
}

func newStandalone() Storage {
	s := &standalone{
		maps: make(map[string]interface{}, 20),
	}
	go s.watchFile()
	return s
}

var _ Storage = &standalone{}

func (s *standalone) watchFile() {
	file, exist := util.ExistFile()
	if !exist {
		return
	}
	ticker := time.NewTicker(2 * time.Second)
	for {
		<-ticker.C
		ips, err := util.ReadClusterConf(file)
		if err != nil {
			global.Log.Error("ticker Watch file failed", zap.String("err", err.Error()))
			return
		}
		s.toMap(ips)
		global.Log.Info("watch file data:", zap.Strings("ips", s.toSlice()))
	}
}

func (s *standalone) toSlice() []string {
	ips := make([]string, 0, 20)
	for k := range s.maps {
		ips = append(ips, k)
	}
	return ips
}

func (s *standalone) toMap(ips []string) {
	for _, v := range ips {
		s.maps[v] = true
	}
}

func (s *standalone) get() ([]string, error) {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()

	if len(s.maps) != 0 {
		return s.toSlice(), nil
	}

	if file, exist := util.ExistFile(); exist {
		ips, err := util.ReadClusterConf(file)
		s.toMap(ips)
		global.Log.Info("form file redis")
		return ips, err
	}

	global.Log.Info("form cache redis")
	return s.toSlice(), nil
}

func (s *standalone) add(ips []string) error {

	for _, v := range ips {
		if _, ok := s.maps[v]; ok {
			return errors.New("contains duplicate data")
		}
	}

	if file, exist := util.ExistFile(); exist {
		err := util.WriteClusterConf(file, ips)
		if err != nil {
			return err
		}
	}

	s.toMap(ips)

	return nil
}

func (s *standalone) delete(ips []string) error {

	if file, exist := util.ExistFile(); exist {
		err := util.ReplaceClusterFile(file, ips)
		if err != nil {
			return err
		}
	}

	for _, v := range ips {
		delete(s.maps, v)
	}

	return nil
}

func (s *standalone) deleteAll() error {
	if file, exist := util.ExistFile(); exist {
		err := util.ReplaceClusterFile(file, s.toSlice())
		if err != nil {
			return err
		}
	}
	s.maps = make(map[string]interface{}, 20)
	return nil
}
