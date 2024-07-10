# LogArchiveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogArchives** | [**[]LogArchiveModel**](LogArchiveModel.md) | List of archived logs data | 
**PageInfo** | [**PageInfoModel**](PageInfoModel.md) |  | 

## Methods

### NewLogArchiveResponse

`func NewLogArchiveResponse(logArchives []LogArchiveModel, pageInfo PageInfoModel, ) *LogArchiveResponse`

NewLogArchiveResponse instantiates a new LogArchiveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogArchiveResponseWithDefaults

`func NewLogArchiveResponseWithDefaults() *LogArchiveResponse`

NewLogArchiveResponseWithDefaults instantiates a new LogArchiveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogArchives

`func (o *LogArchiveResponse) GetLogArchives() []LogArchiveModel`

GetLogArchives returns the LogArchives field if non-nil, zero value otherwise.

### GetLogArchivesOk

`func (o *LogArchiveResponse) GetLogArchivesOk() (*[]LogArchiveModel, bool)`

GetLogArchivesOk returns a tuple with the LogArchives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogArchives

`func (o *LogArchiveResponse) SetLogArchives(v []LogArchiveModel)`

SetLogArchives sets LogArchives field to given value.


### GetPageInfo

`func (o *LogArchiveResponse) GetPageInfo() PageInfoModel`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *LogArchiveResponse) GetPageInfoOk() (*PageInfoModel, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *LogArchiveResponse) SetPageInfo(v PageInfoModel)`

SetPageInfo sets PageInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


