package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const ERR_TXT = "=== hidgen: Error ==="

func usage() {
	fmt.Println("USAGE: hidgen <input.html> <output.html> <Optional: Heading tag to be replaced>")
}

func main() {
	flag.Usage = usage

	if len(os.Args) < 3 {
		fmt.Println(ERR_TXT)
		usage()
		os.Exit(1)
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]
	replaceHeading := "h2, h3"

	if len(os.Args) == 4 {
		replaceHeading = os.Args[3]
	}

	f, err := os.Open(inputPath)

	if err != nil {
		fmt.Println(ERR_TXT)
		log.Fatal(err)
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Println(ERR_TXT)
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(b)))
	if err != nil {
		fmt.Println(ERR_TXT)
		log.Fatal(err)
	}

	doc.Find(replaceHeading).Each(func(i int, s *goquery.Selection) {
		anchor, _ := goquery.NewDocumentFromReader(strings.NewReader("<a></a>"))
		anchor.Find("a").SetAttr("id", s.Text())
		html, _ := anchor.Html()
		s.BeforeHtml(html + "\n")
	})

	html, err := doc.Find("body").Html()
	if err != nil {
		fmt.Println(ERR_TXT)
		log.Fatal(err)
	}

	// "<img src=''/>" -> "<img src='' />"
	replacedHtml := strings.Replace(html, "/>", " />", -1)

	data := []byte(replacedHtml)

	f2, err := os.Create(outputPath)
	if err != nil {
		fmt.Println(ERR_TXT)
		log.Fatal(err)
	}

	defer f2.Close()

	_, err = f2.Write(data)

	if err != nil {
		fmt.Println(ERR_TXT)
		log.Fatal(err)
	}

	fmt.Println("=== hidgen: Successful ===")
}
