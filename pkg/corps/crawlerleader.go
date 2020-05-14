package corps

import "fmt"

func NewCrawlerLeader(domain, starturl string) *CrawlerLeader {
	return &CrawlerLeader{
		Domain:   domain,
		StartUrl: starturl,
		Urls:     nil,
	}
}

func (cl *CrawlerLeader) Assign(ic ICrawler) {
	fmt.Println("OK，安排爬虫干活")
	// for _, url := range cl.Urls {
	// 循环遍历所有团长准备好的链接，把获得的内容保存到本地
	// }
}
