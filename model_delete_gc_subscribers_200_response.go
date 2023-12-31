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

// checks if the DeleteGCSubscribers200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteGCSubscribers200Response{}

// DeleteGCSubscribers200Response struct for DeleteGCSubscribers200Response
type DeleteGCSubscribers200Response struct {
	Data *DeleteGCSubscribers200ResponseData `json:"data,omitempty"`
}

// NewDeleteGCSubscribers200Response instantiates a new DeleteGCSubscribers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteGCSubscribers200Response() *DeleteGCSubscribers200Response {
	this := DeleteGCSubscribers200Response{}
	return &this
}

// NewDeleteGCSubscribers200ResponseWithDefaults instantiates a new DeleteGCSubscribers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteGCSubscribers200ResponseWithDefaults() *DeleteGCSubscribers200Response {
	this := DeleteGCSubscribers200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DeleteGCSubscribers200Response) GetData() DeleteGCSubscribers200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret DeleteGCSubscribers200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteGCSubscribers200Response) GetDataOk() (*DeleteGCSubscribers200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DeleteGCSubscribers200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DeleteGCSubscribers200ResponseData and assigns it to the Data field.
func (o *DeleteGCSubscribers200Response) SetData(v DeleteGCSubscribers200ResponseData) {
	o.Data = &v
}

func (o DeleteGCSubscribers200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteGCSubscribers200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDeleteGCSubscribers200Response struct {
	value *DeleteGCSubscribers200Response
	isSet bool
}

func (v NullableDeleteGCSubscribers200Response) Get() *DeleteGCSubscribers200Response {
	return v.value
}

func (v *NullableDeleteGCSubscribers200Response) Set(val *DeleteGCSubscribers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteGCSubscribers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteGCSubscribers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteGCSubscribers200Response(val *DeleteGCSubscribers200Response) *NullableDeleteGCSubscribers200Response {
	return &NullableDeleteGCSubscribers200Response{value: val, isSet: true}
}

func (v NullableDeleteGCSubscribers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteGCSubscribers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


