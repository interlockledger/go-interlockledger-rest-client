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
	"crypto/rsa"
	"encoding/base64"
)

// Header of the public key.
var PUBLIC_KEY_HEADER = []byte("PubKey!")

func keySuffix(publicKey crypto.PublicKey) string {

	switch publicKey.(type) {
	case *rsa.PublicKey:
		return "#RSA"
	default:
		return ""
	}
}

/*
This is the interface of all reader key. Reader keys are used to decipher some
types of payloads stored inside IL2 blocks.
*/
type ReaderKey interface {
	/*
		Returns the public key hash of this key.
	*/
	PublicKeyHash() string
	/*
		Returns the public key.
	*/
	PublicKey() crypto.PublicKey
	/*
		Returns the encoded public key suitable to be sent to the node.

		It returns the encoded key and a reader key id derived from the public key.
	*/
	EncodedPublicKey() (string, string, error)
	/*
		Unwraps the given wrapped value with the specified private key.
	*/
	Unwrap(enc []byte) ([]byte, error)
	/*
		Returns true if the private key is present.
	*/
	HasPrivateKey() bool
}

/*
Creates a new ReaderKey from a public and private key.
*/
func NewReaderKey(publicKey crypto.PublicKey, privateKey crypto.PrivateKey) (ReaderKey, error) {
	publicKeyHash, err := CreatePublicKeyHash(publicKey)
	if err != nil {
		return nil, err
	}
	return &readerKeyImpl{privateKey: privateKey, publicKey: publicKey, publicKeyHash: publicKeyHash}, nil
}

/*
Helper function that attempts to create a new ReaderKey from the private key. It
will succeed only if the private key format is known to this function.

For now, only RSA keys are supported.
*/
func NewReaderKeyFromPrivateKey(privateKey crypto.PrivateKey) (ReaderKey, error) {
	pk, err := castRSAPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return NewReaderKey(&pk.PublicKey, privateKey)
}

/*
Basic implementation of a ReaderKey.
*/
type readerKeyImpl struct {
	privateKey    crypto.PrivateKey
	publicKey     crypto.PublicKey
	publicKeyHash string
}

func (k *readerKeyImpl) PublicKeyHash() string {
	return k.publicKeyHash
}

func (k *readerKeyImpl) PublicKey() crypto.PublicKey {
	return k.publicKey
}

func (k *readerKeyImpl) EncodedPublicKey() (string, string, error) {
	encBin, err := ConvertPublicKey(k.PublicKey())
	if err != nil {
		return "", "", err
	}
	readerId, err := CreateReaderId(encBin)
	if err != nil {
		return "", "", err
	}
	// Compute the value
	suffix := keySuffix(k.publicKey)
	encBinLen := base64.URLEncoding.EncodedLen(len(encBin))
	headerLen := len(PUBLIC_KEY_HEADER)
	tmp := make([]byte, headerLen+encBinLen+len(suffix))
	copy(tmp[0:], PUBLIC_KEY_HEADER)
	base64.URLEncoding.Encode(tmp[headerLen:], encBin)
	copy(tmp[headerLen+encBinLen:], []byte(suffix))
	return string(tmp), readerId, nil
}

func (k *readerKeyImpl) Unwrap(enc []byte) ([]byte, error) {
	if k.HasPrivateKey() {
		return DecryptWithPrivate(k.privateKey, enc)
	} else {
		return nil, ErrInvalidPrivateKey
	}
}

func (k *readerKeyImpl) HasPrivateKey() bool {
	return k.privateKey != nil
}
