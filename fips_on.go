// +build fips

package fips

/*
#include <stdlib.h>
#include <openssl/crypto.h>
#include <openssl/err.h>
#cgo LDFLAGS: -lcrypto
*/
import "C"
import "errors"

func mode() (ONOFF, error) {
	return ONOFF(C.FIPS_mode()), nil
}

func modeSet(mode ONOFF) (ONOFF, error) {
	o := ONOFF(C.FIPS_mode_set(C.int(mode)))
	if o != mode {
		return o, errors.New(lastError())
	}
	return o, nil
}

func lastError() string {
	buf := C.malloc(1024)
	e := C.ERR_get_error() // a C.ulong
	C.ERR_load_crypto_strings()
	defer C.ERR_free_strings()
	C.ERR_error_string_n(e, (*C.char)(buf), 1024)
	defer C.free(buf)
	return C.GoString((*C.char)(buf))
}
