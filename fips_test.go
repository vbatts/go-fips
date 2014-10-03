package fips

import "testing"

func TestTest(t *testing.T) {
	expected := ON
	o, err := ModeSet(expected)
	if err != nil {
		if err == ErrFipsDisabled {
			// ModeSet will not turn it on if fips is not linked in
			expected = OFF
		} else {
			// the error is something else
			t.Fatal(err)
		}
	}

	if o != expected {
		t.Errorf("expected %q, got %q", expected, o)
	}
}
