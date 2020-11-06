package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/Haydn0230/Link"
	"io/ioutil"
)

type options struct {
	depth int
	url string
}

//stash.skybet.net
func main() {
	//op := getFlags()

	//res, err := http.Get(*urlFlag)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer res.Body.Close()

	//fmt.Printf("\n%v - %v\n", res.Status, res.StatusCode)

	//scanner := bufio.NewScanner(res.Body)
	htmlDoc, err := ioutil.ReadFile("./calhoun.html")
	if err != nil {
		fmt.Printf("Error %+v/n",err)
	}
	r := bytes.NewReader(htmlDoc)
	ll, err := Link.Parse(r)
	if err != nil {
		fmt.Printf("Error %+v/n",err)
	}

	fmt.Printf("\n RESULTS \n %+v",ll)
	//for scanner.Scan() {
	//
	//	fmt.Print(scanner.Text())
	//}

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


