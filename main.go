package main

import (
	"flag"
	"fmt"
	"github.com/Haydn0230/golessons/parse/Link"

	"log"
	"net/http"
)
func main() {
	urlFlag :=  flag.String("url", "", "The site you want to build a site map for")
	//depthFlag := flag.Int("depth", 2, "the depth of the links you want to explore. By default it will go 2 deep")
	flag.Parse()

	fmt.Println(*urlFlag)

	res, err := http.Get( *urlFlag)
	if err != nil{
		log.Fatal(err)
	}
	defer res.Body.Close()


	//Link.HTML()
	fmt.Printf("\n%v - %v\n", res.Status, res.StatusCode)



	//scanner := bufio.NewScanner(res.Body)
	ll, err := Link.HTML(res.Body)
	if err != nil {
		fmt.Printf("Error %v/n",err)
	}

	fmt.Println(ll)
	//for scanner.Scan() {
	//
	//	fmt.Print(scanner.Text())
	//}

}