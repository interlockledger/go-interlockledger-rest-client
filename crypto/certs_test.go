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

package crypto

import (
	"crypto/rsa"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadCertificateWithKey(t *testing.T) {
	certFile := getSampleFile("cert.pem")
	keyFile := getSampleFile("key.pem")

	pair, err := LoadCertificateWithKey(certFile, keyFile)
	assert.Nil(t, err)
	assert.NotNil(t, pair.PrivateKey)
}

func TestLoadCertificate(t *testing.T) {
	certFile := getSampleFile("cert.pem")
	cert, err := LoadCertificate(certFile)
	assert.Nil(t, err)
	assert.NotNil(t, cert)
	assert.Equal(t, 1, len(cert))

	certFile = getSampleFile("certs.pem")
	cert, err = LoadCertificate(certFile)
	assert.Nil(t, err)
	assert.NotNil(t, cert)
	assert.Equal(t, 2, len(cert))

	certFile = getSampleFile("empty-certs.pem")
	cert, err = LoadCertificate(certFile)
	assert.Error(t, err)
	assert.Nil(t, cert)

	certFile = getSampleFile("bad-certs.pem")
	cert, err = LoadCertificate(certFile)
	assert.Error(t, err)
	assert.Nil(t, cert)

	certFile = getSampleFile("bad-entry-certs.pem")
	cert, err = LoadCertificate(certFile)
	assert.Error(t, err)
	assert.Nil(t, cert)

	cert, err = LoadCertificate("this file does not exist.")
	assert.Error(t, err)
	assert.Nil(t, cert)
}

func TestLoadPrivateKey(t *testing.T) {
	keyFile := getSampleFile("key.pem")

	key, err := LoadPrivateKey(keyFile)
	assert.Nil(t, err)
	var exp *rsa.PrivateKey
	assert.IsType(t, exp, key)

	certFile := getSampleFile("cert.pem")
	key, err = LoadPrivateKey(certFile)
	assert.ErrorIs(t, err, ErrInvalidPrivateKey)
	assert.Nil(t, key)
}

func TestLoadCertificateWithKeyFromPKCS12(t *testing.T) {

	cert, err := LoadCertificateWithKeyFromPKCS12(getSampleFile("sample.pfx"), "password")
	assert.Nil(t, err)
	assert.NotNil(t, cert.PrivateKey)
}
