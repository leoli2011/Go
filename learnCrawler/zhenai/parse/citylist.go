package parse

import (
	"learnCrawler/engine"
	"regexp"
	"fmt"
)

const cityRe = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(content []byte) engine.ParseResult  {
	res := regexp.MustCompile(cityRe)
	matcher := res.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}

	limit := 10
	for _, m := range matcher {
		fmt.Printf("city:%v, url:%v", string(m[2]), string(m[1]))
		result.Requests = append(result.Requests, engine.Request{
			Url:string(m[1]),
			ParseFunc:ParsecityUser,
		})
		result.Items = append(result.Items,  "city:"+string(m[2]))
		limit--
		if limit  < 0 {
			break
		}
	}

	return result
}