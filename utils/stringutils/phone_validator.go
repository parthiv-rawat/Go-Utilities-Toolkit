package stringutils

import "regexp"

func IsValidPhone(phone string) bool {
	re := regexp.MustCompile(`^[6-9]\d{9}$`)
	return re.MatchString(phone)
}
