# \LogsAPI

All URIs are relative to *https://api.na-01.cloud.solarwinds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Logs**](LogsAPI.md#Logs) | **Get** /logs | Search logs



## Logs

> LogsResponse Logs(ctx).Filter(filter).Group(group).StartTime(startTime).EndTime(endTime).Direction(direction).PageSize(pageSize).SkipToken(skipToken).Execute()

Search logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	swoclient "github.com/solarwinds/swo-sdk-go/swov1"
)

func main() {
	filter := "filter_example" // string | A search query string. (optional)
	group := "group_example" // string | Search a specific group (optional)
	startTime := "startTime_example" // string | yyyy-MM-ddTHH:mm:ssZ (optional)
	endTime := "endTime_example" // string | yyyy-MM-ddTHH:mm:ssZ (optional)
	direction := "direction_example" // string | Search direction: backward, forward, or tail (optional)
	pageSize := int32(56) // int32 | Number of logs in a response page (optional) (default to 10000)
	skipToken := "skipToken_example" // string | For opaque pagination, with default value : empty string (optional) (default to "")

	configuration := swoclient.NewConfiguration()
	apiClient := swoclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.Logs(context.Background()).Filter(filter).Group(group).StartTime(startTime).EndTime(endTime).Direction(direction).PageSize(pageSize).SkipToken(skipToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.Logs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Logs`: LogsResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.Logs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A search query string. | 
 **group** | **string** | Search a specific group | 
 **startTime** | **string** | yyyy-MM-ddTHH:mm:ssZ | 
 **endTime** | **string** | yyyy-MM-ddTHH:mm:ssZ | 
 **direction** | **string** | Search direction: backward, forward, or tail | 
 **pageSize** | **int32** | Number of logs in a response page | [default to 10000]
 **skipToken** | **string** | For opaque pagination, with default value : empty string | [default to &quot;&quot;]

### Return type

[**LogsResponse**](LogsResponse.md)

### Authorization

[Full Access API Token](../README.md#Full Access API Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

