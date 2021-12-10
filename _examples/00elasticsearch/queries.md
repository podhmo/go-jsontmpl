# query

query package is json builder package for accessing elasticsearch, in internal use.

## SearchQuery

for search query.

```json
{
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
}
```