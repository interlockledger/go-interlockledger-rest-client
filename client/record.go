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
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
	. "github.com/interlockledger/go-interlockledger-rest-client/client/models"
)

// Linger please
var (
	_ context.Context
)

// Base RecordApi record options.
type RecordApiPagingOpts struct {
	Page        optional.Int32
	PageSize    optional.Int32
	LastToFirst optional.Bool
}

// Record API service.
type RecordApiService service

/*
Calls POST /records@{chain}.
*/
func (a *RecordApiService) RecordAdd(ctx context.Context, chain string, localVarPostBody *NewRecordModel) (RecordModel, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")

		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue RecordModel
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/records@{chain}"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", fmt.Sprintf("%v", chain), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		switch localVarHttpResponse.StatusCode {
		case 200:
			var v RecordModel
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		case 400, 401, 403, 404, 422:
			var v map[string]Object
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		default:
			return localVarReturnValue, localVarHttpResponse, newErr
		}
	}
	return localVarReturnValue, localVarHttpResponse, nil
}

/*
Options for POST /records@{chain}/asJson.
*/
type RecordApiRecordAddAsJsonOpts struct {
	ApplicationId optional.Int64
	PayloadTagId  optional.Int64
	Type_         optional.String
}

/*
Calls POST /records@{chain}/asJson.
*/
func (a *RecordApiService) RecordAddAsJson(ctx context.Context, chain string, options *RecordApiRecordAddAsJsonOpts, jsonPayload interface{}) (RecordModel, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue RecordModel
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/records@{chain}/asJson"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", fmt.Sprintf("%v", chain), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if options != nil && options.ApplicationId.IsSet() {
		localVarQueryParams.Add("applicationId", parameterToString(options.ApplicationId.Value(), ""))
	}
	if options != nil && options.PayloadTagId.IsSet() {
		localVarQueryParams.Add("payloadTagId", parameterToString(options.PayloadTagId.Value(), ""))
	}
	if options != nil && options.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(options.Type_.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, jsonPayload, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		switch localVarHttpResponse.StatusCode {
		case 200:
			var v RecordModel
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		case 400, 401, 403, 404, 422:
			var v map[string]Object
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		default:
			return localVarReturnValue, localVarHttpResponse, newErr
		}
	}
	return localVarReturnValue, localVarHttpResponse, nil
}

/*
Calls GET /records@{chain}/{serial}.
*/
func (a *RecordApiService) RecordGet(ctx context.Context, chain string, serial int64) (RecordModel, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue RecordModel
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/records@{chain}/{serial}"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", fmt.Sprintf("%v", chain), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", fmt.Sprintf("%v", serial), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		switch localVarHttpResponse.StatusCode {
		case 200:
			var v RecordModel
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		case 400, 401, 403, 404, 422:
			var v map[string]Object
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		default:
			return localVarReturnValue, localVarHttpResponse, newErr
		}
	}
	return localVarReturnValue, localVarHttpResponse, nil
}

/*
Calls GET /records@{chain}/asJson/{serial}.
*/
func (a *RecordApiService) RecordGetAsJson(ctx context.Context, chain string, serial int64) (RecordModelAsJson, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue RecordModelAsJson
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/records@{chain}/asJson/{serial}"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", fmt.Sprintf("%v", chain), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", fmt.Sprintf("%v", serial), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		switch localVarHttpResponse.StatusCode {
		case 200:
			var v RecordModelAsJson
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		case 400, 401, 403, 404, 422:
			var v map[string]Object
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		default:
			return localVarReturnValue, localVarHttpResponse, newErr
		}
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
Options for GET /records@{chain}.
*/
type RecordApiRecordsListOpts struct {
	RecordApiPagingOpts
	FirstSerial optional.Int64
	LastSerial  optional.Int64
}

