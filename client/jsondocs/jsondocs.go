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

package jsondocs

import (
	"fmt"

	"github.com/interlockledger/go-iltags/utils"
	"github.com/interlockledger/go-interlockledger-rest-client/client/models"
	mycrypto "github.com/interlockledger/go-interlockledger-rest-client/crypto"
)

var (
	// If the cipher scheme is unsupported.
	ErrUnsupportedCipher = fmt.Errorf("unsupported cipher")
	// The key is not available.
	ErrKeyNotAvailable = fmt.Errorf("key not available")
	// The current key is not a reading key for the given entry.
	ErrNotAReadingKey = fmt.Errorf("not a reading key")
)

func decipherJSONCore(key mycrypto.ReaderKey, encKey, encIV, encrypted []byte) (string, error) {
	iv, err := key.Unwrap(encIV)
	if err != nil {
		return "", err
	}
	defer utils.ShredBytes(iv)
	k, err := key.Unwrap(encKey)
	if err != nil {
		return "", err
	}
	defer utils.ShredBytes(k)
	return mycrypto.DecipherJSON(k, iv, encrypted)
}

func decipherJSONProcessParameters(key mycrypto.ReaderKey, params *models.ReadingKeyModel, cipherText string) (string, error) {
	binIV, err := models.DecodeBytes(params.EncryptedIV)
	if err != nil {
		return "", err
	}
	defer utils.ShredBytes(binIV)
	binKey, err := models.DecodeBytes(params.EncryptedKey)
	if err != nil {
		return "", err
	}
	defer utils.ShredBytes(binKey)
	binEnc, err := models.DecodeBytes(cipherText)
	if err != nil {
		return "", err
	}
	return decipherJSONCore(key, binKey, binIV, binEnc)
}

/*
Deciphers JSON received from the server using the specified reader key.
*/
func DecipherJSON(key mycrypto.ReaderKey, json *models.JsonDocumentModel) (string, error) {
	if json.EncryptedJson == nil {
		return "", fmt.Errorf("encryptedJson is not set")
	}
	if json.EncryptedJson.Cipher == nil || *json.EncryptedJson.Cipher != models.AES256_CipherAlgorithm {
		return "", ErrUnsupportedCipher
	}
	k := json.EncryptedJson.FindReadingKey(key.PublicKeyHash())
	if k == nil {
		return "", ErrNotAReadingKey
	}
	return decipherJSONProcessParameters(key, k, json.EncryptedJson.CipherText)
}
