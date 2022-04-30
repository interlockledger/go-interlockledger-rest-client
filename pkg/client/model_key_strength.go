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

/*
 * InterlockLedger Node REST API
 *
 * A node instance inside the peer-to-peer network of the InterlockLedger
 *
 * API version: v7.2.0
 * Contact: core@interlockledger.network
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type KeyStrength string

// List of KeyStrength
const (
	NORMAL_KeyStrength       KeyStrength = "Normal"
	STRONG_KeyStrength       KeyStrength = "Strong"
	EXTRA_STRONG_KeyStrength KeyStrength = "ExtraStrong"
	MEGA_STRONG_KeyStrength  KeyStrength = "MegaStrong"
	SUPER_STRONG_KeyStrength KeyStrength = "SuperStrong"
	HYPER_STRONG_KeyStrength KeyStrength = "HyperStrong"
	ULTRA_STRONG_KeyStrength KeyStrength = "UltraStrong"
)
