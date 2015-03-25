// +build linux

package fips

import (
	"bytes"
	"io/ioutil"
	"os"
)

var (
	kernelCommandLine   = "/proc/cmdline"
	kernelFipsParameter = []byte("fips=")
)

func kernelMode() (ONOFF, error) {
	if _, err := os.Stat(kernelCommandLine); os.IsNotExist(err) {
		return OFF, ErrKernelNotSupported
	}
	fh, err := os.Open(kernelCommandLine)
	if err != nil {
		return OFF, err
	}
	defer fh.Close()

	buf, err := ioutil.ReadAll(fh)
	if err != nil {
		return OFF, err
	}

	enabled := OFF
	for _, chunk := range bytes.Split(buf, []byte(" ")) {
		if bytes.HasPrefix(chunk, kernelFipsParameter) {
			val := bytes.TrimPrefix(chunk, kernelFipsParameter)
			if string(val) == "1" {
				enabled = ON
			} else if string(val) == "0" {
				enabled = OFF
			}
		}
	}

	return enabled, nil
}
