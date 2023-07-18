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

// checks if the Campaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Campaign{}

// Campaign struct for Campaign
type Campaign struct {
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	CampaignID *int32 `json:"CampaignID,omitempty"`
	Views *int32 `json:"views,omitempty"`
	Clicks *int32 `json:"clicks,omitempty"`
	Lists []BounceResultsInnerCampaign `json:"lists,omitempty"`
	StartedAt *string `json:"started_at,omitempty"`
	ToSend *int32 `json:"to_send,omitempty"`
	Sent *int32 `json:"sent,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Type *string `json:"type,omitempty"`
	Name *string `json:"name,omitempty"`
	Subject *string `json:"subject,omitempty"`
	FromEmail *string `json:"from_email,omitempty"`
	Body *string `json:"body,omitempty"`
	SendAt *string `json:"send_at,omitempty"`
	Status *string `json:"status,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Tags []string `json:"tags,omitempty"`
	TemplateId *int32 `json:"template_id,omitempty"`
	Messenger *string `json:"messenger,omitempty"`
}

// NewCampaign instantiates a new Campaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaign() *Campaign {
	this := Campaign{}
	return &this
}

// NewCampaignWithDefaults instantiates a new Campaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignWithDefaults() *Campaign {
	this := Campaign{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Campaign) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Campaign) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Campaign) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Campaign) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Campaign) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Campaign) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Campaign) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Campaign) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Campaign) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCampaignID returns the CampaignID field value if set, zero value otherwise.
func (o *Campaign) GetCampaignID() int32 {
	if o == nil || IsNil(o.CampaignID) {
		var ret int32
		return ret
	}
	return *o.CampaignID
}

// GetCampaignIDOk returns a tuple with the CampaignID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetCampaignIDOk() (*int32, bool) {
	if o == nil || IsNil(o.CampaignID) {
		return nil, false
	}
	return o.CampaignID, true
}

// HasCampaignID returns a boolean if a field has been set.
func (o *Campaign) HasCampaignID() bool {
	if o != nil && !IsNil(o.CampaignID) {
		return true
	}

	return false
}

// SetCampaignID gets a reference to the given int32 and assigns it to the CampaignID field.
func (o *Campaign) SetCampaignID(v int32) {
	o.CampaignID = &v
}

// GetViews returns the Views field value if set, zero value otherwise.
func (o *Campaign) GetViews() int32 {
	if o == nil || IsNil(o.Views) {
		var ret int32
		return ret
	}
	return *o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetViewsOk() (*int32, bool) {
	if o == nil || IsNil(o.Views) {
		return nil, false
	}
	return o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *Campaign) HasViews() bool {
	if o != nil && !IsNil(o.Views) {
		return true
	}

	return false
}

// SetViews gets a reference to the given int32 and assigns it to the Views field.
func (o *Campaign) SetViews(v int32) {
	o.Views = &v
}

// GetClicks returns the Clicks field value if set, zero value otherwise.
func (o *Campaign) GetClicks() int32 {
	if o == nil || IsNil(o.Clicks) {
		var ret int32
		return ret
	}
	return *o.Clicks
}

// GetClicksOk returns a tuple with the Clicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetClicksOk() (*int32, bool) {
	if o == nil || IsNil(o.Clicks) {
		return nil, false
	}
	return o.Clicks, true
}

// HasClicks returns a boolean if a field has been set.
func (o *Campaign) HasClicks() bool {
	if o != nil && !IsNil(o.Clicks) {
		return true
	}

	return false
}

// SetClicks gets a reference to the given int32 and assigns it to the Clicks field.
func (o *Campaign) SetClicks(v int32) {
	o.Clicks = &v
}

// GetLists returns the Lists field value if set, zero value otherwise.
func (o *Campaign) GetLists() []BounceResultsInnerCampaign {
	if o == nil || IsNil(o.Lists) {
		var ret []BounceResultsInnerCampaign
		return ret
	}
	return o.Lists
}

// GetListsOk returns a tuple with the Lists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetListsOk() ([]BounceResultsInnerCampaign, bool) {
	if o == nil || IsNil(o.Lists) {
		return nil, false
	}
	return o.Lists, true
}

// HasLists returns a boolean if a field has been set.
func (o *Campaign) HasLists() bool {
	if o != nil && !IsNil(o.Lists) {
		return true
	}

	return false
}

