package juejin


type Article struct {
	ArticleId int64 `json:"article_id"`
	CategoryId int64 `json:"category_id"`
	BriefContent string `json:"brief_content"`
	CTime int64 `json:"brief_content"`
}
