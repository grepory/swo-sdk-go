# CheckForStringType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** | Defines whether the check should pass only when the string is present on the page (CONTAINS) or only when it is absent (DOES_NOT_CONTAIN).  | 
**Value** | **string** | The string that which will be searched in the page source code. | 

## Methods

### NewCheckForStringType

`func NewCheckForStringType(operator string, value string, ) *CheckForStringType`

NewCheckForStringType instantiates a new CheckForStringType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckForStringTypeWithDefaults

`func NewCheckForStringTypeWithDefaults() *CheckForStringType`

NewCheckForStringTypeWithDefaults instantiates a new CheckForStringType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *CheckForStringType) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CheckForStringType) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CheckForStringType) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *CheckForStringType) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CheckForStringType) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CheckForStringType) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


