# GetMedia200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MediaFileObject**](MediaFileObject.md) |  | [optional] 

## Methods

### NewGetMedia200Response

`func NewGetMedia200Response() *GetMedia200Response`

NewGetMedia200Response instantiates a new GetMedia200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMedia200ResponseWithDefaults

`func NewGetMedia200ResponseWithDefaults() *GetMedia200Response`

NewGetMedia200ResponseWithDefaults instantiates a new GetMedia200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetMedia200Response) GetData() []MediaFileObject`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMedia200Response) GetDataOk() (*[]MediaFileObject, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMedia200Response) SetData(v []MediaFileObject)`

SetData sets Data field to given value.

### HasData

`func (o *GetMedia200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


