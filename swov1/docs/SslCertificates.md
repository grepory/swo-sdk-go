# SslCertificates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Status** | **string** |  | 
**MinTimeValid** | **int32** |  | 
**Certificates** | Pointer to [**SslCertificatesAllOfCertificates**](SslCertificatesAllOfCertificates.md) |  | [optional] 

## Methods

### NewSslCertificates

`func NewSslCertificates(type_ string, status string, minTimeValid int32, ) *SslCertificates`

NewSslCertificates instantiates a new SslCertificates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslCertificatesWithDefaults

`func NewSslCertificatesWithDefaults() *SslCertificates`

NewSslCertificatesWithDefaults instantiates a new SslCertificates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SslCertificates) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SslCertificates) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SslCertificates) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *SslCertificates) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SslCertificates) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SslCertificates) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMinTimeValid

`func (o *SslCertificates) GetMinTimeValid() int32`

GetMinTimeValid returns the MinTimeValid field if non-nil, zero value otherwise.

### GetMinTimeValidOk

`func (o *SslCertificates) GetMinTimeValidOk() (*int32, bool)`

GetMinTimeValidOk returns a tuple with the MinTimeValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTimeValid

`func (o *SslCertificates) SetMinTimeValid(v int32)`

SetMinTimeValid sets MinTimeValid field to given value.


### GetCertificates

`func (o *SslCertificates) GetCertificates() SslCertificatesAllOfCertificates`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *SslCertificates) GetCertificatesOk() (*SslCertificatesAllOfCertificates, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *SslCertificates) SetCertificates(v SslCertificatesAllOfCertificates)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *SslCertificates) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


