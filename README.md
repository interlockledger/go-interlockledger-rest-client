# go-interlockledger-rest-client

Our Go client to the InterlockLedger node REST API.

## Requirements

This library was developed using Go 1.18 but it should work perfectly on older versions.

## How to use it

To see how to use this libraty, check the code `internal/main.go` as it
describes how to instantiate the client and use it to contact the node.

## Notes about the code

This code was originally genearated using the 
[Swagger Codegen](https://github.com/swagger-api/swagger-codegen.git) and has
been manually patched to reflect the needs of the **InterlockLedger** API, 
specially regarding the client certificate authentication.

In order to retain compatibility with future versions of this code, all future
releases will be manually updated as the code generator.

## License

This library is released under a BSD-3-Clause License.



