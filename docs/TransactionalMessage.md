# TransactionalMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriberEmail** | Pointer to **string** |  | [optional] 
**SubscriberId** | Pointer to **int32** |  | [optional] 
**TemplateId** | Pointer to **int32** |  | [optional] 
**FromEmail** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Headers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Messenger** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 

## Methods

### NewTransactionalMessage

`func NewTransactionalMessage() *TransactionalMessage`

NewTransactionalMessage instantiates a new TransactionalMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionalMessageWithDefaults

`func NewTransactionalMessageWithDefaults() *TransactionalMessage`

NewTransactionalMessageWithDefaults instantiates a new TransactionalMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriberEmail

`func (o *TransactionalMessage) GetSubscriberEmail() string`

GetSubscriberEmail returns the SubscriberEmail field if non-nil, zero value otherwise.

### GetSubscriberEmailOk

`func (o *TransactionalMessage) GetSubscriberEmailOk() (*string, bool)`

GetSubscriberEmailOk returns a tuple with the SubscriberEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberEmail

`func (o *TransactionalMessage) SetSubscriberEmail(v string)`

SetSubscriberEmail sets SubscriberEmail field to given value.

### HasSubscriberEmail

`func (o *TransactionalMessage) HasSubscriberEmail() bool`

HasSubscriberEmail returns a boolean if a field has been set.

### GetSubscriberId

`func (o *TransactionalMessage) GetSubscriberId() int32`

GetSubscriberId returns the SubscriberId field if non-nil, zero value otherwise.

### GetSubscriberIdOk

`func (o *TransactionalMessage) GetSubscriberIdOk() (*int32, bool)`

GetSubscriberIdOk returns a tuple with the SubscriberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberId

`func (o *TransactionalMessage) SetSubscriberId(v int32)`

SetSubscriberId sets SubscriberId field to given value.

### HasSubscriberId

`func (o *TransactionalMessage) HasSubscriberId() bool`

HasSubscriberId returns a boolean if a field has been set.

### GetTemplateId

`func (o *TransactionalMessage) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TransactionalMessage) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TransactionalMessage) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *TransactionalMessage) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetFromEmail

`func (o *TransactionalMessage) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *TransactionalMessage) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *TransactionalMessage) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *TransactionalMessage) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### GetData

`func (o *TransactionalMessage) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionalMessage) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionalMessage) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *TransactionalMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeaders

`func (o *TransactionalMessage) GetHeaders() []map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *TransactionalMessage) GetHeadersOk() (*[]map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *TransactionalMessage) SetHeaders(v []map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *TransactionalMessage) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMessenger

`func (o *TransactionalMessage) GetMessenger() string`

GetMessenger returns the Messenger field if non-nil, zero value otherwise.

### GetMessengerOk

`func (o *TransactionalMessage) GetMessengerOk() (*string, bool)`

GetMessengerOk returns a tuple with the Messenger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessenger

`func (o *TransactionalMessage) SetMessenger(v string)`

SetMessenger sets Messenger field to given value.

### HasMessenger

`func (o *TransactionalMessage) HasMessenger() bool`

HasMessenger returns a boolean if a field has been set.

### GetContentType

`func (o *TransactionalMessage) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *TransactionalMessage) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *TransactionalMessage) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *TransactionalMessage) HasContentType() bool`

HasContentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


