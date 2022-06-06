package blog

import (
	"io/ioutil"
	"my-spider/pkg/http"
	"my-spider/pkg/log"
	"my-spider/pkg/util"
)

type JiajunInfo struct {
	CommonInfo
	List []util.HtmlLinkInfo
}

func NewJiajunInfo() JiajunInfo {

	return JiajunInfo{
		CommonInfo: CommonInfo{
			Author:  "Jiajun的编程随想",
			WebUrl:  "https://jiajunhuang.com",
			Page:    "",
			PageUrl: "https://jiajunhuang.com/",
		},
	}

}

var JiajunPage string

func (c *JiajunInfo) SetWebPage() {
	if JiajunPage != "" {
		c.Page = JiajunPage
		c.setBody()
		return
	}

	res := http.Get(c.PageUrl)
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Logger.Log(log.LevelError, "GetWebPage err", err)
	}
	c.Page = string(b)
	c.setBody()

	JiajunPage = c.Page
}

func (c *JiajunInfo) setBody() {
	c.List = util.GetHtmlMatchLink(c.Page, "<a href=\"/articles/2022", 10)
}

func (c *JiajunInfo) GetData() interface{} {
	blogData := ResponseBlogData{
		WebName: c.Author,
		WebUrl:  c.WebUrl,
		List:    nil,
	}

	itemList := make([]Item, 0)

	for _, v := range c.List {
		one := Item{
			Url:          "",
			Title:        v.Content,
			BriefContent: "",
			CreateTime:   v.CreateTime,
		}

		one.Url = c.WebUrl + v.Href
		itemList = append(itemList, one)
	}

	blogData.List = itemList

	return blogData
}
