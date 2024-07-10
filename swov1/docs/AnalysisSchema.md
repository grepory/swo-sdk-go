# AnalysisSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationInMs** | **int32** |  | 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**AnalysisId** | **string** |  | 
**ResolveIp** | Pointer to [**AnalysisSchemaAllOfResolveIp**](AnalysisSchemaAllOfResolveIp.md) |  | [optional] 
**Connection** | Pointer to [**AnalysisSchemaAllOfConnection**](AnalysisSchemaAllOfConnection.md) |  | [optional] 
**Message** | Pointer to [**AnalysisSchemaAllOfMessage**](AnalysisSchemaAllOfMessage.md) |  | [optional] 
**Output** | Pointer to [**AnalysisSchemaAllOfOutput**](AnalysisSchemaAllOfOutput.md) |  | [optional] 
**Traceroute** | Pointer to [**AnalysisSchemaAllOfTraceroute**](AnalysisSchemaAllOfTraceroute.md) |  | [optional] 
**Redirects** | Pointer to [**[]AnalysisSchemaAllOfRedirects**](AnalysisSchemaAllOfRedirects.md) |  | [optional] 

## Methods

### NewAnalysisSchema

`func NewAnalysisSchema(durationInMs int32, analysisId string, ) *AnalysisSchema`

NewAnalysisSchema instantiates a new AnalysisSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisSchemaWithDefaults

`func NewAnalysisSchemaWithDefaults() *AnalysisSchema`

NewAnalysisSchemaWithDefaults instantiates a new AnalysisSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationInMs

`func (o *AnalysisSchema) GetDurationInMs() int32`

GetDurationInMs returns the DurationInMs field if non-nil, zero value otherwise.

### GetDurationInMsOk

`func (o *AnalysisSchema) GetDurationInMsOk() (*int32, bool)`

GetDurationInMsOk returns a tuple with the DurationInMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMs

`func (o *AnalysisSchema) SetDurationInMs(v int32)`

SetDurationInMs sets DurationInMs field to given value.


### GetErrorMessage

`func (o *AnalysisSchema) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AnalysisSchema) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AnalysisSchema) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AnalysisSchema) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetAnalysisId

`func (o *AnalysisSchema) GetAnalysisId() string`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *AnalysisSchema) GetAnalysisIdOk() (*string, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *AnalysisSchema) SetAnalysisId(v string)`

SetAnalysisId sets AnalysisId field to given value.


### GetResolveIp

`func (o *AnalysisSchema) GetResolveIp() AnalysisSchemaAllOfResolveIp`

GetResolveIp returns the ResolveIp field if non-nil, zero value otherwise.

### GetResolveIpOk

`func (o *AnalysisSchema) GetResolveIpOk() (*AnalysisSchemaAllOfResolveIp, bool)`

GetResolveIpOk returns a tuple with the ResolveIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveIp

`func (o *AnalysisSchema) SetResolveIp(v AnalysisSchemaAllOfResolveIp)`

SetResolveIp sets ResolveIp field to given value.

### HasResolveIp

`func (o *AnalysisSchema) HasResolveIp() bool`

HasResolveIp returns a boolean if a field has been set.

### GetConnection

`func (o *AnalysisSchema) GetConnection() AnalysisSchemaAllOfConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *AnalysisSchema) GetConnectionOk() (*AnalysisSchemaAllOfConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *AnalysisSchema) SetConnection(v AnalysisSchemaAllOfConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *AnalysisSchema) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetMessage

`func (o *AnalysisSchema) GetMessage() AnalysisSchemaAllOfMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AnalysisSchema) GetMessageOk() (*AnalysisSchemaAllOfMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AnalysisSchema) SetMessage(v AnalysisSchemaAllOfMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AnalysisSchema) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOutput

`func (o *AnalysisSchema) GetOutput() AnalysisSchemaAllOfOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *AnalysisSchema) GetOutputOk() (*AnalysisSchemaAllOfOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *AnalysisSchema) SetOutput(v AnalysisSchemaAllOfOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *AnalysisSchema) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetTraceroute

`func (o *AnalysisSchema) GetTraceroute() AnalysisSchemaAllOfTraceroute`

GetTraceroute returns the Traceroute field if non-nil, zero value otherwise.

### GetTracerouteOk

`func (o *AnalysisSchema) GetTracerouteOk() (*AnalysisSchemaAllOfTraceroute, bool)`

GetTracerouteOk returns a tuple with the Traceroute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceroute

`func (o *AnalysisSchema) SetTraceroute(v AnalysisSchemaAllOfTraceroute)`

SetTraceroute sets Traceroute field to given value.

### HasTraceroute

`func (o *AnalysisSchema) HasTraceroute() bool`

HasTraceroute returns a boolean if a field has been set.

### GetRedirects

`func (o *AnalysisSchema) GetRedirects() []AnalysisSchemaAllOfRedirects`

GetRedirects returns the Redirects field if non-nil, zero value otherwise.

### GetRedirectsOk

`func (o *AnalysisSchema) GetRedirectsOk() (*[]AnalysisSchemaAllOfRedirects, bool)`

GetRedirectsOk returns a tuple with the Redirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirects

`func (o *AnalysisSchema) SetRedirects(v []AnalysisSchemaAllOfRedirects)`

SetRedirects sets Redirects field to given value.

### HasRedirects

`func (o *AnalysisSchema) HasRedirects() bool`

HasRedirects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


