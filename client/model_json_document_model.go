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
	"time"
)

type JsonDocumentModel struct {
	ApplicationId int64 `json:"applicationId,omitempty"`
	// Chain unique ID
	ChainId      string     `json:"chainId,omitempty"`
	CreatedAt    time.Time  `json:"createdAt,omitempty"`
	Hash         string     `json:"hash,omitempty"`
	Network      *NetworkId `json:"network,omitempty"`
	PayloadTagId int64      `json:"payloadTagId,omitempty"`
	// A universal record reference in the form networkName:chainId@recordSerial
	Reference     string              `json:"reference,omitempty"`
	Serial        int64               `json:"serial,omitempty"`
	Type_         *RecordType         `json:"type,omitempty"`
	Version       int32               `json:"version,omitempty"`
	EncryptedJson *EncryptedTextModel `json:"encryptedJson,omitempty"`
	JsonText      string              `json:"jsonText,omitempty"`
}
