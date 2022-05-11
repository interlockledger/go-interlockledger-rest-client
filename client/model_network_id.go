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

type NetworkId struct {
	AsBytes        string `json:"asBytes,omitempty"`
	AsHex          string `json:"asHex,omitempty"`
	AsULong        int64  `json:"asULong,omitempty"`
	Checksum       int32  `json:"checksum,omitempty"`
	DefaultAddress string `json:"defaultAddress,omitempty"`
	DefaultPort    int32  `json:"defaultPort,omitempty"`
	// Chain unique ID
	Genesis     string `json:"genesis,omitempty"`
	Inactive    bool   `json:"inactive,omitempty"`
	IsCustom    bool   `json:"isCustom,omitempty"`
	IsEmpty     bool   `json:"isEmpty,omitempty"`
	IsError     bool   `json:"isError,omitempty"`
	IsInvalid   bool   `json:"isInvalid,omitempty"`
	IsLocalOnly bool   `json:"isLocalOnly,omitempty"`
	// Chain unique ID
	LicenseSigners        string         `json:"licenseSigners,omitempty"`
	Name                  string         `json:"name,omitempty"`
	ProxySeeder           *NetworkSeed   `json:"proxySeeder,omitempty"`
	Seeders               []NetworkSeed  `json:"seeders,omitempty"`
	TextualRepresentation string         `json:"textualRepresentation,omitempty"`
	Type_                 *NetworkIdType `json:"type,omitempty"`
}
