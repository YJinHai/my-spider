package blog

import "time"

type ResponseBlogData struct {
	WebName string      `json:"web_name"`
	WebUrl  string      `json:"web_url"`
	List    interface{} `json:"list"`
}

type Item struct {
	Url          string    `json:"url"`
	Title        string    `json:"title"`
	BriefContent string    `json:"brief_content"`
	CreateTime   time.Time `json:"create_time"`
}
