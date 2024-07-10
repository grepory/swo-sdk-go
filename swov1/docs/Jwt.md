# Jwt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scopes** | **[]string** |  | 

## Methods

### NewJwt

`func NewJwt(scopes []string, ) *Jwt`

NewJwt instantiates a new Jwt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJwtWithDefaults

`func NewJwtWithDefaults() *Jwt`

NewJwtWithDefaults instantiates a new Jwt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopes

`func (o *Jwt) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *Jwt) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *Jwt) SetScopes(v []string)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


