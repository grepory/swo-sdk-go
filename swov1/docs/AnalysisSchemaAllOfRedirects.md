# AnalysisSchemaAllOfRedirects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**DurationInMs** | **int32** |  | 
**IpLookupDetails** | Pointer to [**AnalysisSchemaAllOfIpLookupDetails**](AnalysisSchemaAllOfIpLookupDetails.md) |  | [optional] 
**ConnectionDetails** | Pointer to [**AnalysisSchemaAllOfConnectionDetails**](AnalysisSchemaAllOfConnectionDetails.md) |  | [optional] 
**RequestDetails** | Pointer to [**HttpMessage**](HttpMessage.md) |  | [optional] 
**ResponseDetails** | Pointer to [**HttpMessage**](HttpMessage.md) |  | [optional] 

## Methods

### NewAnalysisSchemaAllOfRedirects

`func NewAnalysisSchemaAllOfRedirects(url string, durationInMs int32, ) *AnalysisSchemaAllOfRedirects`

NewAnalysisSchemaAllOfRedirects instantiates a new AnalysisSchemaAllOfRedirects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisSchemaAllOfRedirectsWithDefaults

`func NewAnalysisSchemaAllOfRedirectsWithDefaults() *AnalysisSchemaAllOfRedirects`

NewAnalysisSchemaAllOfRedirectsWithDefaults instantiates a new AnalysisSchemaAllOfRedirects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AnalysisSchemaAllOfRedirects) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AnalysisSchemaAllOfRedirects) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AnalysisSchemaAllOfRedirects) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDurationInMs

`func (o *AnalysisSchemaAllOfRedirects) GetDurationInMs() int32`

GetDurationInMs returns the DurationInMs field if non-nil, zero value otherwise.

### GetDurationInMsOk

`func (o *AnalysisSchemaAllOfRedirects) GetDurationInMsOk() (*int32, bool)`

GetDurationInMsOk returns a tuple with the DurationInMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMs

`func (o *AnalysisSchemaAllOfRedirects) SetDurationInMs(v int32)`

SetDurationInMs sets DurationInMs field to given value.


### GetIpLookupDetails

`func (o *AnalysisSchemaAllOfRedirects) GetIpLookupDetails() AnalysisSchemaAllOfIpLookupDetails`

GetIpLookupDetails returns the IpLookupDetails field if non-nil, zero value otherwise.

### GetIpLookupDetailsOk

`func (o *AnalysisSchemaAllOfRedirects) GetIpLookupDetailsOk() (*AnalysisSchemaAllOfIpLookupDetails, bool)`

GetIpLookupDetailsOk returns a tuple with the IpLookupDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpLookupDetails

`func (o *AnalysisSchemaAllOfRedirects) SetIpLookupDetails(v AnalysisSchemaAllOfIpLookupDetails)`

SetIpLookupDetails sets IpLookupDetails field to given value.

### HasIpLookupDetails

`func (o *AnalysisSchemaAllOfRedirects) HasIpLookupDetails() bool`

HasIpLookupDetails returns a boolean if a field has been set.

### GetConnectionDetails

`func (o *AnalysisSchemaAllOfRedirects) GetConnectionDetails() AnalysisSchemaAllOfConnectionDetails`

GetConnectionDetails returns the ConnectionDetails field if non-nil, zero value otherwise.

### GetConnectionDetailsOk

`func (o *AnalysisSchemaAllOfRedirects) GetConnectionDetailsOk() (*AnalysisSchemaAllOfConnectionDetails, bool)`

GetConnectionDetailsOk returns a tuple with the ConnectionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionDetails

`func (o *AnalysisSchemaAllOfRedirects) SetConnectionDetails(v AnalysisSchemaAllOfConnectionDetails)`

SetConnectionDetails sets ConnectionDetails field to given value.

### HasConnectionDetails

`func (o *AnalysisSchemaAllOfRedirects) HasConnectionDetails() bool`

HasConnectionDetails returns a boolean if a field has been set.

### GetRequestDetails

`func (o *AnalysisSchemaAllOfRedirects) GetRequestDetails() HttpMessage`

GetRequestDetails returns the RequestDetails field if non-nil, zero value otherwise.

### GetRequestDetailsOk

`func (o *AnalysisSchemaAllOfRedirects) GetRequestDetailsOk() (*HttpMessage, bool)`

GetRequestDetailsOk returns a tuple with the RequestDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDetails

`func (o *AnalysisSchemaAllOfRedirects) SetRequestDetails(v HttpMessage)`

SetRequestDetails sets RequestDetails field to given value.

### HasRequestDetails

`func (o *AnalysisSchemaAllOfRedirects) HasRequestDetails() bool`

HasRequestDetails returns a boolean if a field has been set.

### GetResponseDetails

`func (o *AnalysisSchemaAllOfRedirects) GetResponseDetails() HttpMessage`

GetResponseDetails returns the ResponseDetails field if non-nil, zero value otherwise.

### GetResponseDetailsOk

`func (o *AnalysisSchemaAllOfRedirects) GetResponseDetailsOk() (*HttpMessage, bool)`

GetResponseDetailsOk returns a tuple with the ResponseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseDetails

`func (o *AnalysisSchemaAllOfRedirects) SetResponseDetails(v HttpMessage)`

SetResponseDetails sets ResponseDetails field to given value.

### HasResponseDetails

`func (o *AnalysisSchemaAllOfRedirects) HasResponseDetails() bool`

HasResponseDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


