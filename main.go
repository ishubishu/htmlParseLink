package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	type htmlParseData struct {
		Href string
		Text string
	}
	var parseData []htmlParseData // slice to hold multiple htmlParseData objects
	file, err := os.Open("ex1.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	z := html.NewTokenizer(file)
	var depth int

	for {
		tokenType := z.Next()
		tn, _ := z.TagName()
		if tokenType == html.ErrorToken {
			fmt.Println(parseData)
			return
		}
		if string(tn) == "a" {
			var newData htmlParseData
			if tokenType == html.StartTagToken {
				key, value, _ := z.TagAttr()
				if string(key) == "href" {
					newData.Href = string(value)
				}
				depth = 1
				text := getText(z, &depth)
				newData.Text = text
				parseData = append(parseData, newData)
			}
		}
	}


}

func getText(z *html.Tokenizer, depth *int) string {
	var text strings.Builder
	for {
		tokenType := z.Next()
		switch tokenType {
		case html.ErrorToken:
			return text.String()
		case html.TextToken:
			text.WriteString(" ")
			text.WriteString(strings.TrimSpace(string(z.Text())))
		case html.StartTagToken:
			*depth++
		case html.EndTagToken:
			*depth--
			if *depth == 0 {
				return text.String()
			}
		}
	}
}


