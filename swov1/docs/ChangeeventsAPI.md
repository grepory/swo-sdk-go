# \ChangeeventsAPI

All URIs are relative to *https://api.na-01.cloud.solarwinds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEvent**](ChangeeventsAPI.md#CreateEvent) | **Post** /changeevents | Create an event



## CreateEvent

> CreateEvent(ctx).ChangeEvent(changeEvent).Execute()

Create an event

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
	changeEvent := *swoclient.NewChangeEvent("Name_example", "Title_example") // ChangeEvent | A JSON object containing a change event

	configuration := swoclient.NewConfiguration()
	apiClient := swoclient.NewAPIClient(configuration)
	r, err := apiClient.ChangeeventsAPI.CreateEvent(context.Background()).ChangeEvent(changeEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeeventsAPI.CreateEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changeEvent** | [**ChangeEvent**](ChangeEvent.md) | A JSON object containing a change event | 

### Return type

 (empty response body)

### Authorization

[Full Access API Token](../README.md#Full Access API Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

