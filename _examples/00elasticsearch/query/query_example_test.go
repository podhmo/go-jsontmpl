package query_test

import (
	"io"
	"m/00elasticsearch/query"
	"os"

	gojsontmpl "github.com/podhmo/go-jsontmpl"
)

func ExampleSearchQueryBuilder() {
	b := query.NewSearchQueryBuilder(gojsontmpl.Default())
	buf, _ := b.Word("え").Size(5).ToReader()
	io.Copy(os.Stdout, buf)

    // Output:
    // {
    //     "size": 5,
    //     "query": {
    //         "bool": {
    //             "should": [
    //                 {
    //                     "match": {
    //                         "word.autocomplete": {
    //                             "query": "え"
    //                         }
    //                     }
    //                 },
    //                 {
    //                     "match": {
    //                         "word.readingform": {
    //                             "query": "え",
    //                             "fuzziness": "AUTO",
    //                             "operator": "and"
    //                         }
    //                     }
    //                 }
    //             ]
    //         }
    //     }
    // }
}