/*
Calls GET /records@{chain}.
*/
func (a *RecordApiService) RecordsList(ctx context.Context, chain string, options *RecordApiRecordsListOpts) (RecordModelPageOf, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue RecordModelPageOf
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/records@{chain}"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", fmt.Sprintf("%v", chain), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if options != nil && options.FirstSerial.IsSet() {
		localVarQueryParams.Add("firstSerial", parameterToString(options.FirstSerial.Value(), ""))
	}
	if options != nil && options.LastSerial.IsSet() {
		localVarQueryParams.Add("lastSerial", parameterToString(options.LastSerial.Value(), ""))
	}
	if options != nil && options.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(options.Page.Value(), ""))
	}
	if options != nil && options.PageSize.IsSet() {
		localVarQueryParams.Add("pageSize", parameterToString(options.PageSize.Value(), ""))
	}
	if options != nil && options.LastToFirst.IsSet() {
		localVarQueryParams.Add("lastToFirst", parameterToString(options.LastToFirst.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		switch localVarHttpResponse.StatusCode {
		case 200:
			var v RecordModelPageOf
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		case 400, 401, 403, 404, 422:
			var v map[string]Object
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		default:
			return localVarReturnValue, localVarHttpResponse, newErr
		}
	}
	return localVarReturnValue, localVarHttpResponse, nil
}

/*
Options for GET /records@{chain}/asJson.
*/
type RecordApiRecordsListAsJsonOpts struct {
	RecordApiPagingOpts
	FirstSerial optional.Int64
	LastSerial  optional.Int64
}

/*
Calls GET /records@{chain}/asJson.
*/
func (a *RecordApiService) RecordsListAsJson(ctx context.Context, chain string, options *RecordApiRecordsListAsJsonOpts) (RecordModelAsJsonPageOf, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue RecordModelAsJsonPageOf
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/records@{chain}/asJson"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", fmt.Sprintf("%v", chain), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if options != nil && options.FirstSerial.IsSet() {
		localVarQueryParams.Add("firstSerial", parameterToString(options.FirstSerial.Value(), ""))
	}
	if options != nil && options.LastSerial.IsSet() {
		localVarQueryParams.Add("lastSerial", parameterToString(options.LastSerial.Value(), ""))
	}
	if options != nil && options.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(options.Page.Value(), ""))
	}
	if options != nil && options.PageSize.IsSet() {
		localVarQueryParams.Add("pageSize", parameterToString(options.PageSize.Value(), ""))
	}
	if options != nil && options.LastToFirst.IsSet() {
		localVarQueryParams.Add("lastToFirst", parameterToString(options.LastToFirst.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		switch localVarHttpResponse.StatusCode {
		case 200:
			var v RecordModelAsJsonPageOf
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		case 400, 401, 403, 404, 422:
			var v map[string]Object
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		default:
			return localVarReturnValue, localVarHttpResponse, newErr
		}
	}
	return localVarReturnValue, localVarHttpResponse, nil
}

/*
Options for GET /records@{chain}/query.
*/
type RecordApiRecordsQueryOpts struct {
	RecordApiPagingOpts
	QueryAsInterlockQL optional.String
	HowMany            optional.Int64
}

/*
Calls GET /records@{chain}/query.
*/
func (a *RecordApiService) RecordsQuery(ctx context.Context, chain string, options *RecordApiRecordsQueryOpts) (RecordModelPageOf, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue RecordModelPageOf
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/records@{chain}/query"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", fmt.Sprintf("%v", chain), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if options != nil && options.QueryAsInterlockQL.IsSet() {
		localVarQueryParams.Add("queryAsInterlockQL", parameterToString(options.QueryAsInterlockQL.Value(), ""))
	}
	if options != nil && options.HowMany.IsSet() {
		localVarQueryParams.Add("howMany", parameterToString(options.HowMany.Value(), ""))
	}
	if options != nil && options.LastToFirst.IsSet() {
		localVarQueryParams.Add("lastToFirst", parameterToString(options.LastToFirst.Value(), ""))
	}
	if options != nil && options.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(options.Page.Value(), ""))
	}
	if options != nil && options.PageSize.IsSet() {
		localVarQueryParams.Add("pageSize", parameterToString(options.PageSize.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		switch localVarHttpResponse.StatusCode {
		case 200:
			var v RecordModelPageOf
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		case 400, 401, 403, 404, 422:
			var v map[string]Object
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		default:
			return localVarReturnValue, localVarHttpResponse, newErr
		}
	}
	return localVarReturnValue, localVarHttpResponse, nil
}

/*
Options for Calls GET /records@{chain}/asJson/query.
*/
type RecordApiRecordsQueryAsJsonOpts struct {
	RecordApiPagingOpts
	QueryAsInterlockQL optional.String
	HowMany            optional.Int64
}

/*
Calls GET /records@{chain}/asJson/query.
*/
func (a *RecordApiService) RecordsQueryAsJson(ctx context.Context, chain string, options *RecordApiRecordsQueryAsJsonOpts) (RecordModelAsJsonPageOf, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue RecordModelAsJsonPageOf
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/records@{chain}/asJson/query"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", fmt.Sprintf("%v", chain), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if options != nil && options.QueryAsInterlockQL.IsSet() {
		localVarQueryParams.Add("queryAsInterlockQL", parameterToString(options.QueryAsInterlockQL.Value(), ""))
	}
	if options != nil && options.HowMany.IsSet() {
		localVarQueryParams.Add("howMany", parameterToString(options.HowMany.Value(), ""))
	}
	if options != nil && options.LastToFirst.IsSet() {
		localVarQueryParams.Add("lastToFirst", parameterToString(options.LastToFirst.Value(), ""))
	}
	if options != nil && options.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(options.Page.Value(), ""))
	}
	if options != nil && options.PageSize.IsSet() {
		localVarQueryParams.Add("pageSize", parameterToString(options.PageSize.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		switch localVarHttpResponse.StatusCode {
		case 200:
			var v RecordModelAsJsonPageOf
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		case 400, 401, 403, 404, 422:
			var v map[string]Object
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		default:
			return localVarReturnValue, localVarHttpResponse, newErr
		}
	}
	return localVarReturnValue, localVarHttpResponse, nil
}
