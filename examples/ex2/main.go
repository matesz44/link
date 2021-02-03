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
	4. [X] Filter links to different domains
	5. [X] Find all the pages recursively (BFS)
	6. [ ] Output XML
*/

func main() {

	urlFlag, depthFlag := argp()
	//fmt.Println(*urlFlag)
	pages := bfs(*urlFlag, *depthFlag)
	//fmt.Println(depthFlag)
	//pages := get(*urlFlag)
	//filter(baseURL)

	for _, page := range pages {
		fmt.Println(page)
	}
}

func argp() (urlFlag *string, depthFlag *int) {
	urlFlag = flag.String("u", "https://m4t3sz.gitlab.io/bsc/", "url you want to crawl")
	depthFlag = flag.Int("d", 3, "depth you want to follow links")
	flag.Parse()
	return urlFlag, depthFlag
}

// empty structs need less mem than bools
type empty struct{}

// all of the urls
func bfs(urlStr string, depth int) []string {
	// with a map you dont need to iterate
	// over the whole thing like with a slice
	seen := make(map[string]empty)
	var q map[string]empty
	nq := map[string]empty{
		urlStr: struct{}{},
	}
	for i := 0; i <= depth; i++ {
		q, nq = nq, make(map[string]empty)
		for url := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = empty{}

			for _, link := range get(url) {
				nq[link] = empty{}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for url := range seen {
		ret = append(ret, url)
	}
	return ret
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
	req := reqURL.String()
	//base := baseURL.String()
	return filter(hrefs(resp.Body, baseURL), withPrefix(req))
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

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		// https://m4t3sz.gitlab.io/bsc --> OK
		// https://m4t3sz.gitlab.io/bsc/writeup --> OK
		// https://twitter.com --> NOPE
		if keepFn(link) {
			ret = append(ret, link)
		}
	}

	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
