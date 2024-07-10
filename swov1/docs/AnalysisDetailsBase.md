# AnalysisDetailsBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationInMs** | **int32** |  | 
**ErrorMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewAnalysisDetailsBase

`func NewAnalysisDetailsBase(durationInMs int32, ) *AnalysisDetailsBase`

NewAnalysisDetailsBase instantiates a new AnalysisDetailsBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisDetailsBaseWithDefaults

`func NewAnalysisDetailsBaseWithDefaults() *AnalysisDetailsBase`

NewAnalysisDetailsBaseWithDefaults instantiates a new AnalysisDetailsBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationInMs

`func (o *AnalysisDetailsBase) GetDurationInMs() int32`

GetDurationInMs returns the DurationInMs field if non-nil, zero value otherwise.

### GetDurationInMsOk

`func (o *AnalysisDetailsBase) GetDurationInMsOk() (*int32, bool)`

GetDurationInMsOk returns a tuple with the DurationInMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMs

`func (o *AnalysisDetailsBase) SetDurationInMs(v int32)`

SetDurationInMs sets DurationInMs field to given value.


### GetErrorMessage

`func (o *AnalysisDetailsBase) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AnalysisDetailsBase) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AnalysisDetailsBase) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AnalysisDetailsBase) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


