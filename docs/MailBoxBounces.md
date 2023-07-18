# MailBoxBounces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**AuthProtocol** | Pointer to **string** |  | [optional] 
**ReturnPath** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**TlsEnabled** | Pointer to **bool** |  | [optional] 
**TlsSkipVerify** | Pointer to **bool** |  | [optional] 
**ScanInterval** | Pointer to **string** |  | [optional] 

## Methods

### NewMailBoxBounces

`func NewMailBoxBounces() *MailBoxBounces`

NewMailBoxBounces instantiates a new MailBoxBounces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailBoxBouncesWithDefaults

`func NewMailBoxBouncesWithDefaults() *MailBoxBounces`

NewMailBoxBouncesWithDefaults instantiates a new MailBoxBounces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *MailBoxBounces) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MailBoxBounces) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MailBoxBounces) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MailBoxBounces) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetEnabled

`func (o *MailBoxBounces) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MailBoxBounces) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MailBoxBounces) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MailBoxBounces) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *MailBoxBounces) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MailBoxBounces) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MailBoxBounces) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MailBoxBounces) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHost

`func (o *MailBoxBounces) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MailBoxBounces) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MailBoxBounces) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *MailBoxBounces) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *MailBoxBounces) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MailBoxBounces) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MailBoxBounces) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *MailBoxBounces) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAuthProtocol

`func (o *MailBoxBounces) GetAuthProtocol() string`

GetAuthProtocol returns the AuthProtocol field if non-nil, zero value otherwise.

### GetAuthProtocolOk

`func (o *MailBoxBounces) GetAuthProtocolOk() (*string, bool)`

GetAuthProtocolOk returns a tuple with the AuthProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProtocol

`func (o *MailBoxBounces) SetAuthProtocol(v string)`

SetAuthProtocol sets AuthProtocol field to given value.

### HasAuthProtocol

`func (o *MailBoxBounces) HasAuthProtocol() bool`

HasAuthProtocol returns a boolean if a field has been set.

### GetReturnPath

`func (o *MailBoxBounces) GetReturnPath() string`

GetReturnPath returns the ReturnPath field if non-nil, zero value otherwise.

### GetReturnPathOk

`func (o *MailBoxBounces) GetReturnPathOk() (*string, bool)`

GetReturnPathOk returns a tuple with the ReturnPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPath

`func (o *MailBoxBounces) SetReturnPath(v string)`

SetReturnPath sets ReturnPath field to given value.

### HasReturnPath

`func (o *MailBoxBounces) HasReturnPath() bool`

HasReturnPath returns a boolean if a field has been set.

### GetUsername

`func (o *MailBoxBounces) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MailBoxBounces) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MailBoxBounces) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MailBoxBounces) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetTlsEnabled

`func (o *MailBoxBounces) GetTlsEnabled() bool`

GetTlsEnabled returns the TlsEnabled field if non-nil, zero value otherwise.

### GetTlsEnabledOk

`func (o *MailBoxBounces) GetTlsEnabledOk() (*bool, bool)`

GetTlsEnabledOk returns a tuple with the TlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnabled

`func (o *MailBoxBounces) SetTlsEnabled(v bool)`

SetTlsEnabled sets TlsEnabled field to given value.

### HasTlsEnabled

`func (o *MailBoxBounces) HasTlsEnabled() bool`

HasTlsEnabled returns a boolean if a field has been set.

### GetTlsSkipVerify

`func (o *MailBoxBounces) GetTlsSkipVerify() bool`

GetTlsSkipVerify returns the TlsSkipVerify field if non-nil, zero value otherwise.

### GetTlsSkipVerifyOk

`func (o *MailBoxBounces) GetTlsSkipVerifyOk() (*bool, bool)`

GetTlsSkipVerifyOk returns a tuple with the TlsSkipVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsSkipVerify

`func (o *MailBoxBounces) SetTlsSkipVerify(v bool)`

SetTlsSkipVerify sets TlsSkipVerify field to given value.

### HasTlsSkipVerify

`func (o *MailBoxBounces) HasTlsSkipVerify() bool`

HasTlsSkipVerify returns a boolean if a field has been set.

### GetScanInterval

`func (o *MailBoxBounces) GetScanInterval() string`

GetScanInterval returns the ScanInterval field if non-nil, zero value otherwise.

### GetScanIntervalOk

`func (o *MailBoxBounces) GetScanIntervalOk() (*string, bool)`

GetScanIntervalOk returns a tuple with the ScanInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanInterval

`func (o *MailBoxBounces) SetScanInterval(v string)`

SetScanInterval sets ScanInterval field to given value.

### HasScanInterval

`func (o *MailBoxBounces) HasScanInterval() bool`

HasScanInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


