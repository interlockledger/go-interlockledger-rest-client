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
	"encoding/json"
	"errors"
	"math"
)

var InvalidTypeConversionError = errors.New("Invalid type conversion.")

type ObjectFieldType int

const (
	ObjectFieldMissing ObjectFieldType = iota
	ObjectFieldNil
	ObjectFieldString
	ObjectFieldNumber
	ObjectFieldBool
	ObjectFieldObject
	ObjectFieldArray
	ObjectFieldUnknown
)

// This type can be used to hold a non structured JSON object. It is an alias to
// map[string]any.
type Object = map[string]any

// This type can be used to hold a non structured JSON array. It is an alias to
// []any.
type Array = []any

/*
Maps the given value into the proper ObjectFieldType. This function may be used
to simplify the implementation of verifiers for unstructured JSON objects.
*/
func MapObjectType(v interface{}) ObjectFieldType {
	if v == nil {
		return ObjectFieldNil
	}
	switch v.(type) {
	case string:
		return ObjectFieldString
	case float64:
		return ObjectFieldNumber
	case bool:
		return ObjectFieldBool
	case map[string]any:
		return ObjectFieldObject
	case []any:
		return ObjectFieldArray
	default:
		return ObjectFieldUnknown
	}
}

/*
Converts the specified field into a string. It fails if v is not a string or a
pointer to it.
*/
func FieldAsString(v any) (*string, error) {
	if v == nil {
		return nil, nil
	}
	switch v.(type) {
	case *string:
		return v.(*string), nil
	case string:
		s := v.(string)
		return &s, nil
	default:
		return nil, InvalidTypeConversionError
	}
}

/*
Converts the specified field into a bool. It fails if v is not a bool or a
pointer to it.
*/
func FieldAsBool(v any) (*bool, error) {
	if v == nil {
		return nil, nil
	}
	switch v.(type) {
	case *bool:
		return v.(*bool), nil
	case bool:
		s := v.(bool)
		return &s, nil
	default:
		return nil, InvalidTypeConversionError
	}
}

/*
Converts the specified field into a number. It fails if v is not a float64 or a
pointer to it.
*/
func FieldAsNumber(v any) (*float64, error) {
	if v == nil {
		return nil, nil
	}
	switch v.(type) {
	case *float64:
		return v.(*float64), nil
	case float64:
		s := v.(float64)
		return &s, nil
	default:
		return nil, InvalidTypeConversionError
	}
}

/*
Converts the specified field into an integer. It fails if v is not a float64 or
a pointer to it or if the value of v has a non zero fractional part.

It is important to notice that, according to
https://pkg.go.dev/encoding/json#Unmarshal, numbers are always encoded to
float64 regardless of being integers or floating points.

This means that integers with more than 53 bits will likely lost their least
significant bits unless those bits are all zeroes. Only values between
-9007199254740992 and 9007199254740992 are guaranteed to be properly converted.
Any value outside of this range is subjected to loose their least significant
bits.
*/
func FieldAsInteger(v any) (*int64, error) {
	f, err := FieldAsNumber(v)
	if err != nil {
		return nil, err
	}
	if f == nil {
		return nil, nil
	}
	i := math.Floor(*f)
	if i != *f {
		return nil, InvalidTypeConversionError
	}
	ret := int64(i)
	return &ret, nil
}

/*
Converts the specified field into an Object. It fails if v is not an Object or a
pointer to it.
*/
func FieldAsObject(v any) (*Object, error) {
	if v == nil {
		return nil, nil
	}
	switch v.(type) {
	case *Object:
		return v.(*Object), nil
	case Object:
		s := v.(Object)
		return &s, nil
	default:
		return nil, InvalidTypeConversionError
	}
}

/*
Converts the specified field into an Array. It fails if v is not an Array or a
pointer to it.
*/
func FieldAsArray(v any) (*Array, error) {
	if v == nil {
		return nil, nil
	}
	switch v.(type) {
	case *Array:
		return v.(*Array), nil
	case Array:
		s := v.(Array)
		return &s, nil
	default:
		return nil, InvalidTypeConversionError
	}
}

// Converts the src into dst using a JSON as an intermediate. It will work as
// long as src can be safely converted into a JSON using the standard package
// `encoding/json`.
func ConvertUsingJSON(src, dst any) error {
	// This implementation just converts the struct to a JSON and then back into
	// a map. In the future, we could optimize this method do make the conversion
	// directly. Unfortunately, most of the code of this package is
	bin, err := json.Marshal(src)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bin, &dst); err != nil {
		return err
	}
	return nil
}

// This helper function converts any structure into an Object. It will work as
// long as src can be safely converted into a JSON using the standard package
// `encoding/json`.
func ToObject(src any) (*Object, error) {
	var dst Object
	if err := ConvertUsingJSON(src, &dst); err != nil {
		return nil, err
	}
	return &dst, nil
}

// Converts this object into the specified struct. It will work as
// long as dst can be safely converted from JSON using the standard package
// `encoding/json`.
func FromObject(src Object, dst any) error {
	return ConvertUsingJSON(src, dst)
}
