package cmpjson_test

import (
	"fmt"

	"github.com/chanced/cmpjson"
)

func ExampleEqual() {
	dataA := []byte(`{
		"fieldA": ["1", "2", "3"],
		"fieldB": "str"
	}`)
	dataB := []byte(`{
		"fieldB": "str",
		"fieldA": ["2","3"]
	}`)
	if ok, diff, err := cmpjson.Equal(dataA, dataB); !ok {
		fmt.Println(diff)

		if err != nil {
			fmt.Println(err)
		}
	}
	// Output:
	// [
	//   {
	//     "op": "remove",
	//     "path": "/fieldA/2"
	//   },
	//   {
	//     "op": "replace",
	//     "path": "/fieldA/0",
	//     "value": "2"
	//   },
	//   {
	//     "op": "replace",
	//     "path": "/fieldA/1",
	//     "value": "3"
	//   }
	// ]
}

func ExampleMustEqual() {
	dataA := map[string]interface{}{
		"fieldA": []interface{}{"1", "2", "3"},
		"fieldB": "initial",
	}
	dataB := map[string]interface{}{
		"fieldA": []interface{}{"1", "2"},
		"fieldB": "newval",
	}

	if ok, diff := cmpjson.MustEqual(dataA, dataB); !ok {
		fmt.Println(diff)
	} else {
		fmt.Println("dataA and dataB were equal")
	}
	// Output:
	// [
	//   {
	//     "op": "remove",
	//     "path": "/fieldA/2"
	//   },
	//   {
	//     "op": "replace",
	//     "path": "/fieldB",
	//     "value": "newval"
	//   }
	// ]
}
