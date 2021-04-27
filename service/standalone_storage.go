package service

import (
	"github/qm012/nacos-adress/global"
	"github/qm012/nacos-adress/util"
	"go.uber.org/zap"
	"strings"
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

	ticker := time.NewTicker(2 * time.Second)
	for {
		file, exist := util.ExistFile()
		if !exist {
			global.Log.Error("file cluster.conf not exists,use cache mode")
			return
		}
		<-ticker.C
		ips, err := util.ReadClusterConf(file)
		if err != nil {
			global.Log.Error("ticker Watch file failed", zap.String("err", err.Error()))
			return
		}
		s.addNewMap(ips)
		global.Log.Debug("watch file data:", zap.Strings("ips", s.toSlice()))
	}
}

func (s *standalone) toSlice() []string {
	ips := make([]string, 0, 20)
	for k := range s.maps {
		ips = append(ips, k)
	}
	return ips
}

func (s *standalone) addMap(ips []string) {
	for _, v := range ips {
		if !util.IsCorrectIpAddress(v) && !strings.Contains(v, ":") {
			global.Log.Error("read cluster config data is not IP address",
				zap.String("IP", v))
			continue
		}
		s.maps[v] = true
	}
}

func (s *standalone) addNewMap(ips []string) {
	s.maps = make(map[string]interface{}, 20)
	for _, v := range ips {
		if len(v) == 0 {
			continue
		}
		if v == "\b" {
			continue
		}
		if !util.IsCorrectIpAddress(v) && !strings.Contains(v, ":") {
			global.Log.Error("read cluster config data is not IP address",
				zap.String("IP", v))
			continue
		}
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
		s.addNewMap(ips)
		return ips, err
	}

	return s.toSlice(), nil
}

func (s *standalone) add(ips []string) error {

	temp := make([]string, 0, len(ips))
	for _, v := range ips {
		if _, ok := s.maps[v]; !ok {
			temp = append(temp, v)
		}
	}

	if file, exist := util.ExistFile(); exist {
		err := util.WriteClusterConf(file, temp)
		if err != nil {
			return err
		}
	}

	s.addMap(temp)

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
