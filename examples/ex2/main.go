package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
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
	6. [X] Output XML
*/

func main() {
	urlFlag, depthFlag, xmlFlag := argp()
	pages := bfs(*urlFlag, *depthFlag)
	if *xmlFlag {
		outXML(pages)
	} else {
		for _, page := range pages {
			fmt.Println(page)
		}
	}
}

func argp() (urlFlag *string, depthFlag *int, xmlFlag *bool) {
	urlFlag = flag.String("u", "https://m4t3sz.gitlab.io/bsc/", "url you want to crawl")
	depthFlag = flag.Int("d", 3, "depth you want to follow links")
	xmlFlag = flag.Bool("x", false, "use this to output xml in a sitemap format")
	flag.Parse()
	return urlFlag, depthFlag, xmlFlag
}

const xmlns = "https://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func outXML(pages []string) {
	toXML := urlset{
		Urls:  make([]loc, len(pages)),
		Xmlns: xmlns,
	}
	for i, page := range pages {
		toXML.Urls[i] = loc{page}
	}

	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXML); err != nil {
		panic(err)
	}
	fmt.Println()
}

// empty structs need less mem than bools
type empty struct{}

// every url
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
		if len(q) == 0 {
			break
		}
		for url := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = empty{}

			for _, link := range get(url) {
				if _, ok := seen[link]; !ok {
					nq[link] = empty{}
				}
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

	// If we pass http:// it usually redirects
	// to https:// so we need the URL after this
	// redirect
	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	req := reqURL.String()
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
