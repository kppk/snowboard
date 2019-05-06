package drafter

/*
#cgo CFLAGS: -I"${SRCDIR}/ext/drafter/src/" -I"${SRCDIR}/ext/drafter/ext/snowcrash/src/"
#cgo darwin LDFLAGS: -L"${SRCDIR}/ext/drafter/build/out/Release/" -ldrafter -lsnowcrash -lmarkdownparser -lsundown -lc++
#cgo linux LDFLAGS: -L"${SRCDIR}/ext/drafter/build/out/Release/" -ldrafter -lsnowcrash -lmarkdownparser -lsundown -lstdc++
#include <stdlib.h>
#include <stdio.h>
#include "drafter.h"
*/
import "C"
import (
	"fmt"
	"io"
	"io/ioutil"
	"unsafe"
)

func Parse(r io.Reader) ([]byte, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	cSource := C.CString(string(b))
	cResult := &C.drafter_result{}
	cOption := C.drafter_parse_options{requireBlueprintName: false}

	code := int(C.drafter_parse_blueprint(cSource, &cResult, cOption))
	if code != 0 {
		return nil, fmt.Errorf("Parse failed with code: %d", code)
	}

	C.free(unsafe.Pointer(cSource))

	return serialize(cResult), nil
}

func Validate(r io.Reader) ([]byte, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	cSource := C.CString(string(b))
	cOption := C.drafter_parse_options{requireBlueprintName: false}
	cResult := &C.drafter_result{}

	code := int(C.drafter_check_blueprint(cSource, &cResult, cOption))
	if code != 0 {
		return nil, fmt.Errorf("Validate failed with code: %d", code)
	}

	C.free(unsafe.Pointer(cSource))

	return serialize(cResult), nil
}

func Version() string {
	return C.GoString(C.drafter_version_string())
}

func serialize(r *C.drafter_result) []byte {
	options := C.drafter_serialize_options{sourcemap: false, format: C.DRAFTER_SERIALIZE_JSON}
	cResult := C.drafter_serialize(r, options)
	results := C.GoString(cResult)

	C.free(unsafe.Pointer(cResult))

	return []byte(results)
}
