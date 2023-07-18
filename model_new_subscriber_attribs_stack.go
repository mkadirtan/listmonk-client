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

// checks if the NewSubscriberAttribsStack type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewSubscriberAttribsStack{}

// NewSubscriberAttribsStack struct for NewSubscriberAttribsStack
type NewSubscriberAttribsStack struct {
	Languages []string `json:"languages,omitempty"`
}

// NewNewSubscriberAttribsStack instantiates a new NewSubscriberAttribsStack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewSubscriberAttribsStack() *NewSubscriberAttribsStack {
	this := NewSubscriberAttribsStack{}
	return &this
}

// NewNewSubscriberAttribsStackWithDefaults instantiates a new NewSubscriberAttribsStack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewSubscriberAttribsStackWithDefaults() *NewSubscriberAttribsStack {
	this := NewSubscriberAttribsStack{}
	return &this
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *NewSubscriberAttribsStack) GetLanguages() []string {
	if o == nil || IsNil(o.Languages) {
		var ret []string
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSubscriberAttribsStack) GetLanguagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *NewSubscriberAttribsStack) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []string and assigns it to the Languages field.
func (o *NewSubscriberAttribsStack) SetLanguages(v []string) {
	o.Languages = v
}

func (o NewSubscriberAttribsStack) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewSubscriberAttribsStack) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Languages) {
		toSerialize["languages"] = o.Languages
	}
	return toSerialize, nil
}

type NullableNewSubscriberAttribsStack struct {
	value *NewSubscriberAttribsStack
	isSet bool
}

func (v NullableNewSubscriberAttribsStack) Get() *NewSubscriberAttribsStack {
	return v.value
}

func (v *NullableNewSubscriberAttribsStack) Set(val *NewSubscriberAttribsStack) {
	v.value = val
	v.isSet = true
}

func (v NullableNewSubscriberAttribsStack) IsSet() bool {
	return v.isSet
}

func (v *NullableNewSubscriberAttribsStack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewSubscriberAttribsStack(val *NewSubscriberAttribsStack) *NullableNewSubscriberAttribsStack {
	return &NullableNewSubscriberAttribsStack{value: val, isSet: true}
}

func (v NullableNewSubscriberAttribsStack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewSubscriberAttribsStack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

