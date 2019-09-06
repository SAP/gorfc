// +build linux,cgo amd64,cgo

// gorfc wraps the SAP NetWeaver RFC library written in C.
// Its provides methods for maintaining a connection to an ABAP backend and calling remote enabled functions from Go.
// The functions of the library take and return Go data types.
package gorfc

/*

#cgo linux CFLAGS: -DNDEBUG -D_LARGEFILE_SOURCE -D_FILE_OFFSET_BITS=64 -DSAPonUNIX
#cgo linux CFLAGS: -DSAPwithUNICODE -D__NO_MATH_INLINES -DSAPwithTHREADS -DSAPonLIN
#cgo linux CFLAGS: -O2 -minline-all-stringops -g -fno-strict-aliasing -fno-omit-frame-pointer
#cgo linux CFLAGS: -m64 -fexceptions -funsigned-char -Wall -Wno-uninitialized -Wno-long-long
#cgo linux CFLAGS: -Wcast-align -pthread -pipe -Wno-unused-variable

#cgo linux CFLAGS: -I/usr/local/sap/nwrfcsdk/include

#cgo linux LDFLAGS: -O2 -minline-all-stringops -g -fno-strict-aliasing -fno-omit-frame-pointer
#cgo linux LDFLAGS: -m64 -fexceptions -funsigned-char -Wall -Wno-uninitialized -Wno-long-long
#cgo linux LDFLAGS: -Wcast-align -pthread

#cgo windows CFLAGS: -DNDEBUG -D_LARGEFILE_SOURCE -D_FILE_OFFSET_BITS=64 -DSAPonWIN
#cgo windows CFLAGS: -DSAPwithUNICODE -D__NO_MATH_INLINES -DSAPwithTHREADS
#cgo windows CFLAGS: -O2 -minline-all-stringops -g -fno-strict-aliasing -fno-omit-frame-pointer
#cgo windows CFLAGS: -m64 -fexceptions -funsigned-char -Wall -Wno-uninitialized -Wno-long-long
#cgo windows CFLAGS: -Wcast-align -pipe -Wunused-variable

#cgo windows CFLAGS: -IC:/nwrfcsdk/include/
#cgo windows LDFLAGS: -LC:/nwrfcsdk/lib/ -lsapnwrfc -llibsapucum

#cgo windows LDFLAGS: -O2 -minline-all-stringops -g -fno-strict-aliasing -fno-omit-frame-pointer
#cgo windows LDFLAGS: -m64 -fexceptions -funsigned-char -Wall -Wno-uninitialized -Wno-long-long
#cgo windows LDFLAGS: -Wcast-align

#include <sapnwrfc.h>

*/
import "C"

//export wrapNwrfcCallback
func wrapNwrfcCallback(conn C.RFC_CONNECTION_HANDLE, params C.RFC_FUNCTION_HANDLE, errorInfo *C.RFC_ERROR_INFO, n C.int) C.RFC_RC {
	var attributes C.RFC_ATTRIBUTES
	cb := callbacks[n]
	args, err := wrapResult(cb.desc, params, C.RFC_EXPORT, false)
	if err != nil {
		return C.RFC_UNKNOWN_ERROR
	}
	rc := C.RfcGetConnectionAttributes(conn, &attributes, errorInfo)
	if rc != C.RFC_OK {
		return rc
	}
	attrs, err := wrapConnectionAttributes(attributes, false)
	if err != nil {
		return C.RFC_UNKNOWN_ERROR
	}
	result, err := cb.fn(attrs, args)
	if err != nil {
		return C.RFC_UNKNOWN_ERROR
	}
	if err := fillParams(result, cb.desc, params); err != nil {
		// TODO(davegalos): implement fillErrorInfo(err)
		return C.RFC_UNKNOWN_ERROR
	}

	return C.RFC_OK
}
