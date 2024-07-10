# EntitiesWebsitesEntityIdGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **string** |  | 
**Name** | **string** | Name of the website, which must be unique within the organization. The website must also not contain any control characters, any white space other than space (U+0020), or any consecutive, leading or trailing spaces.  | 
**Url** | **string** | URL of the website. Must be a valid URL with no leading or trailing white space. Must not contain invalid port number (&gt;65535).  | 
**AvailabilityCheckSettings** | Pointer to [**NullableWebsiteSchemaAvailabilityCheckSettings**](WebsiteSchemaAvailabilityCheckSettings.md) |  | [optional] 
**Rum** | Pointer to [**NullableWebsiteSchemaRum**](WebsiteSchemaRum.md) |  | [optional] 
**NextOnDemandAvailabilityTime** | Pointer to **float32** | Timestamp for when the next on-demand check could be executed.  If at \&quot;0\&quot;, it means you can execute it anytime.  | [optional] 

## Methods

### NewEntitiesWebsitesEntityIdGet200Response

`func NewEntitiesWebsitesEntityIdGet200Response(id string, type_ string, name string, url string, ) *EntitiesWebsitesEntityIdGet200Response`

NewEntitiesWebsitesEntityIdGet200Response instantiates a new EntitiesWebsitesEntityIdGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitiesWebsitesEntityIdGet200ResponseWithDefaults

`func NewEntitiesWebsitesEntityIdGet200ResponseWithDefaults() *EntitiesWebsitesEntityIdGet200Response`

NewEntitiesWebsitesEntityIdGet200ResponseWithDefaults instantiates a new EntitiesWebsitesEntityIdGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntitiesWebsitesEntityIdGet200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitiesWebsitesEntityIdGet200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitiesWebsitesEntityIdGet200Response) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *EntitiesWebsitesEntityIdGet200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntitiesWebsitesEntityIdGet200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntitiesWebsitesEntityIdGet200Response) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *EntitiesWebsitesEntityIdGet200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitiesWebsitesEntityIdGet200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitiesWebsitesEntityIdGet200Response) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *EntitiesWebsitesEntityIdGet200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EntitiesWebsitesEntityIdGet200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EntitiesWebsitesEntityIdGet200Response) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAvailabilityCheckSettings

`func (o *EntitiesWebsitesEntityIdGet200Response) GetAvailabilityCheckSettings() WebsiteSchemaAvailabilityCheckSettings`

GetAvailabilityCheckSettings returns the AvailabilityCheckSettings field if non-nil, zero value otherwise.

### GetAvailabilityCheckSettingsOk

`func (o *EntitiesWebsitesEntityIdGet200Response) GetAvailabilityCheckSettingsOk() (*WebsiteSchemaAvailabilityCheckSettings, bool)`

GetAvailabilityCheckSettingsOk returns a tuple with the AvailabilityCheckSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityCheckSettings

`func (o *EntitiesWebsitesEntityIdGet200Response) SetAvailabilityCheckSettings(v WebsiteSchemaAvailabilityCheckSettings)`

SetAvailabilityCheckSettings sets AvailabilityCheckSettings field to given value.

### HasAvailabilityCheckSettings

`func (o *EntitiesWebsitesEntityIdGet200Response) HasAvailabilityCheckSettings() bool`

HasAvailabilityCheckSettings returns a boolean if a field has been set.

### SetAvailabilityCheckSettingsNil

`func (o *EntitiesWebsitesEntityIdGet200Response) SetAvailabilityCheckSettingsNil(b bool)`

 SetAvailabilityCheckSettingsNil sets the value for AvailabilityCheckSettings to be an explicit nil

### UnsetAvailabilityCheckSettings
`func (o *EntitiesWebsitesEntityIdGet200Response) UnsetAvailabilityCheckSettings()`

UnsetAvailabilityCheckSettings ensures that no value is present for AvailabilityCheckSettings, not even an explicit nil
### GetRum

`func (o *EntitiesWebsitesEntityIdGet200Response) GetRum() WebsiteSchemaRum`

GetRum returns the Rum field if non-nil, zero value otherwise.

### GetRumOk

`func (o *EntitiesWebsitesEntityIdGet200Response) GetRumOk() (*WebsiteSchemaRum, bool)`

GetRumOk returns a tuple with the Rum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRum

`func (o *EntitiesWebsitesEntityIdGet200Response) SetRum(v WebsiteSchemaRum)`

SetRum sets Rum field to given value.

### HasRum

`func (o *EntitiesWebsitesEntityIdGet200Response) HasRum() bool`

HasRum returns a boolean if a field has been set.

### SetRumNil

`func (o *EntitiesWebsitesEntityIdGet200Response) SetRumNil(b bool)`

 SetRumNil sets the value for Rum to be an explicit nil

### UnsetRum
`func (o *EntitiesWebsitesEntityIdGet200Response) UnsetRum()`

UnsetRum ensures that no value is present for Rum, not even an explicit nil
### GetNextOnDemandAvailabilityTime

`func (o *EntitiesWebsitesEntityIdGet200Response) GetNextOnDemandAvailabilityTime() float32`

GetNextOnDemandAvailabilityTime returns the NextOnDemandAvailabilityTime field if non-nil, zero value otherwise.

### GetNextOnDemandAvailabilityTimeOk

`func (o *EntitiesWebsitesEntityIdGet200Response) GetNextOnDemandAvailabilityTimeOk() (*float32, bool)`

GetNextOnDemandAvailabilityTimeOk returns a tuple with the NextOnDemandAvailabilityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOnDemandAvailabilityTime

`func (o *EntitiesWebsitesEntityIdGet200Response) SetNextOnDemandAvailabilityTime(v float32)`

SetNextOnDemandAvailabilityTime sets NextOnDemandAvailabilityTime field to given value.

### HasNextOnDemandAvailabilityTime

`func (o *EntitiesWebsitesEntityIdGet200Response) HasNextOnDemandAvailabilityTime() bool`

HasNextOnDemandAvailabilityTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


