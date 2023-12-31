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

// checks if the ImportStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportStatus{}

// ImportStatus struct for ImportStatus
type ImportStatus struct {
	Data *ImportStatusData `json:"data,omitempty"`
}

// NewImportStatus instantiates a new ImportStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportStatus() *ImportStatus {
	this := ImportStatus{}
	return &this
}

// NewImportStatusWithDefaults instantiates a new ImportStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportStatusWithDefaults() *ImportStatus {
	this := ImportStatus{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ImportStatus) GetData() ImportStatusData {
	if o == nil || IsNil(o.Data) {
		var ret ImportStatusData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportStatus) GetDataOk() (*ImportStatusData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ImportStatus) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ImportStatusData and assigns it to the Data field.
func (o *ImportStatus) SetData(v ImportStatusData) {
	o.Data = &v
}

func (o ImportStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableImportStatus struct {
	value *ImportStatus
	isSet bool
}

func (v NullableImportStatus) Get() *ImportStatus {
	return v.value
}

func (v *NullableImportStatus) Set(val *ImportStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableImportStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableImportStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportStatus(val *ImportStatus) *NullableImportStatus {
	return &NullableImportStatus{value: val, isSet: true}
}

func (v NullableImportStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


