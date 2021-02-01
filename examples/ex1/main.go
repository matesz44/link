package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/matesz44/link"
)

func main() {
	HTML := flag.String("html", "index.html", "html file to parse")
	flag.Parse()
	file, err := os.Open(*HTML)
	if err != nil {
		panic(err)
	}

	links, err := link.Parse(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
