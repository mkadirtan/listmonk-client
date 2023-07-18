# Campaign

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

## Methods

### NewCampaign

`func NewCampaign() *Campaign`

NewCampaign instantiates a new Campaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignWithDefaults

`func NewCampaignWithDefaults() *Campaign`

NewCampaignWithDefaults instantiates a new Campaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Campaign) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Campaign) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Campaign) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Campaign) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Campaign) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Campaign) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Campaign) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Campaign) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Campaign) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Campaign) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Campaign) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Campaign) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCampaignID

`func (o *Campaign) GetCampaignID() int32`

GetCampaignID returns the CampaignID field if non-nil, zero value otherwise.

### GetCampaignIDOk

`func (o *Campaign) GetCampaignIDOk() (*int32, bool)`

GetCampaignIDOk returns a tuple with the CampaignID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignID

`func (o *Campaign) SetCampaignID(v int32)`

SetCampaignID sets CampaignID field to given value.

### HasCampaignID

`func (o *Campaign) HasCampaignID() bool`

HasCampaignID returns a boolean if a field has been set.

### GetViews

`func (o *Campaign) GetViews() int32`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *Campaign) GetViewsOk() (*int32, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *Campaign) SetViews(v int32)`

SetViews sets Views field to given value.

### HasViews

`func (o *Campaign) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetClicks

`func (o *Campaign) GetClicks() int32`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *Campaign) GetClicksOk() (*int32, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *Campaign) SetClicks(v int32)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *Campaign) HasClicks() bool`

HasClicks returns a boolean if a field has been set.

### GetLists

`func (o *Campaign) GetLists() []BounceResultsInnerCampaign`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *Campaign) GetListsOk() (*[]BounceResultsInnerCampaign, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *Campaign) SetLists(v []BounceResultsInnerCampaign)`

SetLists sets Lists field to given value.

### HasLists

`func (o *Campaign) HasLists() bool`

HasLists returns a boolean if a field has been set.

### GetStartedAt

`func (o *Campaign) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Campaign) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Campaign) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Campaign) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetToSend

`func (o *Campaign) GetToSend() int32`

GetToSend returns the ToSend field if non-nil, zero value otherwise.

### GetToSendOk

`func (o *Campaign) GetToSendOk() (*int32, bool)`

GetToSendOk returns a tuple with the ToSend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToSend

`func (o *Campaign) SetToSend(v int32)`

SetToSend sets ToSend field to given value.

### HasToSend

`func (o *Campaign) HasToSend() bool`

HasToSend returns a boolean if a field has been set.

### GetSent

`func (o *Campaign) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *Campaign) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *Campaign) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *Campaign) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetUuid

`func (o *Campaign) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Campaign) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Campaign) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Campaign) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *Campaign) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Campaign) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Campaign) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Campaign) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Campaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Campaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Campaign) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Campaign) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubject

`func (o *Campaign) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Campaign) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Campaign) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Campaign) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetFromEmail

`func (o *Campaign) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *Campaign) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *Campaign) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *Campaign) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### GetBody

`func (o *Campaign) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Campaign) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Campaign) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *Campaign) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetSendAt

`func (o *Campaign) GetSendAt() string`

GetSendAt returns the SendAt field if non-nil, zero value otherwise.

### GetSendAtOk

`func (o *Campaign) GetSendAtOk() (*string, bool)`

GetSendAtOk returns a tuple with the SendAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAt

`func (o *Campaign) SetSendAt(v string)`

SetSendAt sets SendAt field to given value.

### HasSendAt

`func (o *Campaign) HasSendAt() bool`

HasSendAt returns a boolean if a field has been set.

### GetStatus

`func (o *Campaign) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Campaign) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Campaign) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Campaign) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContentType

`func (o *Campaign) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *Campaign) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *Campaign) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *Campaign) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetTags

`func (o *Campaign) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Campaign) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Campaign) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Campaign) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplateId

`func (o *Campaign) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *Campaign) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *Campaign) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *Campaign) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetMessenger

`func (o *Campaign) GetMessenger() string`

GetMessenger returns the Messenger field if non-nil, zero value otherwise.

### GetMessengerOk

`func (o *Campaign) GetMessengerOk() (*string, bool)`

GetMessengerOk returns a tuple with the Messenger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessenger

`func (o *Campaign) SetMessenger(v string)`

SetMessenger sets Messenger field to given value.

### HasMessenger

`func (o *Campaign) HasMessenger() bool`

HasMessenger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


