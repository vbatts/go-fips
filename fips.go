/*
see http://www.openssl.org/docs/fips/UserGuide-2.0.pdf
to set up an environment where fips mode can be enabled
*/
package fips

import (
	"errors"
)

var ErrFipsDisabled = errors.New("not built with fips tags")

const (
	OFF ONOFF = iota
	ON
)

type ONOFF int

func (oo ONOFF) String() string {
	if oo == ON {
		return "ON"
	}
	return "OFF"
}
