package main

import "encoding/json"

// ArrayOrString is a custom type that allows us to support
// unmarshaling properties as either an ArrayOrString
// either Values or Value will be populated so consumers should check for
// the length of Values before proceeding with Value
type ArrayOrString struct {
	Values []interface{}
	Value  string
}

// UnmarshalJSON fulfills the interface so we can use json.Unmarshal
func (as *ArrayOrString) UnmarshalJSON(b []byte) error {
	if b[0] == '[' {
		return json.Unmarshal(b, &as.Values)
	}
	return json.Unmarshal(b, &as.Value)
}
