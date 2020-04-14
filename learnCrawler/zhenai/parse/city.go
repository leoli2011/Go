package parse

import (
	"learnCrawler/engine"
	"regexp"
)

const userreg = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`
func ParsecityUser(content []byte) engine.ParseResult  {
	res := regexp.MustCompile(userreg)
	matcher := res.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	var limit = 10
	for _, m := range matcher {
		result.Requests = append(result.Requests, engine.Request{
			string(m[1]),
			Parseprofile,
		})
		name := string(m[2])
		result.Items = append(result.Items, "user:"+name)
		limit--
		if limit < 0 {
			break
		}
	}
	return result
}
