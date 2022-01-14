# cmpjson

A simple package for Go that compares JSON for testing purposes.

To determine whether json is
equal,[github.com/evanphx/json-patch/v5](https://github.com/evanphx/json-patch)
is utilized. If the json is not equal, a diff is created utilizing
[github.com/wI2L/jsondiff](https://github.com/wI2L/jsondiff).

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
	if ok, diff, err := cmpjson.Equal(dataA, dataB); !ok {
		t.Error(diff, err)
	}
	// alternatively:
	if ok, diff := cmpjson.MustEqual(dataA, dataB); !ok {
		t.Error("dataA and dataB were not equal:", diff)
	}
}


```

Prints:

```

=== RUN   TestSomething
    /cmpjson/examples_test.go:84: [
          {
            "op": "remove",
            "path": "/fieldA/2"
          },
          {
            "op": "replace",
            "path": "/fieldA/0",
            "value": "2"
          },
          {
            "op": "replace",
            "path": "/fieldA/1",
            "value": "3"
          }
        ] <nil>
    /cmpjson/examples_test.go:88: dataA and dataB were not equal: [
          {
            "op": "remove",
            "path": "/fieldA/2"
          },
          {
            "op": "replace",
            "path": "/fieldA/0",
            "value": "2"
          },
          {
            "op": "replace",
            "path": "/fieldA/1",
            "value": "3"
          }
        ]
--- FAIL: TestSomething (0.00s)
FAIL
FAIL	github.com/chanced/cmpjson	0.245s
```

[https://play.golang.org/p/VMrggnGiydw](https://play.golang.org/p/VMrggnGiydw)

## License

MIT
