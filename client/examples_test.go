// BSD 3-Clause License
//
// Copyright (c) 2022, InterlockLedger
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package client

import (
	"context"
	"fmt"
	"os"
)

func ExampleNewAPIClient() {
	// Creates a new configuration.
	configuration := NewConfiguration()
	// Set the name of the server here
	configuration.BasePath = "https://your_node_server:port/"
	// Sets the required client certificate. Without it, you will not be able to
	// access the node as client certificate is the only authentication method
	// accepted by the client.
	err := configuration.SetClientCertificate("cert.pem", "key.pem")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load the client certificate: %v\n", err)
		os.Exit(1)
	}
	// Create the new client
	client := NewAPIClient(configuration)

	// Query the version of the server.
	version, _, err := client.NodeApi.ApiVersion(context.TODO())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to query the server's version: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("The server's version is %s\n", version)
}
