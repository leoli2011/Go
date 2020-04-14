package parse

import (
	"learnCrawler/engine"
	"regexp"
	"strconv"
	"learnCrawler/model"
)

var agere = regexp.MustCompile(`data-v-8b1eac0c>([\d+])cm</div>`)
var hightre = regexp.MustCompile(`data-v-8b1eac0c>([\d+])岁</div>`)

func Parseprofile(content []byte) engine.ParseResult  {
	 age, err := strconv.Atoi(extrareg(content, agere))
	 hight, err := strconv.Atoi(extrareg(content, hightre))
	profile := model.Profile{}
	if err != nil {
		profile.Age = age
		profile.Height = hight
	}


	result := engine.ParseResult{Items: []interface{} {profile }}
	return result
}

func extrareg(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)
	if len(match) >=2 {
		return string (match[1])
	} else {
		return ""
	}

}