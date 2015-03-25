// see http://www.openssl.org/docs/fips/UserGuide-2.0.pdf
// to set up an environment where fips mode can be enabled

package fips

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

// Mode checks whether is FIPS mode is on
func Mode() (ONOFF, error) {
	return mode()
}

// ModeSet attempts to turn on FIPS for the context of this executable
func ModeSet(mode ONOFF) (ONOFF, error) {
	return modeSet(mode)
}

// LastError is empty when fips is not built, or
//  error:[error code]:[library name]:[function name]:[reason string]
// This error code can also be read with `openssl errstr <error code>`
func LastError() string {
	return lastError()
}
