package fips

import "errors"

var (
	// ErrFipsDisabled is returned when this package is built without fips build
	// tag
	ErrFipsDisabled = errors.New("not built with fips tags")

	// ErrKernelNotSupported is whether the kernel can be checked for fips mode
	ErrKernelNotSupported = errors.New("No FIPS check for this kernel")
)
