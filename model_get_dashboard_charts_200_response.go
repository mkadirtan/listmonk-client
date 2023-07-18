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

// checks if the GetDashboardCharts200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDashboardCharts200Response{}

// GetDashboardCharts200Response struct for GetDashboardCharts200Response
type GetDashboardCharts200Response struct {
	Data *DashboardChart `json:"data,omitempty"`
}

// NewGetDashboardCharts200Response instantiates a new GetDashboardCharts200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDashboardCharts200Response() *GetDashboardCharts200Response {
	this := GetDashboardCharts200Response{}
	return &this
}

// NewGetDashboardCharts200ResponseWithDefaults instantiates a new GetDashboardCharts200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDashboardCharts200ResponseWithDefaults() *GetDashboardCharts200Response {
	this := GetDashboardCharts200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetDashboardCharts200Response) GetData() DashboardChart {
	if o == nil || IsNil(o.Data) {
		var ret DashboardChart
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDashboardCharts200Response) GetDataOk() (*DashboardChart, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetDashboardCharts200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DashboardChart and assigns it to the Data field.
func (o *GetDashboardCharts200Response) SetData(v DashboardChart) {
	o.Data = &v
}

func (o GetDashboardCharts200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDashboardCharts200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetDashboardCharts200Response struct {
	value *GetDashboardCharts200Response
	isSet bool
}

func (v NullableGetDashboardCharts200Response) Get() *GetDashboardCharts200Response {
	return v.value
}

func (v *NullableGetDashboardCharts200Response) Set(val *GetDashboardCharts200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDashboardCharts200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDashboardCharts200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDashboardCharts200Response(val *GetDashboardCharts200Response) *NullableGetDashboardCharts200Response {
	return &NullableGetDashboardCharts200Response{value: val, isSet: true}
}

func (v NullableGetDashboardCharts200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDashboardCharts200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


