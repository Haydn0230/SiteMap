package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/Haydn0230/Link"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type options struct {
	depth int
	url string
}

type SiteMap struct {
	url string `json:"url"`
	pages []SiteMap `json:"pages"`
}

func main() {
	//op := getFlags()

	//res, err := http.Get(*urlFlag)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer res.Body.Close()

	//fmt.Printf("\n%v - %v\n", res.Status, res.StatusCode)

	//scanner := bufio.NewScanner(res.Body)
	//urlFlag := "./calhoun.html"

	//sitemap := SiteMap{
	//	url:   urlFlag,
	//	pages: buildSiteMap(urlFlag),
	//}


	//fmt.Printf("\n RESULTS \n %+v", links)

	//for scanner.Scan() {
	//
	//	fmt.Print(scanner.Text())
	//}

}

func buildSiteMap(count, depth int, url string) []SiteMap {
	count ++
	if count == depth {
		return []SiteMap{}
	}

	links, err := linksFromFiles(url)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	var sitemaps = make([]SiteMap, 0)
	for _, link := range filter(links) {
		sitemap := SiteMap{
			url:   link,
			pages: buildSiteMap(count, depth, link),
		}
		sitemaps = append(sitemaps, sitemap)
	}

	return sitemaps
}




func getFlags() options {
	op := options{
		url: *flag.String("url", "", "The site you want to build a site map for"),
		depth: *flag.Int("depth", 2, "the depth of the links you want to explore. By default it will go 2 deep"),
		}
	//depthFlag := flag.Int("depth", 2, "the depth of the links you want to explore. By default it will go 2 deep")
	flag.Parse()
	return op
}

func linksFromFiles(path string) ([]Link.Link, error){
	htmlDoc, err := ioutil.ReadFile("./testBuildSiteMapper.html")
	if err != nil {
		return []Link.Link{}, err
	}
	r := bytes.NewReader(htmlDoc)

	return Link.Parse(r)
}

func HTML(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fmt.Printf("\n%v - %v\n", res.Status, res.StatusCode)
}

func filter(links []Link.Link) []string {
	filteredLinks :=make([]string,0,len(links))
	for _, l := range links {
		url, _ := url.Parse(l.Href)

		if url == nil {
			continue
		}

		if strings.Compare(url.Host, "www.calhoun.io") == 0 {
			filteredLinks = append(filteredLinks, l.Href)
		}

		if url.Host == "" && url.Path != "" {
			filteredLinks = append(filteredLinks, l.Href)
		}
	}
	return filteredLinks
}

//lb, err := json.Marshal(links)
//if err !=nil {
//fmt.Printf("Error %+v/n",err)
//}
//
//ioutil.WriteFile("./testFilterData.json", lb, 0777)


