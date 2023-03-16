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
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"

	"github.com/interlockledger/go-interlockledger-rest-client/crypto"
)

// Configuration of the client.
type Configuration struct {
	// Base path of the server API.
	BasePath string `json:"basePath,omitempty"`
	// The name of the host.
	Host string `json:"host,omitempty"`
	// The scheme to be used.
	Scheme string `json:"scheme,omitempty"`
	// Set of headers to be sent on all requests.
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	// The client's user agent.
	UserAgent string `json:"userAgent,omitempty"`
	// If true, the server connections will not be validated.
	NoServerVerification bool `json:"noServerVerification"`
	// Path to the client certificate file (PEM).
	CertFile string `json:"certFile,omitempty"`
	// Path to the client key file (PEM).
	KeyFile string `json:"keyFile,omitempty"`
	// Path to the client PFX certificate/key file.
	PFXFile string `json:"pfxFile,omitempty"`
	// Password of the client PFX certificate/key file.
	PFXPassword string `json:"pfxPassword,omitempty"`
	// The client associated with this configuration.
	HTTPClient *http.Client `json:"-"`
	// The set of client certificates to be used. It will be initialized according
	// to the current configuration if necessary.
	ClientCertificates []tls.Certificate `json:"-"`
	// The certificate pool to be used. It will be automatically initialized
	// when required.
	CertPool *x509.CertPool `json:"-"`
}

// Creates a new client configuration.
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "/",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}
	return cfg
}

// Adds custom headers to all connections from this client.
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

/*
Loads the client certificate required to access the API and sets the
`http.Client`.
*/
func (c *Configuration) SetClientCertificate(certificateFile string, keyFile string) error {
	return c.SetClientCertificateEx(certificateFile, keyFile, c.NoServerVerification)
}

/*
Loads the client certificate required to access the API and sets the
`http.Client`.
*/
func (c *Configuration) SetClientCertificateEx(
	certificateFile string, keyFile string, noServerVerification bool) error {
	cert, err := crypto.LoadCertificateWithKey(certificateFile, keyFile)
	if err != nil {
		return err
	}
	return c.SetClientCert(cert, noServerVerification)
}

/*
Loads the client certificate required to access the API and sets the
`http.Client`.
*/
func (c *Configuration) SetClientCertificatePKCS12(file string, password string) error {
	return c.SetClientCertificatePKCS12Ex(file, password, c.NoServerVerification)
}

/*
Loads the client certificate required to access the API and sets the
`http.Client`.
*/
func (c *Configuration) SetClientCertificatePKCS12Ex(
	file string, password string, noServerVerification bool) error {
	cert, err := crypto.LoadCertificateWithKeyFromPKCS12(file, password)
	if err != nil {
		return err
	}
	return c.SetClientCert(cert, noServerVerification)
}

/*
Loads the client certificate required to access the API and sets the
`http.Client`.
*/
func (c *Configuration) SetClientCert(cert tls.Certificate, noServerVerification bool) error {
	c.ClientCertificates = []tls.Certificate{cert}
	return c.createHTTPClient(noServerVerification)
}

/*
Loads the client certificate required to access the API and sets the
`http.Client`.
*/
func (c *Configuration) createHTTPClient(noServerVerification bool) error {
	if c.CertPool == nil {
		pool, err := x509.SystemCertPool()
		if err != nil {
			return err
		}
		c.CertPool = pool
	}
	tlsConfig := &tls.Config{
		Certificates:       c.ClientCertificates,
		RootCAs:            c.CertPool,
		InsecureSkipVerify: noServerVerification,
	}
	c.HTTPClient = &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}}
	return nil
}

/*
Initializes the inner HTTPClient using the current parameters.

It will load the client certificate from either PFXFile (preffered) or
CertFile/KeyFile parameters and enable or disable the server
verification based on the value of NoServerVerification.

If called more than once, the previos HTTPClient will be replaced by the new one.

On success, if the fields CertPool and ClientCertificates are not initialized,
they will be initialized according to the current configuration.

If fails if there is no valid client certificate to load.
*/
func (c *Configuration) Init() error {

	if c.ClientCertificates != nil {
		return c.createHTTPClient(c.NoServerVerification)
	} else if c.PFXFile != "" {
		return c.SetClientCertificatePKCS12Ex(c.PFXFile, c.PFXPassword, c.NoServerVerification)
	} else if c.CertFile != "" {
		return c.SetClientCertificateEx(c.CertFile, c.KeyFile, c.NoServerVerification)
	}
	return fmt.Errorf("no client certificate set")
}
