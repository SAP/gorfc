# Run go fmt against code
fmt:
	go fmt ./...
	goimports -w -local github.com/sap/gorfc .
