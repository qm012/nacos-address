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
	file, err := os.Open("./cluster.conf")
	if err != nil {
		global.Log.Fatal("open file failed", zap.String("err", err.Error()))
		return nil, false
	}
	return file, true
}

func ReplaceClusterFile(file *os.File, ips []string) error {
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		global.Log.Fatal("readAll file data failed", zap.String("err", err.Error()))
		return err
	}
	str := string(data)
	for _, v := range ips {
		if strings.Contains(str, v) {
			str = strings.ReplaceAll(str, v, "")
		}
	}
	return nil
}
func WriteClusterConf(file *os.File, ips []string) (err error) {
	defer file.Close()
	for _, v := range ips {
		ip := fmt.Sprintf("%v\n", v)
		_, err := file.WriteString(ip)
		if err != nil {
			global.Log.Fatal("write file failed", zap.String("err", err.Error()))
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
			global.Log.Fatal("read line file failed", zap.String("err", err.Error()))
			return ips, err
		}
		instance := strings.TrimSpace(string(line))
		if strings.HasPrefix(instance, "#") {
			continue
		}
		if strings.Contains(instance, "#") {
			//192.168.71.52:8848 # Instance A
			instance = strings.TrimSpace(instance[:strings.Index(instance, "#")])
		}
		index := strings.Index(instance, ",")
		if index > 0 {
			//support the format: ip1:port,ip2:port  # multi inline
			ips = append(ips, strings.Split(instance, ",")...)
			continue
		}
		if SliceContains(ips, instance) {
			continue
		}
		ips = append(ips, instance)
	}
	return
}
