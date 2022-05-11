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

// Node details
type NodeDetailsModel struct {
	// List of owned chains, only the ids
	Chains []string `json:"chains,omitempty"`
	// Mapping color
	Color string `json:"color,omitempty"`
	// Dictionary with values for extensions on node configuration
	Extensions map[string]string `json:"extensions,omitempty"`
	// Unique node id
	Id string `json:"id"`
	// Node name
	Name string `json:"name,omitempty"`
	// Network this node participates on
	Network string `json:"network,omitempty"`
	// Node owner id [Optional]
	OwnerId string `json:"ownerId,omitempty"`
	// Node owner name [Optional]
	OwnerName string `json:"ownerName,omitempty"`
	// Node url for Peer2Peer protocol
	PeerAddress string `json:"peerAddress,omitempty"`
	// List of active roles running in the node
	Roles            []string          `json:"roles,omitempty"`
	SoftwareVersions *SoftwareVersions `json:"softwareVersions,omitempty"`
}
