/*
Listmonk

The API collection for listmonk

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DashboardCountData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardCountData{}

// DashboardCountData struct for DashboardCountData
type DashboardCountData struct {
	Subscribers *DashboardCountDataSubscribers `json:"subscribers,omitempty"`
	Lists *DashboardCountDataLists `json:"lists,omitempty"`
	Campaigns *DashboardCountDataCampaigns `json:"campaigns,omitempty"`
	Messages *int32 `json:"messages,omitempty"`
}

// NewDashboardCountData instantiates a new DashboardCountData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardCountData() *DashboardCountData {
	this := DashboardCountData{}
	return &this
}

// NewDashboardCountDataWithDefaults instantiates a new DashboardCountData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardCountDataWithDefaults() *DashboardCountData {
	this := DashboardCountData{}
	return &this
}

// GetSubscribers returns the Subscribers field value if set, zero value otherwise.
func (o *DashboardCountData) GetSubscribers() DashboardCountDataSubscribers {
	if o == nil || IsNil(o.Subscribers) {
		var ret DashboardCountDataSubscribers
		return ret
	}
	return *o.Subscribers
}

// GetSubscribersOk returns a tuple with the Subscribers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardCountData) GetSubscribersOk() (*DashboardCountDataSubscribers, bool) {
	if o == nil || IsNil(o.Subscribers) {
		return nil, false
	}
	return o.Subscribers, true
}

// HasSubscribers returns a boolean if a field has been set.
func (o *DashboardCountData) HasSubscribers() bool {
	if o != nil && !IsNil(o.Subscribers) {
		return true
	}

	return false
}

// SetSubscribers gets a reference to the given DashboardCountDataSubscribers and assigns it to the Subscribers field.
func (o *DashboardCountData) SetSubscribers(v DashboardCountDataSubscribers) {
	o.Subscribers = &v
}

// GetLists returns the Lists field value if set, zero value otherwise.
func (o *DashboardCountData) GetLists() DashboardCountDataLists {
	if o == nil || IsNil(o.Lists) {
		var ret DashboardCountDataLists
		return ret
	}
	return *o.Lists
}

// GetListsOk returns a tuple with the Lists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardCountData) GetListsOk() (*DashboardCountDataLists, bool) {
	if o == nil || IsNil(o.Lists) {
		return nil, false
	}
	return o.Lists, true
}

// HasLists returns a boolean if a field has been set.
func (o *DashboardCountData) HasLists() bool {
	if o != nil && !IsNil(o.Lists) {
		return true
	}

	return false
}

// SetLists gets a reference to the given DashboardCountDataLists and assigns it to the Lists field.
func (o *DashboardCountData) SetLists(v DashboardCountDataLists) {
	o.Lists = &v
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise.
func (o *DashboardCountData) GetCampaigns() DashboardCountDataCampaigns {
	if o == nil || IsNil(o.Campaigns) {
		var ret DashboardCountDataCampaigns
		return ret
	}
	return *o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardCountData) GetCampaignsOk() (*DashboardCountDataCampaigns, bool) {
	if o == nil || IsNil(o.Campaigns) {
		return nil, false
	}
	return o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *DashboardCountData) HasCampaigns() bool {
	if o != nil && !IsNil(o.Campaigns) {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given DashboardCountDataCampaigns and assigns it to the Campaigns field.
func (o *DashboardCountData) SetCampaigns(v DashboardCountDataCampaigns) {
	o.Campaigns = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *DashboardCountData) GetMessages() int32 {
	if o == nil || IsNil(o.Messages) {
		var ret int32
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardCountData) GetMessagesOk() (*int32, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *DashboardCountData) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given int32 and assigns it to the Messages field.
func (o *DashboardCountData) SetMessages(v int32) {
	o.Messages = &v
}

func (o DashboardCountData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardCountData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subscribers) {
		toSerialize["subscribers"] = o.Subscribers
	}
	if !IsNil(o.Lists) {
		toSerialize["lists"] = o.Lists
	}
	if !IsNil(o.Campaigns) {
		toSerialize["campaigns"] = o.Campaigns
	}
	if !IsNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	return toSerialize, nil
}

type NullableDashboardCountData struct {
	value *DashboardCountData
	isSet bool
}

func (v NullableDashboardCountData) Get() *DashboardCountData {
	return v.value
}

func (v *NullableDashboardCountData) Set(val *DashboardCountData) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardCountData) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardCountData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardCountData(val *DashboardCountData) *NullableDashboardCountData {
	return &NullableDashboardCountData{value: val, isSet: true}
}

func (v NullableDashboardCountData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardCountData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

