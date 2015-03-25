// see http://www.openssl.org/docs/fips/UserGuide-2.0.pdf
// to set up an environment where fips mode can be enabled

package fips

import "errors"

// ErrFipsDisabled is returned when this package is built without fips build tag
var ErrFipsDisabled = errors.New("not built with fips tags")

// ONOFF is either on or off
type ONOFF int

const (
	// OFF is off
	OFF ONOFF = iota
	// ON is on
	ON
)

// String is for the Stringer infterface
func (oo ONOFF) String() string {
	if oo == ON {
		return "ON"
	}
	return "OFF"
}
