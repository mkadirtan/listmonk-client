# GetLists200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]List**](List.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**PerPage** | Pointer to **int32** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetLists200ResponseData

`func NewGetLists200ResponseData() *GetLists200ResponseData`

NewGetLists200ResponseData instantiates a new GetLists200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLists200ResponseDataWithDefaults

`func NewGetLists200ResponseDataWithDefaults() *GetLists200ResponseData`

NewGetLists200ResponseDataWithDefaults instantiates a new GetLists200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *GetLists200ResponseData) GetResults() []List`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetLists200ResponseData) GetResultsOk() (*[]List, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetLists200ResponseData) SetResults(v []List)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetLists200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotal

`func (o *GetLists200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetLists200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetLists200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetLists200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPerPage

`func (o *GetLists200ResponseData) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *GetLists200ResponseData) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *GetLists200ResponseData) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *GetLists200ResponseData) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetPage

`func (o *GetLists200ResponseData) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetLists200ResponseData) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetLists200ResponseData) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetLists200ResponseData) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


