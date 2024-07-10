# MeasurementGrouping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]Attribute**](Attribute.md) |  | [optional] 
**Measurements** | [**[]Measurement**](Measurement.md) |  | 

## Methods

### NewMeasurementGrouping

`func NewMeasurementGrouping(measurements []Measurement, ) *MeasurementGrouping`

NewMeasurementGrouping instantiates a new MeasurementGrouping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasurementGroupingWithDefaults

`func NewMeasurementGroupingWithDefaults() *MeasurementGrouping`

NewMeasurementGroupingWithDefaults instantiates a new MeasurementGrouping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *MeasurementGrouping) GetAttributes() []Attribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MeasurementGrouping) GetAttributesOk() (*[]Attribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MeasurementGrouping) SetAttributes(v []Attribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MeasurementGrouping) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetMeasurements

`func (o *MeasurementGrouping) GetMeasurements() []Measurement`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *MeasurementGrouping) GetMeasurementsOk() (*[]Measurement, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *MeasurementGrouping) SetMeasurements(v []Measurement)`

SetMeasurements sets Measurements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


