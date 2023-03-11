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
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"golang.org/x/crypto/pkcs12"
)

/*
Loads a certificate with its private key. Both the certificate and the key must
be in PEM format.
*/
func LoadCertificateWithKey(certificateFile string, keyFile string) (tls.Certificate, error) {
	return tls.LoadX509KeyPair(certificateFile, keyFile)
}

/*
Loads all certificates inside the specified file and returns them as a list. The
certificates must be in PEM format.

It fails if there is no certificates to load or if one of them is invalid.
*/
func LoadCertificate(certificateFile string) ([]*x509.Certificate, error) {
	bytes, err := os.ReadFile(certificateFile)
	if err != nil {
		return nil, err
	}
	return ParseCertificate(bytes)
}

/*
Parses all certificates inside the specified file and returns them as a list. The
certificates must be in PEM format.

It fails if there is no certificates to load or if one of them is invalid.
*/
func ParseCertificate(bytes []byte) ([]*x509.Certificate, error) {
	certs := make([]*x509.Certificate, 0, 1)
	for {
		pemBlock, rest := pem.Decode(bytes)
		if pemBlock == nil {
			break
		} else if pemBlock.Type != "CERTIFICATE" {
			return nil, ErrInvalidCertificateFile
		}
		if cert, err := x509.ParseCertificate(pemBlock.Bytes); err != nil {
			return nil, err
		} else {
			certs = append(certs, cert)
		}
		if len(rest) == 0 {
			break
		}
		bytes = rest
	}
	if len(certs) == 0 {
		return nil, fmt.Errorf("no certifcates in the file: %w", ErrInvalidCertificateFile)
	}
	return certs, nil
}

/*
Loads the private key from a PEM file. It can load RSA, ECDSA and EdDSA keys.
*/
func LoadPrivateKey(keyFile string) (crypto.PrivateKey, error) {
	bytes, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	return ParsePrivateKey(bytes)
}

/*
Parses the private key from a PEM file. It can load RSA, ECDSA and EdDSA keys.
*/
func ParsePrivateKey(bytes []byte) (crypto.PrivateKey, error) {
	pemBlock, _ := pem.Decode(bytes)
	if pemBlock == nil || pemBlock.Type != "PRIVATE KEY" {
		return nil, ErrInvalidPrivateKey
	}
	// This code is based on the code inside standard library tls.
	if key, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes); err == nil {
		return key, nil
	}
	if key, err := x509.ParsePKCS8PrivateKey(pemBlock.Bytes); err == nil {
		switch key.(type) {
		case *rsa.PrivateKey, *ecdsa.PrivateKey, ed25519.PrivateKey:
			return key, nil
		default:
			return ErrInvalidPrivateKey, nil
		}
	}
	if key, err := x509.ParseECPrivateKey(pemBlock.Bytes); err == nil {
		return key, nil
	}
	return nil, ErrInvalidPrivateKey
}

/*
Loads a certificate with its private key from a PKCS #12 file.
*/
func LoadCertificateWithKeyFromPKCS12(file string, password string) (tls.Certificate, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return tls.Certificate{}, err
	}
	return ParseCertificateWithKeyFromPKCS12(bytes, password)
}

/*
Parses a certificate with its private key from a PKCS #12 file.
*/
func ParseCertificateWithKeyFromPKCS12(bytes []byte, password string) (tls.Certificate, error) {
	blocks, err := pkcs12.ToPEM(bytes, password)
	if err != nil {
		return tls.Certificate{}, err
	}
	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}
	return tls.X509KeyPair(pemData, pemData)
}
