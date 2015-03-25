// +build !fips

package fips

func mode() (ONOFF, error) {
	return OFF, ErrFipsDisabled
}

func modeSet(mode ONOFF) (ONOFF, error) {
	return OFF, ErrFipsDisabled
}

func lastError() string {
	return ""
}
