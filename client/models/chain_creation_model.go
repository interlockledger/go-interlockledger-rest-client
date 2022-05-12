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

// Chain creation parameters
type ChainCreationModel struct {
	// List of additional apps (only the numeric ids)
	AdditionalApps []int64 `json:"additionalApps,omitempty"`
	// API certificates to authorize with corresponding permissions
	ApiCertificates []CertificatePermitModel `json:"apiCertificates,omitempty"`
	// Description (perhaps intended primary usage) [Optional]
	Description string `json:"description,omitempty"`
	// Emergency closing key password
	EmergencyClosingKeyPassword string       `json:"emergencyClosingKeyPassword"`
	EmergencyClosingKeyStrength *KeyStrength `json:"emergencyClosingKeyStrength,omitempty"`
	KeysAlgorithm               *Algorithm   `json:"keysAlgorithm,omitempty"`
	// App/Key management key password
	ManagementKeyPassword string       `json:"managementKeyPassword"`
	ManagementKeyStrength *KeyStrength `json:"managementKeyStrength,omitempty"`
	// Name
	Name                 string       `json:"name"`
	OperatingKeyStrength *KeyStrength `json:"operatingKeyStrength,omitempty"`
	// Parent record Id [Optional]
	Parent string `json:"parent,omitempty"`
}
