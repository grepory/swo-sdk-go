# WebsiteSchemaAvailabilityCheckSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckForString** | Pointer to [**NullableCheckForStringType**](CheckForStringType.md) |  | [optional] 
**TestIntervalInSeconds** | **int32** | Configure how often availability tests should be performed.  Provide a number of seconds that is divisible by 60 and no greater than 14400 (4 hours).  | 
**Protocols** | [**[]WebsiteProtocol**](WebsiteProtocol.md) | Configure which protocols need availability tests to be performed.  At least one protocol must be provided.  | 
**PlatformOptions** | Pointer to [**ProbePlatformOptions**](ProbePlatformOptions.md) |  | [optional] 
**TestFrom** | [**TestFrom**](TestFrom.md) |  | 
**Ssl** | Pointer to [**NullableSslMonitoring**](SslMonitoring.md) |  | [optional] 
**CustomHeaders** | Pointer to [**[]CustomHeadersInner**](CustomHeadersInner.md) | Configure custom request headers to be sent with each availability test. It is possible to provide multiple headers with the same name.  If omitted, set to null or set to an empty array, no custom headers will be sent.  | [optional] 
**AllowInsecureRenegotiation** | Pointer to **bool** | Allow insecure SSL renegotiation which introduces a security risk in the communication process.  Checking this option could lead to exposing to unauthorized entities and possibility of unauthorized access, interception, or manipulation of sensitive data, compromising the integrity and security of the communication channel.  Available only with HTTPS check.  If omitted or set to null, insecure SSL renegotiation won&#39;t be allowed.  | [optional] 

## Methods

### NewWebsiteSchemaAvailabilityCheckSettings

`func NewWebsiteSchemaAvailabilityCheckSettings(testIntervalInSeconds int32, protocols []WebsiteProtocol, testFrom TestFrom, ) *WebsiteSchemaAvailabilityCheckSettings`

NewWebsiteSchemaAvailabilityCheckSettings instantiates a new WebsiteSchemaAvailabilityCheckSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteSchemaAvailabilityCheckSettingsWithDefaults

`func NewWebsiteSchemaAvailabilityCheckSettingsWithDefaults() *WebsiteSchemaAvailabilityCheckSettings`

NewWebsiteSchemaAvailabilityCheckSettingsWithDefaults instantiates a new WebsiteSchemaAvailabilityCheckSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckForString

`func (o *WebsiteSchemaAvailabilityCheckSettings) GetCheckForString() CheckForStringType`

GetCheckForString returns the CheckForString field if non-nil, zero value otherwise.

### GetCheckForStringOk

`func (o *WebsiteSchemaAvailabilityCheckSettings) GetCheckForStringOk() (*CheckForStringType, bool)`

GetCheckForStringOk returns a tuple with the CheckForString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckForString

`func (o *WebsiteSchemaAvailabilityCheckSettings) SetCheckForString(v CheckForStringType)`

SetCheckForString sets CheckForString field to given value.

### HasCheckForString

`func (o *WebsiteSchemaAvailabilityCheckSettings) HasCheckForString() bool`

HasCheckForString returns a boolean if a field has been set.

### SetCheckForStringNil

`func (o *WebsiteSchemaAvailabilityCheckSettings) SetCheckForStringNil(b bool)`

 SetCheckForStringNil sets the value for CheckForString to be an explicit nil

### UnsetCheckForString
`func (o *WebsiteSchemaAvailabilityCheckSettings) UnsetCheckForString()`

UnsetCheckForString ensures that no value is present for CheckForString, not even an explicit nil
### GetTestIntervalInSeconds

`func (o *WebsiteSchemaAvailabilityCheckSettings) GetTestIntervalInSeconds() int32`

GetTestIntervalInSeconds returns the TestIntervalInSeconds field if non-nil, zero value otherwise.

### GetTestIntervalInSecondsOk

`func (o *WebsiteSchemaAvailabilityCheckSettings) GetTestIntervalInSecondsOk() (*int32, bool)`

GetTestIntervalInSecondsOk returns a tuple with the TestIntervalInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestIntervalInSeconds

`func (o *WebsiteSchemaAvailabilityCheckSettings) SetTestIntervalInSeconds(v int32)`

SetTestIntervalInSeconds sets TestIntervalInSeconds field to given value.


### GetProtocols

