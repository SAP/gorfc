# SAP NW RFC Connector for GO

The _gorfc_ package provides bindings for SAP NW RFC Library, for a comfortable way of calling remote enabled ABAP function modules (RFMs) from GO.

Current release is fully functional on Linux and experimental on Windows, see the [Issue #1](https://github.com/SAP/gorfc/issues/1).

## Table of contents

* [Platforms and Prerequisites](#platforms)
* [Install](#install)
	* [GO](#install-go)
	* [SAP NW RFC Library](#install-rfcsdk)
	* [gorfc](#install-gorfc)
* [Getting Started](#getting-started)
* [To Do](#todo)
* [References](#references)

## <a name="platforms"></a> Platforms and Prerequisites

The SAP NW RFC Library is a prerequsite for using the GO RFC connector and must be installed on a same system. It is available on platforms supported by GO, except OSX.

A prerequisite to download _SAP NW RFC Library_ is having a **customer or partner account** on _SAP Service Marketplace_ . If you are SAP employee please check SAP OSS note [1037575 - Software download authorizations for SAP employees](http://service.sap.com/sap/support/notes/1037575).

_SAP NW RFC Library_ is fully backwards compatible, supporting all NetWeaver systems, from today, down to release R/3 4.0. You can always use the newest version released on Service Marketplace and connect to older systems as well.

## <a name="install"></a>Install

To start using SAP NW RFC Connector for GO, you shall:

1. Install and Configure GO
2. Install SAP NW RFC Library for your platform
3. Install GORFC package

### <a name="install-go"></a>Install and Configure GO

If you are new to GO, the GO distribution shall be installed first, following [GO Installation](#ref1) and [GO Configuration](#ref2) instructions. See also [GO Environment Variables](#ref3).

#### Windows Config Example

After running the [MSI installer](https://golang.org/dl/), the default C:\Go folder is created and the _GOROOT_ system variable is set to C:\Go\.

Create the GO work environment directory:

```shell
cd c:\
mkdir workspace
```

Set the environment user varialbes GOPATH and GOBIN, add the bin subdirectories to PATH and restart the Windows shell.

```shell
GOPATH = C:\workspace
GOBIN = %GOPATH%\bin
PATH = %GOROOT%\bin;%GOBIN%:%PATH%
```

See also [GO on Windows Example](#ref4).

#### Linux

The work environment setup works the same way like on Windows and [these instructions](https://github.com/golang/go/wiki/Ubuntu) describe the installation on Ubuntu Linux for example.

### <a name="install-rfcsdk"></a>Install SAP NW RFC Library

To obtain and install _SAP NW RFC Library_ from _SAP Service Marketplace_, you can follow [the same instructions as for Python or nodejs RFC connectors](http://sap.github.io/PyRFC/install.html#install-c-connector).

### <a name="install-gorfc"></a>Install GORFC

To install _gorfc_ and dependencies, run following commands:

```bash
export CGO_CFLAGS="-I $SAPNWRFC_HOME/include"
export CGO_LDFLAGS="-L $SAPNWRFC_HOME/lib"
go get github.com/stretchr/testify
go get github.com/sap/gorfc
cd $GOPATH/src/github.com/sap/gorfc
go build
go install
```

To test the installation, run the example provided:

```bash
cd $GOPATH/src/github.com/sap/gorfc/example
go run hello_gorfc.go
```

## <a name="getting-started"></a>Getting Started

See the _hello_gorfc.go_ example and _gorfc_test.go_ unit tests.

The GO RFC Connector follows the same principles and the implementation model of [Python](https://github.com/SAP/PyRFC) and [nodejs](https://github.com/SAP/node-rfc) RFC connectors and you may check examples and documentation there as well.

```go
package main

import (
    "fmt"
    "github.com/sap/gorfc/gorfc"
    "github.com/stretchr/testify/assert"
    "reflect"
    "testing"
    "time"
)

func abapSystem() gorfc.ConnectionParameter {
    return gorfc.ConnectionParameter{
        Dest:      "I64",
        Client:    "800",
        User:      "demo",
        Passwd:    "welcome",
        Lang:      "EN",
        Ashost:    "11.111.11.111",
        Sysnr:     "00",
        Saprouter: "/H/222.22.222.22/S/2222/W/xxxxx/H/222.22.222.222/H/",
    }   
}

func main() {
    c, _ := gorfc.Connection(abapSystem())
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
    //  assert.Equal(t, importStruct["RFCHEX3"], echoStruct["RFCHEX3"])
    assert.Equal(t, importStruct["RFCTIME"].(time.Time).Format("150405"), echoStruct["RFCTIME"].(time.Time).Format("15.
    assert.Equal(t, importStruct["RFCDATE"].(time.Time).Format("20060102"), e/Users/d037732/Downloads/gorfc/README.mdchoStruct["RFCDATE"].(time.Time).Format(".
    assert.Equal(t, importStruct["RFCDATA1"], echoStruct["RFCDATA1"])
    assert.Equal(t, importStruct["RFCDATA2"], echoStruct["RFCDATA2"])

    fmt.Println(reflect.TypeOf(importStruct["RFCDATE"]))
    fmt.Println(reflect.TypeOf(importStruct["RFCTIME"]))

    c.Close()
```

## <a name="todo"></a>To Do

* Improve the documentation
* Fix Windows compiler flags

## <a name="references"></a>References

* <a name="ref1">[GO Installation](https://golang.org/doc/install)
* <a name="ref2">[GO Configuration](https://golang.org/doc/code.html)
* <a name="ref3">[GO Environment Variables](https://golang.org/cmd/go/#hdr-Environment_variables)
* <a name="ref4">[GO on Windows Example](http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/)
* <a name="ref5">[Another GO on Windows Example](https://github.com/abourget/getting-started-with-golang/blob/master/Getting_Started_for_Windows.md)
