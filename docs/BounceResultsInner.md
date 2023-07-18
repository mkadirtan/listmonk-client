# BounceResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**SubscriberUuid** | Pointer to **string** |  | [optional] 
**SubscriberId** | Pointer to **int32** |  | [optional] 
**Campaign** | Pointer to [**BounceResultsInnerCampaign**](BounceResultsInnerCampaign.md) |  | [optional] 
**CampaignUuid** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewBounceResultsInner

`func NewBounceResultsInner() *BounceResultsInner`

NewBounceResultsInner instantiates a new BounceResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBounceResultsInnerWithDefaults

`func NewBounceResultsInnerWithDefaults() *BounceResultsInner`

NewBounceResultsInnerWithDefaults instantiates a new BounceResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BounceResultsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BounceResultsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BounceResultsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BounceResultsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *BounceResultsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BounceResultsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BounceResultsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BounceResultsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSource

`func (o *BounceResultsInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BounceResultsInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BounceResultsInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BounceResultsInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetMeta

`func (o *BounceResultsInner) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BounceResultsInner) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BounceResultsInner) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BounceResultsInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BounceResultsInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BounceResultsInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BounceResultsInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BounceResultsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *BounceResultsInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BounceResultsInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BounceResultsInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BounceResultsInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSubscriberUuid

`func (o *BounceResultsInner) GetSubscriberUuid() string`

GetSubscriberUuid returns the SubscriberUuid field if non-nil, zero value otherwise.

### GetSubscriberUuidOk

`func (o *BounceResultsInner) GetSubscriberUuidOk() (*string, bool)`

GetSubscriberUuidOk returns a tuple with the SubscriberUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberUuid

`func (o *BounceResultsInner) SetSubscriberUuid(v string)`

SetSubscriberUuid sets SubscriberUuid field to given value.

### HasSubscriberUuid

`func (o *BounceResultsInner) HasSubscriberUuid() bool`

HasSubscriberUuid returns a boolean if a field has been set.

### GetSubscriberId

`func (o *BounceResultsInner) GetSubscriberId() int32`

GetSubscriberId returns the SubscriberId field if non-nil, zero value otherwise.

### GetSubscriberIdOk

`func (o *BounceResultsInner) GetSubscriberIdOk() (*int32, bool)`

GetSubscriberIdOk returns a tuple with the SubscriberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberId

`func (o *BounceResultsInner) SetSubscriberId(v int32)`

SetSubscriberId sets SubscriberId field to given value.

### HasSubscriberId

`func (o *BounceResultsInner) HasSubscriberId() bool`

HasSubscriberId returns a boolean if a field has been set.

### GetCampaign

`func (o *BounceResultsInner) GetCampaign() BounceResultsInnerCampaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *BounceResultsInner) GetCampaignOk() (*BounceResultsInnerCampaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *BounceResultsInner) SetCampaign(v BounceResultsInnerCampaign)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *BounceResultsInner) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetCampaignUuid

`func (o *BounceResultsInner) GetCampaignUuid() string`

GetCampaignUuid returns the CampaignUuid field if non-nil, zero value otherwise.

### GetCampaignUuidOk

`func (o *BounceResultsInner) GetCampaignUuidOk() (*string, bool)`

GetCampaignUuidOk returns a tuple with the CampaignUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignUuid

`func (o *BounceResultsInner) SetCampaignUuid(v string)`

SetCampaignUuid sets CampaignUuid field to given value.

### HasCampaignUuid

`func (o *BounceResultsInner) HasCampaignUuid() bool`

HasCampaignUuid returns a boolean if a field has been set.

### GetTotal

`func (o *BounceResultsInner) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BounceResultsInner) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BounceResultsInner) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BounceResultsInner) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


