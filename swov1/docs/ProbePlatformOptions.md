# ProbePlatformOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProbePlatforms** | [**[]ProbePlatform**](ProbePlatform.md) |  | 
**TestFromAll** | Pointer to **bool** |  | [optional] 

## Methods

### NewProbePlatformOptions

`func NewProbePlatformOptions(probePlatforms []ProbePlatform, ) *ProbePlatformOptions`

NewProbePlatformOptions instantiates a new ProbePlatformOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProbePlatformOptionsWithDefaults

`func NewProbePlatformOptionsWithDefaults() *ProbePlatformOptions`

NewProbePlatformOptionsWithDefaults instantiates a new ProbePlatformOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProbePlatforms

`func (o *ProbePlatformOptions) GetProbePlatforms() []ProbePlatform`

GetProbePlatforms returns the ProbePlatforms field if non-nil, zero value otherwise.

### GetProbePlatformsOk

`func (o *ProbePlatformOptions) GetProbePlatformsOk() (*[]ProbePlatform, bool)`

GetProbePlatformsOk returns a tuple with the ProbePlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbePlatforms

`func (o *ProbePlatformOptions) SetProbePlatforms(v []ProbePlatform)`

SetProbePlatforms sets ProbePlatforms field to given value.


### GetTestFromAll

`func (o *ProbePlatformOptions) GetTestFromAll() bool`

GetTestFromAll returns the TestFromAll field if non-nil, zero value otherwise.

### GetTestFromAllOk

`func (o *ProbePlatformOptions) GetTestFromAllOk() (*bool, bool)`

GetTestFromAllOk returns a tuple with the TestFromAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestFromAll

`func (o *ProbePlatformOptions) SetTestFromAll(v bool)`

SetTestFromAll sets TestFromAll field to given value.

### HasTestFromAll

`func (o *ProbePlatformOptions) HasTestFromAll() bool`

HasTestFromAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


