# \LogsArchivesAPI

All URIs are relative to *https://api.na-01.cloud.solarwinds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLogArchives**](LogsArchivesAPI.md#GetLogArchives) | **Get** /logs/archives | Retrieve data on archived logs



## GetLogArchives

> LogArchiveResponse GetLogArchives(ctx).StartTime(startTime).EndTime(endTime).PageSize(pageSize).SkipToken(skipToken).Execute()

Retrieve data on archived logs



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
	startTime := "2024-01-23T00:00:00Z" // string | Time must be within 1 year from the current date. <br><br> ISO 8601 standard <br> Format: <i>yyyy-MM-ddTHH:mm:ssZ</i>
	endTime := "2024-01-23T10:30:00Z" // string | Time must be within 1 year from the current date. <br><br> ISO 8601 standard <br> Format: <i>yyyy-MM-ddTHH:mm:ssZ</i>
	pageSize := int32(56) // int32 | Maximum number of records to be fetched per request (optional) (default to 1000)
	skipToken := "skipToken_example" // string | For opaque pagination (optional)

	configuration := swoclient.NewConfiguration()
	apiClient := swoclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsArchivesAPI.GetLogArchives(context.Background()).StartTime(startTime).EndTime(endTime).PageSize(pageSize).SkipToken(skipToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesAPI.GetLogArchives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogArchives`: LogArchiveResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsArchivesAPI.GetLogArchives`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogArchivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **string** | Time must be within 1 year from the current date. &lt;br&gt;&lt;br&gt; ISO 8601 standard &lt;br&gt; Format: &lt;i&gt;yyyy-MM-ddTHH:mm:ssZ&lt;/i&gt; | 
 **endTime** | **string** | Time must be within 1 year from the current date. &lt;br&gt;&lt;br&gt; ISO 8601 standard &lt;br&gt; Format: &lt;i&gt;yyyy-MM-ddTHH:mm:ssZ&lt;/i&gt; | 
 **pageSize** | **int32** | Maximum number of records to be fetched per request | [default to 1000]
 **skipToken** | **string** | For opaque pagination | 

### Return type

[**LogArchiveResponse**](LogArchiveResponse.md)

### Authorization

[Full Access API Token](../README.md#Full Access API Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

