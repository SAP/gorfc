package gorfc

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//
// Helper Functions
//

func isValueInList(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

//
// NW RFC Lib Version
//
func TestNWRFCLibVersion(t *testing.T) {
	major, minor, patchlevel := GetNWRFCLibVersion()
	assert.Equal(t, uint(7420), major) // adapt to your NW RFC Lib version
	assert.Equal(t, uint(0), minor)
	assert.Equal(t, uint(0), patchlevel)
}

//
// Connection Tests
//
func TestConnect(t *testing.T) {
	fmt.Println("Connection test: Open and Close")
	c, err := ConnectionFromParams(abapSystem())
	if err != nil {
		t.SkipNow()
	}
	assert.NotNil(t, c)
	assert.Nil(t, err)
	c.Close()
}

func TestConnectionAttributes(t *testing.T) {
	fmt.Println("Connection test: Parameters")
	c, err := ConnectionFromParams(abapSystem())
	if err != nil {
		t.SkipNow()
	}

	a, err := c.GetConnectionAttributes()
	paramNames := []string{
		"Dest",
		"Host",
		"PartnerHost",
		"SysNumber",
		"SysId",
		"Client",
		"User",
		"Language",
		"Trace",
		"IsoLanguage",
		"Codepage",
		"PartnerCodepage",
		"RfcRole",
		"Type",
		"PartnerType",
		"Rel",
		"PartnerRel",
		"KernelRel",
		"CpicConvId",
		"ProgName",
		"PartnerBytesPerChar",
		"PartnerSystemCodepage",
		"Reserved"}

	s := reflect.ValueOf(&a).Elem()
	// check if all parameters returned
	assert.Equal(t, 23, s.NumField())
	for i := 0; i < s.NumField(); i++ {
		pname := s.Type().Field(i).Name
		assert.Equal(t, true, isValueInList(pname, paramNames), pname)
		//f := s.Field(i)
		//fmt.Println(i, f.Type(), f, s.Type().Field(i).Name)
	}
	// check some parameters
	assert.Equal(t, strings.ToUpper(abapSystem().User), a.User)
	assert.Equal(t, abapSystem().Sysnr, a.SysNumber)
	assert.Equal(t, abapSystem().Client, a.Client)
	c.Close()
}

func TestPing(t *testing.T) {
	fmt.Println("Connection test: Ping")
	c, err := ConnectionFromParams(abapSystem())
	if err != nil {
		return
	}
	err = c.Ping()
	assert.Nil(t, err)
	c.Close()
}

func TestAlive(t *testing.T) {
	fmt.Println("Connection test: Alive")
	c, err := ConnectionFromParams(abapSystem())
	if err != nil {
		return
	}
	a := c.Alive()
	assert.True(t, a)
	c.Close()
	a = c.Alive()
	assert.False(t, a)
}

func TestReopen(t *testing.T) {
	fmt.Println("Connection test: Reopen")
	c, err := ConnectionFromParams(abapSystem())
	if err != nil {
		return
	}
	err = c.Reopen()
	assert.Nil(t, err)
	c.Close()
}

func TestConnectFromDest(t *testing.T) {
	c, err := ConnectionFromDest("I64_2")
	if err != nil {
		return
	}
	assert.NotNil(t, c)
	assert.Nil(t, err)
	c.Close()
}

//
// Connection Errors
//

func TestWrongUserConnect(t *testing.T) {
	fmt.Println("Connection Error: Logon")
	a := abapSystem()
	a.User = "@!n0user"
	c, err := ConnectionFromParams(a)
	if err != nil {
		return
	}
	assert.Nil(t, c)
	assert.NotNil(t, err)
	assert.Equal(t, "Connection could not be opened", err.(*RfcError).Description)
	assert.Equal(t, "Name or password is incorrect (repeat logon)", err.(*RfcError).ErrorInfo.Message)
	assert.Equal(t, "RFC_LOGON_FAILURE", err.(*RfcError).ErrorInfo.Code)
	assert.Equal(t, "RFC_LOGON_FAILURE", err.(*RfcError).ErrorInfo.Key)
}

func TestMissingAshostConnect(t *testing.T) {
	fmt.Println("Connection Error: Connection parameter missing")
	a := abapSystem()
	a.Ashost = ""
	c, err := ConnectionFromParams(a)
	if err != nil {
		return
	}
	assert.Nil(t, c)
	assert.NotNil(t, err)
	assert.Equal(t, "Connection could not be opened", err.(*RfcError).Description)
	assert.Equal(t, "Parameter ASHOST, GWHOST or MSHOST is missing.", err.(*RfcError).ErrorInfo.Message)
	assert.Equal(t, "RFC_INVALID_PARAMETER", err.(*RfcError).ErrorInfo.Code)
	assert.Equal(t, "RFC_INVALID_PARAMETER", err.(*RfcError).ErrorInfo.Key)
}

func TestWrongParameter(t *testing.T) {
	fmt.Println("Connection Error: Invoke with wrong parameter")
	type importStruct struct {
		XXX string
	}
	c, err := ConnectionFromParams(abapSystem())
	if err != nil {
		return
	}
	r, err := c.Call("STFC_CONNECTION", importStruct{"wrong param"})
	assert.Equal(t, map[string]interface{}(nil), r)
	assert.Equal(t, "RFC_INVALID_PARAMETER", err.(*RfcError).ErrorInfo.Code) // todo: should be "20" ??
	assert.Equal(t, "RFC_INVALID_PARAMETER", err.(*RfcError).ErrorInfo.Key)
	assert.Equal(t, "field 'XXX' not found", err.(*RfcError).ErrorInfo.Message)
	c.Close()
}

//
// STFC Tests
//

func TestFunctionDescription(t *testing.T) {
	c, err := ConnectionFromParams(abapSystem())
	if err != nil {
		return
	}
	d, err := c.GetFunctionDescription("STFC_CONNECTION")
	assert.Nil(t, err)
	assert.Equal(t, "ECHOTEXT", d.Parameters[0].Name)
	assert.Equal(t, "RESPTEXT", d.Parameters[1].Name)
	assert.Equal(t, "REQUTEXT", d.Parameters[2].Name)
	c.Close()
}

func TestFunctionCall(t *testing.T) {
	c, err := ConnectionFromParams(abapSystem())
	if err != nil {
		return
	}
	type importedStruct struct {
		RFCFLOAT float64
		RFCCHAR1 string
		RFCCHAR2 string
		RFCCHAR4 string
		RFCINT1  int
		RFCINT2  int
		RFCINT4  int
		RFCHEX3  []byte
		RFCTIME  time.Time
		RFCDATE  time.Time
		RFCDATA1 string
		RFCDATA2 string
	}
	type parameter struct {
		IMPORTSTRUCT importedStruct
		RFCTABLE     []importedStruct
	}
	importStruct := importedStruct{1.23456789, "A", "BC", "DEFG", 1, 2, 345, []byte{0, 11, 12}, time.Now(), time.Now(), "HELLÖ SÄP", "DATA222"}
	params := parameter{importStruct, []importedStruct{importStruct}}
	r, err := c.Call("STFC_STRUCTURE", params)
	assert.Nil(t, err)
	assert.NotNil(t, r)

	assert.Nil(t, r["IMPORTSTUCT"])

	assert.NotNil(t, r["ECHOSTRUCT"])
	echoStruct := r["ECHOSTRUCT"].(map[string]interface{})
	assert.Equal(t, importStruct.RFCFLOAT, echoStruct["RFCFLOAT"])
	assert.Equal(t, importStruct.RFCCHAR1, echoStruct["RFCCHAR1"])
	assert.Equal(t, importStruct.RFCCHAR2, echoStruct["RFCCHAR2"])
	assert.Equal(t, importStruct.RFCCHAR4, echoStruct["RFCCHAR4"])
	assert.Equal(t, importStruct.RFCINT1, echoStruct["RFCINT1"])
	assert.Equal(t, importStruct.RFCINT2, echoStruct["RFCINT2"])
	assert.Equal(t, importStruct.RFCINT4, echoStruct["RFCINT4"])
	//assert.Equal(t, importStruct.RFCHEX3, echoStruct["RFCHEX3"])
	assert.Equal(t, importStruct.RFCTIME.Format("150405"), echoStruct["RFCTIME"].(time.Time).Format("150405"))
	assert.Equal(t, importStruct.RFCDATE.Format("20060102"), echoStruct["RFCDATE"].(time.Time).Format("20060102"))
	assert.Equal(t, importStruct.RFCDATA1, echoStruct["RFCDATA1"])
	assert.Equal(t, importStruct.RFCDATA2, echoStruct["RFCDATA2"])

	assert.NotNil(t, r["RFCTABLE"])
	echoTableLine := r["RFCTABLE"].([]interface{})[0].(map[string]interface{})
	assert.Equal(t, importStruct.RFCFLOAT, echoTableLine["RFCFLOAT"])
	assert.Equal(t, importStruct.RFCCHAR1, echoTableLine["RFCCHAR1"])
	assert.Equal(t, importStruct.RFCCHAR2, echoTableLine["RFCCHAR2"])
	assert.Equal(t, importStruct.RFCCHAR4, echoTableLine["RFCCHAR4"])
	assert.Equal(t, importStruct.RFCINT1, echoTableLine["RFCINT1"])
	assert.Equal(t, importStruct.RFCINT2, echoTableLine["RFCINT2"])
	assert.Equal(t, importStruct.RFCINT4, echoTableLine["RFCINT4"])
	//assert.Equal(t, importStruct.RFCHEX3, echoTableLine["RFCHEX3"])
	assert.Equal(t, importStruct.RFCTIME.Format("150405"), echoTableLine["RFCTIME"].(time.Time).Format("150405"))
	assert.Equal(t, importStruct.RFCDATE.Format("20060102"), echoTableLine["RFCDATE"].(time.Time).Format("20060102"))
	assert.Equal(t, importStruct.RFCDATA1, echoTableLine["RFCDATA1"])
	assert.Equal(t, importStruct.RFCDATA2, echoTableLine["RFCDATA2"])
	c.Close()
}

func TestStructPassedAsMap(t *testing.T) {
	c, err := ConnectionFromParams(abapSystem())
	if err != nil {
		return
	}

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
	//assert.Equal(t, importStruct["RFCHEX3"], echoStruct["RFCHEX3"])
	assert.Equal(t, importStruct["RFCTIME"].(time.Time).Format("150405"), echoStruct["RFCTIME"].(time.Time).Format("150405"))
	assert.Equal(t, importStruct["RFCDATE"].(time.Time).Format("20060102"), echoStruct["RFCDATE"].(time.Time).Format("20060102"))
	assert.Equal(t, importStruct["RFCDATA1"], echoStruct["RFCDATA1"])
	assert.Equal(t, importStruct["RFCDATA2"], echoStruct["RFCDATA2"])
	c.Close()
}

func TestConfigParameter(t *testing.T) {
	//rstrip = false
	c, err := ConnectionFromParams(abapSystem())
	if err != nil {
		return
	}
	c.RStrip(false)
	r, _ := c.Call("STFC_CONNECTION", map[string]interface{}{"REQUTEXT": "HELLÖ SÄP"})
	assert.Equal(t, 255, len(reflect.ValueOf(r["ECHOTEXT"]).String()))
	assert.Equal(t, "HELLÖ SÄP", strings.TrimSpace(reflect.ValueOf(r["ECHOTEXT"]).String()))

	//returnImportParams = true
	c, _ = ConnectionFromParams(abapSystem())
	c.ReturnImportParams(true)
	r, _ = c.Call("STFC_CONNECTION", map[string]interface{}{"REQUTEXT": "HELLÖ SÄP"})
	assert.Equal(t, "HELLÖ SÄP", r["REQUTEXT"])
	c.Close()
}

func TestInvalidParameterFunctionCall(t *testing.T) {
	c, err := ConnectionFromParams(abapSystem())
	if err != nil {
		return
	}

	r, err := c.Call("STFC_CONNECTION", map[string]interface{}{"XXX": "wrongParameter"})
	assert.Nil(t, r)
	assert.NotNil(t, err)
	assert.Equal(t, "Could not get the parameter description for \"XXX\"", err.(*RfcError).Description)
	assert.Equal(t, "field 'XXX' not found", err.(*RfcError).ErrorInfo.Message)
	assert.Equal(t, "RFC_INVALID_PARAMETER", err.(*RfcError).ErrorInfo.Code)
	assert.Equal(t, "RFC_INVALID_PARAMETER", err.(*RfcError).ErrorInfo.Key)
	c.Close()
}

func TestErrorFunctionCall(t *testing.T) {
	c, err := ConnectionFromParams(abapSystem())
	if err != nil {
		return
	}

	r, err := c.Call("RFC_RAISE_ERROR", map[string]interface{}{"MESSAGETYPE": "A"})
	assert.Nil(t, r)
	assert.NotNil(t, err)
	assert.Equal(t, "Could not invoke function \"RFC_RAISE_ERROR\"", err.(*RfcError).Description)
	assert.Equal(t, "Function not supported", err.(*RfcError).ErrorInfo.Message)
	assert.Equal(t, "RFC_ABAP_MESSAGE", err.(*RfcError).ErrorInfo.Code)
	assert.Equal(t, "Function not supported", err.(*RfcError).ErrorInfo.Key)
	assert.Equal(t, "SR", err.(*RfcError).ErrorInfo.AbapMsgClass)
	assert.Equal(t, "A", err.(*RfcError).ErrorInfo.AbapMsgType)
	assert.Equal(t, "006", err.(*RfcError).ErrorInfo.AbapMsgNumber)
	assert.Equal(t, "STRING", err.(*RfcError).ErrorInfo.AbapMsgV1)
	c.Close()
}

func abapSystem() ConnectionParameter {
	return ConnectionParameter{
		Dest:   "I64",
		Client: "800",
		User:   "demo",
		Passwd: "welcome",
		Lang:   "EN",
		Ashost: "10.117.24.158	",
		Sysnr:     "00",
		Saprouter: "/H/203.13.155.17/W/xjkb3d/H/172.19.138.120/H/",
	}
}
