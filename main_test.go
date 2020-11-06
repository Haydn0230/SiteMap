package main

import (
	"encoding/json"
	"fmt"
	"github.com/Haydn0230/Link"
	"io/ioutil"
	"log"
	"testing"
)



func Test_filter(t *testing.T) {
	testData, err := ioutil.ReadFile("./testFilterData.json")
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	var links []Link.Link

	err = json.Unmarshal(testData, &links)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	fl := filter(links)

	fmt.Printf("\n RESULTS \n %+v", fl)

}

func Test_BuildSiteMapper(t *testing.T) {
	testUrl := "./testBuildSiteMapper.json"
	sitemap := SiteMap{
		url:   testUrl,
		pages: buildSiteMap(0, 3, testUrl),
	}
	// TODO: get it to write to Json
	xb, err := json.Marshal(sitemap)
	if err !=nil {
		fmt.Printf("Error %+v/n",err)
	}

	ioutil.WriteFile("./output.json", xb, 0777)
	//fmt.Printf("\n RESULTS \n %+v", xb)

}