package blog

import (
	"io/ioutil"
	"my-spider/pkg/http"
	"my-spider/pkg/log"
	"my-spider/pkg/util"
)

type WeixinInfo struct {
	CommonInfo
	List []util.HtmlLinkInfo
}

func NewWeixinInfo() WeixinInfo {

	return WeixinInfo{
		CommonInfo: CommonInfo{
			Author:  "鸟窝1",
			WebUrl:  "https://Weixin.com",
			Page:    "",
			PageUrl: "https://Weixin.com/categories/Go",
			//PageUrl: "https://Weixin.com",
		},
	}

}

var WeixinPage string

func (c *WeixinInfo) SetWebPage() {
	if WeixinPage != "" {
		c.Page = WeixinPage
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

	WeixinPage = c.Page
}

func (c *WeixinInfo) setBody() {
	c.List = util.GetHtmlMatchDataLinkTitle(c.Page, "<li class=\"album__list-item js_album_item js_wx_tap_highlight wx_tap_cell\"\n\t\t\t\t\t\tdata-msgid=\"", 10)
}

func (c *WeixinInfo) GetData() interface{} {
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
