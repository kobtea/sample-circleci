package greeting

import "testing"

func TestSay(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			"foo",
			"hello foo",
		},
	}
	for _, test := range tests {
		got := Say(test.input)
		if got != test.expect {
			t.Errorf("expect %s, but got %s", test.expect, got)
		}
	}
}
