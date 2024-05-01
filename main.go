package htmlparser

import (
	"os"
	"strings"

	"golang.org/x/net/html"
)

type HTMLParseData struct {
	Href string
	Text string
}

func ParseHTMLFile(filePath string) ([]HTMLParseData, error) {
	var parseData []HTMLParseData
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	z := html.NewTokenizer(file)
	var depth int

	for {
		tokenType := z.Next()
		tn, _ := z.TagName()
		if tokenType == html.ErrorToken {
			return parseData, nil
		}
		if string(tn) == "a" {
			var newData HTMLParseData
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
