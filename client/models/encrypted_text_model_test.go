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

package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptedTextModelFindReadingKey(t *testing.T) {
	var m EncryptedTextModel

	assert.Nil(t, m.FindReadingKey("hash1"))

	m.ReadingKeys = append(m.ReadingKeys, ReadingKeyModel{PublicKeyHash: "hash1"})
	m.ReadingKeys = append(m.ReadingKeys, ReadingKeyModel{PublicKeyHash: "hash2"})
	m.ReadingKeys = append(m.ReadingKeys, ReadingKeyModel{PublicKeyHash: "hash3"})

	assert.Equal(t, "hash1", m.FindReadingKey("hash1").PublicKeyHash)
	assert.Equal(t, "hash2", m.FindReadingKey("hash2").PublicKeyHash)
	assert.Equal(t, "hash3", m.FindReadingKey("hash3").PublicKeyHash)
	assert.Nil(t, m.FindReadingKey("hash4"))
}
