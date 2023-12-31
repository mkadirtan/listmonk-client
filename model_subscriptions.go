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

// checks if the Subscriptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Subscriptions{}

// Subscriptions struct for Subscriptions
type Subscriptions struct {
	SubscriptionStatus *string `json:"subscription_status,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
}

// NewSubscriptions instantiates a new Subscriptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptions() *Subscriptions {
	this := Subscriptions{}
	return &this
}

// NewSubscriptionsWithDefaults instantiates a new Subscriptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionsWithDefaults() *Subscriptions {
	this := Subscriptions{}
	return &this
}

// GetSubscriptionStatus returns the SubscriptionStatus field value if set, zero value otherwise.
func (o *Subscriptions) GetSubscriptionStatus() string {
	if o == nil || IsNil(o.SubscriptionStatus) {
		var ret string
		return ret
	}
	return *o.SubscriptionStatus
}

// GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscriptions) GetSubscriptionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionStatus) {
		return nil, false
	}
	return o.SubscriptionStatus, true
}

// HasSubscriptionStatus returns a boolean if a field has been set.
func (o *Subscriptions) HasSubscriptionStatus() bool {
	if o != nil && !IsNil(o.SubscriptionStatus) {
		return true
	}

	return false
}

// SetSubscriptionStatus gets a reference to the given string and assigns it to the SubscriptionStatus field.
func (o *Subscriptions) SetSubscriptionStatus(v string) {
	o.SubscriptionStatus = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Subscriptions) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscriptions) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Subscriptions) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Subscriptions) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Subscriptions) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscriptions) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Subscriptions) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Subscriptions) SetType(v string) {
	o.Type = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Subscriptions) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscriptions) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Subscriptions) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Subscriptions) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o Subscriptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Subscriptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionStatus) {
		toSerialize["subscription_status"] = o.SubscriptionStatus
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableSubscriptions struct {
	value *Subscriptions
	isSet bool
}

func (v NullableSubscriptions) Get() *Subscriptions {
	return v.value
}

func (v *NullableSubscriptions) Set(val *Subscriptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptions(val *Subscriptions) *NullableSubscriptions {
	return &NullableSubscriptions{value: val, isSet: true}
}

func (v NullableSubscriptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


