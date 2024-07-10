/*
SolarWinds Observability REST API

[Rest API Documentation](https://documentation.solarwinds.com/en/success_center/observability/content/api/api-swagger.htm)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package swov1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// LogsAPIService LogsAPI service
type LogsAPIService service

type ApiLogsRequest struct {
	ctx context.Context
	ApiService *LogsAPIService
	filter *string
	group *string
	startTime *string
	endTime *string
	direction *string
	pageSize *int32
	skipToken *string
}

// A search query string.
func (r ApiLogsRequest) Filter(filter string) ApiLogsRequest {
	r.filter = &filter
	return r
}

// Search a specific group
func (r ApiLogsRequest) Group(group string) ApiLogsRequest {
	r.group = &group
	return r
}

// yyyy-MM-ddTHH:mm:ssZ
func (r ApiLogsRequest) StartTime(startTime string) ApiLogsRequest {
	r.startTime = &startTime
	return r
}

// yyyy-MM-ddTHH:mm:ssZ
func (r ApiLogsRequest) EndTime(endTime string) ApiLogsRequest {
	r.endTime = &endTime
	return r
}

// Search direction: backward, forward, or tail
func (r ApiLogsRequest) Direction(direction string) ApiLogsRequest {
	r.direction = &direction
	return r
}

// Number of logs in a response page
func (r ApiLogsRequest) PageSize(pageSize int32) ApiLogsRequest {
	r.pageSize = &pageSize
	return r
}

// For opaque pagination, with default value : empty string
func (r ApiLogsRequest) SkipToken(skipToken string) ApiLogsRequest {
	r.skipToken = &skipToken
	return r
}

func (r ApiLogsRequest) Execute() (*LogsResponse, *http.Response, error) {
	return r.ApiService.LogsExecute(r)
}

/*
Logs Search logs

Search logs within a time period

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLogsRequest
*/
func (a *LogsAPIService) Logs(ctx context.Context) ApiLogsRequest {
	return ApiLogsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LogsResponse
func (a *LogsAPIService) LogsExecute(r ApiLogsRequest) (*LogsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LogsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsAPIService.Logs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
	if r.group != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group", r.group, "")
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "")
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "")
	}
	if r.direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	} else {
		var defaultValue int32 = 10000
		r.pageSize = &defaultValue
	}
	if r.skipToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skipToken", r.skipToken, "")
	} else {
		var defaultValue string = ""
		r.skipToken = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}