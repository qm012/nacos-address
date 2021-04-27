package util

import (
	"net"
)

func IsCorrectIpAddress(ipAddress string) bool {
	if len(ipAddress) == 0 {
		return false
	}
	IP := net.ParseIP(ipAddress)
	if IP == nil {
		return false
	}
	return true
}
