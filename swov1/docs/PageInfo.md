# PageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrevPage** | Pointer to **string** |  | [optional] 
**NextPage** | Pointer to **string** |  | [optional] 

## Methods

### NewPageInfo

`func NewPageInfo() *PageInfo`

NewPageInfo instantiates a new PageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageInfoWithDefaults

`func NewPageInfoWithDefaults() *PageInfo`

NewPageInfoWithDefaults instantiates a new PageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrevPage

`func (o *PageInfo) GetPrevPage() string`

GetPrevPage returns the PrevPage field if non-nil, zero value otherwise.

### GetPrevPageOk

`func (o *PageInfo) GetPrevPageOk() (*string, bool)`

GetPrevPageOk returns a tuple with the PrevPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPage

`func (o *PageInfo) SetPrevPage(v string)`

SetPrevPage sets PrevPage field to given value.

### HasPrevPage

`func (o *PageInfo) HasPrevPage() bool`

HasPrevPage returns a boolean if a field has been set.

### GetNextPage

`func (o *PageInfo) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *PageInfo) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *PageInfo) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *PageInfo) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


