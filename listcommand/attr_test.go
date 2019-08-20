package listcommand

import (
	"testing"
)

func TestValueStringOrStringArray_UnmarshalJSON(t *testing.T) {
	// make our test cases
	// in the table-driven testing format
	cases := []struct {
		have []byte
		want string
	}{
		{
			have: []byte(`["string","array"]`),
			want: "string array",
		},
		{
			have: []byte(`"single string"`),
			want: "single string",
		},
	}

	// iterate the test cases
	// and assert that they do what we want
	for _, tc := range cases {
		var vs valueStringOrStringArray
		err := vs.UnmarshalJSON(tc.have)
		if err != nil {
			t.Fatalf("valueStringOrStringArray failed to unmarshall: %s", err)
		}

		if (string)(vs) != tc.want {
			t.Errorf("valueStringOrStringArray unmarshall expected: %s, got: %s", tc.want, (string)(vs))
		}
	}
}

func TestAttr_Attr(t *testing.T) {
	// run a list command for a specific attr and make sure we got the right one
	a := Attr{}
	attrs, err := a.Attr("version", false)
	if err != nil {
		t.Fatalf("list attr attr=version failed with error %s", err)
	}
	if attrs[0].AttrName != "version" {
		t.Errorf("list attr attr=version failed. listed attr=version but got attr=%s", attrs[0].AttrName)
	}
}
