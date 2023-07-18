# ServerConfigData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messengers** | Pointer to **[]string** |  | [optional] 
**Langs** | Pointer to [**[]ServerConfigDataLangsInner**](ServerConfigDataLangsInner.md) |  | [optional] 
**Lang** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **string** |  | [optional] 
**NeedsRestart** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewServerConfigData

`func NewServerConfigData() *ServerConfigData`

NewServerConfigData instantiates a new ServerConfigData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigDataWithDefaults

`func NewServerConfigDataWithDefaults() *ServerConfigData`

NewServerConfigDataWithDefaults instantiates a new ServerConfigData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessengers

`func (o *ServerConfigData) GetMessengers() []string`

GetMessengers returns the Messengers field if non-nil, zero value otherwise.

### GetMessengersOk

`func (o *ServerConfigData) GetMessengersOk() (*[]string, bool)`

GetMessengersOk returns a tuple with the Messengers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessengers

`func (o *ServerConfigData) SetMessengers(v []string)`

SetMessengers sets Messengers field to given value.

### HasMessengers

`func (o *ServerConfigData) HasMessengers() bool`

HasMessengers returns a boolean if a field has been set.

### GetLangs

`func (o *ServerConfigData) GetLangs() []ServerConfigDataLangsInner`

GetLangs returns the Langs field if non-nil, zero value otherwise.

### GetLangsOk

`func (o *ServerConfigData) GetLangsOk() (*[]ServerConfigDataLangsInner, bool)`

GetLangsOk returns a tuple with the Langs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangs

`func (o *ServerConfigData) SetLangs(v []ServerConfigDataLangsInner)`

SetLangs sets Langs field to given value.

### HasLangs

`func (o *ServerConfigData) HasLangs() bool`

HasLangs returns a boolean if a field has been set.

### GetLang

`func (o *ServerConfigData) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *ServerConfigData) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *ServerConfigData) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *ServerConfigData) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetUpdate

`func (o *ServerConfigData) GetUpdate() string`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ServerConfigData) GetUpdateOk() (*string, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ServerConfigData) SetUpdate(v string)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ServerConfigData) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetNeedsRestart

`func (o *ServerConfigData) GetNeedsRestart() bool`

GetNeedsRestart returns the NeedsRestart field if non-nil, zero value otherwise.

### GetNeedsRestartOk

`func (o *ServerConfigData) GetNeedsRestartOk() (*bool, bool)`

GetNeedsRestartOk returns a tuple with the NeedsRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsRestart

`func (o *ServerConfigData) SetNeedsRestart(v bool)`

SetNeedsRestart sets NeedsRestart field to given value.

### HasNeedsRestart

`func (o *ServerConfigData) HasNeedsRestart() bool`

HasNeedsRestart returns a boolean if a field has been set.

### GetVersion

`func (o *ServerConfigData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServerConfigData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServerConfigData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServerConfigData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


