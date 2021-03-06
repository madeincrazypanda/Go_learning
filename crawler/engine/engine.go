package engine

import (
	"Go_learning/crawler/fetcher"
	"log"
)

func Run(seeds ...Request){
	var requests []Request
	for _,r :=range seeds{
		requests=append(requests,r)
	}
	for len(requests)>0{
		r:=requests[0]
		requests=requests[1:]
		body,err:=fetcher.Fetch(r.Url)
		if err != nil{
			log.Printf("Fetcher:error"+"fetching URL %s: %v",r.Url,err)
			continue
		}
		log.Printf("Fetcher:"+"URL %s",r.Url)
		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...) // 把所有数据送进去
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}

	}
}
