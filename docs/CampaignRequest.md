# CampaignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Lists** | Pointer to **[]float32** |  | [optional] 
**FromEmail** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Messenger** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**SendLater** | Pointer to **bool** |  | [optional] 
**SendAt** | Pointer to [**CampaignRequestSendAt**](CampaignRequestSendAt.md) |  | [optional] 

## Methods

### NewCampaignRequest

`func NewCampaignRequest() *CampaignRequest`

NewCampaignRequest instantiates a new CampaignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignRequestWithDefaults

`func NewCampaignRequestWithDefaults() *CampaignRequest`

NewCampaignRequestWithDefaults instantiates a new CampaignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubject

`func (o *CampaignRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CampaignRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CampaignRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CampaignRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetLists

`func (o *CampaignRequest) GetLists() []float32`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *CampaignRequest) GetListsOk() (*[]float32, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *CampaignRequest) SetLists(v []float32)`

SetLists sets Lists field to given value.

### HasLists

`func (o *CampaignRequest) HasLists() bool`

HasLists returns a boolean if a field has been set.

### GetFromEmail

`func (o *CampaignRequest) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *CampaignRequest) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *CampaignRequest) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *CampaignRequest) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### GetContentType

`func (o *CampaignRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CampaignRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CampaignRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *CampaignRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetMessenger

`func (o *CampaignRequest) GetMessenger() string`

GetMessenger returns the Messenger field if non-nil, zero value otherwise.

### GetMessengerOk

`func (o *CampaignRequest) GetMessengerOk() (*string, bool)`

GetMessengerOk returns a tuple with the Messenger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessenger

`func (o *CampaignRequest) SetMessenger(v string)`

SetMessenger sets Messenger field to given value.

### HasMessenger

`func (o *CampaignRequest) HasMessenger() bool`

HasMessenger returns a boolean if a field has been set.

### GetType

`func (o *CampaignRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CampaignRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTags

`func (o *CampaignRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CampaignRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CampaignRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CampaignRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSendLater

`func (o *CampaignRequest) GetSendLater() bool`

GetSendLater returns the SendLater field if non-nil, zero value otherwise.

### GetSendLaterOk

`func (o *CampaignRequest) GetSendLaterOk() (*bool, bool)`

GetSendLaterOk returns a tuple with the SendLater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendLater

`func (o *CampaignRequest) SetSendLater(v bool)`

SetSendLater sets SendLater field to given value.

### HasSendLater

`func (o *CampaignRequest) HasSendLater() bool`

HasSendLater returns a boolean if a field has been set.

### GetSendAt

`func (o *CampaignRequest) GetSendAt() CampaignRequestSendAt`

GetSendAt returns the SendAt field if non-nil, zero value otherwise.

### GetSendAtOk

`func (o *CampaignRequest) GetSendAtOk() (*CampaignRequestSendAt, bool)`

GetSendAtOk returns a tuple with the SendAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAt

`func (o *CampaignRequest) SetSendAt(v CampaignRequestSendAt)`

SetSendAt sets SendAt field to given value.

### HasSendAt

`func (o *CampaignRequest) HasSendAt() bool`

HasSendAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


