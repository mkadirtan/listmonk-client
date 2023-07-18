# UpdateCampaignArchiveByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archive** | Pointer to **bool** |  | [optional] 
**ArchiveTemplateId** | Pointer to **int32** |  | [optional] 
**ArchiveMeta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateCampaignArchiveByIdRequest

`func NewUpdateCampaignArchiveByIdRequest() *UpdateCampaignArchiveByIdRequest`

NewUpdateCampaignArchiveByIdRequest instantiates a new UpdateCampaignArchiveByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCampaignArchiveByIdRequestWithDefaults

`func NewUpdateCampaignArchiveByIdRequestWithDefaults() *UpdateCampaignArchiveByIdRequest`

NewUpdateCampaignArchiveByIdRequestWithDefaults instantiates a new UpdateCampaignArchiveByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchive

`func (o *UpdateCampaignArchiveByIdRequest) GetArchive() bool`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *UpdateCampaignArchiveByIdRequest) GetArchiveOk() (*bool, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *UpdateCampaignArchiveByIdRequest) SetArchive(v bool)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *UpdateCampaignArchiveByIdRequest) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetArchiveTemplateId

`func (o *UpdateCampaignArchiveByIdRequest) GetArchiveTemplateId() int32`

GetArchiveTemplateId returns the ArchiveTemplateId field if non-nil, zero value otherwise.

### GetArchiveTemplateIdOk

`func (o *UpdateCampaignArchiveByIdRequest) GetArchiveTemplateIdOk() (*int32, bool)`

GetArchiveTemplateIdOk returns a tuple with the ArchiveTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveTemplateId

`func (o *UpdateCampaignArchiveByIdRequest) SetArchiveTemplateId(v int32)`

SetArchiveTemplateId sets ArchiveTemplateId field to given value.

### HasArchiveTemplateId

`func (o *UpdateCampaignArchiveByIdRequest) HasArchiveTemplateId() bool`

HasArchiveTemplateId returns a boolean if a field has been set.

### GetArchiveMeta

`func (o *UpdateCampaignArchiveByIdRequest) GetArchiveMeta() map[string]interface{}`

GetArchiveMeta returns the ArchiveMeta field if non-nil, zero value otherwise.

### GetArchiveMetaOk

`func (o *UpdateCampaignArchiveByIdRequest) GetArchiveMetaOk() (*map[string]interface{}, bool)`

GetArchiveMetaOk returns a tuple with the ArchiveMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveMeta

`func (o *UpdateCampaignArchiveByIdRequest) SetArchiveMeta(v map[string]interface{})`

SetArchiveMeta sets ArchiveMeta field to given value.

### HasArchiveMeta

`func (o *UpdateCampaignArchiveByIdRequest) HasArchiveMeta() bool`

HasArchiveMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


