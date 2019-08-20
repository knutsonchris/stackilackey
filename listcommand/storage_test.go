package listcommand

import (
	"testing"
)

func TestFlexString_UnmarshalJSON(t *testing.T) {
	// make our test cases
	// in the table-driven testing format
	cases := []struct {
		have []byte
		want string
	}{
		{
			have: []byte("\"test string\""),
			want: "test string",
		},
		{
			have: []byte("1"),
			want: "1",
		},
	}

	// iterate the test cases
	// and assert that they do what we want
	for _, tc := range cases {
		var fs FlexString
		err := fs.UnmarshalJSON(tc.have)
		if err != nil {
			t.Fatalf("FlexString failed to unmarshall: %s", err)
		}

		if (string)(fs) != tc.want {
			t.Errorf("FlexString unmarshall expected: %s, got: %s", tc.want, (string)(fs))
		}
	}
}
