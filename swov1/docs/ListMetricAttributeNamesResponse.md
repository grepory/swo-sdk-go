# ListMetricAttributeNamesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | **[]string** |  | 
**PageInfo** | Pointer to [**PageInfo**](PageInfo.md) |  | [optional] 

## Methods

### NewListMetricAttributeNamesResponse

`func NewListMetricAttributeNamesResponse(names []string, ) *ListMetricAttributeNamesResponse`

NewListMetricAttributeNamesResponse instantiates a new ListMetricAttributeNamesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMetricAttributeNamesResponseWithDefaults

`func NewListMetricAttributeNamesResponseWithDefaults() *ListMetricAttributeNamesResponse`

NewListMetricAttributeNamesResponseWithDefaults instantiates a new ListMetricAttributeNamesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *ListMetricAttributeNamesResponse) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *ListMetricAttributeNamesResponse) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *ListMetricAttributeNamesResponse) SetNames(v []string)`

SetNames sets Names field to given value.


### GetPageInfo

`func (o *ListMetricAttributeNamesResponse) GetPageInfo() PageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *ListMetricAttributeNamesResponse) GetPageInfoOk() (*PageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *ListMetricAttributeNamesResponse) SetPageInfo(v PageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *ListMetricAttributeNamesResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


