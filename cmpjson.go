package cmpjson

import (
	"encoding/json"

	"github.com/google/go-cmp/cmp"
	"github.com/pmezard/go-difflib/difflib"
)

// Equal checks the equality json by unmarshaling each into interface{},
// and comparing the results with go-cmp
//
// It panics if there is an error marshaling or unmarshaling to keep the
// interface clean & simple.
//
// The intended purpose is for unit testing equality of json. It is not intended
// to be used outside of that purpose.
func Equal(a []byte, b []byte) bool {
	var av, bv interface{}
	err := json.Unmarshal(a, &av)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &bv)
	if err != nil {
		panic(err)
	}
	return cmp.Equal(av, bv)
}

func Diff(a []byte, b []byte) string {
	var av, bv interface{}
	err := json.Unmarshal(a, &av)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &bv)
	if err != nil {
		panic(err)
	}
	var ad, bd []byte
	ad, err = json.MarshalIndent(av, "", "  ")
	if err != nil {
		panic(err)
	}
	bd, err = json.MarshalIndent(bv, "", "  ")
	if err != nil {
		panic(err)
	}

	df := difflib.UnifiedDiff{
		A: difflib.SplitLines(string(ad)),
		B: difflib.SplitLines(string(bd)),
	}
	str, err := difflib.GetUnifiedDiffString(df)
	if err != nil {
		panic(err)
	}
	return str
}
