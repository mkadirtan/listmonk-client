# SMTPTest

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
**StrEmailHeaders** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 

## Methods

### NewSMTPTest

`func NewSMTPTest() *SMTPTest`

NewSMTPTest instantiates a new SMTPTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMTPTestWithDefaults

`func NewSMTPTestWithDefaults() *SMTPTest`

NewSMTPTestWithDefaults instantiates a new SMTPTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *SMTPTest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SMTPTest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SMTPTest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SMTPTest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetEnabled

`func (o *SMTPTest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SMTPTest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SMTPTest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SMTPTest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHost

`func (o *SMTPTest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SMTPTest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SMTPTest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SMTPTest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHelloHostname

`func (o *SMTPTest) GetHelloHostname() string`

GetHelloHostname returns the HelloHostname field if non-nil, zero value otherwise.

### GetHelloHostnameOk

`func (o *SMTPTest) GetHelloHostnameOk() (*string, bool)`

GetHelloHostnameOk returns a tuple with the HelloHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloHostname

`func (o *SMTPTest) SetHelloHostname(v string)`

SetHelloHostname sets HelloHostname field to given value.

### HasHelloHostname

`func (o *SMTPTest) HasHelloHostname() bool`

HasHelloHostname returns a boolean if a field has been set.

### GetPort

`func (o *SMTPTest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SMTPTest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SMTPTest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SMTPTest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAuthProtocol

`func (o *SMTPTest) GetAuthProtocol() string`

GetAuthProtocol returns the AuthProtocol field if non-nil, zero value otherwise.

### GetAuthProtocolOk

`func (o *SMTPTest) GetAuthProtocolOk() (*string, bool)`

GetAuthProtocolOk returns a tuple with the AuthProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProtocol

`func (o *SMTPTest) SetAuthProtocol(v string)`

SetAuthProtocol sets AuthProtocol field to given value.

### HasAuthProtocol

`func (o *SMTPTest) HasAuthProtocol() bool`

HasAuthProtocol returns a boolean if a field has been set.

### GetUsername

`func (o *SMTPTest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SMTPTest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SMTPTest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SMTPTest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmailHeaders

`func (o *SMTPTest) GetEmailHeaders() []map[string]interface{}`

GetEmailHeaders returns the EmailHeaders field if non-nil, zero value otherwise.

### GetEmailHeadersOk

`func (o *SMTPTest) GetEmailHeadersOk() (*[]map[string]interface{}, bool)`

GetEmailHeadersOk returns a tuple with the EmailHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailHeaders

`func (o *SMTPTest) SetEmailHeaders(v []map[string]interface{})`

SetEmailHeaders sets EmailHeaders field to given value.

### HasEmailHeaders

`func (o *SMTPTest) HasEmailHeaders() bool`

HasEmailHeaders returns a boolean if a field has been set.

### GetMaxConns

`func (o *SMTPTest) GetMaxConns() int32`

GetMaxConns returns the MaxConns field if non-nil, zero value otherwise.

### GetMaxConnsOk

`func (o *SMTPTest) GetMaxConnsOk() (*int32, bool)`

GetMaxConnsOk returns a tuple with the MaxConns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConns

`func (o *SMTPTest) SetMaxConns(v int32)`

SetMaxConns sets MaxConns field to given value.

### HasMaxConns

`func (o *SMTPTest) HasMaxConns() bool`

HasMaxConns returns a boolean if a field has been set.

### GetMaxMsgRetries

`func (o *SMTPTest) GetMaxMsgRetries() int32`

GetMaxMsgRetries returns the MaxMsgRetries field if non-nil, zero value otherwise.

### GetMaxMsgRetriesOk

`func (o *SMTPTest) GetMaxMsgRetriesOk() (*int32, bool)`

GetMaxMsgRetriesOk returns a tuple with the MaxMsgRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgRetries

`func (o *SMTPTest) SetMaxMsgRetries(v int32)`

SetMaxMsgRetries sets MaxMsgRetries field to given value.

### HasMaxMsgRetries

`func (o *SMTPTest) HasMaxMsgRetries() bool`

HasMaxMsgRetries returns a boolean if a field has been set.

### GetIdleTimeout

`func (o *SMTPTest) GetIdleTimeout() string`

GetIdleTimeout returns the IdleTimeout field if non-nil, zero value otherwise.

### GetIdleTimeoutOk

`func (o *SMTPTest) GetIdleTimeoutOk() (*string, bool)`

GetIdleTimeoutOk returns a tuple with the IdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeout

`func (o *SMTPTest) SetIdleTimeout(v string)`

SetIdleTimeout sets IdleTimeout field to given value.

### HasIdleTimeout

`func (o *SMTPTest) HasIdleTimeout() bool`

HasIdleTimeout returns a boolean if a field has been set.

### GetWaitTimeout

`func (o *SMTPTest) GetWaitTimeout() string`

GetWaitTimeout returns the WaitTimeout field if non-nil, zero value otherwise.

### GetWaitTimeoutOk

`func (o *SMTPTest) GetWaitTimeoutOk() (*string, bool)`

GetWaitTimeoutOk returns a tuple with the WaitTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTimeout

`func (o *SMTPTest) SetWaitTimeout(v string)`

SetWaitTimeout sets WaitTimeout field to given value.

### HasWaitTimeout

`func (o *SMTPTest) HasWaitTimeout() bool`

HasWaitTimeout returns a boolean if a field has been set.

### GetTlsType

`func (o *SMTPTest) GetTlsType() string`

GetTlsType returns the TlsType field if non-nil, zero value otherwise.

### GetTlsTypeOk

`func (o *SMTPTest) GetTlsTypeOk() (*string, bool)`

GetTlsTypeOk returns a tuple with the TlsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsType

`func (o *SMTPTest) SetTlsType(v string)`

SetTlsType sets TlsType field to given value.

### HasTlsType

`func (o *SMTPTest) HasTlsType() bool`

HasTlsType returns a boolean if a field has been set.

### GetTlsSkipVerify

`func (o *SMTPTest) GetTlsSkipVerify() bool`

GetTlsSkipVerify returns the TlsSkipVerify field if non-nil, zero value otherwise.

### GetTlsSkipVerifyOk

`func (o *SMTPTest) GetTlsSkipVerifyOk() (*bool, bool)`

GetTlsSkipVerifyOk returns a tuple with the TlsSkipVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsSkipVerify

`func (o *SMTPTest) SetTlsSkipVerify(v bool)`

SetTlsSkipVerify sets TlsSkipVerify field to given value.

### HasTlsSkipVerify

`func (o *SMTPTest) HasTlsSkipVerify() bool`

HasTlsSkipVerify returns a boolean if a field has been set.

### GetStrEmailHeaders

`func (o *SMTPTest) GetStrEmailHeaders() string`

GetStrEmailHeaders returns the StrEmailHeaders field if non-nil, zero value otherwise.

### GetStrEmailHeadersOk

`func (o *SMTPTest) GetStrEmailHeadersOk() (*string, bool)`

GetStrEmailHeadersOk returns a tuple with the StrEmailHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrEmailHeaders

`func (o *SMTPTest) SetStrEmailHeaders(v string)`

SetStrEmailHeaders sets StrEmailHeaders field to given value.

### HasStrEmailHeaders

`func (o *SMTPTest) HasStrEmailHeaders() bool`

HasStrEmailHeaders returns a boolean if a field has been set.

### GetPassword

`func (o *SMTPTest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SMTPTest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SMTPTest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SMTPTest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEmail

`func (o *SMTPTest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SMTPTest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SMTPTest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SMTPTest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


