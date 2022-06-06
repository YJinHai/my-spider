package blog

type CommonInfo struct {
	Author  string `json:"author,omitempty"`
	WebUrl  string `json:"web_url,omitempty"`
	Page    string `json:"page,omitempty"`
	PageUrl string `json:"page_url,omitempty"`
}

func NewCommonInfo(author string, webUrl string, page string) *CommonInfo {
	return &CommonInfo{Author: author, WebUrl: webUrl, Page: page}
}

func (c CommonInfo) GetAuthor() string {
	return c.Author
}

func (c CommonInfo) GetWebUrl() string {
	return c.WebUrl
}

func (c CommonInfo) GetPage() string {
	return c.Page
}
