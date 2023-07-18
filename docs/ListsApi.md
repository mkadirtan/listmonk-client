# \ListsApi

All URIs are relative to *http://localhost:9000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateList**](ListsApi.md#CreateList) | **Post** /lists | 
[**DeleteListById**](ListsApi.md#DeleteListById) | **Delete** /lists/{list_id} | 
[**GetListById**](ListsApi.md#GetListById) | **Get** /lists/{list_id} | 
[**GetLists**](ListsApi.md#GetLists) | **Get** /lists | 
[**UpdateListById**](ListsApi.md#UpdateListById) | **Put** /lists/{list_id} | 



## CreateList

> CreateList200Response CreateList(ctx).NewList(newList).Execute()





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
    newList := *openapiclient.NewNewList() // NewList | new list info (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsApi.CreateList(context.Background()).NewList(newList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsApi.CreateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateList`: CreateList200Response
    fmt.Fprintf(os.Stdout, "Response from `ListsApi.CreateList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newList** | [**NewList**](NewList.md) | new list info | 

### Return type

[**CreateList200Response**](CreateList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteListById

> GetHealthCheck200Response DeleteListById(ctx, listId).Execute()





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
    listId := int32(56) // int32 | The id value of the lists you want to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsApi.DeleteListById(context.Background(), listId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsApi.DeleteListById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteListById`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `ListsApi.DeleteListById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int32** | The id value of the lists you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListByIdRequest struct via the builder pattern


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


## GetListById

> CreateList200Response GetListById(ctx, listId).Execute()





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
    listId := int32(56) // int32 | The id value of the list you want to retreive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsApi.GetListById(context.Background(), listId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsApi.GetListById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListById`: CreateList200Response
    fmt.Fprintf(os.Stdout, "Response from `ListsApi.GetListById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int32** | The id value of the list you want to retreive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateList200Response**](CreateList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLists

> GetLists200Response GetLists(ctx).Page(page).PerPage(perPage).Query(query).OrderBy(orderBy).Order(order).Minimal(minimal).Execute()





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
    page := int32(56) // int32 | total number of pages (optional)
    perPage := int32(56) // int32 | number of items per page (optional)
    query := "query_example" // string | Optional string to search a list by name. (optional)
    orderBy := "orderBy_example" // string | Field to sort results by. name|status|created_at|updated_at (optional)
    order := "order_example" // string | ASC|DESC Sort by ascending or descending order. (optional)
    minimal := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsApi.GetLists(context.Background()).Page(page).PerPage(perPage).Query(query).OrderBy(orderBy).Order(order).Minimal(minimal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsApi.GetLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLists`: GetLists200Response
    fmt.Fprintf(os.Stdout, "Response from `ListsApi.GetLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | total number of pages | 
 **perPage** | **int32** | number of items per page | 
 **query** | **string** | Optional string to search a list by name. | 
 **orderBy** | **string** | Field to sort results by. name|status|created_at|updated_at | 
 **order** | **string** | ASC|DESC Sort by ascending or descending order. | 
 **minimal** | **bool** |  | 

### Return type

[**GetLists200Response**](GetLists200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateListById

> CreateList200Response UpdateListById(ctx, listId).List(list).Execute()





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
    listId := int32(56) // int32 | The id value of the list you want to update
    list := *openapiclient.NewList() // List | updated list field values (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsApi.UpdateListById(context.Background(), listId).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsApi.UpdateListById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateListById`: CreateList200Response
    fmt.Fprintf(os.Stdout, "Response from `ListsApi.UpdateListById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int32** | The id value of the list you want to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | [**List**](List.md) | updated list field values | 

### Return type

[**CreateList200Response**](CreateList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

