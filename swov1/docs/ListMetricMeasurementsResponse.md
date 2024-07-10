# ListMetricMeasurementsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groupings** | [**[]MeasurementGrouping**](MeasurementGrouping.md) |  | 
**BucketSizeInSeconds** | **int32** |  | 
**PageInfo** | Pointer to [**PageInfo**](PageInfo.md) |  | [optional] 

## Methods

### NewListMetricMeasurementsResponse

`func NewListMetricMeasurementsResponse(groupings []MeasurementGrouping, bucketSizeInSeconds int32, ) *ListMetricMeasurementsResponse`

NewListMetricMeasurementsResponse instantiates a new ListMetricMeasurementsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMetricMeasurementsResponseWithDefaults

`func NewListMetricMeasurementsResponseWithDefaults() *ListMetricMeasurementsResponse`

NewListMetricMeasurementsResponseWithDefaults instantiates a new ListMetricMeasurementsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupings

`func (o *ListMetricMeasurementsResponse) GetGroupings() []MeasurementGrouping`

GetGroupings returns the Groupings field if non-nil, zero value otherwise.

### GetGroupingsOk

`func (o *ListMetricMeasurementsResponse) GetGroupingsOk() (*[]MeasurementGrouping, bool)`

GetGroupingsOk returns a tuple with the Groupings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupings

`func (o *ListMetricMeasurementsResponse) SetGroupings(v []MeasurementGrouping)`

SetGroupings sets Groupings field to given value.


### GetBucketSizeInSeconds

`func (o *ListMetricMeasurementsResponse) GetBucketSizeInSeconds() int32`

GetBucketSizeInSeconds returns the BucketSizeInSeconds field if non-nil, zero value otherwise.

### GetBucketSizeInSecondsOk

`func (o *ListMetricMeasurementsResponse) GetBucketSizeInSecondsOk() (*int32, bool)`

GetBucketSizeInSecondsOk returns a tuple with the BucketSizeInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketSizeInSeconds

`func (o *ListMetricMeasurementsResponse) SetBucketSizeInSeconds(v int32)`

SetBucketSizeInSeconds sets BucketSizeInSeconds field to given value.


### GetPageInfo

`func (o *ListMetricMeasurementsResponse) GetPageInfo() PageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *ListMetricMeasurementsResponse) GetPageInfoOk() (*PageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *ListMetricMeasurementsResponse) SetPageInfo(v PageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *ListMetricMeasurementsResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


