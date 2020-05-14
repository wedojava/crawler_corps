package corps

import "fmt"

type Corps struct {
	CLeaders []CrawlerLeader
}

type CrawlerLeader struct {
	Domain   string
	StartUrl string
	Urls     []string
}

type Crawler struct {
	Url      string
	Raw      string
	Title    string
	Content  string
	Datetime string
}

type ICorps interface {
	Fire()
}

type ICrawlerLeader interface {
	GetUrls()
	Assign()
}

type ICrawler interface {
	GetTitle()
	GetContent()
	GetDatetime()
}

func (c *Corps) Fire(ic ICrawlerLeader) error {
	fmt.Println("由Corps调用针对不同网站的爬虫")
	// 先获得url列表
	ic.GetUrls()
	// 然后安排爬虫
	ic.Assign()
	return nil
}
