# cmpjson
A simple package for Go that compares JSON

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
