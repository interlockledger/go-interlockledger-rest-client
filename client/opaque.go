// BSD 3-Clause License
//
// Copyright (c) 2022-2023, InterlockLedger
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
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/interlockledger/go-interlockledger-rest-client/client/models"
)

// Node API service.
type OpaqueService service

/*
Calls POST /opaque/{chain}.
*/
func (a *OpaqueService) Create(ctx context.Context,
	chain string, appId int64, payloadType int64, payload io.Reader, lastChangedRecordSerial int64) (models.OpaqueRecordModel, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.OpaqueRecordModel
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/opaque/" + chain

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// Set the content type to application/octet-stream
	localVarHeaderParams["Content-Type"] = "application/octet-stream"

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	// Query parameters
	localVarQueryParams.Add("appId", strconv.FormatInt(appId, 10))
	localVarQueryParams.Add("payloadTypeId", strconv.FormatInt(payloadType, 10))
	if lastChangedRecordSerial > 0 {
		localVarQueryParams.Add("lastChangedRecordSerial", strconv.FormatInt(lastChangedRecordSerial, 10))
	}

	// body params
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, payload, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
		case 201:
			var v models.ChainCreatedModel
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		case 400, 401, 403, 404, 422:
			var v map[string]models.Object
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
Calls GET /opaque/{chain}@{serial}. It returns the current payload, the appId
(reserved for future uses), the payloadTypeId and the actual response.
*/
func (a *OpaqueService) Get(ctx context.Context,
	chain string, serial int64) ([]byte, int64, int64, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/opaque/" + chain + "@" +
		strconv.FormatInt(serial, 10)

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
	localVarHttpHeaderAccepts := []string{"application/octet-stream", "application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	// body params
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, 0, 0, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return nil, 0, 0, localVarHttpResponse, err
	}

	// Get the payloadTypeId
	var typeId int64
	typeIdStr := localVarHttpResponse.Header.Get("x-payload-type-id")
	if typeIdStr != "" {
		typeId, err = strconv.ParseInt(typeIdStr, 10, 64)
		if err != nil {
			return nil, 0, 0, localVarHttpResponse, err
		}
	}
	// TODO
	var appId int64

	// Read the body
	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return nil, 0, 0, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		return localVarBody, appId, typeId, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		switch localVarHttpResponse.StatusCode {
		case 201:
			var v models.ChainCreatedModel
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return nil, 0, 0, localVarHttpResponse, newErr
			}
			newErr.model = v
			return nil, 0, 0, localVarHttpResponse, newErr
		case 400, 401, 403, 404, 422:
			var v map[string]models.Object
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return nil, 0, 0, localVarHttpResponse, newErr
			}
			newErr.model = v
			return nil, 0, 0, localVarHttpResponse, newErr
		default:
			return nil, 0, 0, localVarHttpResponse, newErr
		}
	}
	return nil, 0, 0, localVarHttpResponse, nil
}

/*
Calls GET /opaque/{chain}@{serial}. It returns the current payload, the lastChangedRecordSerial
(reserved for future uses) and the actual response.
*/
func (a *OpaqueService) Query(ctx context.Context,
	chain string, appId int64, payloadTypeIds []int64, howMany int64, lastToFirst bool, page int, pageSize int) ([]byte, int64, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/opaque/" + chain + "/query"

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
	localVarHttpHeaderAccepts := []string{"application/octet-stream", "application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	localVarQueryParams.Add("appId", strconv.FormatInt(appId, 10))
	for _, payloadTypeId := range payloadTypeIds {
		localVarQueryParams.Add("payloadTypeIds", strconv.FormatInt(payloadTypeId, 10))

	}
	localVarQueryParams.Add("howMany", strconv.FormatInt(howMany, 10))
	localVarQueryParams.Add("lastToFirst", strconv.FormatBool(lastToFirst))
	if page > 0 {
		localVarQueryParams.Add("page", strconv.FormatInt(int64(page), 10))
	}
	if pageSize > 0 {
		localVarQueryParams.Add("pageSize", strconv.FormatInt((int64(pageSize)), 10))
	}

	// body params
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, 0, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return nil, 0, localVarHttpResponse, err
	}

	var lastChangedRecordSerial int64

	// Read the body
	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return nil, lastChangedRecordSerial, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		return localVarBody, lastChangedRecordSerial, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		switch localVarHttpResponse.StatusCode {
		case 201:
			var v models.ChainCreatedModel
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return nil, 0, localVarHttpResponse, newErr
			}
			newErr.model = v
			return nil, 0, localVarHttpResponse, newErr
		case 400, 401, 403, 404, 422:
			var v map[string]models.Object
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return nil, 0, localVarHttpResponse, newErr
			}
			newErr.model = v
			return nil, 0, localVarHttpResponse, newErr
		default:
			return nil, 0, localVarHttpResponse, newErr
		}
	}
	return nil, lastChangedRecordSerial, localVarHttpResponse, nil
}
