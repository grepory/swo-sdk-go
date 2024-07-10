# GetMetricResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Units** | Pointer to **string** |  | [optional] 
**Formula** | Pointer to **string** |  | [optional] 
**LastReportedTime** | Pointer to **string** |  | [optional] 

## Methods

### NewGetMetricResponse

`func NewGetMetricResponse(name string, ) *GetMetricResponse`

NewGetMetricResponse instantiates a new GetMetricResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMetricResponseWithDefaults

`func NewGetMetricResponseWithDefaults() *GetMetricResponse`

NewGetMetricResponseWithDefaults instantiates a new GetMetricResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetMetricResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetMetricResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetMetricResponse) SetName(v string)`

SetName sets Name field to given value.


### GetUnits

`func (o *GetMetricResponse) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *GetMetricResponse) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *GetMetricResponse) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *GetMetricResponse) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetFormula

`func (o *GetMetricResponse) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *GetMetricResponse) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *GetMetricResponse) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *GetMetricResponse) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetLastReportedTime

`func (o *GetMetricResponse) GetLastReportedTime() string`

GetLastReportedTime returns the LastReportedTime field if non-nil, zero value otherwise.

### GetLastReportedTimeOk

`func (o *GetMetricResponse) GetLastReportedTimeOk() (*string, bool)`

GetLastReportedTimeOk returns a tuple with the LastReportedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedTime

`func (o *GetMetricResponse) SetLastReportedTime(v string)`

SetLastReportedTime sets LastReportedTime field to given value.

### HasLastReportedTime

`func (o *GetMetricResponse) HasLastReportedTime() bool`

HasLastReportedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


