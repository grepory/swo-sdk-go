# HttpMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationInMs** | **int32** |  | 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Headers** | [**[]HttpHeader**](HttpHeader.md) |  | 
**Body** | Pointer to **string** |  | [optional] 

## Methods

### NewHttpMessage

`func NewHttpMessage(durationInMs int32, headers []HttpHeader, ) *HttpMessage`

NewHttpMessage instantiates a new HttpMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpMessageWithDefaults

`func NewHttpMessageWithDefaults() *HttpMessage`

NewHttpMessageWithDefaults instantiates a new HttpMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationInMs

`func (o *HttpMessage) GetDurationInMs() int32`

GetDurationInMs returns the DurationInMs field if non-nil, zero value otherwise.

### GetDurationInMsOk

`func (o *HttpMessage) GetDurationInMsOk() (*int32, bool)`

GetDurationInMsOk returns a tuple with the DurationInMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMs

`func (o *HttpMessage) SetDurationInMs(v int32)`

SetDurationInMs sets DurationInMs field to given value.


### GetErrorMessage

`func (o *HttpMessage) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *HttpMessage) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *HttpMessage) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *HttpMessage) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetHeaders

`func (o *HttpMessage) GetHeaders() []HttpHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HttpMessage) GetHeadersOk() (*[]HttpHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HttpMessage) SetHeaders(v []HttpHeader)`

SetHeaders sets Headers field to given value.


### GetBody

`func (o *HttpMessage) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *HttpMessage) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *HttpMessage) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *HttpMessage) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


