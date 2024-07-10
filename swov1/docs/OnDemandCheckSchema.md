# OnDemandCheckSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**NextOnDemandAvailabilityTime** | Pointer to **string** | Time when it will be possible to schedule another on-demand check | [optional] 
**OnDemandCheckStatus** | Pointer to [**OnDemandCheckStatus**](OnDemandCheckStatus.md) |  | [optional] 

## Methods

### NewOnDemandCheckSchema

`func NewOnDemandCheckSchema(id string, ) *OnDemandCheckSchema`

NewOnDemandCheckSchema instantiates a new OnDemandCheckSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnDemandCheckSchemaWithDefaults

`func NewOnDemandCheckSchemaWithDefaults() *OnDemandCheckSchema`

NewOnDemandCheckSchemaWithDefaults instantiates a new OnDemandCheckSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OnDemandCheckSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OnDemandCheckSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OnDemandCheckSchema) SetId(v string)`

SetId sets Id field to given value.


### GetNextOnDemandAvailabilityTime

`func (o *OnDemandCheckSchema) GetNextOnDemandAvailabilityTime() string`

GetNextOnDemandAvailabilityTime returns the NextOnDemandAvailabilityTime field if non-nil, zero value otherwise.

### GetNextOnDemandAvailabilityTimeOk

`func (o *OnDemandCheckSchema) GetNextOnDemandAvailabilityTimeOk() (*string, bool)`

GetNextOnDemandAvailabilityTimeOk returns a tuple with the NextOnDemandAvailabilityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOnDemandAvailabilityTime

`func (o *OnDemandCheckSchema) SetNextOnDemandAvailabilityTime(v string)`

SetNextOnDemandAvailabilityTime sets NextOnDemandAvailabilityTime field to given value.

### HasNextOnDemandAvailabilityTime

`func (o *OnDemandCheckSchema) HasNextOnDemandAvailabilityTime() bool`

HasNextOnDemandAvailabilityTime returns a boolean if a field has been set.

### GetOnDemandCheckStatus

`func (o *OnDemandCheckSchema) GetOnDemandCheckStatus() OnDemandCheckStatus`

GetOnDemandCheckStatus returns the OnDemandCheckStatus field if non-nil, zero value otherwise.

### GetOnDemandCheckStatusOk

`func (o *OnDemandCheckSchema) GetOnDemandCheckStatusOk() (*OnDemandCheckStatus, bool)`

GetOnDemandCheckStatusOk returns a tuple with the OnDemandCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemandCheckStatus

`func (o *OnDemandCheckSchema) SetOnDemandCheckStatus(v OnDemandCheckStatus)`

SetOnDemandCheckStatus sets OnDemandCheckStatus field to given value.

### HasOnDemandCheckStatus

`func (o *OnDemandCheckSchema) HasOnDemandCheckStatus() bool`

HasOnDemandCheckStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


