# NewSubscriber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Lists** | Pointer to **[]int32** |  | [optional] 
**ListUuids** | Pointer to **[]string** |  | [optional] 
**PreconfirmSubscriptions** | Pointer to **bool** |  | [optional] 
**Attribs** | Pointer to [**NewSubscriberAttribs**](NewSubscriberAttribs.md) |  | [optional] 

## Methods

### NewNewSubscriber

`func NewNewSubscriber() *NewSubscriber`

NewNewSubscriber instantiates a new NewSubscriber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewSubscriberWithDefaults

`func NewNewSubscriberWithDefaults() *NewSubscriber`

NewNewSubscriberWithDefaults instantiates a new NewSubscriber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *NewSubscriber) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NewSubscriber) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NewSubscriber) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *NewSubscriber) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *NewSubscriber) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewSubscriber) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewSubscriber) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NewSubscriber) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *NewSubscriber) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NewSubscriber) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NewSubscriber) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NewSubscriber) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLists

`func (o *NewSubscriber) GetLists() []int32`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *NewSubscriber) GetListsOk() (*[]int32, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *NewSubscriber) SetLists(v []int32)`

SetLists sets Lists field to given value.

### HasLists

`func (o *NewSubscriber) HasLists() bool`

HasLists returns a boolean if a field has been set.

### GetListUuids

`func (o *NewSubscriber) GetListUuids() []string`

GetListUuids returns the ListUuids field if non-nil, zero value otherwise.

### GetListUuidsOk

`func (o *NewSubscriber) GetListUuidsOk() (*[]string, bool)`

GetListUuidsOk returns a tuple with the ListUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListUuids

`func (o *NewSubscriber) SetListUuids(v []string)`

SetListUuids sets ListUuids field to given value.

### HasListUuids

`func (o *NewSubscriber) HasListUuids() bool`

HasListUuids returns a boolean if a field has been set.

### GetPreconfirmSubscriptions

`func (o *NewSubscriber) GetPreconfirmSubscriptions() bool`

GetPreconfirmSubscriptions returns the PreconfirmSubscriptions field if non-nil, zero value otherwise.

### GetPreconfirmSubscriptionsOk

`func (o *NewSubscriber) GetPreconfirmSubscriptionsOk() (*bool, bool)`

GetPreconfirmSubscriptionsOk returns a tuple with the PreconfirmSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconfirmSubscriptions

`func (o *NewSubscriber) SetPreconfirmSubscriptions(v bool)`

SetPreconfirmSubscriptions sets PreconfirmSubscriptions field to given value.

### HasPreconfirmSubscriptions

`func (o *NewSubscriber) HasPreconfirmSubscriptions() bool`

HasPreconfirmSubscriptions returns a boolean if a field has been set.

### GetAttribs

`func (o *NewSubscriber) GetAttribs() NewSubscriberAttribs`

GetAttribs returns the Attribs field if non-nil, zero value otherwise.

### GetAttribsOk

`func (o *NewSubscriber) GetAttribsOk() (*NewSubscriberAttribs, bool)`

GetAttribsOk returns a tuple with the Attribs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribs

`func (o *NewSubscriber) SetAttribs(v NewSubscriberAttribs)`

SetAttribs sets Attribs field to given value.

### HasAttribs

`func (o *NewSubscriber) HasAttribs() bool`

HasAttribs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


