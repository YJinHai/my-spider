package juejin

type Spider struct {
	CUrl string
}

func NewSpider(CUrl string) *Spider {
	return &Spider{CUrl: CUrl}
}

func (s *Spider)GetRecommendFollowFeedLastList() {

}
