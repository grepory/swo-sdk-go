# ListMetricsForEntityTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Metrics** | [**[]GetMetricResponse**](GetMetricResponse.md) |  | 

## Methods

### NewListMetricsForEntityTypeResponse

`func NewListMetricsForEntityTypeResponse(type_ string, metrics []GetMetricResponse, ) *ListMetricsForEntityTypeResponse`

NewListMetricsForEntityTypeResponse instantiates a new ListMetricsForEntityTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMetricsForEntityTypeResponseWithDefaults

`func NewListMetricsForEntityTypeResponseWithDefaults() *ListMetricsForEntityTypeResponse`

NewListMetricsForEntityTypeResponseWithDefaults instantiates a new ListMetricsForEntityTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ListMetricsForEntityTypeResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListMetricsForEntityTypeResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListMetricsForEntityTypeResponse) SetType(v string)`

SetType sets Type field to given value.


### GetMetrics

`func (o *ListMetricsForEntityTypeResponse) GetMetrics() []GetMetricResponse`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ListMetricsForEntityTypeResponse) GetMetricsOk() (*[]GetMetricResponse, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ListMetricsForEntityTypeResponse) SetMetrics(v []GetMetricResponse)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


