package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	ChannelTitle string  `xml:"title"`
	ChannelLink  string  `xml:"link"`
	ItemsList    []Items `xml:"item"`
}

type Items struct {
	ItemTitle      string `xml:"title"`
	ItemLink       string `xml:"link"`
	Categories []string `xml:"category"`
}

func main() {
	var info RSS

	data := readDevtoPosts()

	err := xml.Unmarshal(data, &info)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	printInfo(info)
}

func getDevtoPosts() *http.Response {

	resp, err := http.Get("https://dev.to//rss")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return resp
}

func readDevtoPosts() []byte {
	resp := getDevtoPosts()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}

func printInfo(parsedXML RSS) {
	fmt.Println("--------Posts of today at Dev.to--------")
	fmt.Println("----------------------------------------")

	for post := range parsedXML.Channel.ItemsList {
		numOfPost := post + 1

		fmt.Println("Post number", numOfPost)
		fmt.Println("Post title:", parsedXML.Channel.ItemsList[post].ItemTitle)
		fmt.Println("Post link:", parsedXML.Channel.ItemsList[post].ItemLink)

		if len(parsedXML.Channel.ItemsList[post].Categories) <= 0 {
			fmt.Println("Has no categories defined...")
		}

		for category := range parsedXML.Channel.ItemsList[post].Categories {
			numOfCategory := category + 1

			fmt.Println("Category number", numOfCategory, "is", parsedXML.Channel.ItemsList[post].Categories[category])
		}

		fmt.Println("-----------------------------------")
	}
}
