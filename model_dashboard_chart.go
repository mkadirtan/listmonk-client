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

// checks if the DashboardChart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardChart{}

// DashboardChart struct for DashboardChart
type DashboardChart struct {
	LinkClicks []DashboardChartLinkClicksInner `json:"link_clicks,omitempty"`
	CampaignViews []DashboardChartLinkClicksInner `json:"campaign_views,omitempty"`
}

// NewDashboardChart instantiates a new DashboardChart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardChart() *DashboardChart {
	this := DashboardChart{}
	return &this
}

// NewDashboardChartWithDefaults instantiates a new DashboardChart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardChartWithDefaults() *DashboardChart {
	this := DashboardChart{}
	return &this
}

// GetLinkClicks returns the LinkClicks field value if set, zero value otherwise.
func (o *DashboardChart) GetLinkClicks() []DashboardChartLinkClicksInner {
	if o == nil || IsNil(o.LinkClicks) {
		var ret []DashboardChartLinkClicksInner
		return ret
	}
	return o.LinkClicks
}

// GetLinkClicksOk returns a tuple with the LinkClicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardChart) GetLinkClicksOk() ([]DashboardChartLinkClicksInner, bool) {
	if o == nil || IsNil(o.LinkClicks) {
		return nil, false
	}
	return o.LinkClicks, true
}

// HasLinkClicks returns a boolean if a field has been set.
func (o *DashboardChart) HasLinkClicks() bool {
	if o != nil && !IsNil(o.LinkClicks) {
		return true
	}

	return false
}

// SetLinkClicks gets a reference to the given []DashboardChartLinkClicksInner and assigns it to the LinkClicks field.
func (o *DashboardChart) SetLinkClicks(v []DashboardChartLinkClicksInner) {
	o.LinkClicks = v
}

// GetCampaignViews returns the CampaignViews field value if set, zero value otherwise.
func (o *DashboardChart) GetCampaignViews() []DashboardChartLinkClicksInner {
	if o == nil || IsNil(o.CampaignViews) {
		var ret []DashboardChartLinkClicksInner
		return ret
	}
	return o.CampaignViews
}

// GetCampaignViewsOk returns a tuple with the CampaignViews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardChart) GetCampaignViewsOk() ([]DashboardChartLinkClicksInner, bool) {
	if o == nil || IsNil(o.CampaignViews) {
		return nil, false
	}
	return o.CampaignViews, true
}

// HasCampaignViews returns a boolean if a field has been set.
func (o *DashboardChart) HasCampaignViews() bool {
	if o != nil && !IsNil(o.CampaignViews) {
		return true
	}

	return false
}

// SetCampaignViews gets a reference to the given []DashboardChartLinkClicksInner and assigns it to the CampaignViews field.
func (o *DashboardChart) SetCampaignViews(v []DashboardChartLinkClicksInner) {
	o.CampaignViews = v
}

func (o DashboardChart) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardChart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LinkClicks) {
		toSerialize["link_clicks"] = o.LinkClicks
	}
	if !IsNil(o.CampaignViews) {
		toSerialize["campaign_views"] = o.CampaignViews
	}
	return toSerialize, nil
}

type NullableDashboardChart struct {
	value *DashboardChart
	isSet bool
}

func (v NullableDashboardChart) Get() *DashboardChart {
	return v.value
}

func (v *NullableDashboardChart) Set(val *DashboardChart) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardChart) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardChart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardChart(val *DashboardChart) *NullableDashboardChart {
	return &NullableDashboardChart{value: val, isSet: true}
}

func (v NullableDashboardChart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardChart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


