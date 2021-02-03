package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/matesz44/link"
)

/*
	PLAN:
	1. [X] GET req to url
	2. [X] Parse links on the page with `link`
	3. [X] Build absolute URL's from relative links
	4. [/] Filter links to different domains
	5. [ ] Find all the pages recursively (BFS)
	6. [ ] Output XML
*/

func main() {
	urlFlag := flag.String("u", "https://m4t3sz.gitlab.io/bsc/", "url you want to crawl")
	flag.Parse()

	//fmt.Println(*urlFlag)
	pages := get(*urlFlag)

	for _, page := range pages {
		fmt.Println(page)
	}

}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//	io.Copy(os.Stdout, resp.Body)

	// If we pass http:// it usually redirects
	// to https:// so we need the URL after this
	// redirect
	reqURL := resp.Request.URL
	//fmt.Println(reqURL.String())
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	//fmt.Println("Request URL: ", reqURL.String())
	//fmt.Println("Base URL:", baseURL.String())
	return hrefs(resp.Body, baseURL)
}

func hrefs(r io.Reader, baseURL *url.URL) []string {
	base := baseURL.String()
	links, _ := link.Parse(r)

	/*
		Possible URL's:
		/some-path
		http://some-link.com/some-path
		https://some-link.com/some-path
		//some-link.com/some-path
		#fragment
		mailto:asd@asd.asd
	*/

	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "//"):
			ret = append(ret, baseURL.Scheme+":"+l.Href)
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}
