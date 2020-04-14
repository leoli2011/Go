package engine

import (
	"log"
	"learnCrawler/fetch"
)

func Run(seed... Request)  {
	var requests []Request
	for _,r := range seed {
		requests = append(requests, r)
	}

	for (len(requests) > 0) {
		r := requests[0]
		requests = requests[1:]
		log.Printf("fetching %s", r.Url)

		body, err := fetch.Fetch(r.Url)
		if err != nil {
			log.Printf("fetch : error " + "fetch Url %s, %v", r.Url, err)
			continue
		}
		ParseResult := r.ParseFunc(body)
		requests = append(requests, ParseResult.Requests...)

		for _, item := range ParseResult.Items{
			log.Printf("go get item %v", item)
		}
	}
}
