package client

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestAPIClient_ToGenericSwaggerError(t *testing.T) {
	var c APIClient

	e1 := GenericSwaggerError{}
	e := c.ToGenericSwaggerError(e1)
	assert.Equal(t, e, &e1)

	e = c.ToGenericSwaggerError(&e1)
	assert.Equal(t, e, &e1)

	e = c.ToGenericSwaggerError(fmt.Errorf("dummy"))
	assert.Nil(t, e)
}

func TestGetHeaderInt64(t *testing.T) {
	h := make(http.Header)

	h.Set("a", "9223372036854775807")
	h.Set("b", "-9223372036854775808")
	h.Set("c", "X")

	v, err := GetHeaderInt64(h, "z", 123)
	assert.Nil(t, err)
	assert.Equal(t, int64(123), v)

	v, err = GetHeaderInt64(h, "a", 123)
	assert.Nil(t, err)
	assert.Equal(t, int64(9223372036854775807), v)

	v, err = GetHeaderInt64(h, "b", 123)
	assert.Nil(t, err)
	assert.Equal(t, int64(-9223372036854775808), v)

	_, err = GetHeaderInt64(h, "c", 123)
	assert.Error(t, err)
}
