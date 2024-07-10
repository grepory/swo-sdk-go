# \EntitiesAPI

All URIs are relative to *https://api.na-01.cloud.solarwinds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEntities**](EntitiesAPI.md#GetEntities) | **Get** /entities | Get entities by type
[**GetEntityById**](EntitiesAPI.md#GetEntityById) | **Get** /entities/{id} | Get entity by id



## GetEntities

> GetEntities(ctx).Type_(type_).Name(name).PageSize(pageSize).SkipToken(skipToken).Execute()

Get entities by type

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
	type_ := "type__example" // string | The entities type
	name := "name_example" // string | The entities name (optional) (default to "")
	pageSize := int32(56) // int32 | Number of entities in a response page (optional) (default to 50)
	skipToken := "skipToken_example" // string | For opaque pagination (optional) (default to "")

	configuration := swoclient.NewConfiguration()
	apiClient := swoclient.NewAPIClient(configuration)
	r, err := apiClient.EntitiesAPI.GetEntities(context.Background()).Type_(type_).Name(name).PageSize(pageSize).SkipToken(skipToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitiesAPI.GetEntities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | The entities type | 
 **name** | **string** | The entities name | [default to &quot;&quot;]
 **pageSize** | **int32** | Number of entities in a response page | [default to 50]
 **skipToken** | **string** | For opaque pagination | [default to &quot;&quot;]

### Return type

 (empty response body)

### Authorization

[Full Access API Token](../README.md#Full Access API Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntityById

> GetEntityById(ctx, id).Execute()

Get entity by id

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
	id := "id_example" // string | The entities ID

	configuration := swoclient.NewConfiguration()
	apiClient := swoclient.NewAPIClient(configuration)
	r, err := apiClient.EntitiesAPI.GetEntityById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitiesAPI.GetEntityById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The entities ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Full Access API Token](../README.md#Full Access API Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

