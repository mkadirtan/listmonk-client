# GetServerConfig200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ServerConfig**](ServerConfig.md) |  | [optional] 

## Methods

### NewGetServerConfig200Response

`func NewGetServerConfig200Response() *GetServerConfig200Response`

NewGetServerConfig200Response instantiates a new GetServerConfig200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerConfig200ResponseWithDefaults

`func NewGetServerConfig200ResponseWithDefaults() *GetServerConfig200Response`

NewGetServerConfig200ResponseWithDefaults instantiates a new GetServerConfig200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetServerConfig200Response) GetData() ServerConfig`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetServerConfig200Response) GetDataOk() (*ServerConfig, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetServerConfig200Response) SetData(v ServerConfig)`

SetData sets Data field to given value.

### HasData

`func (o *GetServerConfig200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


