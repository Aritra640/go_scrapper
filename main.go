package main

import (
	"fmt"
	"log"
	"os"

	//"io/ioutil" - scrapped after go v 1.6
	"strconv"

	"encoding/json"

	"github.com/gocolly/colly/v2"
)

type fact struct {
	ID   int    //'json:"id"'
	Desc string //'json:"description"'
}

func main() {

	allFacts := make([]fact, 0)
	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)
	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		factID, err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			log.Println("Could not get id")
		}

		factDesc := element.Text
		new_fact := fact{
			ID:   factID,
			Desc: factDesc,
		}

		allFacts = append(allFacts, new_fact)

		//fmt.Println(new_fact.Desc)
		//fmt.Println(allFacts)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit("https://www.factretriever.com/rhino-facts")

	//enc := json.NewEncoder(os.Stdout)
	//enc.SetIndent("", " ")
	//enc.Encode(allFacts)

	writeJSON(allFacts)
}

// func writeJSON(data []fact) {
// 	file, err := json.MarshalIndent(data, "", " ")
// 	if err != nil {
// 		log.Println("error : unable to create a json file")
// 		return
// 	}
// 	_ = os.WriteFile("rhinofacts.json", file, 0644)
// }

func writeJSON(data []fact) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("error : unable to create a json file")
		return
	}
	f_ins, e := os.Create(("rhino_facts.json"))
	if e != nil {
		log.Println("error : unable to write json file")
	}

	l, e1 := f_ins.Write(file)
	if e1 != nil {
		log.Println("error : unable to access writer")
	}
	log.Println(l, "json-file written successfully")
	l_err := f_ins.Close()
	if l_err != nil {
		log.Println("error : file instance open")
		return
	}
	return
}
