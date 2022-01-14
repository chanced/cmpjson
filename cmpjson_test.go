package cmpjson_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/stretchr/testify/require"
)

type badMarshaler struct{}

func (b badMarshaler) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("failing")
}

func TestMustEqual(t *testing.T) {
	assert := require.New(t)
	defer func() { recover() }()
	cmpjson.MustEqual(badMarshaler{}, badMarshaler{})
	assert.FailNow("failed to panic")
}

func TestEqual(t *testing.T) {
	assert := require.New(t)
	a := []byte(`{ "v": "v"}`)
	b := []byte(`{ "v": "v"}`)
	assert.True(cmpjson.Equal(a, b))

	a = []byte(`{ "v": "v"}`)
	b = []byte(`{ "b": "v"}`)
	assert.False(cmpjson.Equal(a, b))

	a = []byte(`{ "t": true, "v": {"s": ["a"]}}`)
	b = []byte(`{ "t": true, "v": {"s": ["a"]}}`)
	assert.True(cmpjson.Equal(a, b))

	a = []byte(`{ 
			"t": true, 
			"v": {"s": ["a"] }
		}`)
	b = []byte(`{ 
			"t": true, 
			"v": { "x": ["a"] }
		}`)
	var m map[string]interface{}
	if err := json.Unmarshal(a, &m); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(b, &m); err != nil {
		panic(err)
	}
	diff, err := cmpjson.Diff(a, b)
	assert.NoError(err)
	assert.NotEmpty(diff)

	ok, d, err := cmpjson.Equal(a, b)
	assert.NoError(err)
	assert.False(ok)
	assert.NotEmpty(d)
	assert.False(cmpjson.MustEqual(a, b))
	assert.False(cmpjson.MustEqual(a, b))
	assert.NotEmpty(cmpjson.Diff(a, b))
	a = []byte(`[1,2,3,4]`)
	b = []byte(`[1,2,5,4]`)
	assert.False(cmpjson.MustEqual(a, b))
	assert.NotEmpty(cmpjson.MustDiff(a, b))
}
