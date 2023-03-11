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
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"

	"github.com/interlockledger/go-iltags/tags"
	"github.com/interlockledger/go-iltags/tags/impl"
	"github.com/interlockledger/go-iltags/utils"
)

var (
	// The private key is invalid.
	ErrInvalidPrivateKey = fmt.Errorf("invalid private key")
	// The public key is invalid.
	ErrInvalidPublicKey = fmt.Errorf("invalid public key")
	// The size of the IV is invalid.
	ErrInvalidBlockCipherIv = fmt.Errorf("invalid block size IV")
	// The encrypted message is invalid.
	ErrInvalidEncryptedMessage = fmt.Errorf("invalid encrypted message")
	// The message padding is invalid.
	ErrInvalidPadding = fmt.Errorf("invalid message padding")
	// Invalid certificate file.
	ErrInvalidCertificateFile = fmt.Errorf("invalid cetificate file")
)

// Try to cast the given private key into a RSA private key.
func castRSAPrivateKey(privateKey crypto.PrivateKey) (*rsa.PrivateKey, error) {
	if k, ok := privateKey.(*rsa.PrivateKey); ok {
		return k, nil
	}
	return nil, ErrInvalidPrivateKey
}

// Try to cast the given public key into a RSA private key.
func castRSAPublicKey(publicKey crypto.PublicKey) (*rsa.PublicKey, error) {
	if k, ok := publicKey.(*rsa.PublicKey); ok {
		return k, nil
	}
	return nil, ErrInvalidPublicKey
}

// Deciphers the message using the specified private key.
func DecryptRSAWithPrivate(privateKey *rsa.PrivateKey, encrypted []byte) ([]byte, error) {
	return rsa.DecryptOAEP(crypto.SHA1.New(), nil, privateKey, encrypted, nil)
}

// Deciphers the message using the specified private key.
func DecryptWithPrivate(privateKey crypto.PrivateKey, encrypted []byte) ([]byte, error) {
	// This code is based on the python version of this code.
	pkey, err := castRSAPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return DecryptRSAWithPrivate(pkey, encrypted)
}

/*
Remove the ISO 10126 padding and return a subslice of plain that contains the
actual data. Due to the way this padding works, it can also removes the PKCS#5,
PKCS#7 and ANSI X9.23 padding as they are special cases of ISO 10126 padding.
*/
func RemoveISO10126Padding(blockSize int, plain []byte) ([]byte, error) {
	paddingSize := int(plain[len(plain)-1])
	if paddingSize == 0 || paddingSize > blockSize {
		return nil, ErrInvalidPadding
	}
	return plain[0 : len(plain)-paddingSize], nil
}

/*
Remove the padding zero padding and return a subslice of plain that contains the
actual data.
*/
func RemoveZeroPadding(plain []byte) []byte {
	lastIndex := -1
	for i := len(plain) - 1; i >= 0; i-- {
		if plain[i] != 0 {
			lastIndex = i
			break
		}
	}
	return plain[0 : lastIndex+1]
}

// Decipher the specified block using the specified key and IV.
func DecipherAESCBC(key, iv, encrypted []byte) ([]byte, int, error) {
	bc, err := aes.NewCipher(key)
	if err != nil {
		return nil, 0, err
	}
	if len(iv)%bc.BlockSize() != 0 {
		return nil, 0, ErrInvalidBlockCipherIv
	}
	if len(encrypted)%bc.BlockSize() != 0 {
		return nil, 0, ErrInvalidEncryptedMessage
	}
	dec := make([]byte, len(encrypted))
	bm := cipher.NewCBCDecrypter(bc, iv)
	bm.CryptBlocks(dec, encrypted)
	return dec, bc.BlockSize(), nil
}

// Decipher the specified block using the specified key and IV.
func DecipherJSON(key, iv, encrypted []byte) (string, error) {
	plain, _, err := DecipherAESCBC(key, iv, encrypted)
	if err != nil {
		return "", err
	}
	defer utils.ShredBytes(plain)
	plain = RemoveZeroPadding(plain)
	return impl.DeserializeStdStringTag(bytes.NewReader(plain))
}

func convertRSAPublicKey(publicKey *rsa.PublicKey) ([]byte, error) {
	nTag := impl.NewStdBytesTag()
	nTag.Payload = publicKey.N.Bytes()
	eTag := impl.NewStdBytesTag()
	eTag.Payload = big.NewInt(int64(publicKey.E)).Bytes()
	rootTag := impl.NewILTagSequenceTag(tags.TagID(40))
	rootTag.Payload = []tags.ILTag{nTag, eTag}
	return tags.ILTagToBytes(rootTag)
}

// Converts the given public key into a format suitable for use with IL2 API.
func ConvertPublicKey(publicKey crypto.PublicKey) ([]byte, error) {
	// For now, only RSA is supported.
	p, err := castRSAPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	return convertRSAPublicKey(p)
}

/*
Computes the public key hash according to the IL2 standard.
*/
func CreatePublicKeyHashBin(publicKey []byte) (string, error) {
	sha256 := crypto.SHA256.New()
	n, err := sha256.Write(publicKey)
	if err != nil {
		return "", err
	}
	if n != len(publicKey) {
		return "", fmt.Errorf("unable to hash the key")
	}
	h := sha256.Sum(nil)
	s := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(h)
	sb := strings.Builder{}
	sb.WriteString(s)
	sb.WriteString("#SHA256")
	return sb.String(), nil
}

/*
Computes the public key hash according to the IL2 standard.
*/
func CreatePublicKeyHash(publicKey crypto.PublicKey) (string, error) {
	// Convert
	bin, err := ConvertPublicKey(publicKey)
	if err != nil {
		return "", err
	}
	return CreatePublicKeyHashBin(bin)
}

/*
Creates a probably unique ReaderId from a string of bytes. It may be a key, the
certificate payload or even a random value.
*/
func CreateReaderId(data []byte) (string, error) {

	sha1 := crypto.SHA1.New()
	n, err := sha1.Write(data)
	if err != nil {
		return "", err
	}
	if n != len(data) {
		return "", fmt.Errorf("unable to hash the certificate")
	}
	h := sha1.Sum(nil)
	s := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(h)
	sb := strings.Builder{}
	sb.WriteString("Key!")
	sb.WriteString(s)
	sb.WriteString("#SHA1")
	return sb.String(), nil
}
