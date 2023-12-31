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

// checks if the GetSubscribers200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSubscribers200Response{}

// GetSubscribers200Response struct for GetSubscribers200Response
type GetSubscribers200Response struct {
	Data *GetSubscribers200ResponseData `json:"data,omitempty"`
}

// NewGetSubscribers200Response instantiates a new GetSubscribers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubscribers200Response() *GetSubscribers200Response {
	this := GetSubscribers200Response{}
	return &this
}

// NewGetSubscribers200ResponseWithDefaults instantiates a new GetSubscribers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubscribers200ResponseWithDefaults() *GetSubscribers200Response {
	this := GetSubscribers200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetSubscribers200Response) GetData() GetSubscribers200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GetSubscribers200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubscribers200Response) GetDataOk() (*GetSubscribers200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetSubscribers200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetSubscribers200ResponseData and assigns it to the Data field.
func (o *GetSubscribers200Response) SetData(v GetSubscribers200ResponseData) {
	o.Data = &v
}

func (o GetSubscribers200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSubscribers200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetSubscribers200Response struct {
	value *GetSubscribers200Response
	isSet bool
}

func (v NullableGetSubscribers200Response) Get() *GetSubscribers200Response {
	return v.value
}

func (v *NullableGetSubscribers200Response) Set(val *GetSubscribers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubscribers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubscribers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSubscribers200Response(val *GetSubscribers200Response) *NullableGetSubscribers200Response {
	return &NullableGetSubscribers200Response{value: val, isSet: true}
}

func (v NullableGetSubscribers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSubscribers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


