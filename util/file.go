package util

import (
	"bufio"
	"fmt"
	"github/qm012/nacos-adress/global"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ExistFile() (*os.File, bool) {
	filePath := "./config/cluster.conf"
	_, err := os.Lstat(filePath)
	if err != nil {
		global.Log.Debug("file not exists")
		return nil, os.IsExist(err)
	}
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		global.Log.Error("open file failed", zap.String("err", err.Error()))
		return nil, false
	}
	return file, true
}

func ReplaceClusterFile(file *os.File, ips []string) error {
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		global.Log.Error("readAll file data failed", zap.String("err", err.Error()))
		return err
	}
	str := string(data)
	for _, v := range ips {
		if strings.Contains(str, v) {
			str = strings.ReplaceAll(str, "\n"+v, "")
		}
	}

	if err := os.Truncate("./config/cluster.conf", 0); err != nil {
		global.Log.Error("file truncate failed", zap.Error(err))
		return err
	}
	_, err = file.WriteString(str)
	return err
}

func WriteClusterConf(file *os.File, ips []string) (err error) {
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	var first = strings.HasSuffix(string(all), "\n")

	for _, v := range ips {
		var ip = fmt.Sprintf("%s\n", v)
		if !first {
			ip = fmt.Sprintf("\n%s\n", v)
			first = true
		}
		_, err := file.WriteString(ip)
		if err != nil {
			global.Log.Error("write file failed", zap.String("err", err.Error()))
			return err
		}
	}
	return nil
}

func ReadClusterConf(file *os.File) (ips []string, err error) {
	defer file.Close()
	reader := bufio.NewReader(file)
	ips = make([]string, 0, 20)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			global.Log.Error("read line file failed", zap.String("err", err.Error()))
			return ips, err
		}
		instance := strings.TrimSpace(string(line))
		if strings.HasPrefix(instance, "#") || strings.Contains(instance, ",") {
			continue
		}
		if strings.Contains(instance, "#") {
			//192.168.71.52:8848 # Instance A
			instance = strings.TrimSpace(instance[:strings.Index(instance, "#")])
		}
		if SliceContains(ips, instance) {
			continue
		}
		ips = append(ips, instance)
	}
	return
}
