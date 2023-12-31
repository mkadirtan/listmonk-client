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

// checks if the GetServerConfig200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetServerConfig200Response{}

// GetServerConfig200Response struct for GetServerConfig200Response
type GetServerConfig200Response struct {
	Data *ServerConfig `json:"data,omitempty"`
}

// NewGetServerConfig200Response instantiates a new GetServerConfig200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServerConfig200Response() *GetServerConfig200Response {
	this := GetServerConfig200Response{}
	return &this
}

// NewGetServerConfig200ResponseWithDefaults instantiates a new GetServerConfig200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServerConfig200ResponseWithDefaults() *GetServerConfig200Response {
	this := GetServerConfig200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetServerConfig200Response) GetData() ServerConfig {
	if o == nil || IsNil(o.Data) {
		var ret ServerConfig
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServerConfig200Response) GetDataOk() (*ServerConfig, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetServerConfig200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ServerConfig and assigns it to the Data field.
func (o *GetServerConfig200Response) SetData(v ServerConfig) {
	o.Data = &v
}

func (o GetServerConfig200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetServerConfig200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetServerConfig200Response struct {
	value *GetServerConfig200Response
	isSet bool
}

func (v NullableGetServerConfig200Response) Get() *GetServerConfig200Response {
	return v.value
}

func (v *NullableGetServerConfig200Response) Set(val *GetServerConfig200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServerConfig200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServerConfig200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServerConfig200Response(val *GetServerConfig200Response) *NullableGetServerConfig200Response {
	return &NullableGetServerConfig200Response{value: val, isSet: true}
}

func (v NullableGetServerConfig200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServerConfig200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


