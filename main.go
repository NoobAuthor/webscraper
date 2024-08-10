package main

import (
	"NoobAuthor/webscraper/scraper"
	"NoobAuthor/webscraper/utils"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <url>")
	}

	url := os.Args[1]

	utils.InitLogger()

	links, err := scraper.FetchAndParse(url)
	if err != nil {
		log.Fatalf("Failed to fetch and parse the page: %v", err)
	}

	deadLinks := scraper.CheckLinks(links)

	if len(deadLinks) == 0 {
		fmt.Println("No dead links found")
	} else {
		fmt.Println("Dead links foung:")
		for _, link := range deadLinks {
			fmt.Println(link)
		}
	}
}
