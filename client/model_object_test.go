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
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const SAMPLE_JSON string = "{\"int\": 1234, \"float\": 12345.6, \"string\": \"this is a string\"," +
	" \"bool\" : true, \"null\": null, \"array\": [1, true], \"object\":{}}"

func TestMapObjectType(t *testing.T) {

	assert.Equal(t, ObjectFieldNil, MapObjectType(nil))
	assert.Equal(t, ObjectFieldString, MapObjectType(""))
	assert.Equal(t, ObjectFieldBool, MapObjectType(true))
	assert.Equal(t, ObjectFieldNumber, MapObjectType(float64(1)))
	assert.Equal(t, ObjectFieldObject, MapObjectType(Object{}))
	assert.Equal(t, ObjectFieldObject, MapObjectType(map[string]any{}))
	assert.Equal(t, ObjectFieldArray, MapObjectType([]any{}))
	assert.Equal(t, ObjectFieldArray, MapObjectType(Array{}))
}

func TestFieldAsString(t *testing.T) {

	v, err := FieldAsString(nil)
	assert.Nil(t, err)
	assert.Nil(t, v)
	s := ""
	v, err = FieldAsString(&s)
	assert.Nil(t, err)
	assert.Equal(t, "", *v)

	v, err = FieldAsString(s)
	assert.Nil(t, err)
	assert.Equal(t, "", *v)

	v, err = FieldAsString(1)
	assert.ErrorIs(t, err, InvalidTypeConversionError)
	assert.Nil(t, v)
}

func TestFieldAsBool(t *testing.T) {

	v, err := FieldAsBool(nil)
	assert.Nil(t, err)
	assert.Nil(t, v)
	s := true
	v, err = FieldAsBool(&s)
	assert.Nil(t, err)
	assert.Equal(t, true, *v)

	v, err = FieldAsBool(s)
	assert.Nil(t, err)
	assert.Equal(t, true, *v)

	v, err = FieldAsBool(1)
	assert.ErrorIs(t, err, InvalidTypeConversionError)
	assert.Nil(t, v)
}

func TestFieldAsNumber(t *testing.T) {

	v, err := FieldAsNumber(nil)
	assert.Nil(t, err)
	assert.Nil(t, v)
	s := float64(1)
	v, err = FieldAsNumber(&s)
	assert.Nil(t, err)
	assert.Equal(t, float64(1), *v)

	v, err = FieldAsNumber(s)
	assert.Nil(t, err)
	assert.Equal(t, float64(1), *v)

	v, err = FieldAsNumber(1)
	assert.ErrorIs(t, err, InvalidTypeConversionError)
	assert.Nil(t, v)
}

func TestFieldAsObject(t *testing.T) {

	v, err := FieldAsObject(nil)
	assert.Nil(t, err)
	assert.Nil(t, v)
	s := Object{}
	v, err = FieldAsObject(&s)
	assert.Nil(t, err)
	assert.Equal(t, Object{}, *v)

	v, err = FieldAsObject(s)
	assert.Nil(t, err)
	assert.Equal(t, Object{}, *v)

	v, err = FieldAsObject(1)
	assert.ErrorIs(t, err, InvalidTypeConversionError)
	assert.Nil(t, v)
}

func TestFieldAsArray(t *testing.T) {

	v, err := FieldAsArray(nil)
	assert.Nil(t, err)
	assert.Nil(t, v)
	s := Array{}
	v, err = FieldAsArray(&s)
	assert.Nil(t, err)
	assert.Equal(t, Array{}, *v)

	v, err = FieldAsArray(s)
	assert.Nil(t, err)
	assert.Equal(t, Array{}, *v)

	v, err = FieldAsArray(1)
	assert.ErrorIs(t, err, InvalidTypeConversionError)
	assert.Nil(t, v)
}

func TestFieldAsInteger(t *testing.T) {

	v, err := FieldAsInteger(nil)
	assert.Nil(t, err)
	assert.Nil(t, v)

	s := float64(12345)
	v, err = FieldAsInteger(&s)
	assert.Nil(t, err)
	assert.Equal(t, int64(12345), *v)

	v, err = FieldAsInteger(s)
	assert.Nil(t, err)
	assert.Equal(t, int64(12345), *v)

	// Ensure maximum precision
	s = float64(9007199254740992)
	v, err = FieldAsInteger(s)
	assert.Nil(t, err)
	assert.Equal(t, int64(9007199254740991), *v)

	s = float64(-9007199254740991)
	v, err = FieldAsInteger(s)
	assert.Nil(t, err)
	assert.Equal(t, int64(-9007199254740991), *v)

	// Testing a conversion to a variable with 54 bits
	s = float64(18014398509481984)
	v, err = FieldAsInteger(s)
	assert.Nil(t, err)
	assert.Equal(t, int64(9007199254740992), *v)

	v, err = FieldAsInteger(1)
	assert.ErrorIs(t, err, InvalidTypeConversionError)
	assert.Nil(t, v)

	s = float64(12345.1)
	v, err = FieldAsInteger(s)
	assert.ErrorIs(t, err, InvalidTypeConversionError)
	assert.Nil(t, v)
}

func TestUnstructuredDeserializationTypes(t *testing.T) {
	var o Object

	// This unit test is designed to verify if the JSON deserialization behavior
	// changed from the expected behavior.
	assert.Nil(t, json.Unmarshal([]byte(SAMPLE_JSON), &o))

	f := float64(1)
	assert.IsType(t, f, o["int"])
	assert.IsType(t, f, o["float"])

	s := ""
	assert.IsType(t, s, o["string"])

	b := true
	assert.IsType(t, b, o["bool"])

	assert.Nil(t, o["null"])

	a := []any{}
	assert.IsType(t, a, o["array"])

	assert.IsType(t, o, o["object"])
}
