// +build !fips

package fips

// Mode checks whether is FIPS mode is on
func Mode() (ONOFF, error) {
	return OFF, ErrFipsDisabled
}

// ModeSet attempts to turn on FIPS for the context of this executable
func ModeSet(mode ONOFF) (ONOFF, error) {
	return OFF, ErrFipsDisabled
}

// LastError is empty when fips is not built
func LastError() string {
	return ""
}
