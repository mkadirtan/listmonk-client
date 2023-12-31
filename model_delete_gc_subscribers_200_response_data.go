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

// checks if the DeleteGCSubscribers200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteGCSubscribers200ResponseData{}

// DeleteGCSubscribers200ResponseData struct for DeleteGCSubscribers200ResponseData
type DeleteGCSubscribers200ResponseData struct {
	Count *int32 `json:"count,omitempty"`
}

// NewDeleteGCSubscribers200ResponseData instantiates a new DeleteGCSubscribers200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteGCSubscribers200ResponseData() *DeleteGCSubscribers200ResponseData {
	this := DeleteGCSubscribers200ResponseData{}
	return &this
}

// NewDeleteGCSubscribers200ResponseDataWithDefaults instantiates a new DeleteGCSubscribers200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteGCSubscribers200ResponseDataWithDefaults() *DeleteGCSubscribers200ResponseData {
	this := DeleteGCSubscribers200ResponseData{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *DeleteGCSubscribers200ResponseData) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteGCSubscribers200ResponseData) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *DeleteGCSubscribers200ResponseData) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *DeleteGCSubscribers200ResponseData) SetCount(v int32) {
	o.Count = &v
}

func (o DeleteGCSubscribers200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteGCSubscribers200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableDeleteGCSubscribers200ResponseData struct {
	value *DeleteGCSubscribers200ResponseData
	isSet bool
}

func (v NullableDeleteGCSubscribers200ResponseData) Get() *DeleteGCSubscribers200ResponseData {
	return v.value
}

func (v *NullableDeleteGCSubscribers200ResponseData) Set(val *DeleteGCSubscribers200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteGCSubscribers200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteGCSubscribers200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteGCSubscribers200ResponseData(val *DeleteGCSubscribers200ResponseData) *NullableDeleteGCSubscribers200ResponseData {
	return &NullableDeleteGCSubscribers200ResponseData{value: val, isSet: true}
}

func (v NullableDeleteGCSubscribers200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteGCSubscribers200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


