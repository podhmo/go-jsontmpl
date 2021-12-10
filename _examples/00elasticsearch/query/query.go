package query

import (
	"bytes"
	"strconv"
	"strings"

	gojsontmpl "github.com/podhmo/go-jsontmpl"
)

// SearchQueryBuilder for search query.
type SearchQueryBuilder struct {
	builder *gojsontmpl.Builder
	Params  struct {
		// Word: query
		Word string `target:"${Word}"`

		// Size: number of limit
		Size int `target:"${Size:int}"`
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

func NewSearchQueryBuilder(c *gojsontmpl.Config) *SearchQueryBuilder {
	return &SearchQueryBuilder{
		builder: c.NewBuilder(queryTmpl),
	}
}

func (b *SearchQueryBuilder) ToReader() (*bytes.Buffer, error) {
	replacer := strings.NewReplacer(
		`"${Size:int}"`, strconv.Itoa(b.Params.Size),
		`"${Word}"`, strconv.Quote(b.Params.Word),
	)
	return b.builder.ToReader(replacer.Replace)
}

func (b *SearchQueryBuilder) Size(v int) *SearchQueryBuilder {
	copied := b.Params
	copied.Size = v
	return &SearchQueryBuilder{builder: b.builder,
		Params: copied,
	}
}

func (b *SearchQueryBuilder) Word(v string) *SearchQueryBuilder {
	copied := b.Params
	copied.Word = v
	return &SearchQueryBuilder{builder: b.builder,
		Params: copied,
	}
}
