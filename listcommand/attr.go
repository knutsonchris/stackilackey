package listcommand

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
)

// valueStringOrStringArray will contain the value of an attr, sometimes it will be a string and other times a []string
// it has a custom unmarshal function that will convert the []string to a space dilimeted string, something that Stacki can accept as a substitute
type valueStringOrStringArray string

// Attr represents an attribute stored in the stacki database
type Attr struct {
	Scope    string                   `json:"scope"`
	Type     string                   `json:"type"`
	AttrName string                   `json:"attr"`
	Value    valueStringOrStringArray `json:"value"`
}

/*
Attr will list the set of global attributes.
Parameters

	[attr=string]

	A shell syntax glob pattern to specify to attributes to
	be listed.

	[shadow=boolean]

	Specifies if shadow attributes are listed, the default
	is True.
*/
func (a *Attr) Attr(attrName string, shadow bool) ([]Attr, error) {
	var shadowstr string
	if shadow == true {
		shadowstr = "true"
	} else {
		shadowstr = "false"
	}
	args := []interface{}{attrName, shadowstr}
	c := fmt.Sprintf("list attr attr='%s' shadow='%s'", args...)
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	attrs := []Attr{}
	err = json.Unmarshal(b, &attrs)
	return attrs, err
}

// UnmarshalJSON is a valueStringOrStringArray reciever that so we can survive recieving a string (like we are expecting) or an array of strings
// the Stacki API can be rather unpredictable
func (v *valueStringOrStringArray) UnmarshalJSON(b []byte) error {

	// if it is just a string like we want then no problem
	var str string
	err := json.Unmarshal(b, &str)
	if err == nil {
		return json.Unmarshal(b, (*string)(v))
	}
	// if it is not, take each string in the array and create a space-delimeted list
	strSlice := make([]string, 0)
	err = json.Unmarshal(b, &strSlice)
	if err != nil {
		// is we end up here then that means it was neither a string or []string
		return err
	}
	*v = valueStringOrStringArray(strings.Join(strSlice, " "))
	return nil
}
