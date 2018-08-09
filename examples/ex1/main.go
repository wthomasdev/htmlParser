package main

import (
	"fmt"
	"gophercises/link"
	"strings"
)

var exampleHTML = `<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page <span> some span </span> </a>
  <a href="/second-page">A second link to another page</a>
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
