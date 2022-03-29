package util

import (
	"regexp"
)

type HtmlLinkInfo struct {
	Html    string
	HtmlPre string // html中内容前面的代码
	Href    string // 链接
	Title   string
	Content string
}

func MdMatchStringByHtmlHrefTag(data string) string {
	pat := regexp.MustCompile(`([\s\S])+`)
	res := pat.FindString(data)
	return res
}

func HtmlMatchLink(data string) [][]string {
	pat := regexp.MustCompile(`(<a[\s\S]*?href="(.*?)".*?>)([\s\S]*?)</a>`)
	res := pat.FindAllStringSubmatch(data, -1)
	return res
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
