// Package serve_gorfc recreates the STFC_DEEP_TABLE example from the nwrfc docs in go.
package main

import (
	"fmt"

	"google3/base/go/log"
	"google3/third_party/golang/gorfc/gorfc/gorfc"
)

func serve(_ gorfc.ConnectionAttributes, args map[string]interface{}) (interface{}, error) {
	fmt.Printf("%v\n", args)
	return map[string]interface{}{}, nil
}

func main() {
	conn, err := gorfc.ConnectionFromParams(gorfc.ConnectionParameter{
		Client: "000",
		User:   "user",
		Passwd: "****",
		Lang:   "DE",
		Ashost: "binmain",
		Sysnr:  "53",
	})
	if err != nil {
		log.Exit(err)
	}

	desc, err := conn.GetFunctionDescription("STFC_DEEP_TABLE")
	if err != nil {
		log.Exit(err)
	}
	conn.Close()

	if err := gorfc.InstallServerFunction(desc, serve); err != nil {
		log.Exit(err)
	}

	serverConn, err := gorfc.ServerFromParams(gorfc.ServerParameter{
		Program_id: "MY_SERVER",
		Gwhost:     "binmain",
		Gwserv:     "sapgw53",
	})
	if err != nil {
		log.Exit(err)
	}
	defer serverConn.Close()

	for {
		serverConn.ListenAndDispatch(120)
	}

	//	printfU(cU("Starting to listen...\n\n"));
	//	while(RFC_OK == rc || RFC_RETRY == rc || RFC_ABAP_EXCEPTION == rc){
	//		rc = RfcListenAndDispatch(serverHandle, 120, &errorInfo);
	//		printfU(cU("RfcListenAndDispatch() returned %s\n"), RfcGetRcAsString(rc));
	//		switch (rc){
	//		case RFC_RETRY: // This only notifies us, that no request came in within the
	//			timeout period.
	//			// We just continue our loop.
	//			printfU(cU("No request within 120s.\n"));
	//			break;
	//		case RFC_ABAP_EXCEPTION: // Our function module implementation has returned RFC_ABAP_EXCEPTION.
	//			// This is equivalent to an ABAP function module throwing an ABAP Exception.
	//			// The Exception has been returned to R/3 and our connection is still open.
	//			// So we just loop around.
	//			printfU(cU("ABAP_EXCEPTION in implementing function: %s\n"), errorInfo.key);
	//			break;

	//		case RFC_NOT_FOUND: // R/3 tried to invoke a function module, for which we did not supply
	//			// an implementation. R/3 has been notified of this through a SYSTEM_FAILURE,
	//			// so we need to refresh our connection.
	//			printfU(cU("Unknown function module: %s\n"), errorInfo.message);
	//		case RFC_EXTERNAL_FAILURE: // Our function module implementation raised a SYSTEM_FAILURE. In this case
	//			// the connection needs to be refreshed as well.
	//			printfU(cU("SYSTEM_FAILURE has been sent to backend.\n\n"));
	//		case RFC_ABAP_MESSAGE: // And in this case a fresh connection is needed as well
	//			serverHandle = RfcRegisterServer(serverCon, 3, &errorInfo);
	//			rc = errorInfo.code;
	//			break;
	//		}
	//		// This allows us to shutdown the RFC Server from R/3. The implementation of STFC_DEEP_TABLE
	//		// will set listening to false, if IMPORT_TAB-C == STOP.
	//		if (!listening){
	//			RfcCloseConnection(serverHandle, NULL);
	//			break;
	//		}
	//	}
	//	return 0;
}
