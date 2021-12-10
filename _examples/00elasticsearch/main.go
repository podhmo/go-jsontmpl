package main

import (
	"io"
	"m/00elasticsearch/query"
	"os"

	gojsontmpl "github.com/podhmo/go-jsontmpl"
)

// TODO: code generation

func main() {
	c := gojsontmpl.Default()
	// TODO: required
	// TODO: gentle error message
	// TODO: with default
	b := query.NewSearchQueryBuilder(c).Size(10).Word("え")
	buf, err := b.ToReader()
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, buf)
}
