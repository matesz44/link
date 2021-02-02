package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/matesz44/link"
)

/*
	PLAN:
	1. GET req to url
	2. Parse links on the page with `link`
	3. Build absolute URL's from relative links
	4. FFilter links to different domains
	5. Find all the pages recursively (BFS)
	6. Output XML
*/

func main() {
	urlFlag := flag.String("u", "https://m4t3sz.gitlab.io/bsc/", "url you want to crawl")
	flag.Parse()

	//fmt.Println(*urlFlag)

	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//	io.Copy(os.Stdout, resp.Body)

	/*
		Possible URL's:
		/some-path
		http://some-link.com/some-path
		https://some-link.com/some-path
		//some-link.com/some-path
		#fragment
		mailto:asd@asd.asd
	*/

	// If we pass http:// it usually redirects
	// to https:// so we need the URL after this
	// redirect
	reqURL := resp.Request.URL
	//fmt.Println(reqURL.String())
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()
	//fmt.Println("Request URL: ", reqURL.String())
	//fmt.Println("Base URL:", base)

	links, _ := link.Parse(resp.Body)

	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "//"):
			hrefs = append(hrefs, "https:"+l.Href)
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}
	for _, href := range hrefs {
		fmt.Println(href)
	}

}
