# SAP NW RFC Connector for GO

**this is only a fork, but it seems that [origin](https://github.com/SAP/gorfc) is very *static***  

**all credit for this code goes to [@bsrdjan](https://github.com/bsrdjan)**  

The **saprfc** package provides bindings for **SAP NW RFC Library**, for an easy way of interacting with SAP systems

The goal of this fork is the best possible compatibility with Linux, if you want work with Windows use the [original package](https://github.com/SAP/gorfc) 

## Table of contents

* [Platforms and Prerequisites](#platforms)
* [Install](#install)
* [Getting Started](#getting-started)
* [To Do](#todo)
* [References](#references)

## Platforms and Prerequisites

The SAP NW RFC Library is a prerequsite for using the GO RFC connector and must be installed on a same system. It is available on platforms supported by GO, except OSX.

A prerequisite to download _SAP NW RFC Library_ is having a **customer or partner account** on _SAP Service Marketplace_ . If you are SAP employee please check SAP OSS note [1037575 - Software download authorizations for SAP employees](http://service.sap.com/sap/support/notes/1037575).

_SAP NW RFC Library_ is fully backwards compatible, supporting all NetWeaver systems, from today, down to release R/3 4.0. You can always use the newest version released on Service Marketplace and connect to older systems as well.

## Install

To start using SAP NW RFC Connector for GO, you shall:

1. [Install and Configure Golang](https://golang.org/doc/install)
2. Install SAP NW RFC Library for your platform
3. Install SAPRFC package

### Install SAP NW RFC Library

To obtain and install _SAP NW RFC Library_ from _SAP Service Marketplace_, you can follow [the same instructions as for Python or nodejs RFC connectors](http://sap.github.io/PyRFC/install.html#install-c-connector).
The Download is [here](https://launchpad.support.sap.com/#/softwarecenter/template/products/%20_APP=00200682500000001943&_EVENT=DISPHIER&HEADER=Y&FUNCTIONBAR=N&EVENT=TREE&NE=NAVIGATE&ENR=01200314690200010197&V=MAINT&TA=ACTUAL&PAGE=SEARCH), but the SAP page is the worst, maybe it's better to search for a torrent or ask a friend at SAP.

### Install SAPRFC

To install _saprfc_ and dependencies, run following commands:

```bash
export CGO_CFLAGS="-I $SAPNWRFC_HOME/include"
export CGO_LDFLAGS="-L $SAPNWRFC_HOME/lib"
go get simonwaldherr.de/go/saprfc
cd $GOPATH/src/simonwaldherr.de/go/saprfc
go build
go install
```

## Getting Started

See the _hello_gorfc.go_ example and _saprfc_test.go_ unit tests.

The GO RFC Connector follows the same principles and the implementation model of [Python](https://github.com/SAP/PyRFC) and [nodejs](https://github.com/SAP/node-rfc) RFC connectors and you may check examples and documentation there as well.

```go
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
```

## To Do

* Improve the documentation
* Fix Windows compiler flags

## References

* [GO Installation](https://golang.org/doc/install)
* [GO Configuration](https://golang.org/doc/code.html)
* [GO Environment Variables](https://golang.org/cmd/go/#hdr-Environment_variables)
* [GO on Windows Example](http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/)
* [Another GO on Windows Example](https://github.com/abourget/getting-started-with-golang/blob/master/Getting_Started_for_Windows.md)