`func (o *WebsiteSchemaAvailabilityCheckSettings) GetProtocols() []WebsiteProtocol`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *WebsiteSchemaAvailabilityCheckSettings) GetProtocolsOk() (*[]WebsiteProtocol, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *WebsiteSchemaAvailabilityCheckSettings) SetProtocols(v []WebsiteProtocol)`

SetProtocols sets Protocols field to given value.


### GetPlatformOptions

`func (o *WebsiteSchemaAvailabilityCheckSettings) GetPlatformOptions() ProbePlatformOptions`

GetPlatformOptions returns the PlatformOptions field if non-nil, zero value otherwise.

### GetPlatformOptionsOk

`func (o *WebsiteSchemaAvailabilityCheckSettings) GetPlatformOptionsOk() (*ProbePlatformOptions, bool)`

GetPlatformOptionsOk returns a tuple with the PlatformOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformOptions

`func (o *WebsiteSchemaAvailabilityCheckSettings) SetPlatformOptions(v ProbePlatformOptions)`

SetPlatformOptions sets PlatformOptions field to given value.

### HasPlatformOptions

`func (o *WebsiteSchemaAvailabilityCheckSettings) HasPlatformOptions() bool`

HasPlatformOptions returns a boolean if a field has been set.

### GetTestFrom

`func (o *WebsiteSchemaAvailabilityCheckSettings) GetTestFrom() TestFrom`

GetTestFrom returns the TestFrom field if non-nil, zero value otherwise.

### GetTestFromOk

`func (o *WebsiteSchemaAvailabilityCheckSettings) GetTestFromOk() (*TestFrom, bool)`

GetTestFromOk returns a tuple with the TestFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestFrom

`func (o *WebsiteSchemaAvailabilityCheckSettings) SetTestFrom(v TestFrom)`

SetTestFrom sets TestFrom field to given value.


### GetSsl

`func (o *WebsiteSchemaAvailabilityCheckSettings) GetSsl() SslMonitoring`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *WebsiteSchemaAvailabilityCheckSettings) GetSslOk() (*SslMonitoring, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *WebsiteSchemaAvailabilityCheckSettings) SetSsl(v SslMonitoring)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *WebsiteSchemaAvailabilityCheckSettings) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### SetSslNil

`func (o *WebsiteSchemaAvailabilityCheckSettings) SetSslNil(b bool)`

 SetSslNil sets the value for Ssl to be an explicit nil

### UnsetSsl
`func (o *WebsiteSchemaAvailabilityCheckSettings) UnsetSsl()`

UnsetSsl ensures that no value is present for Ssl, not even an explicit nil
### GetCustomHeaders

`func (o *WebsiteSchemaAvailabilityCheckSettings) GetCustomHeaders() []CustomHeadersInner`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *WebsiteSchemaAvailabilityCheckSettings) GetCustomHeadersOk() (*[]CustomHeadersInner, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *WebsiteSchemaAvailabilityCheckSettings) SetCustomHeaders(v []CustomHeadersInner)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *WebsiteSchemaAvailabilityCheckSettings) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### SetCustomHeadersNil

`func (o *WebsiteSchemaAvailabilityCheckSettings) SetCustomHeadersNil(b bool)`

 SetCustomHeadersNil sets the value for CustomHeaders to be an explicit nil

### UnsetCustomHeaders
`func (o *WebsiteSchemaAvailabilityCheckSettings) UnsetCustomHeaders()`

UnsetCustomHeaders ensures that no value is present for CustomHeaders, not even an explicit nil
### GetAllowInsecureRenegotiation

`func (o *WebsiteSchemaAvailabilityCheckSettings) GetAllowInsecureRenegotiation() bool`

GetAllowInsecureRenegotiation returns the AllowInsecureRenegotiation field if non-nil, zero value otherwise.

### GetAllowInsecureRenegotiationOk

`func (o *WebsiteSchemaAvailabilityCheckSettings) GetAllowInsecureRenegotiationOk() (*bool, bool)`

GetAllowInsecureRenegotiationOk returns a tuple with the AllowInsecureRenegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecureRenegotiation

`func (o *WebsiteSchemaAvailabilityCheckSettings) SetAllowInsecureRenegotiation(v bool)`

SetAllowInsecureRenegotiation sets AllowInsecureRenegotiation field to given value.

### HasAllowInsecureRenegotiation

`func (o *WebsiteSchemaAvailabilityCheckSettings) HasAllowInsecureRenegotiation() bool`

HasAllowInsecureRenegotiation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