// SetLists gets a reference to the given []BounceResultsInnerCampaign and assigns it to the Lists field.
func (o *Campaign) SetLists(v []BounceResultsInnerCampaign) {
	o.Lists = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *Campaign) GetStartedAt() string {
	if o == nil || IsNil(o.StartedAt) {
		var ret string
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetStartedAtOk() (*string, bool) {
	if o == nil || IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *Campaign) HasStartedAt() bool {
	if o != nil && !IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given string and assigns it to the StartedAt field.
func (o *Campaign) SetStartedAt(v string) {
	o.StartedAt = &v
}

// GetToSend returns the ToSend field value if set, zero value otherwise.
func (o *Campaign) GetToSend() int32 {
	if o == nil || IsNil(o.ToSend) {
		var ret int32
		return ret
	}
	return *o.ToSend
}

// GetToSendOk returns a tuple with the ToSend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetToSendOk() (*int32, bool) {
	if o == nil || IsNil(o.ToSend) {
		return nil, false
	}
	return o.ToSend, true
}

// HasToSend returns a boolean if a field has been set.
func (o *Campaign) HasToSend() bool {
	if o != nil && !IsNil(o.ToSend) {
		return true
	}

	return false
}

// SetToSend gets a reference to the given int32 and assigns it to the ToSend field.
func (o *Campaign) SetToSend(v int32) {
	o.ToSend = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *Campaign) GetSent() int32 {
	if o == nil || IsNil(o.Sent) {
		var ret int32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetSentOk() (*int32, bool) {
	if o == nil || IsNil(o.Sent) {
		return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *Campaign) HasSent() bool {
	if o != nil && !IsNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given int32 and assigns it to the Sent field.
func (o *Campaign) SetSent(v int32) {
	o.Sent = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Campaign) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Campaign) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Campaign) SetUuid(v string) {
	o.Uuid = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Campaign) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Campaign) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Campaign) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Campaign) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Campaign) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Campaign) SetName(v string) {
	o.Name = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *Campaign) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *Campaign) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *Campaign) SetSubject(v string) {
	o.Subject = &v
}

// GetFromEmail returns the FromEmail field value if set, zero value otherwise.
func (o *Campaign) GetFromEmail() string {
	if o == nil || IsNil(o.FromEmail) {
		var ret string
		return ret
	}
	return *o.FromEmail
}

// GetFromEmailOk returns a tuple with the FromEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetFromEmailOk() (*string, bool) {
	if o == nil || IsNil(o.FromEmail) {
		return nil, false
	}
	return o.FromEmail, true
}

// HasFromEmail returns a boolean if a field has been set.
func (o *Campaign) HasFromEmail() bool {
	if o != nil && !IsNil(o.FromEmail) {
		return true
	}

	return false
}

// SetFromEmail gets a reference to the given string and assigns it to the FromEmail field.
func (o *Campaign) SetFromEmail(v string) {
	o.FromEmail = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *Campaign) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *Campaign) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *Campaign) SetBody(v string) {
	o.Body = &v
}

// GetSendAt returns the SendAt field value if set, zero value otherwise.
func (o *Campaign) GetSendAt() string {
	if o == nil || IsNil(o.SendAt) {
		var ret string
		return ret
	}
	return *o.SendAt
}

// GetSendAtOk returns a tuple with the SendAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetSendAtOk() (*string, bool) {
	if o == nil || IsNil(o.SendAt) {
		return nil, false
	}
	return o.SendAt, true
}

// HasSendAt returns a boolean if a field has been set.
func (o *Campaign) HasSendAt() bool {
	if o != nil && !IsNil(o.SendAt) {
		return true
	}

	return false
}

// SetSendAt gets a reference to the given string and assigns it to the SendAt field.
func (o *Campaign) SetSendAt(v string) {
	o.SendAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Campaign) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Campaign) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Campaign) SetStatus(v string) {
	o.Status = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *Campaign) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *Campaign) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *Campaign) SetContentType(v string) {
	o.ContentType = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Campaign) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Campaign) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Campaign) SetTags(v []string) {
	o.Tags = v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *Campaign) GetTemplateId() int32 {
	if o == nil || IsNil(o.TemplateId) {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *Campaign) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *Campaign) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetMessenger returns the Messenger field value if set, zero value otherwise.
func (o *Campaign) GetMessenger() string {
	if o == nil || IsNil(o.Messenger) {
		var ret string
		return ret
	}
	return *o.Messenger
}

// GetMessengerOk returns a tuple with the Messenger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetMessengerOk() (*string, bool) {
	if o == nil || IsNil(o.Messenger) {
		return nil, false
	}
	return o.Messenger, true
}

// HasMessenger returns a boolean if a field has been set.
func (o *Campaign) HasMessenger() bool {
	if o != nil && !IsNil(o.Messenger) {
		return true
	}

	return false
}

// SetMessenger gets a reference to the given string and assigns it to the Messenger field.
func (o *Campaign) SetMessenger(v string) {
	o.Messenger = &v
}

func (o Campaign) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Campaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.CampaignID) {
		toSerialize["CampaignID"] = o.CampaignID
	}
	if !IsNil(o.Views) {
		toSerialize["views"] = o.Views
	}
	if !IsNil(o.Clicks) {
		toSerialize["clicks"] = o.Clicks
	}
	if !IsNil(o.Lists) {
		toSerialize["lists"] = o.Lists
	}
	if !IsNil(o.StartedAt) {
		toSerialize["started_at"] = o.StartedAt
	}
	if !IsNil(o.ToSend) {
		toSerialize["to_send"] = o.ToSend
	}
	if !IsNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.FromEmail) {
		toSerialize["from_email"] = o.FromEmail
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.SendAt) {
		toSerialize["send_at"] = o.SendAt
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ContentType) {
		toSerialize["content_type"] = o.ContentType
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if !IsNil(o.Messenger) {
		toSerialize["messenger"] = o.Messenger
	}
	return toSerialize, nil
}

type NullableCampaign struct {
	value *Campaign
	isSet bool
}

func (v NullableCampaign) Get() *Campaign {
	return v.value
}

func (v *NullableCampaign) Set(val *Campaign) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaign(val *Campaign) *NullableCampaign {
	return &NullableCampaign{value: val, isSet: true}
}

func (v NullableCampaign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


