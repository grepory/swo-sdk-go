# ListMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricsInfo** | [**[]GetMetricResponse**](GetMetricResponse.md) |  | 
**PageInfo** | Pointer to [**PageInfo**](PageInfo.md) |  | [optional] 

## Methods

### NewListMetricsResponse

`func NewListMetricsResponse(metricsInfo []GetMetricResponse, ) *ListMetricsResponse`

NewListMetricsResponse instantiates a new ListMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMetricsResponseWithDefaults

`func NewListMetricsResponseWithDefaults() *ListMetricsResponse`

NewListMetricsResponseWithDefaults instantiates a new ListMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricsInfo

`func (o *ListMetricsResponse) GetMetricsInfo() []GetMetricResponse`

GetMetricsInfo returns the MetricsInfo field if non-nil, zero value otherwise.

### GetMetricsInfoOk

`func (o *ListMetricsResponse) GetMetricsInfoOk() (*[]GetMetricResponse, bool)`

GetMetricsInfoOk returns a tuple with the MetricsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsInfo

`func (o *ListMetricsResponse) SetMetricsInfo(v []GetMetricResponse)`

SetMetricsInfo sets MetricsInfo field to given value.


### GetPageInfo

`func (o *ListMetricsResponse) GetPageInfo() PageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *ListMetricsResponse) GetPageInfoOk() (*PageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *ListMetricsResponse) SetPageInfo(v PageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *ListMetricsResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


