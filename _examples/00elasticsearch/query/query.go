package query

import (
	"bytes"
	"strconv"
	"strings"

	gojsontmpl "github.com/podhmo/go-jsontmpl"
)

type QueryBuilder struct {
	builder *gojsontmpl.Builder
	Params  struct {
		Word string `target:"${Word}"`
		Size int    `target:"${Size:int}"`
	}
}

const queryTmpl = `{
    "size": "${Size:int}",
    "query": {
        "bool": {
            "should": [
                {
                    "match": {
                        "word.autocomplete": {
                            "query": "${Word}"
                        }
                    }
                },
                {
                    "match": {
                        "word.readingform": {
                            "query": "${Word}",
                            "fuzziness": "AUTO",
                            "operator": "and"
                        }
                    }
                }
            ]
        }
    }
}`

func NewQueryBuilder(c *gojsontmpl.Config) *QueryBuilder {
	return &QueryBuilder{
		builder: c.NewBuilder(queryTmpl),
	}
}

func (b *QueryBuilder) ToReader() (*bytes.Buffer, error) {
	replacer := strings.NewReplacer(
		`"${Size:int}"`, strconv.Itoa(b.Params.Size),
		`"${Word}"`, strconv.Quote(b.Params.Word),
	)
	return b.builder.ToReader(replacer.Replace)
}

func (b *QueryBuilder) Size(v int) *QueryBuilder {
	copied := b.Params
	copied.Size = v
	return &QueryBuilder{builder: b.builder,
		Params: copied,
	}
}

func (b *QueryBuilder) Word(v string) *QueryBuilder {
	copied := b.Params
	copied.Word = v
	return &QueryBuilder{builder: b.builder,
		Params: copied,
	}
}
