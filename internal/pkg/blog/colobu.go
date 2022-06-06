package blog

import (
	"io/ioutil"
	"my-spider/pkg/http"
	"my-spider/pkg/log"
	"my-spider/pkg/util"
)

type ColobuInfo struct {
	CommonInfo
	List []util.HtmlLinkInfo
}

func NewColobuInfo() ColobuInfo {

	return ColobuInfo{
		CommonInfo: CommonInfo{
			Author:  "鸟窝",
			WebUrl:  "https://colobu.com",
			Page:    "",
			PageUrl: "https://colobu.com/categories/Go",
			//PageUrl: "https://colobu.com",
		},
	}

}

var Page string

func (c *ColobuInfo) SetWebPage() {
	if Page != "" {
		c.Page = Page
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

	Page = c.Page
}

func (c *ColobuInfo) setBody() {
	c.List = util.GetHtmlMatchLink(c.Page, "<a class=\"article-title\" href=\"/2022/", 10)
}

func (c *ColobuInfo) GetData() interface{} {
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
