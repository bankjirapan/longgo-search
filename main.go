package main

import (
	"flag"
	"log"

	"longgo-search.com/api"
	"longgo-search.com/crawler"
)

func main() {

	mode := flag.String("mode", "server", "Mode to run the application: 'server' or 'crawler'")
	url := flag.String("url", "", "URL to crawl (required if mode is 'crawler')")
	flag.Parse()

	switch *mode {
	case "server":
		log.Println("Starting HTTP Server...")
		api.StartServer()
	case "crawler":
		if *url == "" {
			log.Fatal("URL is required in crawler mode")
		}
		log.Println("Starting Crawler with URL:", *url)
		crawler.Crawler(*url)
	default:
		log.Fatalf("Invalid mode: %s. Use 'server' or 'crawler'", *mode)
	}
}
