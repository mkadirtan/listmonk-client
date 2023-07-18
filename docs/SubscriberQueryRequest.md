# SubscriberQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** |  | [optional] 
**Ids** | Pointer to **[]int32** | The ids of the subscribers to be modified. | [optional] 
**Action** | Pointer to **string** | Whether to add, remove, or unsubscribe the users. | [optional] 
**TargetListIds** | Pointer to **[]int32** | The ids of the lists to be modified. | [optional] 
**Status** | Pointer to **string** | confirmed, unconfirmed, or unsubscribed status. | [optional] 

## Methods

### NewSubscriberQueryRequest

`func NewSubscriberQueryRequest() *SubscriberQueryRequest`

NewSubscriberQueryRequest instantiates a new SubscriberQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriberQueryRequestWithDefaults

`func NewSubscriberQueryRequestWithDefaults() *SubscriberQueryRequest`

NewSubscriberQueryRequestWithDefaults instantiates a new SubscriberQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *SubscriberQueryRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SubscriberQueryRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SubscriberQueryRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SubscriberQueryRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetIds

`func (o *SubscriberQueryRequest) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *SubscriberQueryRequest) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *SubscriberQueryRequest) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *SubscriberQueryRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetAction

`func (o *SubscriberQueryRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SubscriberQueryRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SubscriberQueryRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SubscriberQueryRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetTargetListIds

`func (o *SubscriberQueryRequest) GetTargetListIds() []int32`

GetTargetListIds returns the TargetListIds field if non-nil, zero value otherwise.

### GetTargetListIdsOk

`func (o *SubscriberQueryRequest) GetTargetListIdsOk() (*[]int32, bool)`

GetTargetListIdsOk returns a tuple with the TargetListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetListIds

`func (o *SubscriberQueryRequest) SetTargetListIds(v []int32)`

SetTargetListIds sets TargetListIds field to given value.

### HasTargetListIds

`func (o *SubscriberQueryRequest) HasTargetListIds() bool`

HasTargetListIds returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriberQueryRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriberQueryRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriberQueryRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriberQueryRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


