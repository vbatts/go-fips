package fips

// KernelMode checks whether fips flags are present for the running kernel
//
// This is presently only for Linux kernels
func KernelMode() (ONOFF, error) {
	return kernelMode()
}
