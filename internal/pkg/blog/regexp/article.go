package regexp

import (
	"fmt"
	"regexp"
)

func getArticleBodyPreContent(str string) string {
	var del = regexp.MustCompile(`<[\s\S]*>`)
	// 标题
	var rgxTitle = regexp.MustCompile(`"post-title" itemprop="name headline">\s+<a[^>]+>(.+)</a>`)
	titles := rgxTitle.FindAllStringSubmatch(str, 10)
	for _, sub := range titles {
		if len(sub) < 2 {
			continue
		}

		s := del.ReplaceAllString(sub[1], "")
		fmt.Println(s)
	}
	// 内容
	var rgx = regexp.MustCompile(`itemprop="articleBody">\s+<p>(.+)</p>`)
	rs := rgx.FindAllStringSubmatch(str, 10)

	for _, sub := range rs {
		if len(sub) < 2 {
			continue
		}
		s := del.ReplaceAllString(sub[1], "")
		fmt.Println(s)
	}

	return ""
}

func getAllContent(str string) string {
	var del = regexp.MustCompile(`<[\s\S]*>`)
	// 标题
	var rgxTitle = regexp.MustCompile(`"post-title" itemprop="name headline">\s+<a[^>]+>(.+)</a>`)
	titles := rgxTitle.FindAllStringSubmatch(str, 10)
	for _, sub := range titles {
		if len(sub) < 2 {
			continue
		}

		s := del.ReplaceAllString(sub[1], "")
		fmt.Println(s)
	}
	// 内容
	var rgx = regexp.MustCompile(`itemprop="articleBody">\s+<p>(.+)</p>`)
	rs := rgx.FindAllStringSubmatch(str, 10)

	for _, sub := range rs {
		if len(sub) < 2 {
			continue
		}
		s := del.ReplaceAllString(sub[1], "")
		fmt.Println(s)
	}

	return ""
}
