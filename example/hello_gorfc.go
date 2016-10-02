package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	saprfc "simonwaldherr.de/go/saprfc"
)

func abapSystem() saprfc.ConnectionParameter {
	return saprfc.ConnectionParameter{
		Dest:      "I64",
		Client:    "800",
		User:      "demo",
		Passwd:    "welcome",
		Lang:      "EN",
		Ashost:    "10.117.24.158",
		Sysnr:     "00",
		Saprouter: "/H/203.13.155.17/S/3299/W/xjkb3d/H/172.19.137.194/H/",
	}
}

func main() {
	c, _ := saprfc.ConnectionFromParams(abapSystem())
	var t *testing.T

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
			"RFCDATA1": "Hellö SÄP",
			"RFCDATA2": "DATA222",
		},
	}
	r, _ := c.Call("STFC_STRUCTURE", params)

	assert.NotNil(t, r["ECHOSTRUCT"])
	importStruct := params["IMPORTSTRUCT"].(map[string]interface{})
	echoStruct := r["ECHOSTRUCT"].(map[string]interface{})
	assert.Equal(t, importStruct["RFCFLOAT"], echoStruct["RFCFLOAT"])
	assert.Equal(t, importStruct["RFCCHAR1"], echoStruct["RFCCHAR1"])
	assert.Equal(t, importStruct["RFCCHAR2"], echoStruct["RFCCHAR2"])
	assert.Equal(t, importStruct["RFCCHAR4"], echoStruct["RFCCHAR4"])
	assert.Equal(t, importStruct["RFCINT1"], echoStruct["RFCINT1"])
	assert.Equal(t, importStruct["RFCINT2"], echoStruct["RFCINT2"])
	assert.Equal(t, importStruct["RFCINT4"], echoStruct["RFCINT4"])
	//	assert.Equal(t, importStruct["RFCHEX3"], echoStruct["RFCHEX3"])
	assert.Equal(t, importStruct["RFCTIME"].(time.Time).Format("150405"), echoStruct["RFCTIME"].(time.Time).Format("150405"))
	assert.Equal(t, importStruct["RFCDATE"].(time.Time).Format("20060102"), echoStruct["RFCDATE"].(time.Time).Format("20060102"))
	assert.Equal(t, importStruct["RFCDATA1"], echoStruct["RFCDATA1"])
	assert.Equal(t, importStruct["RFCDATA2"], echoStruct["RFCDATA2"])

	fmt.Println(reflect.TypeOf(importStruct["RFCDATE"]))
	fmt.Println(reflect.TypeOf(importStruct["RFCTIME"]))

	c.Close()
}
