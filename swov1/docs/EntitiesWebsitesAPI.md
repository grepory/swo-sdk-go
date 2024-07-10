# \EntitiesWebsitesAPI

All URIs are relative to *https://api.na-01.cloud.solarwinds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntitiesWebsitesEntityIdDelete**](EntitiesWebsitesAPI.md#EntitiesWebsitesEntityIdDelete) | **Delete** /entities/websites/{entityId} | Delete website
[**EntitiesWebsitesEntityIdGet**](EntitiesWebsitesAPI.md#EntitiesWebsitesEntityIdGet) | **Get** /entities/websites/{entityId} | Get website
[**EntitiesWebsitesEntityIdPut**](EntitiesWebsitesAPI.md#EntitiesWebsitesEntityIdPut) | **Put** /entities/websites/{entityId} | Update website
[**EntitiesWebsitesPost**](EntitiesWebsitesAPI.md#EntitiesWebsitesPost) | **Post** /entities/websites | Create website



## EntitiesWebsitesEntityIdDelete

> EntityIdSchema EntitiesWebsitesEntityIdDelete(ctx, entityId).Execute()

Delete website

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
	entityId := "e-1448474379026206720" // string | 

	configuration := swoclient.NewConfiguration()
	apiClient := swoclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitiesWebsitesAPI.EntitiesWebsitesEntityIdDelete(context.Background(), entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitiesWebsitesAPI.EntitiesWebsitesEntityIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitiesWebsitesEntityIdDelete`: EntityIdSchema
	fmt.Fprintf(os.Stdout, "Response from `EntitiesWebsitesAPI.EntitiesWebsitesEntityIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitiesWebsitesEntityIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntityIdSchema**](EntityIdSchema.md)

### Authorization

[Full Access API Token](../README.md#Full Access API Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitiesWebsitesEntityIdGet

> EntitiesWebsitesEntityIdGet200Response EntitiesWebsitesEntityIdGet(ctx, entityId).Execute()

Get website

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
	entityId := "e-1448474379026206720" // string | 

	configuration := swoclient.NewConfiguration()
	apiClient := swoclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitiesWebsitesAPI.EntitiesWebsitesEntityIdGet(context.Background(), entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitiesWebsitesAPI.EntitiesWebsitesEntityIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitiesWebsitesEntityIdGet`: EntitiesWebsitesEntityIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EntitiesWebsitesAPI.EntitiesWebsitesEntityIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitiesWebsitesEntityIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntitiesWebsitesEntityIdGet200Response**](EntitiesWebsitesEntityIdGet200Response.md)

### Authorization

[Full Access API Token](../README.md#Full Access API Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitiesWebsitesEntityIdPut

> EntityIdSchema EntitiesWebsitesEntityIdPut(ctx, entityId).WebsiteSchema(websiteSchema).Execute()

Update website

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
	entityId := "e-1448474379026206720" // string | 
	websiteSchema := *swoclient.NewWebsiteSchema("solarwinds.com", "https://www.solarwinds.com") // WebsiteSchema | A JSON object containing Website configuration

	configuration := swoclient.NewConfiguration()
	apiClient := swoclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitiesWebsitesAPI.EntitiesWebsitesEntityIdPut(context.Background(), entityId).WebsiteSchema(websiteSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitiesWebsitesAPI.EntitiesWebsitesEntityIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitiesWebsitesEntityIdPut`: EntityIdSchema
	fmt.Fprintf(os.Stdout, "Response from `EntitiesWebsitesAPI.EntitiesWebsitesEntityIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitiesWebsitesEntityIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **websiteSchema** | [**WebsiteSchema**](WebsiteSchema.md) | A JSON object containing Website configuration | 

### Return type

[**EntityIdSchema**](EntityIdSchema.md)

### Authorization

[Full Access API Token](../README.md#Full Access API Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitiesWebsitesPost

> EntityIdSchema EntitiesWebsitesPost(ctx).WebsiteSchema(websiteSchema).Execute()

Create website

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
	websiteSchema := *swoclient.NewWebsiteSchema("solarwinds.com", "https://www.solarwinds.com") // WebsiteSchema | A JSON object containing Website configuration

	configuration := swoclient.NewConfiguration()
	apiClient := swoclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitiesWebsitesAPI.EntitiesWebsitesPost(context.Background()).WebsiteSchema(websiteSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitiesWebsitesAPI.EntitiesWebsitesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitiesWebsitesPost`: EntityIdSchema
	fmt.Fprintf(os.Stdout, "Response from `EntitiesWebsitesAPI.EntitiesWebsitesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntitiesWebsitesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **websiteSchema** | [**WebsiteSchema**](WebsiteSchema.md) | A JSON object containing Website configuration | 

### Return type

[**EntityIdSchema**](EntityIdSchema.md)

### Authorization

[Full Access API Token](../README.md#Full Access API Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

