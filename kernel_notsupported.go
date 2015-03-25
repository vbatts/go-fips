// +build !linux

package fips

func kernelMode() (ONOFF, error) {
	return OFF, ErrKernelNotSupported
}
