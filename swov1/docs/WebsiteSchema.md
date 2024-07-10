# WebsiteSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the website, which must be unique within the organization. The website must also not contain any control characters, any white space other than space (U+0020), or any consecutive, leading or trailing spaces.  | 
**Url** | **string** | URL of the website. Must be a valid URL with no leading or trailing white space. Must not contain invalid port number (&gt;65535).  | 
**AvailabilityCheckSettings** | Pointer to [**NullableWebsiteSchemaAvailabilityCheckSettings**](WebsiteSchemaAvailabilityCheckSettings.md) |  | [optional] 
**Rum** | Pointer to [**NullableWebsiteSchemaRum**](WebsiteSchemaRum.md) |  | [optional] 

## Methods

### NewWebsiteSchema

`func NewWebsiteSchema(name string, url string, ) *WebsiteSchema`

NewWebsiteSchema instantiates a new WebsiteSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteSchemaWithDefaults

`func NewWebsiteSchemaWithDefaults() *WebsiteSchema`

NewWebsiteSchemaWithDefaults instantiates a new WebsiteSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebsiteSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebsiteSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebsiteSchema) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *WebsiteSchema) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebsiteSchema) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebsiteSchema) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAvailabilityCheckSettings

`func (o *WebsiteSchema) GetAvailabilityCheckSettings() WebsiteSchemaAvailabilityCheckSettings`

GetAvailabilityCheckSettings returns the AvailabilityCheckSettings field if non-nil, zero value otherwise.

### GetAvailabilityCheckSettingsOk

`func (o *WebsiteSchema) GetAvailabilityCheckSettingsOk() (*WebsiteSchemaAvailabilityCheckSettings, bool)`

GetAvailabilityCheckSettingsOk returns a tuple with the AvailabilityCheckSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityCheckSettings

`func (o *WebsiteSchema) SetAvailabilityCheckSettings(v WebsiteSchemaAvailabilityCheckSettings)`

SetAvailabilityCheckSettings sets AvailabilityCheckSettings field to given value.

### HasAvailabilityCheckSettings

`func (o *WebsiteSchema) HasAvailabilityCheckSettings() bool`

HasAvailabilityCheckSettings returns a boolean if a field has been set.

### SetAvailabilityCheckSettingsNil

`func (o *WebsiteSchema) SetAvailabilityCheckSettingsNil(b bool)`

 SetAvailabilityCheckSettingsNil sets the value for AvailabilityCheckSettings to be an explicit nil

### UnsetAvailabilityCheckSettings
`func (o *WebsiteSchema) UnsetAvailabilityCheckSettings()`

UnsetAvailabilityCheckSettings ensures that no value is present for AvailabilityCheckSettings, not even an explicit nil
### GetRum

`func (o *WebsiteSchema) GetRum() WebsiteSchemaRum`

GetRum returns the Rum field if non-nil, zero value otherwise.

### GetRumOk

`func (o *WebsiteSchema) GetRumOk() (*WebsiteSchemaRum, bool)`

GetRumOk returns a tuple with the Rum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRum

`func (o *WebsiteSchema) SetRum(v WebsiteSchemaRum)`

SetRum sets Rum field to given value.

### HasRum

`func (o *WebsiteSchema) HasRum() bool`

HasRum returns a boolean if a field has been set.

### SetRumNil

`func (o *WebsiteSchema) SetRumNil(b bool)`

 SetRumNil sets the value for Rum to be an explicit nil

### UnsetRum
`func (o *WebsiteSchema) UnsetRum()`

UnsetRum ensures that no value is present for Rum, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


