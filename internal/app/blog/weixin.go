package blog

import "my-spider/internal/pkg/blog"

func TigerbWeixinPage() interface{} {
	obj := blog.WeixinInfo{
		CommonInfo: blog.CommonInfo{
			Author:  "TIGERB",
			WebUrl:  "https://Weixin.com",
			Page:    "",
			PageUrl: "https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzA5MDEwMDYyOA==&action=getalbum&album_id=1708285814280863746&scene=173&from_msgid=2454620078&from_itemidx=1&count=3&nolastread=1&scene=21",
			//PageUrl: "https://Weixin.com",
		},
	}
	obj.SetWebPage()
	retData := obj.GetData()
	return retData
}
