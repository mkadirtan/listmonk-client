# \BouncesApi

All URIs are relative to *http://localhost:9000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBounceById**](BouncesApi.md#DeleteBounceById) | **Delete** /bounces/{id} | 
[**DeleteBounces**](BouncesApi.md#DeleteBounces) | **Delete** /bounces | 
[**GetBounceById**](BouncesApi.md#GetBounceById) | **Get** /bounces/{id} | 
[**GetBounces**](BouncesApi.md#GetBounces) | **Get** /bounces | 



## DeleteBounceById

> GetHealthCheck200Response DeleteBounceById(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(56) // int32 | The id value of the bounce you want to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BouncesApi.DeleteBounceById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BouncesApi.DeleteBounceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBounceById`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `BouncesApi.DeleteBounceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id value of the bounce you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBounceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetHealthCheck200Response**](GetHealthCheck200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBounces

> GetHealthCheck200Response DeleteBounces(ctx).All(all).Id(id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    all := true // bool | flag for multiple bounce record deletion (optional)
    id := "id_example" // string | list of bounce ids to delete (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BouncesApi.DeleteBounces(context.Background()).All(all).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BouncesApi.DeleteBounces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBounces`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `BouncesApi.DeleteBounces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBouncesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **bool** | flag for multiple bounce record deletion | 
 **id** | **string** | list of bounce ids to delete | 

### Return type

[**GetHealthCheck200Response**](GetHealthCheck200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBounceById

> GetBounceById200Response GetBounceById(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(56) // int32 | The id value of the bounce you want to retreive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BouncesApi.GetBounceById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BouncesApi.GetBounceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBounceById`: GetBounceById200Response
    fmt.Fprintf(os.Stdout, "Response from `BouncesApi.GetBounceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id value of the bounce you want to retreive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBounceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetBounceById200Response**](GetBounceById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBounces

> GetBounces200Response GetBounces(ctx).CampaignId(campaignId).Page(page).PerPage(perPage).Source(source).OrderBy(orderBy).Order(order).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    campaignId := int32(56) // int32 | bounce record retrieval of particular campaign (optional)
    page := int32(56) // int32 | total number of pages (optional)
    perPage := int32(56) // int32 | number of items per page (optional)
    source := "source_example" // string |  (optional)
    orderBy := "orderBy_example" // string |  (optional)
    order := "order_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BouncesApi.GetBounces(context.Background()).CampaignId(campaignId).Page(page).PerPage(perPage).Source(source).OrderBy(orderBy).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BouncesApi.GetBounces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBounces`: GetBounces200Response
    fmt.Fprintf(os.Stdout, "Response from `BouncesApi.GetBounces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBouncesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **int32** | bounce record retrieval of particular campaign | 
 **page** | **int32** | total number of pages | 
 **perPage** | **int32** | number of items per page | 
 **source** | **string** |  | 
 **orderBy** | **string** |  | 
 **order** | **string** |  | 

### Return type

[**GetBounces200Response**](GetBounces200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

