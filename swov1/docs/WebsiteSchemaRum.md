# WebsiteSchemaRum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApdexTimeInSeconds** | Pointer to **int32** |  | [optional] 
**Snippet** | Pointer to **string** |  | [optional] [readonly] 
**Spa** | **bool** |  | 

## Methods

### NewWebsiteSchemaRum

`func NewWebsiteSchemaRum(spa bool, ) *WebsiteSchemaRum`

NewWebsiteSchemaRum instantiates a new WebsiteSchemaRum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteSchemaRumWithDefaults

`func NewWebsiteSchemaRumWithDefaults() *WebsiteSchemaRum`

NewWebsiteSchemaRumWithDefaults instantiates a new WebsiteSchemaRum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApdexTimeInSeconds

`func (o *WebsiteSchemaRum) GetApdexTimeInSeconds() int32`

GetApdexTimeInSeconds returns the ApdexTimeInSeconds field if non-nil, zero value otherwise.

### GetApdexTimeInSecondsOk

`func (o *WebsiteSchemaRum) GetApdexTimeInSecondsOk() (*int32, bool)`

GetApdexTimeInSecondsOk returns a tuple with the ApdexTimeInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApdexTimeInSeconds

`func (o *WebsiteSchemaRum) SetApdexTimeInSeconds(v int32)`

SetApdexTimeInSeconds sets ApdexTimeInSeconds field to given value.

### HasApdexTimeInSeconds

`func (o *WebsiteSchemaRum) HasApdexTimeInSeconds() bool`

HasApdexTimeInSeconds returns a boolean if a field has been set.

### GetSnippet

`func (o *WebsiteSchemaRum) GetSnippet() string`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *WebsiteSchemaRum) GetSnippetOk() (*string, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *WebsiteSchemaRum) SetSnippet(v string)`

SetSnippet sets Snippet field to given value.

### HasSnippet

`func (o *WebsiteSchemaRum) HasSnippet() bool`

HasSnippet returns a boolean if a field has been set.

### GetSpa

`func (o *WebsiteSchemaRum) GetSpa() bool`

GetSpa returns the Spa field if non-nil, zero value otherwise.

### GetSpaOk

`func (o *WebsiteSchemaRum) GetSpaOk() (*bool, bool)`

GetSpaOk returns a tuple with the Spa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpa

`func (o *WebsiteSchemaRum) SetSpa(v bool)`

SetSpa sets Spa field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


