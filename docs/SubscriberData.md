# SubscriberData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to [**[]SubscriberProfile**](SubscriberProfile.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]Subscriptions**](Subscriptions.md) |  | [optional] 
**CampaignViews** | Pointer to **[]map[string]interface{}** |  | [optional] 
**LinkClicks** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewSubscriberData

`func NewSubscriberData() *SubscriberData`

NewSubscriberData instantiates a new SubscriberData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriberDataWithDefaults

`func NewSubscriberDataWithDefaults() *SubscriberData`

NewSubscriberDataWithDefaults instantiates a new SubscriberData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SubscriberData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SubscriberData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SubscriberData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SubscriberData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetProfile

`func (o *SubscriberData) GetProfile() []SubscriberProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SubscriberData) GetProfileOk() (*[]SubscriberProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SubscriberData) SetProfile(v []SubscriberProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SubscriberData) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSubscriptions

`func (o *SubscriberData) GetSubscriptions() []Subscriptions`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *SubscriberData) GetSubscriptionsOk() (*[]Subscriptions, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *SubscriberData) SetSubscriptions(v []Subscriptions)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *SubscriberData) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetCampaignViews

`func (o *SubscriberData) GetCampaignViews() []map[string]interface{}`

GetCampaignViews returns the CampaignViews field if non-nil, zero value otherwise.

### GetCampaignViewsOk

`func (o *SubscriberData) GetCampaignViewsOk() (*[]map[string]interface{}, bool)`

GetCampaignViewsOk returns a tuple with the CampaignViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignViews

`func (o *SubscriberData) SetCampaignViews(v []map[string]interface{})`

SetCampaignViews sets CampaignViews field to given value.

### HasCampaignViews

`func (o *SubscriberData) HasCampaignViews() bool`

HasCampaignViews returns a boolean if a field has been set.

### GetLinkClicks

`func (o *SubscriberData) GetLinkClicks() []map[string]interface{}`

GetLinkClicks returns the LinkClicks field if non-nil, zero value otherwise.

### GetLinkClicksOk

`func (o *SubscriberData) GetLinkClicksOk() (*[]map[string]interface{}, bool)`

GetLinkClicksOk returns a tuple with the LinkClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkClicks

`func (o *SubscriberData) SetLinkClicks(v []map[string]interface{})`

SetLinkClicks sets LinkClicks field to given value.

### HasLinkClicks

`func (o *SubscriberData) HasLinkClicks() bool`

HasLinkClicks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


