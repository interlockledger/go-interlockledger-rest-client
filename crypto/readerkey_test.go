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
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewReaderKey(t *testing.T) {
	var pubKey crypto.PublicKey
	var privKey crypto.PublicKey

	pubKey = &rsa.PublicKey{N: big.NewInt(12345679), E: 3}
	privKey = &rsa.PrivateKey{}
	rk, err := NewReaderKey(pubKey, privKey)
	assert.Nil(t, err)
	assert.NotNil(t, rk)

	pubKey = &ecdsa.PublicKey{}
	privKey = &ecdsa.PrivateKey{}
	rk, err = NewReaderKey(pubKey, privKey)
	assert.Error(t, err)
	assert.Nil(t, rk)
}

func TestNewReaderKeyFromPrivateKey(t *testing.T) {
	certFile := getSampleFile("cert.pem")
	keyFile := getSampleFile("key.pem")

	pair, err := LoadCertificateWithKey(certFile, keyFile)
	require.Nil(t, err)

	rk, err := NewReaderKeyFromPrivateKey(pair.PrivateKey)
	assert.Nil(t, err)
	assert.NotNil(t, rk)

	rk, err = NewReaderKeyFromPrivateKey(&ecdsa.PrivateKey{})
	assert.Error(t, err)
	assert.Nil(t, rk)
}

func ExampleNewReaderKeyFromPrivateKey() {
	pair, err := LoadCertificateWithKey("cert.pem", "key.pem")
	if err != nil {
		fmt.Printf("Unable to load the key pair: %v\n", err.Error())
		os.Exit(1)
	}
	rk, err := NewReaderKeyFromPrivateKey(pair.PrivateKey)
	if err != nil {
		fmt.Printf("Unable create the ReaderKey: %v\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("The public key hash is: %s\n", rk.PublicKeyHash())
}

func createTestReaderKey(t *testing.T) ReaderKey {
	certFile := getSampleFile("cert.pem")
	keyFile := getSampleFile("key.pem")

	pair, err := LoadCertificateWithKey(certFile, keyFile)
	require.Nil(t, err)
	rk, err := NewReaderKeyFromPrivateKey(pair.PrivateKey)
	require.Nil(t, err)
	return rk
}

func TestReaderKeyPublicKeyHash(t *testing.T) {

	rk := createTestReaderKey(t)
	assert.Equal(t, "mQ9L7BboXNmIDL_4HbNHjWGW1tMQUMelTzbnK4lHg-E#SHA256", rk.PublicKeyHash())
}

func TestReaderKeyEncodedPublicKey(t *testing.T) {

	rk := createTestReaderKey(t)
	key, id, err := rk.EncodedPublicKey()
	assert.Nil(t, err)
	assert.Equal(t, "KPgQEPgImalwNPIEhRmEy_69WCGSVidHhP8_MqzOZSBn2AwZsjdDBy84F6TfomZQfCPTf4ND7j8FpHhOpASK5gjBnhAo65QxKYytxsyUXYJ1CO9hB55XJgn7q0B94SYhaO7Va5N2ixailbLbP9D9rnkVuVDp2rPjeahVKC4lYgL8Od_9dvOw5C_emFphJszKKx9-_0ot_bNefcMa_jKJuhwuaNkaArepMfQT6-TvVukTR31BRxSg5wrMNDOrLClKFHTtPtApzDW0sMuCyNKmOOHGwB4jzDDVxqcu46h7BEIKxPYtMCTCqs0M1i5gwZmNSaFu6m3RqrLvttMC2-vVJadL0QvTPRADAQAB", key)
	assert.Equal(t, "Key!sunrzo4_cbjau-a_02egWdYtBJ4#SHA1", id)
}

func TestUnwrap(t *testing.T) {

	rk := createTestReaderKey(t)
	b, err := rk.Unwrap(SAMPLE_ENC_IV)
	assert.Nil(t, err)
	assert.Equal(t, SAMPLE_IV, b)
}
