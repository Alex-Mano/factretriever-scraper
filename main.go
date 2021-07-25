package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// Info si the struct where sata gets stored
type Info struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

var urlPath string
var sub string
var userResp string

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
		fmt.Print("Create JSON file? ")
		fmt.Scanln(&userResp)
		if strings.ToLower(userResp) == "y"{
			writeJSON(allInfos)
			fmt.Println(sub +"-facts.json created!")
		}
	}
}

func writeJSON(data []Info) {
	dataFile, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Could not create JSON")
	}
	ioutil.WriteFile(sub +"-facts.json", dataFile, 0666)
}