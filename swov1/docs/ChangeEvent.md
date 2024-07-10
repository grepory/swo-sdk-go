# ChangeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**Title** | **string** |  | 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ParentEventId** | Pointer to **int64** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewChangeEvent

`func NewChangeEvent(name string, title string, ) *ChangeEvent`

NewChangeEvent instantiates a new ChangeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeEventWithDefaults

`func NewChangeEventWithDefaults() *ChangeEvent`

NewChangeEventWithDefaults instantiates a new ChangeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChangeEvent) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChangeEvent) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChangeEvent) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ChangeEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ChangeEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChangeEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChangeEvent) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *ChangeEvent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChangeEvent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChangeEvent) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTimestamp

`func (o *ChangeEvent) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ChangeEvent) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ChangeEvent) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ChangeEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetSource

`func (o *ChangeEvent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ChangeEvent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ChangeEvent) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ChangeEvent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDescription

`func (o *ChangeEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChangeEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChangeEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChangeEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentEventId

`func (o *ChangeEvent) GetParentEventId() int64`

GetParentEventId returns the ParentEventId field if non-nil, zero value otherwise.

### GetParentEventIdOk

`func (o *ChangeEvent) GetParentEventIdOk() (*int64, bool)`

GetParentEventIdOk returns a tuple with the ParentEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEventId

`func (o *ChangeEvent) SetParentEventId(v int64)`

SetParentEventId sets ParentEventId field to given value.

### HasParentEventId

`func (o *ChangeEvent) HasParentEventId() bool`

HasParentEventId returns a boolean if a field has been set.

### GetTags

`func (o *ChangeEvent) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ChangeEvent) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ChangeEvent) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ChangeEvent) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLinks

`func (o *ChangeEvent) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ChangeEvent) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ChangeEvent) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ChangeEvent) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


