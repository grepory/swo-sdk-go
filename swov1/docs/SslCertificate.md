# SslCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ValidTo** | **string** |  | 
**Certificate** | **string** |  | 

## Methods

### NewSslCertificate

`func NewSslCertificate(name string, validTo string, certificate string, ) *SslCertificate`

NewSslCertificate instantiates a new SslCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslCertificateWithDefaults

`func NewSslCertificateWithDefaults() *SslCertificate`

NewSslCertificateWithDefaults instantiates a new SslCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SslCertificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SslCertificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SslCertificate) SetName(v string)`

SetName sets Name field to given value.


### GetValidTo

`func (o *SslCertificate) GetValidTo() string`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *SslCertificate) GetValidToOk() (*string, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *SslCertificate) SetValidTo(v string)`

SetValidTo sets ValidTo field to given value.


### GetCertificate

`func (o *SslCertificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SslCertificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SslCertificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


