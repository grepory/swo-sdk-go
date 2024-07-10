# \MetadataAPI

All URIs are relative to *https://api.na-01.cloud.solarwinds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEntityTypes**](MetadataAPI.md#ListEntityTypes) | **Get** /metadata/entities/types | List all entity types
[**ListMetricsForEntityType**](MetadataAPI.md#ListMetricsForEntityType) | **Get** /metadata/entities/types/{type}/metrics | List metrics metadata for an entity type



## ListEntityTypes

> ListEntityTypesResponse ListEntityTypes(ctx).Execute()

List all entity types



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

	configuration := swoclient.NewConfiguration()
	apiClient := swoclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataAPI.ListEntityTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.ListEntityTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEntityTypes`: ListEntityTypesResponse
	fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.ListEntityTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListEntityTypesRequest struct via the builder pattern


### Return type

[**ListEntityTypesResponse**](ListEntityTypesResponse.md)

### Authorization

[Full Access API Token](../README.md#Full Access API Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetricsForEntityType

> ListMetricsForEntityTypeResponse ListMetricsForEntityType(ctx, type_).StartTime(startTime).EndTime(endTime).Execute()

List metrics metadata for an entity type



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
	type_ := "type__example" // string | entity type
	startTime := "startTime_example" // string | yyyy-MM-ddTHH:mm:ssZ (optional)
	endTime := "endTime_example" // string | yyyy-MM-ddTHH:mm:ssZ (optional)

	configuration := swoclient.NewConfiguration()
	apiClient := swoclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataAPI.ListMetricsForEntityType(context.Background(), type_).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.ListMetricsForEntityType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMetricsForEntityType`: ListMetricsForEntityTypeResponse
	fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.ListMetricsForEntityType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | entity type | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMetricsForEntityTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **string** | yyyy-MM-ddTHH:mm:ssZ | 
 **endTime** | **string** | yyyy-MM-ddTHH:mm:ssZ | 

### Return type

[**ListMetricsForEntityTypeResponse**](ListMetricsForEntityTypeResponse.md)

### Authorization

[Full Access API Token](../README.md#Full Access API Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

