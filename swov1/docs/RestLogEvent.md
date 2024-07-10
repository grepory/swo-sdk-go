# RestLogEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Time** | **string** |  | 
**Message** | **string** |  | 
**Hostname** | **string** |  | 
**Severity** | **string** |  | 
**Program** | **string** |  | 

## Methods

### NewRestLogEvent

`func NewRestLogEvent(id string, time string, message string, hostname string, severity string, program string, ) *RestLogEvent`

NewRestLogEvent instantiates a new RestLogEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestLogEventWithDefaults

`func NewRestLogEventWithDefaults() *RestLogEvent`

NewRestLogEventWithDefaults instantiates a new RestLogEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestLogEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestLogEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestLogEvent) SetId(v string)`

SetId sets Id field to given value.


### GetTime

`func (o *RestLogEvent) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *RestLogEvent) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *RestLogEvent) SetTime(v string)`

SetTime sets Time field to given value.


### GetMessage

`func (o *RestLogEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RestLogEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RestLogEvent) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetHostname

`func (o *RestLogEvent) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *RestLogEvent) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *RestLogEvent) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetSeverity

`func (o *RestLogEvent) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *RestLogEvent) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *RestLogEvent) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetProgram

`func (o *RestLogEvent) GetProgram() string`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *RestLogEvent) GetProgramOk() (*string, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *RestLogEvent) SetProgram(v string)`

SetProgram sets Program field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


