# Subscriber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Attribs** | Pointer to [**SubscriberProfileAttribs**](SubscriberProfileAttribs.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Lists** | Pointer to [**[]SubscriberListsInner**](SubscriberListsInner.md) |  | [optional] 

## Methods

### NewSubscriber

`func NewSubscriber() *Subscriber`

NewSubscriber instantiates a new Subscriber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriberWithDefaults

`func NewSubscriberWithDefaults() *Subscriber`

NewSubscriberWithDefaults instantiates a new Subscriber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subscriber) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscriber) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscriber) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Subscriber) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Subscriber) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Subscriber) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Subscriber) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Subscriber) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Subscriber) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Subscriber) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Subscriber) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Subscriber) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUuid

`func (o *Subscriber) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Subscriber) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Subscriber) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Subscriber) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetEmail

`func (o *Subscriber) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Subscriber) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Subscriber) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Subscriber) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *Subscriber) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Subscriber) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Subscriber) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Subscriber) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAttribs

`func (o *Subscriber) GetAttribs() SubscriberProfileAttribs`

GetAttribs returns the Attribs field if non-nil, zero value otherwise.

### GetAttribsOk

`func (o *Subscriber) GetAttribsOk() (*SubscriberProfileAttribs, bool)`

GetAttribsOk returns a tuple with the Attribs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribs

`func (o *Subscriber) SetAttribs(v SubscriberProfileAttribs)`

SetAttribs sets Attribs field to given value.

### HasAttribs

`func (o *Subscriber) HasAttribs() bool`

HasAttribs returns a boolean if a field has been set.

### GetStatus

`func (o *Subscriber) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Subscriber) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Subscriber) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Subscriber) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLists

`func (o *Subscriber) GetLists() []SubscriberListsInner`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *Subscriber) GetListsOk() (*[]SubscriberListsInner, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *Subscriber) SetLists(v []SubscriberListsInner)`

SetLists sets Lists field to given value.

### HasLists

`func (o *Subscriber) HasLists() bool`

HasLists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


