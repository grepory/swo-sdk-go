# InnerError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | A more specific error code than was provided by the containing error. | 
**Message** | Pointer to **string** | A human-readable representation of the error. | [optional] 
**Target** | Pointer to **string** | The target of the error. | [optional] 
**Innererror** | Pointer to [**InnerError**](InnerError.md) |  | [optional] 

## Methods

### NewInnerError

`func NewInnerError(code string, ) *InnerError`

NewInnerError instantiates a new InnerError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInnerErrorWithDefaults

`func NewInnerErrorWithDefaults() *InnerError`

NewInnerErrorWithDefaults instantiates a new InnerError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *InnerError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InnerError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InnerError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *InnerError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InnerError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InnerError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InnerError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTarget

`func (o *InnerError) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *InnerError) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *InnerError) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *InnerError) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetInnererror

`func (o *InnerError) GetInnererror() InnerError`

GetInnererror returns the Innererror field if non-nil, zero value otherwise.

### GetInnererrorOk

`func (o *InnerError) GetInnererrorOk() (*InnerError, bool)`

GetInnererrorOk returns a tuple with the Innererror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnererror

`func (o *InnerError) SetInnererror(v InnerError)`

SetInnererror sets Innererror field to given value.

### HasInnererror

`func (o *InnerError) HasInnererror() bool`

HasInnererror returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


