# ListMetricAttributeValuesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Values** | **[]string** |  | 
**PageInfo** | Pointer to [**PageInfo**](PageInfo.md) |  | [optional] 

## Methods

### NewListMetricAttributeValuesResponse

`func NewListMetricAttributeValuesResponse(name string, values []string, ) *ListMetricAttributeValuesResponse`

NewListMetricAttributeValuesResponse instantiates a new ListMetricAttributeValuesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMetricAttributeValuesResponseWithDefaults

`func NewListMetricAttributeValuesResponseWithDefaults() *ListMetricAttributeValuesResponse`

NewListMetricAttributeValuesResponseWithDefaults instantiates a new ListMetricAttributeValuesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListMetricAttributeValuesResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListMetricAttributeValuesResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListMetricAttributeValuesResponse) SetName(v string)`

SetName sets Name field to given value.


### GetValues

`func (o *ListMetricAttributeValuesResponse) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ListMetricAttributeValuesResponse) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ListMetricAttributeValuesResponse) SetValues(v []string)`

SetValues sets Values field to given value.


### GetPageInfo

`func (o *ListMetricAttributeValuesResponse) GetPageInfo() PageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *ListMetricAttributeValuesResponse) GetPageInfoOk() (*PageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *ListMetricAttributeValuesResponse) SetPageInfo(v PageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *ListMetricAttributeValuesResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


