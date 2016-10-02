package main

import (
	"fmt"
	saprfc "simonwaldherr.de/go/saprfc"
	"strings"
)

var SAPconnection *saprfc.Connection

func abapSystem() saprfc.ConnectionParameter {
	return saprfc.ConnectionParameter{
		Dest:      "I64",
		Client:    "400",
		User:      "demo",
		Passwd:    "welcome",
		Lang:      "DE",
		Ashost:    "10.117.24.158",
		Sysnr:     "00",
		Saprouter: "/H/203.13.155.17/S/3299/W/xjkb3d/H/172.19.137.194/H/",
	}
}

func connect() error {
	var err error
	SAPconnection, err = saprfc.ConnectionFromParams(abapSystem())
	return err
}

func close() {
	SAPconnection.Close()
}

func request() []string {
	params := map[string]interface{}{
		"QUERY_TABLE": "USR01",
		"DELIMITER":   ";",
		"NO_DATA":     "",
		"ROWSKIPS":    0,
		"ROWCOUNT":    0,
	}
	r, err := SAPconnection.Call("RFC_READ_TABLE", params)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	var ret []string

	echoStruct := r["DATA"].([]interface{})
	for _, value := range echoStruct {
		values := value.(map[string]interface{})
		for _, val := range values {
			valstr := strings.Split(fmt.Sprint("%s", val), ";")
			ret = append(ret, strings.TrimSpace(valstr[1]))
		}
	}
	return ret
}

func main() {
	connect()

	user := request()

	for _, usr := range user {
		fmt.Println(usr)
	}

	close()
}
