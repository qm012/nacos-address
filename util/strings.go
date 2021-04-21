package util

func SliceContains(strings []string, ip string) bool {

	for _, v := range strings {
		if v == ip {
			return true
		}
	}
	return false
}
