package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/sap/gorfc/gorfc"
)

func abapSystem() gorfc.ConnectionParameters {
	return gorfc.ConnectionParameters{
		"user":   "demo",
		"passwd": "welcome",
		"ashost": "10.68.110.51",
		"sysnr":  "00",
		"client": "620",
		"lang":   "EN",
	}
}

func main() {
	c, err := gorfc.ConnectionFromParams(abapSystem())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected:", c.Alive())

	attrs, _ := c.GetConnectionAttributes()
	fmt.Println("Connection attributes", attrs)

	params := map[string]interface{}{
		"IMPORTSTRUCT": map[string]interface{}{
			"RFCFLOAT": 1.23456789,
			"RFCCHAR1": "A",
			"RFCCHAR2": "BC",
			"RFCCHAR4": "ÄBC",
			"RFCINT1":  0xfe,
			"RFCINT2":  0x7ffe,
			"RFCINT4":  999999999,
			"RFCHEX3":  []byte{255, 254, 253},
			"RFCTIME":  time.Now(),
			"RFCDATE":  time.Now(),
			"RFCDATA1": "HELLÖ SÄP",
			"RFCDATA2": "DATA222",
		},
	}
	r, _ := c.Call("STFC_STRUCTURE", params)

	fmt.Println(r["ECHOSTRUCT"])

	importStruct := params["IMPORTSTRUCT"].(map[string]interface{})
	echoStruct := r["ECHOSTRUCT"].(map[string]interface{})
	fmt.Println(echoStruct)

	// empty time
	fmt.Println(importStruct["RFCDATE"], reflect.TypeOf(importStruct["RFCDATE"]))
	fmt.Println(echoStruct["RFCDATE"], reflect.TypeOf(echoStruct["RFCDATE"]))

	// empty date
	fmt.Println(importStruct["RFCTIME"], reflect.TypeOf(importStruct["RFCTIME"]))
	fmt.Println(echoStruct["RFCTIME"], reflect.TypeOf(echoStruct["RFCTIME"]))

	c.Close()
}
