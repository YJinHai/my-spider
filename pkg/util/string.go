package util

import (
	"strings"
)

func CUrlToSlice(cUrl string) []string {
	cUrl = strings.Replace(cUrl, "\n ","",-1)
	ss := strings.Split(cUrl, "\\")
	if len(ss) < 1 {
		return nil
	}
	endIndex := strings.Count(cUrl,"-H ")
	rets := make([]string, endIndex+1)
	rets[0] = strings.Replace(ss[0], "curl ","",-1)
	rets[0] = strings.Replace(rets[0], "'","",-1)
	for i,_ := range ss {
		if i == 0 {
			continue
		}

		rets[i] = strings.Replace(ss[i], " -H ","",-1)
		rets[i] = strings.Replace(rets[i], "'","",-1)
		if i == endIndex {
			return rets
		}
	}

	return nil
}

func GetCUrlAddr(cUrl string) string {
	cUrl = strings.Replace(cUrl, "\n ","",-1)
	ss := strings.Split(cUrl, "\\")
	if len(ss) < 1 {
		return ""
	}
	addr := strings.Replace(ss[0], "curl ","",-1)
	addr = strings.Replace(addr, "'","",-1)

	return addr
}

func GetCurlHeader(cUrl string) map[string]string {
	cUrl = strings.Replace(cUrl, "\n ","",-1)
	ss := strings.Split(cUrl, "\\")
	if len(ss) < 1 {
		return nil
	}
	endIndex := strings.Count(cUrl,"-H ")
	rets := make([]string, endIndex+1)
	m := make(map[string]string)
	for i,_ := range ss {
		if i == 0 {
			continue
		}

		rets[i] = strings.Replace(ss[i], " -H ","",-1)
		rets[i] = strings.Replace(rets[i], "'","",-1)
		ms := strings.SplitN(rets[i],":",2)
		if len(ms) > 1 {
			m[ms[0]] = ms[1]
		}
		if i == endIndex {
			return m
		}
	}

	return nil
}

func GetCUrlPost(cUrl string) string {
	cUrl = strings.Replace(cUrl, "\n ","",-1)
	ss := strings.Split(cUrl, "\\")
	if len(ss) < 1 {
		return ""
	}
	endIndex := strings.Count(cUrl,"-H ")

	for i,_ := range ss {
		if i == 0 || i <= endIndex{
			continue
		}

		temps := strings.Replace(ss[i], " --data-raw ","",-1)
		if temps == "" {
			continue
		}
		temps = strings.Replace(temps, "'","",-1)
		return temps
	}

	return ""
}
