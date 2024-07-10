# \MetricsAPI

All URIs are relative to *https://api.na-01.cloud.solarwinds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetric**](MetricsAPI.md#GetMetric) | **Get** /metrics/{name} | Get metric info
[**ListMetricAttributeNames**](MetricsAPI.md#ListMetricAttributeNames) | **Get** /metrics/{name}/attributes | List metric attribute names
[**ListMetricAttributeValues**](MetricsAPI.md#ListMetricAttributeValues) | **Get** /metrics/{name}/attributes/{attributeName} | List metric attribute values
[**ListMetricMeasurements**](MetricsAPI.md#ListMetricMeasurements) | **Get** /metrics/{name}/measurements | List metric&#39;s measurement values for a metric
[**ListMetrics**](MetricsAPI.md#ListMetrics) | **Get** /metrics | List metrics



## GetMetric

> GetMetricResponse GetMetric(ctx, name).Execute()

Get metric info



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
	name := "name_example" // string | metric name

	configuration := swoclient.NewConfiguration()
	apiClient := swoclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GetMetric(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetric`: GetMetricResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | metric name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMetricResponse**](GetMetricResponse.md)

### Authorization

[Full Access API Token](../README.md#Full Access API Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetricAttributeNames

> ListMetricAttributeNamesResponse ListMetricAttributeNames(ctx, name).StartTime(startTime).EndTime(endTime).PageSize(pageSize).SkipToken(skipToken).Execute()

List metric attribute names



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
	name := "name_example" // string | metric name
	startTime := "startTime_example" // string | yyyy-MM-ddTHH:mm:ssZ (optional)
	endTime := "endTime_example" // string | yyyy-MM-ddTHH:mm:ssZ (optional)
	pageSize := int32(56) // int32 | Number of attribute names in a response page (optional) (default to 100)
	skipToken := "skipToken_example" // string | For opaque pagination, with default value : empty string (optional) (default to "")

	configuration := swoclient.NewConfiguration()
	apiClient := swoclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.ListMetricAttributeNames(context.Background(), name).StartTime(startTime).EndTime(endTime).PageSize(pageSize).SkipToken(skipToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.ListMetricAttributeNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMetricAttributeNames`: ListMetricAttributeNamesResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.ListMetricAttributeNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | metric name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMetricAttributeNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **string** | yyyy-MM-ddTHH:mm:ssZ | 
 **endTime** | **string** | yyyy-MM-ddTHH:mm:ssZ | 
 **pageSize** | **int32** | Number of attribute names in a response page | [default to 100]
 **skipToken** | **string** | For opaque pagination, with default value : empty string | [default to &quot;&quot;]

### Return type

[**ListMetricAttributeNamesResponse**](ListMetricAttributeNamesResponse.md)

### Authorization

[Full Access API Token](../README.md#Full Access API Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetricAttributeValues

> ListMetricAttributeValuesResponse ListMetricAttributeValues(ctx, name, attributeName).StartTime(startTime).EndTime(endTime).PageSize(pageSize).SkipToken(skipToken).Execute()

List metric attribute values



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
	name := "name_example" // string | metric name
	attributeName := "attributeName_example" // string | attribute name
	startTime := "startTime_example" // string | yyyy-MM-ddTHH:mm:ssZ (optional)
	endTime := "endTime_example" // string | yyyy-MM-ddTHH:mm:ssZ (optional)
	pageSize := int32(56) // int32 | Number of attribute values in a response page (optional) (default to 100)
	skipToken := "skipToken_example" // string | For opaque pagination, with default value : empty string (optional) (default to "")

	configuration := swoclient.NewConfiguration()
	apiClient := swoclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.ListMetricAttributeValues(context.Background(), name, attributeName).StartTime(startTime).EndTime(endTime).PageSize(pageSize).SkipToken(skipToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.ListMetricAttributeValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMetricAttributeValues`: ListMetricAttributeValuesResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.ListMetricAttributeValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | metric name | 
**attributeName** | **string** | attribute name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMetricAttributeValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTime** | **string** | yyyy-MM-ddTHH:mm:ssZ | 
 **endTime** | **string** | yyyy-MM-ddTHH:mm:ssZ | 
 **pageSize** | **int32** | Number of attribute values in a response page | [default to 100]
 **skipToken** | **string** | For opaque pagination, with default value : empty string | [default to &quot;&quot;]

### Return type

[**ListMetricAttributeValuesResponse**](ListMetricAttributeValuesResponse.md)

### Authorization

[Full Access API Token](../README.md#Full Access API Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetricMeasurements

> ListMetricMeasurementsResponse ListMetricMeasurements(ctx, name).Filter(filter).GroupBy(groupBy).AggregateBy(aggregateBy).BucketSizeInSeconds(bucketSizeInSeconds).PreGroupBy(preGroupBy).PreGroupByMethod(preGroupByMethod).StartTime(startTime).EndTime(endTime).PageSize(pageSize).SkipToken(skipToken).Execute()

List metric's measurement values for a metric



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
	name := "name_example" // string | metric name
	filter := "filter_example" // string | Query to filter the measurement values. e.g  id:[id1,id2] category:moderate (optional)
	groupBy := "groupBy_example" // string | CSV string of attribute names to group measurements by. e.g id,category (optional)
	aggregateBy := "aggregateBy_example" // string | One of COUNT, MIN, MAX, AVG, SUM, or LAST (optional) (default to "AVG")
	bucketSizeInSeconds := int32(56) // int32 | Resolution size of an aggregation bucket in seconds (optional)
	preGroupBy := "preGroupBy_example" // string | Secondary grouping to allow aggregating data points inside individual buckets.    Has to be set together with `preGroupByMethod` (optional)
	preGroupByMethod := "preGroupByMethod_example" // string | Secondary aggregation to allow aggregating data points inside individual buckets.    Has to be set together with `preGroupBy` (optional)
	startTime := "startTime_example" // string | yyyy-MM-ddTHH:mm:ssZ (optional)
	endTime := "endTime_example" // string | yyyy-MM-ddTHH:mm:ssZ (optional)
	pageSize := int32(56) // int32 | Number of groupings in a response page (optional) (default to 100)
	skipToken := "skipToken_example" // string | For opaque pagination, with default value : empty string (optional) (default to "")

	configuration := swoclient.NewConfiguration()
	apiClient := swoclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.ListMetricMeasurements(context.Background(), name).Filter(filter).GroupBy(groupBy).AggregateBy(aggregateBy).BucketSizeInSeconds(bucketSizeInSeconds).PreGroupBy(preGroupBy).PreGroupByMethod(preGroupByMethod).StartTime(startTime).EndTime(endTime).PageSize(pageSize).SkipToken(skipToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.ListMetricMeasurements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMetricMeasurements`: ListMetricMeasurementsResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.ListMetricMeasurements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | metric name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMetricMeasurementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Query to filter the measurement values. e.g  id:[id1,id2] category:moderate | 
 **groupBy** | **string** | CSV string of attribute names to group measurements by. e.g id,category | 
 **aggregateBy** | **string** | One of COUNT, MIN, MAX, AVG, SUM, or LAST | [default to &quot;AVG&quot;]
 **bucketSizeInSeconds** | **int32** | Resolution size of an aggregation bucket in seconds | 
 **preGroupBy** | **string** | Secondary grouping to allow aggregating data points inside individual buckets.    Has to be set together with &#x60;preGroupByMethod&#x60; | 
 **preGroupByMethod** | **string** | Secondary aggregation to allow aggregating data points inside individual buckets.    Has to be set together with &#x60;preGroupBy&#x60; | 
 **startTime** | **string** | yyyy-MM-ddTHH:mm:ssZ | 
 **endTime** | **string** | yyyy-MM-ddTHH:mm:ssZ | 
 **pageSize** | **int32** | Number of groupings in a response page | [default to 100]
 **skipToken** | **string** | For opaque pagination, with default value : empty string | [default to &quot;&quot;]

### Return type

[**ListMetricMeasurementsResponse**](ListMetricMeasurementsResponse.md)

### Authorization

[Full Access API Token](../README.md#Full Access API Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetrics

> ListMetricsResponse ListMetrics(ctx).Name(name).StartTime(startTime).EndTime(endTime).PageSize(pageSize).SkipToken(skipToken).Execute()

List metrics



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
	name := "name_example" // string | metric name (optional)
	startTime := "startTime_example" // string | yyyy-MM-ddTHH:mm:ssZ (optional)
	endTime := "endTime_example" // string | yyyy-MM-ddTHH:mm:ssZ (optional)
	pageSize := int32(56) // int32 | Number of metrics in a response page (optional) (default to 100)
	skipToken := "skipToken_example" // string | For opaque pagination, with default value : empty string (optional) (default to "")

	configuration := swoclient.NewConfiguration()
	apiClient := swoclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.ListMetrics(context.Background()).Name(name).StartTime(startTime).EndTime(endTime).PageSize(pageSize).SkipToken(skipToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.ListMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMetrics`: ListMetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.ListMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | metric name | 
 **startTime** | **string** | yyyy-MM-ddTHH:mm:ssZ | 
 **endTime** | **string** | yyyy-MM-ddTHH:mm:ssZ | 
 **pageSize** | **int32** | Number of metrics in a response page | [default to 100]
 **skipToken** | **string** | For opaque pagination, with default value : empty string | [default to &quot;&quot;]

### Return type

[**ListMetricsResponse**](ListMetricsResponse.md)

### Authorization

[Full Access API Token](../README.md#Full Access API Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

