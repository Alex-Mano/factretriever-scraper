package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// info wtf
type Info struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

var urlPath string
var sub string

func main() {
	
	fmt.Print("Subject: ")
	fmt.Scanln(&sub)
	sub = strings.ToLower(sub)
	urlPath = "https://www.factretriever.com/"+ sub +"-facts"

	allInfos :=  make([]Info, 0)

	c := colly.NewCollector(
		colly.AllowedDomains("www.factretriever.com"),
	)

	c.OnHTML(".factsList li", func(e *colly.HTMLElement)  {
		infoID, err := strconv.Atoi(e.Attr("id"))
		if err != nil{
			log.Println("err: ", err)
		}

		infoDesc := e.Text

		info := Info{
			ID:          infoID,
			Description: infoDesc,
		}


		allInfos = append(allInfos, info)
		

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting ", r.URL.String())
	})

	c.Visit(urlPath)

	if len(allInfos) <= 0 {
		if sub[len(sub)-1:]== "s"{
			fmt.Println("Try using the singular form of the word!")
		}
		fmt.Println("\nNo macth for", sub + "!")
	}else{	
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", " ")
		enc.Encode(allInfos)
	}
	

}