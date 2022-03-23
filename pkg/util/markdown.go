package util

import (
	"regexp"
	"strings"
)

func MdMatchStringByStartAndEnd(data, start, end string) string {
	pat := regexp.MustCompile(start + `([\s\S])+` + end)
	res := pat.FindString(data)
	res = strings.Replace(res, start, "", -1)
	res = strings.Replace(res, end, "", -1)
	return res
}
