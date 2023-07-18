# CampaignContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**CampaignID** | Pointer to **int32** |  | [optional] 
**Views** | Pointer to **int32** |  | [optional] 
**Clicks** | Pointer to **int32** |  | [optional] 
**Lists** | Pointer to [**[]BounceResultsInnerCampaign**](BounceResultsInnerCampaign.md) |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**ToSend** | Pointer to **int32** |  | [optional] 
**Sent** | Pointer to **int32** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**FromEmail** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**SendAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TemplateId** | Pointer to **int32** |  | [optional] 
**Messenger** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 

## Methods

### NewCampaignContentRequest

`func NewCampaignContentRequest() *CampaignContentRequest`

NewCampaignContentRequest instantiates a new CampaignContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignContentRequestWithDefaults

`func NewCampaignContentRequestWithDefaults() *CampaignContentRequest`

NewCampaignContentRequestWithDefaults instantiates a new CampaignContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignContentRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignContentRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignContentRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CampaignContentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CampaignContentRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CampaignContentRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CampaignContentRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CampaignContentRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CampaignContentRequest) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CampaignContentRequest) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CampaignContentRequest) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CampaignContentRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCampaignID

`func (o *CampaignContentRequest) GetCampaignID() int32`

GetCampaignID returns the CampaignID field if non-nil, zero value otherwise.

### GetCampaignIDOk

`func (o *CampaignContentRequest) GetCampaignIDOk() (*int32, bool)`

GetCampaignIDOk returns a tuple with the CampaignID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignID

`func (o *CampaignContentRequest) SetCampaignID(v int32)`

SetCampaignID sets CampaignID field to given value.

### HasCampaignID

`func (o *CampaignContentRequest) HasCampaignID() bool`

HasCampaignID returns a boolean if a field has been set.

### GetViews

`func (o *CampaignContentRequest) GetViews() int32`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *CampaignContentRequest) GetViewsOk() (*int32, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *CampaignContentRequest) SetViews(v int32)`

SetViews sets Views field to given value.

### HasViews

`func (o *CampaignContentRequest) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetClicks

`func (o *CampaignContentRequest) GetClicks() int32`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *CampaignContentRequest) GetClicksOk() (*int32, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *CampaignContentRequest) SetClicks(v int32)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *CampaignContentRequest) HasClicks() bool`

HasClicks returns a boolean if a field has been set.

### GetLists

`func (o *CampaignContentRequest) GetLists() []BounceResultsInnerCampaign`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *CampaignContentRequest) GetListsOk() (*[]BounceResultsInnerCampaign, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *CampaignContentRequest) SetLists(v []BounceResultsInnerCampaign)`

SetLists sets Lists field to given value.

### HasLists

`func (o *CampaignContentRequest) HasLists() bool`

HasLists returns a boolean if a field has been set.

### GetStartedAt

`func (o *CampaignContentRequest) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CampaignContentRequest) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CampaignContentRequest) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CampaignContentRequest) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetToSend

`func (o *CampaignContentRequest) GetToSend() int32`

GetToSend returns the ToSend field if non-nil, zero value otherwise.

### GetToSendOk

`func (o *CampaignContentRequest) GetToSendOk() (*int32, bool)`

GetToSendOk returns a tuple with the ToSend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToSend

`func (o *CampaignContentRequest) SetToSend(v int32)`

SetToSend sets ToSend field to given value.

### HasToSend

`func (o *CampaignContentRequest) HasToSend() bool`

HasToSend returns a boolean if a field has been set.

### GetSent

`func (o *CampaignContentRequest) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *CampaignContentRequest) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *CampaignContentRequest) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *CampaignContentRequest) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetUuid

`func (o *CampaignContentRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CampaignContentRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CampaignContentRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CampaignContentRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *CampaignContentRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignContentRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignContentRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CampaignContentRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *CampaignContentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignContentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignContentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignContentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubject

`func (o *CampaignContentRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CampaignContentRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CampaignContentRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CampaignContentRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetFromEmail

`func (o *CampaignContentRequest) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *CampaignContentRequest) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *CampaignContentRequest) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *CampaignContentRequest) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### GetBody

`func (o *CampaignContentRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CampaignContentRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CampaignContentRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CampaignContentRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetSendAt

`func (o *CampaignContentRequest) GetSendAt() string`

GetSendAt returns the SendAt field if non-nil, zero value otherwise.

### GetSendAtOk

`func (o *CampaignContentRequest) GetSendAtOk() (*string, bool)`

GetSendAtOk returns a tuple with the SendAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAt

`func (o *CampaignContentRequest) SetSendAt(v string)`

SetSendAt sets SendAt field to given value.

### HasSendAt

`func (o *CampaignContentRequest) HasSendAt() bool`

HasSendAt returns a boolean if a field has been set.

### GetStatus

`func (o *CampaignContentRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CampaignContentRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CampaignContentRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CampaignContentRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContentType

`func (o *CampaignContentRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CampaignContentRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CampaignContentRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *CampaignContentRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetTags

`func (o *CampaignContentRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CampaignContentRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CampaignContentRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CampaignContentRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplateId

`func (o *CampaignContentRequest) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CampaignContentRequest) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CampaignContentRequest) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *CampaignContentRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetMessenger

`func (o *CampaignContentRequest) GetMessenger() string`

GetMessenger returns the Messenger field if non-nil, zero value otherwise.

### GetMessengerOk

`func (o *CampaignContentRequest) GetMessengerOk() (*string, bool)`

GetMessengerOk returns a tuple with the Messenger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessenger

`func (o *CampaignContentRequest) SetMessenger(v string)`

SetMessenger sets Messenger field to given value.

### HasMessenger

`func (o *CampaignContentRequest) HasMessenger() bool`

HasMessenger returns a boolean if a field has been set.

### GetFrom

`func (o *CampaignContentRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CampaignContentRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CampaignContentRequest) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CampaignContentRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *CampaignContentRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CampaignContentRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CampaignContentRequest) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CampaignContentRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


