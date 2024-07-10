# ValidationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SslCertificates** | Pointer to [**SslCertificates**](SslCertificates.md) |  | [optional] 
**CheckForString** | Pointer to [**CheckForString**](CheckForString.md) |  | [optional] 

## Methods

### NewValidationSchema

`func NewValidationSchema() *ValidationSchema`

NewValidationSchema instantiates a new ValidationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationSchemaWithDefaults

`func NewValidationSchemaWithDefaults() *ValidationSchema`

NewValidationSchemaWithDefaults instantiates a new ValidationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSslCertificates

`func (o *ValidationSchema) GetSslCertificates() SslCertificates`

GetSslCertificates returns the SslCertificates field if non-nil, zero value otherwise.

### GetSslCertificatesOk

`func (o *ValidationSchema) GetSslCertificatesOk() (*SslCertificates, bool)`

GetSslCertificatesOk returns a tuple with the SslCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificates

`func (o *ValidationSchema) SetSslCertificates(v SslCertificates)`

SetSslCertificates sets SslCertificates field to given value.

### HasSslCertificates

`func (o *ValidationSchema) HasSslCertificates() bool`

HasSslCertificates returns a boolean if a field has been set.

### GetCheckForString

`func (o *ValidationSchema) GetCheckForString() CheckForString`

GetCheckForString returns the CheckForString field if non-nil, zero value otherwise.

### GetCheckForStringOk

`func (o *ValidationSchema) GetCheckForStringOk() (*CheckForString, bool)`

GetCheckForStringOk returns a tuple with the CheckForString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckForString

`func (o *ValidationSchema) SetCheckForString(v CheckForString)`

SetCheckForString sets CheckForString field to given value.

### HasCheckForString

`func (o *ValidationSchema) HasCheckForString() bool`

HasCheckForString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


