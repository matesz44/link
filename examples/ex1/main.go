package main

import (
	"fmt"
	"strings"

	"github.com/matesz44/link"
)

var exampleHTML = `
<html>
<body>
	<h1>Hello!</h1>
	<a href="/other-page">
		A link to another page
		<span> Some span </span>
	</a>
	<a href="/second-page">A link to a second page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
