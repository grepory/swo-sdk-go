# SslMonitoring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** | Whether SSL monitoring is enabled for the website.  If set to false, SSL monitoring will be disabled, but the other settings will be remembered in case you re-enable it later.  If omitted, the previous setting will stay in effect. If there is no previous setting, the value will default to false.  | [optional] 
**DaysPriorToExpiration** | Pointer to **NullableInt32** | Number of days before the expiration date an SSL certificate will be considered &#39;expiring.&#39;  | [optional] 
**IgnoreIntermediateCertificates** | Pointer to **NullableBool** | Use this option to limit the certificate expiration check to only the first certificate in the chain (normally the host certificate). This way you will not be warned about impending expiration of intermediate or root Certification Authority certificates in the chain.  This option does not affect any other certificate validity checks besides expiration.  If omitted, the previous setting will stay in effect. If there is no previous setting, the value will default to false.  | [optional] 

## Methods

### NewSslMonitoring

`func NewSslMonitoring() *SslMonitoring`

NewSslMonitoring instantiates a new SslMonitoring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslMonitoringWithDefaults

`func NewSslMonitoringWithDefaults() *SslMonitoring`

NewSslMonitoringWithDefaults instantiates a new SslMonitoring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SslMonitoring) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SslMonitoring) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SslMonitoring) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SslMonitoring) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *SslMonitoring) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *SslMonitoring) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetDaysPriorToExpiration

`func (o *SslMonitoring) GetDaysPriorToExpiration() int32`

GetDaysPriorToExpiration returns the DaysPriorToExpiration field if non-nil, zero value otherwise.

### GetDaysPriorToExpirationOk

`func (o *SslMonitoring) GetDaysPriorToExpirationOk() (*int32, bool)`

GetDaysPriorToExpirationOk returns a tuple with the DaysPriorToExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysPriorToExpiration

`func (o *SslMonitoring) SetDaysPriorToExpiration(v int32)`

SetDaysPriorToExpiration sets DaysPriorToExpiration field to given value.

### HasDaysPriorToExpiration

`func (o *SslMonitoring) HasDaysPriorToExpiration() bool`

HasDaysPriorToExpiration returns a boolean if a field has been set.

### SetDaysPriorToExpirationNil

`func (o *SslMonitoring) SetDaysPriorToExpirationNil(b bool)`

 SetDaysPriorToExpirationNil sets the value for DaysPriorToExpiration to be an explicit nil

### UnsetDaysPriorToExpiration
`func (o *SslMonitoring) UnsetDaysPriorToExpiration()`

UnsetDaysPriorToExpiration ensures that no value is present for DaysPriorToExpiration, not even an explicit nil
### GetIgnoreIntermediateCertificates

`func (o *SslMonitoring) GetIgnoreIntermediateCertificates() bool`

GetIgnoreIntermediateCertificates returns the IgnoreIntermediateCertificates field if non-nil, zero value otherwise.

### GetIgnoreIntermediateCertificatesOk

`func (o *SslMonitoring) GetIgnoreIntermediateCertificatesOk() (*bool, bool)`

GetIgnoreIntermediateCertificatesOk returns a tuple with the IgnoreIntermediateCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreIntermediateCertificates

`func (o *SslMonitoring) SetIgnoreIntermediateCertificates(v bool)`

SetIgnoreIntermediateCertificates sets IgnoreIntermediateCertificates field to given value.

### HasIgnoreIntermediateCertificates

`func (o *SslMonitoring) HasIgnoreIntermediateCertificates() bool`

HasIgnoreIntermediateCertificates returns a boolean if a field has been set.

### SetIgnoreIntermediateCertificatesNil

`func (o *SslMonitoring) SetIgnoreIntermediateCertificatesNil(b bool)`

 SetIgnoreIntermediateCertificatesNil sets the value for IgnoreIntermediateCertificates to be an explicit nil

### UnsetIgnoreIntermediateCertificates
`func (o *SslMonitoring) UnsetIgnoreIntermediateCertificates()`

UnsetIgnoreIntermediateCertificates ensures that no value is present for IgnoreIntermediateCertificates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


