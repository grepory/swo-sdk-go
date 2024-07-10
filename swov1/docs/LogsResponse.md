# LogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | [**[]RestLogEvent**](RestLogEvent.md) |  | 
**PageInfo** | Pointer to [**PageInfo**](PageInfo.md) |  | [optional] 

## Methods

### NewLogsResponse

`func NewLogsResponse(logs []RestLogEvent, ) *LogsResponse`

NewLogsResponse instantiates a new LogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsResponseWithDefaults

`func NewLogsResponseWithDefaults() *LogsResponse`

NewLogsResponseWithDefaults instantiates a new LogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *LogsResponse) GetLogs() []RestLogEvent`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *LogsResponse) GetLogsOk() (*[]RestLogEvent, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *LogsResponse) SetLogs(v []RestLogEvent)`

SetLogs sets Logs field to given value.


### GetPageInfo

`func (o *LogsResponse) GetPageInfo() PageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *LogsResponse) GetPageInfoOk() (*PageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *LogsResponse) SetPageInfo(v PageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *LogsResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


