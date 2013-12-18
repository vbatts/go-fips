// +build !fips

package fips

func Mode() (ONOFF, error) {
	return OFF, ErrFipsDisabled
}

func ModeSet(mode ONOFF) (ONOFF, error) {
	return OFF, ErrFipsDisabled
}

func LastError() string {
	return ""
}
