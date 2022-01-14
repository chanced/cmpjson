package cmpjson

import (
	"encoding/json"
	"fmt"

	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/wI2L/jsondiff"
)

// MustEqual checks the equality of json.
//
// It panics if there is an error marshaling or unmarshaling to keep the
// interface clean & simple.
//
// The intended purpose is for unit testing equality of json. It is not intended
// to be used outside of that purpose.
func MustEqual(a interface{}, b interface{}) (bool, string) {
	r, diff, err := Equal(a, b)
	_ = diff
	if err != nil {
		panic("cmpjson:" + err.Error())
	}
	return r, diff
}

// Equal checks the equality of json.
//
// The intended purpose is for unit testing equality of json. It is not intended
// to be used outside of that purpose.
func Equal(a interface{}, b interface{}) (bool, string, error) {
	ab, err := marshal(a)
	if err != nil {
		return false, "", err
	}
	bb, err := marshal(b)
	if err != nil {
		return false, "", err
	}
	eq := jsonpatch.Equal(ab, bb)
	if eq {
		return true, "", nil
	}
	d, err := Diff(ab, bb)
	return false, string(d), err
}

// MustDiff returns a JSON patch of the differences between of a and b.
//
// MustDiff panics if there is an error marshaling a or b (if necessary) or
// performing the diff
func MustDiff(a interface{}, b interface{}) []byte {
	d, err := Diff(a, b)
	if err != nil {
		panic("cmpjson:" + err.Error())
	}
	return d
}

func Diff(a interface{}, b interface{}) ([]byte, error) {
	ab, err := marshal(a)
	if err != nil {
		return nil, err
	}
	bb, err := marshal(b)
	if err != nil {
		return nil, err
	}
	p, err := jsondiff.CompareJSON(ab, bb)
	if err != nil {
		return nil, fmt.Errorf("cmpjson: failed to compare JSON %s and %s\n\t%w", a, b, err)
	}
	pb, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("cmpjson: failed to Marshal jsondiff.Patch %v\n\t%w", p, err)
	}
	return pb, err
}

func marshal(v interface{}) ([]byte, error) {
	var b []byte
	var err error
	switch t := v.(type) {
	case []byte:
		return t, nil
	case json.Marshaler:
		b, err = t.MarshalJSON()
	default:
		b, err = json.Marshal(v)

	}
	if err != nil {
		return nil, fmt.Errorf("cmpjson: failed to Marshal JSON of %T %v\n\t%w", v, v, err)
	}
	return b, nil
}
