package util

import (
	"regexp"
	"strings"
	"time"
)

type HtmlLinkInfo struct {
	Html       string
	HtmlPre    string // html中内容前面的代码
	Href       string // 链接
	Title      string
	Content    string
	CreateTime time.Time
}

func MdMatchStringByHtmlHrefTag(data string) string {
	pat := regexp.MustCompile(`([\s\S])+`)
	res := pat.FindString(data)
	return res
}

func GetHtmlMatchLink(data, matchTag string, limit int) []HtmlLinkInfo {
	res := HtmlMatchLink(data)
	timeList := HtmlMatchTime(data)

	sourcePre := matchTag
	newRes := res[:0]
	for _, v := range res {

		s := v.HtmlPre
		if len(s) > len(sourcePre) {
			s = s[:len(sourcePre)]
			if s == sourcePre {
				newRes = append(newRes, v)
			}
		}

	}

	if len(timeList) == len(newRes) {
		for i := 0; i < len(newRes); i++ {
			newRes[i].CreateTime = timeList[i].CreateTime
		}
	}

	if len(newRes) < limit {
		limit = len(newRes)
	}

	return newRes[:limit]
}

func GetHtmlMatchTitle(data, matchTag string, limit int) []HtmlLinkInfo {
	res := HtmlMatchTitle(data)
	timeList := HtmlMatchTime(data)

	sourcePre := matchTag
	newRes := res[:0]
	for _, v := range res {

		s := v.HtmlPre
		if len(s) > len(sourcePre) {
			s = s[:len(sourcePre)]
			if s == sourcePre {
				newRes = append(newRes, v)
			}
		}

	}

	if len(timeList) == len(newRes) {
		for i := 0; i < len(newRes); i++ {
			newRes[i].CreateTime = timeList[i].CreateTime
		}
	}

	return newRes[:limit]
}

func HtmlMatchLink(data string) []HtmlLinkInfo {
	pat := regexp.MustCompile(`(<a[\s\S]*?href="(.*?)".*?>)([\s\S]*?)</a>`)
	res := pat.FindAllStringSubmatch(data, -1)
	titles := make([]HtmlLinkInfo, 0)
	for _, v := range res {
		if len(v) < 4 {
			continue
		}
		one := HtmlLinkInfo{
			Html:    v[0],
			HtmlPre: v[1],
			Href:    v[2],
			Content: v[3],
		}
		titles = append(titles, one)
	}

	return titles
}

func HtmlMatchTitle(data string) []HtmlLinkInfo {
	pat := regexp.MustCompile(`(<a[\s\S]*?href="(.*?)"[\s\S]*?title="(.*?)".*?>)([\s\S]*?)</a>`)
	res := pat.FindAllStringSubmatch(data, -1)

	titles := make([]HtmlLinkInfo, len(res))
	for _, v := range res {
		if len(v) < 5 {
			continue
		}
		one := HtmlLinkInfo{
			Html:    v[0],
			HtmlPre: v[1],
			Href:    v[2],
			Title:   v[3],
			Content: v[4],
		}
		titles = append(titles, one)
	}

	return titles
}

func HtmlMatchTime(data string) []HtmlLinkInfo {
	pat := regexp.MustCompile(`(<time[\s\S]*?datetime="(.*?)".*?>)([\s\S]*?)</time>`)
	res := pat.FindAllStringSubmatch(data, -1)
	titles := make([]HtmlLinkInfo, 0)

	for _, v := range res {
		if len(v) < 4 {
			continue
		}
		one := HtmlLinkInfo{
			Html:    v[0],
			HtmlPre: v[1],
			Content: v[3],
		}

		t, err := time.Parse(time.RFC3339Nano, string(v[2]))
		if err != nil {
			continue
		}
		one.CreateTime = t
		titles = append(titles, one)
	}

	return titles
}

func HtmlMatchContent(data string) []HtmlLinkInfo {
	pat := regexp.MustCompile(`(<a[\s\S]*?href="(.*?)"[\s\S]*?title="(.*?)".*?>)([\s\S]*?)</a>`)
	res := pat.FindAllStringSubmatch(data, -1)

	titles := make([]HtmlLinkInfo, len(res))
	for _, v := range res {
		if len(v) < 5 {
			continue
		}
		one := HtmlLinkInfo{
			Html:    v[0],
			HtmlPre: v[1],
			Href:    v[2],
			Content: v[3],
		}
		titles = append(titles, one)
	}

	return titles
}

// 拼接href地址
func JoinHttpUrl(subUrl, url string) string {
	if !strings.Contains(url, "https://") {
		return ""
	}

	if strings.Contains(subUrl, "https://") {
		return subUrl
	}

	url += subUrl

	return url
}
