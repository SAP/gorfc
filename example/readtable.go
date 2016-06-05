package main

import (
	"fmt"
	"simonwaldherr.de/go/golibs/arg"
	saprfc "simonwaldherr.de/go/saprfc"
	"time"
)

var SAPconnection *saprfc.Connection

func printTable(table string) {
	params := map[string]interface{}{
		"QUERY_TABLE": table,
		"DELIMITER":   ";",
		"NO_DATA":     "",
		"ROWSKIPS":    0,
		"ROWCOUNT":    0,
	}
	r, err := SAPconnection.Call("RFC_READ_TABLE", params)
	if err != nil {
		fmt.Println(err)
		return
	}

	echoStruct := r["DATA"].([]interface{})
	for _, value := range echoStruct {
		values := value.(map[string]interface{})
		for _, val := range values {
			fmt.Println(val)
		}
	}
	return
}

func main() {
	//select * from
	arg.String("table", "USR01", "read from table", time.Second*55)
	arg.String("dest", "", "destination system", time.Second*55)
	arg.String("client", "", "client", time.Second*55)
	arg.String("user", "RFC", "username", time.Second*55)
	arg.String("pass", "", "password", time.Second*55)
	arg.String("lang", "DE", "language", time.Second*55)
	arg.String("host", "127.0.0.1", "SAP server", time.Second*55)
	arg.String("sysnr", "00", "SysNr", time.Second*5)
	arg.String("router", "/H/127.0.0.1/H/", "SAP router", time.Second*55)
	arg.Parse()

	SAPconnection, _ = saprfc.ConnectionFromParams(saprfc.ConnectionParameter{
		Dest:      arg.Get("dest").(string),
		Client:    arg.Get("client").(string),
		User:      arg.Get("user").(string),
		Passwd:    arg.Get("pass").(string),
		Lang:      arg.Get("lang").(string),
		Ashost:    arg.Get("host").(string),
		Sysnr:     arg.Get("sysnr").(string),
		Saprouter: arg.Get("router").(string),
	})

	printTable(arg.Get("table").(string))

	SAPconnection.Close()
}
