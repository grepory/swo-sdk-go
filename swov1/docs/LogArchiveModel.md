# LogArchiveModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the file | 
**Name** | **string** | Name of the file | 
**DownloadUrl** | **string** | URL location of the file | 
**ArchivedTimestamp** | **string** | Epoc timestamp (in seconds) of archival | 
**ArchiveSize** | **float64** | Size of file in bytes | 

## Methods

### NewLogArchiveModel

`func NewLogArchiveModel(id string, name string, downloadUrl string, archivedTimestamp string, archiveSize float64, ) *LogArchiveModel`

NewLogArchiveModel instantiates a new LogArchiveModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogArchiveModelWithDefaults

`func NewLogArchiveModelWithDefaults() *LogArchiveModel`

NewLogArchiveModelWithDefaults instantiates a new LogArchiveModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogArchiveModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogArchiveModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogArchiveModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LogArchiveModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogArchiveModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogArchiveModel) SetName(v string)`

SetName sets Name field to given value.


### GetDownloadUrl

`func (o *LogArchiveModel) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *LogArchiveModel) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *LogArchiveModel) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### GetArchivedTimestamp

`func (o *LogArchiveModel) GetArchivedTimestamp() string`

GetArchivedTimestamp returns the ArchivedTimestamp field if non-nil, zero value otherwise.

### GetArchivedTimestampOk

`func (o *LogArchiveModel) GetArchivedTimestampOk() (*string, bool)`

GetArchivedTimestampOk returns a tuple with the ArchivedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedTimestamp

`func (o *LogArchiveModel) SetArchivedTimestamp(v string)`

SetArchivedTimestamp sets ArchivedTimestamp field to given value.


### GetArchiveSize

`func (o *LogArchiveModel) GetArchiveSize() float64`

GetArchiveSize returns the ArchiveSize field if non-nil, zero value otherwise.

### GetArchiveSizeOk

`func (o *LogArchiveModel) GetArchiveSizeOk() (*float64, bool)`

GetArchiveSizeOk returns a tuple with the ArchiveSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveSize

`func (o *LogArchiveModel) SetArchiveSize(v float64)`

SetArchiveSize sets ArchiveSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


