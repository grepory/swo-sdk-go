# TestFrom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ProbeLocationType**](ProbeLocationType.md) |  | 
**Values** | **[]string** | A list of probe location values of the selected &#x60;type&#x60;. At least one value matching an existing probe must be provided.  | 

## Methods

### NewTestFrom

`func NewTestFrom(type_ ProbeLocationType, values []string, ) *TestFrom`

NewTestFrom instantiates a new TestFrom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestFromWithDefaults

`func NewTestFromWithDefaults() *TestFrom`

NewTestFromWithDefaults instantiates a new TestFrom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TestFrom) GetType() ProbeLocationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestFrom) GetTypeOk() (*ProbeLocationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestFrom) SetType(v ProbeLocationType)`

SetType sets Type field to given value.


### GetValues

`func (o *TestFrom) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TestFrom) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TestFrom) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


