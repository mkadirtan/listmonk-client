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

// checks if the GetLogs200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLogs200Response{}

// GetLogs200Response struct for GetLogs200Response
type GetLogs200Response struct {
	Data []string `json:"data,omitempty"`
}

// NewGetLogs200Response instantiates a new GetLogs200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLogs200Response() *GetLogs200Response {
	this := GetLogs200Response{}
	return &this
}

// NewGetLogs200ResponseWithDefaults instantiates a new GetLogs200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLogs200ResponseWithDefaults() *GetLogs200Response {
	this := GetLogs200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetLogs200Response) GetData() []string {
	if o == nil || IsNil(o.Data) {
		var ret []string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogs200Response) GetDataOk() ([]string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetLogs200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *GetLogs200Response) SetData(v []string) {
	o.Data = v
}

func (o GetLogs200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLogs200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetLogs200Response struct {
	value *GetLogs200Response
	isSet bool
}

func (v NullableGetLogs200Response) Get() *GetLogs200Response {
	return v.value
}

func (v *NullableGetLogs200Response) Set(val *GetLogs200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLogs200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLogs200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLogs200Response(val *GetLogs200Response) *NullableGetLogs200Response {
	return &NullableGetLogs200Response{value: val, isSet: true}
}

func (v NullableGetLogs200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLogs200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


