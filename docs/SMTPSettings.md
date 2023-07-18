# SMTPSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**HelloHostname** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**AuthProtocol** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**EmailHeaders** | Pointer to **[]map[string]interface{}** |  | [optional] 
**MaxConns** | Pointer to **int32** |  | [optional] 
**MaxMsgRetries** | Pointer to **int32** |  | [optional] 
**IdleTimeout** | Pointer to **string** |  | [optional] 
**WaitTimeout** | Pointer to **string** |  | [optional] 
**TlsType** | Pointer to **string** |  | [optional] 
**TlsSkipVerify** | Pointer to **bool** |  | [optional] 

## Methods

### NewSMTPSettings

`func NewSMTPSettings() *SMTPSettings`

NewSMTPSettings instantiates a new SMTPSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMTPSettingsWithDefaults

`func NewSMTPSettingsWithDefaults() *SMTPSettings`

NewSMTPSettingsWithDefaults instantiates a new SMTPSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *SMTPSettings) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SMTPSettings) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SMTPSettings) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SMTPSettings) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetEnabled

`func (o *SMTPSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SMTPSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SMTPSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SMTPSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHost

`func (o *SMTPSettings) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SMTPSettings) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SMTPSettings) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SMTPSettings) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHelloHostname

`func (o *SMTPSettings) GetHelloHostname() string`

GetHelloHostname returns the HelloHostname field if non-nil, zero value otherwise.

### GetHelloHostnameOk

`func (o *SMTPSettings) GetHelloHostnameOk() (*string, bool)`

GetHelloHostnameOk returns a tuple with the HelloHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloHostname

`func (o *SMTPSettings) SetHelloHostname(v string)`

SetHelloHostname sets HelloHostname field to given value.

### HasHelloHostname

`func (o *SMTPSettings) HasHelloHostname() bool`

HasHelloHostname returns a boolean if a field has been set.

### GetPort

`func (o *SMTPSettings) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SMTPSettings) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SMTPSettings) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SMTPSettings) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAuthProtocol

`func (o *SMTPSettings) GetAuthProtocol() string`

GetAuthProtocol returns the AuthProtocol field if non-nil, zero value otherwise.

### GetAuthProtocolOk

`func (o *SMTPSettings) GetAuthProtocolOk() (*string, bool)`

GetAuthProtocolOk returns a tuple with the AuthProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProtocol

`func (o *SMTPSettings) SetAuthProtocol(v string)`

SetAuthProtocol sets AuthProtocol field to given value.

### HasAuthProtocol

`func (o *SMTPSettings) HasAuthProtocol() bool`

HasAuthProtocol returns a boolean if a field has been set.

### GetUsername

`func (o *SMTPSettings) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SMTPSettings) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SMTPSettings) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SMTPSettings) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmailHeaders

`func (o *SMTPSettings) GetEmailHeaders() []map[string]interface{}`

GetEmailHeaders returns the EmailHeaders field if non-nil, zero value otherwise.

### GetEmailHeadersOk

`func (o *SMTPSettings) GetEmailHeadersOk() (*[]map[string]interface{}, bool)`

GetEmailHeadersOk returns a tuple with the EmailHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailHeaders

`func (o *SMTPSettings) SetEmailHeaders(v []map[string]interface{})`

SetEmailHeaders sets EmailHeaders field to given value.

### HasEmailHeaders

`func (o *SMTPSettings) HasEmailHeaders() bool`

HasEmailHeaders returns a boolean if a field has been set.

### GetMaxConns

`func (o *SMTPSettings) GetMaxConns() int32`

GetMaxConns returns the MaxConns field if non-nil, zero value otherwise.

### GetMaxConnsOk

`func (o *SMTPSettings) GetMaxConnsOk() (*int32, bool)`

GetMaxConnsOk returns a tuple with the MaxConns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConns

`func (o *SMTPSettings) SetMaxConns(v int32)`

SetMaxConns sets MaxConns field to given value.

### HasMaxConns

`func (o *SMTPSettings) HasMaxConns() bool`

HasMaxConns returns a boolean if a field has been set.

### GetMaxMsgRetries

`func (o *SMTPSettings) GetMaxMsgRetries() int32`

GetMaxMsgRetries returns the MaxMsgRetries field if non-nil, zero value otherwise.

### GetMaxMsgRetriesOk

`func (o *SMTPSettings) GetMaxMsgRetriesOk() (*int32, bool)`

GetMaxMsgRetriesOk returns a tuple with the MaxMsgRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgRetries

`func (o *SMTPSettings) SetMaxMsgRetries(v int32)`

SetMaxMsgRetries sets MaxMsgRetries field to given value.

### HasMaxMsgRetries

`func (o *SMTPSettings) HasMaxMsgRetries() bool`

HasMaxMsgRetries returns a boolean if a field has been set.

### GetIdleTimeout

`func (o *SMTPSettings) GetIdleTimeout() string`

GetIdleTimeout returns the IdleTimeout field if non-nil, zero value otherwise.

### GetIdleTimeoutOk

`func (o *SMTPSettings) GetIdleTimeoutOk() (*string, bool)`

GetIdleTimeoutOk returns a tuple with the IdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeout

`func (o *SMTPSettings) SetIdleTimeout(v string)`

SetIdleTimeout sets IdleTimeout field to given value.

### HasIdleTimeout

`func (o *SMTPSettings) HasIdleTimeout() bool`

HasIdleTimeout returns a boolean if a field has been set.

### GetWaitTimeout

`func (o *SMTPSettings) GetWaitTimeout() string`

GetWaitTimeout returns the WaitTimeout field if non-nil, zero value otherwise.

### GetWaitTimeoutOk

`func (o *SMTPSettings) GetWaitTimeoutOk() (*string, bool)`

GetWaitTimeoutOk returns a tuple with the WaitTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTimeout

`func (o *SMTPSettings) SetWaitTimeout(v string)`

SetWaitTimeout sets WaitTimeout field to given value.

### HasWaitTimeout

`func (o *SMTPSettings) HasWaitTimeout() bool`

HasWaitTimeout returns a boolean if a field has been set.

### GetTlsType

`func (o *SMTPSettings) GetTlsType() string`

GetTlsType returns the TlsType field if non-nil, zero value otherwise.

### GetTlsTypeOk

`func (o *SMTPSettings) GetTlsTypeOk() (*string, bool)`

GetTlsTypeOk returns a tuple with the TlsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsType

`func (o *SMTPSettings) SetTlsType(v string)`

SetTlsType sets TlsType field to given value.

### HasTlsType

`func (o *SMTPSettings) HasTlsType() bool`

HasTlsType returns a boolean if a field has been set.

### GetTlsSkipVerify

`func (o *SMTPSettings) GetTlsSkipVerify() bool`

GetTlsSkipVerify returns the TlsSkipVerify field if non-nil, zero value otherwise.

### GetTlsSkipVerifyOk

`func (o *SMTPSettings) GetTlsSkipVerifyOk() (*bool, bool)`

GetTlsSkipVerifyOk returns a tuple with the TlsSkipVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsSkipVerify

`func (o *SMTPSettings) SetTlsSkipVerify(v bool)`

SetTlsSkipVerify sets TlsSkipVerify field to given value.

### HasTlsSkipVerify

`func (o *SMTPSettings) HasTlsSkipVerify() bool`

HasTlsSkipVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


