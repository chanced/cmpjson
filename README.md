# cmpjson
A simple package for Go that compares JSON.  

To determine whether json is Equal, cmpjson first marshals them into `interface{}` and then performs the equality check using [github.com/google/go-cmp](https://github.com/google/go-cmp). To format the diff [github.com/pmezard/go-difflib](https://github.com/pmezard/go-difflib) is utilized on json that is first unmarshaled into `interface{}` and then marshaled again with indendentions.

## Usage

```go
package main

import (
	"testing"

	"github.com/chanced/cmpjson"
)

func TestSomething(t *testing.T) {
	dataA := []byte(`{
		"fieldA": ["1", "2", "3"],
		"fieldB": "str"
	}`)
	dataB := []byte(`{
		"fieldB": "str",
		"fieldA": ["2","3"] 
	}`)
	if !cmpjson.Equal(dataA, dataB) {
		t.Error(cmpjson.Diff(dataA, dataB))
	}

}


```
Prints:
```
=== RUN   TestSomething
    prog.go:19: @@ -3 +2,0 @@
        -    "1",
        
--- FAIL: TestSomething (0.00s)
FAIL
```
[https://play.golang.org/p/VMrggnGiydw](https://play.golang.org/p/VMrggnGiydw)

## TODO

Provide better context or formatting for situations like above. Unfortunately the way encoding/json marshal indents, simple differences like the array above can be vague. 

