package validate

import (
	"github.com/go-playground/validator/v10"
	"github/qm012/nacos-adress/global"
	"github/qm012/nacos-adress/util"
	"go.uber.org/zap"
	"strings"
)

func VerifyIp(fl validator.FieldLevel) bool {
	ipSliceLen := fl.Field().Len()
	if ipSliceLen == 0 {
		global.Log.Error("ip address list is null")
		return false
	}
	for i := 0; i < ipSliceLen; i++ {
		value := fl.Field().Index(i).String()
		if strings.Contains(value, ":") {
			value = strings.Split(value, ":")[0]
		}
		if !util.IsCorrectIpAddress(value) {
			global.Log.Info("Not a correct IP address", zap.String("ip", value))
			return false
		}
	}
	return true
}
