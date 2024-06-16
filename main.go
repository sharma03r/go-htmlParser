package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var raw = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Test HTML</title>
</head>
<body>
    <h1>Welcome to the Test Page</h1>
    <p>This is a simple paragraph for testing purposes.</p>
    <a href="https://www.example.com">Visit Example.com</a>
		<img src="#"/>
</body>
</html>
`

func visit(n *html.Node, words, pic *int) {
	if n.Type == html.TextNode {
		*words += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		*pic++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, words, pic)
	}
}
func countWordsAndImages(doc *html.Node) (int, int) {
	var words, pic int
	visit(doc, &words, &pic)
	return words, pic
}
func main() {
	doc, err := html.Parse(bytes.NewReader([]byte(raw)))

	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %s\n", err)
		os.Exit(-1)
	}

	words, images := countWordsAndImages(doc)

	fmt.Printf("%d words and %d images", words, images)

}
