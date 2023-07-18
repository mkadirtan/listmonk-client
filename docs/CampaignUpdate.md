# CampaignUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Lists** | Pointer to **[]int32** |  | [optional] 
**FromEmail** | Pointer to **string** |  | [optional] 
**Messenger** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**SendLater** | Pointer to **bool** |  | [optional] 
**SendAt** | Pointer to **map[string]interface{}** |  | [optional] 
**Headers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TemplateId** | Pointer to **int32** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**Altbody** | Pointer to **string** |  | [optional] 
**Archive** | Pointer to **bool** |  | [optional] 
**ArchiveTemplateId** | Pointer to **int32** |  | [optional] 
**ArchiveMeta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCampaignUpdate

`func NewCampaignUpdate() *CampaignUpdate`

NewCampaignUpdate instantiates a new CampaignUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignUpdateWithDefaults

`func NewCampaignUpdateWithDefaults() *CampaignUpdate`

NewCampaignUpdateWithDefaults instantiates a new CampaignUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubject

`func (o *CampaignUpdate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CampaignUpdate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CampaignUpdate) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CampaignUpdate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetLists

`func (o *CampaignUpdate) GetLists() []int32`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *CampaignUpdate) GetListsOk() (*[]int32, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *CampaignUpdate) SetLists(v []int32)`

SetLists sets Lists field to given value.

### HasLists

`func (o *CampaignUpdate) HasLists() bool`

HasLists returns a boolean if a field has been set.

### GetFromEmail

`func (o *CampaignUpdate) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *CampaignUpdate) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *CampaignUpdate) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *CampaignUpdate) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### GetMessenger

`func (o *CampaignUpdate) GetMessenger() string`

GetMessenger returns the Messenger field if non-nil, zero value otherwise.

### GetMessengerOk

`func (o *CampaignUpdate) GetMessengerOk() (*string, bool)`

GetMessengerOk returns a tuple with the Messenger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessenger

`func (o *CampaignUpdate) SetMessenger(v string)`

SetMessenger sets Messenger field to given value.

### HasMessenger

`func (o *CampaignUpdate) HasMessenger() bool`

HasMessenger returns a boolean if a field has been set.

### GetType

`func (o *CampaignUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CampaignUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTags

`func (o *CampaignUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CampaignUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CampaignUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CampaignUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSendLater

`func (o *CampaignUpdate) GetSendLater() bool`

GetSendLater returns the SendLater field if non-nil, zero value otherwise.

### GetSendLaterOk

`func (o *CampaignUpdate) GetSendLaterOk() (*bool, bool)`

GetSendLaterOk returns a tuple with the SendLater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendLater

`func (o *CampaignUpdate) SetSendLater(v bool)`

SetSendLater sets SendLater field to given value.

### HasSendLater

`func (o *CampaignUpdate) HasSendLater() bool`

HasSendLater returns a boolean if a field has been set.

### GetSendAt

`func (o *CampaignUpdate) GetSendAt() map[string]interface{}`

GetSendAt returns the SendAt field if non-nil, zero value otherwise.

### GetSendAtOk

`func (o *CampaignUpdate) GetSendAtOk() (*map[string]interface{}, bool)`

GetSendAtOk returns a tuple with the SendAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAt

`func (o *CampaignUpdate) SetSendAt(v map[string]interface{})`

SetSendAt sets SendAt field to given value.

### HasSendAt

`func (o *CampaignUpdate) HasSendAt() bool`

HasSendAt returns a boolean if a field has been set.

### GetHeaders

`func (o *CampaignUpdate) GetHeaders() []map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *CampaignUpdate) GetHeadersOk() (*[]map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *CampaignUpdate) SetHeaders(v []map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *CampaignUpdate) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetTemplateId

`func (o *CampaignUpdate) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CampaignUpdate) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CampaignUpdate) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *CampaignUpdate) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetContentType

`func (o *CampaignUpdate) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CampaignUpdate) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CampaignUpdate) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *CampaignUpdate) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetBody

`func (o *CampaignUpdate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CampaignUpdate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CampaignUpdate) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CampaignUpdate) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetAltbody

`func (o *CampaignUpdate) GetAltbody() string`

GetAltbody returns the Altbody field if non-nil, zero value otherwise.

### GetAltbodyOk

`func (o *CampaignUpdate) GetAltbodyOk() (*string, bool)`

GetAltbodyOk returns a tuple with the Altbody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltbody

`func (o *CampaignUpdate) SetAltbody(v string)`

SetAltbody sets Altbody field to given value.

### HasAltbody

`func (o *CampaignUpdate) HasAltbody() bool`

HasAltbody returns a boolean if a field has been set.

### GetArchive

`func (o *CampaignUpdate) GetArchive() bool`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *CampaignUpdate) GetArchiveOk() (*bool, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *CampaignUpdate) SetArchive(v bool)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *CampaignUpdate) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetArchiveTemplateId

`func (o *CampaignUpdate) GetArchiveTemplateId() int32`

GetArchiveTemplateId returns the ArchiveTemplateId field if non-nil, zero value otherwise.

### GetArchiveTemplateIdOk

`func (o *CampaignUpdate) GetArchiveTemplateIdOk() (*int32, bool)`

GetArchiveTemplateIdOk returns a tuple with the ArchiveTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveTemplateId

`func (o *CampaignUpdate) SetArchiveTemplateId(v int32)`

SetArchiveTemplateId sets ArchiveTemplateId field to given value.

### HasArchiveTemplateId

`func (o *CampaignUpdate) HasArchiveTemplateId() bool`

HasArchiveTemplateId returns a boolean if a field has been set.

### GetArchiveMeta

`func (o *CampaignUpdate) GetArchiveMeta() map[string]interface{}`

GetArchiveMeta returns the ArchiveMeta field if non-nil, zero value otherwise.

### GetArchiveMetaOk

`func (o *CampaignUpdate) GetArchiveMetaOk() (*map[string]interface{}, bool)`

GetArchiveMetaOk returns a tuple with the ArchiveMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveMeta

`func (o *CampaignUpdate) SetArchiveMeta(v map[string]interface{})`

SetArchiveMeta sets ArchiveMeta field to given value.

### HasArchiveMeta

`func (o *CampaignUpdate) HasArchiveMeta() bool`

HasArchiveMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


