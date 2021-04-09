package cmpjson_test

import (
	"fmt"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/stretchr/testify/require"
)

func TestCompareObject(t *testing.T) {
	assert := require.New(t)
	a := []byte(`{ "v": "v"}`)
	b := []byte(`{ "v": "v"}`)
	assert.True(cmpjson.Equal(a, b))
	a = []byte(`{ "v": "v"}`)
	b = []byte(`{ "b": "v"}`)
	assert.False(cmpjson.Equal(a, b))
	a = []byte(`{ "t": true, "v": 
	{"s": 
	["a"]}}`)
	b = []byte(`{ "t": true, 
		"v": {"s": ["a"]}}`)
	assert.True(cmpjson.Equal(a, b))
	a = []byte(`{ "t": true, "v": {"s": ["a"]}}`)
	b = []byte(`{ "t": true, "v": {"x": ["a"]}}`)
	assert.False(cmpjson.Equal(a, b), cmpjson.Diff(a, b))
	assert.NotEmpty(cmpjson.Diff(a, b))

	a = []byte(`[1,2,3,4]`)
	b = []byte(`[1,2,5,4]`)
	assert.False(cmpjson.Equal(a, b), cmpjson.Diff(a, b))
	assert.NotEmpty(cmpjson.Diff(a, b))
	fmt.Println(cmpjson.Diff(a, b))
}
func TestCompareArray(t *testing.T) {

}
