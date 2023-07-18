# \SubscribersApi

All URIs are relative to *http://localhost:9000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlocklistSubscribersQuery**](SubscribersApi.md#BlocklistSubscribersQuery) | **Put** /subscribers/query/blocklist | 
[**CreateSubscriber**](SubscribersApi.md#CreateSubscriber) | **Post** /subscribers | 
[**DeleteSubscriberBouncesById**](SubscribersApi.md#DeleteSubscriberBouncesById) | **Delete** /subscribers/{id}/bounces | 
[**DeleteSubscriberById**](SubscribersApi.md#DeleteSubscriberById) | **Delete** /subscribers/{id} | 
[**DeleteSubscriberByList**](SubscribersApi.md#DeleteSubscriberByList) | **Delete** /subscribers | 
[**DeleteSubscriberByQuery**](SubscribersApi.md#DeleteSubscriberByQuery) | **Post** /subscribers/query/delete | 
[**ExportSubscriberDataByID**](SubscribersApi.md#ExportSubscriberDataByID) | **Get** /subscribers/{id}/export | 
[**GetSubscriberBouncesById**](SubscribersApi.md#GetSubscriberBouncesById) | **Get** /subscribers/{id}/bounces | 
[**GetSubscriberById**](SubscribersApi.md#GetSubscriberById) | **Get** /subscribers/{id} | 
[**GetSubscribers**](SubscribersApi.md#GetSubscribers) | **Get** /subscribers | 
[**ManageBlocklistBySubscriberList**](SubscribersApi.md#ManageBlocklistBySubscriberList) | **Put** /subscribers/blocklist | 
[**ManageBlocklistSubscribersById**](SubscribersApi.md#ManageBlocklistSubscribersById) | **Put** /subscribers/{id}/blocklist | 
[**ManageSubscriberListById**](SubscribersApi.md#ManageSubscriberListById) | **Put** /subscribers/lists/{id} | 
[**ManageSubscriberLists**](SubscribersApi.md#ManageSubscriberLists) | **Put** /subscribers/lists | 
[**ManageSubscriberListsByQuery**](SubscribersApi.md#ManageSubscriberListsByQuery) | **Put** /subscribers/query/lists | 
[**SubscriberSendOptinById**](SubscribersApi.md#SubscriberSendOptinById) | **Post** /subscribers/{id}/optin | 
[**UpdateSubscriberById**](SubscribersApi.md#UpdateSubscriberById) | **Put** /subscribers/{id} | 



## BlocklistSubscribersQuery

> GetHealthCheck200Response BlocklistSubscribersQuery(ctx).Body(body).Execute()





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
    body := "body_example" // string | Arbitrary SQL expression. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscribersApi.BlocklistSubscribersQuery(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.BlocklistSubscribersQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlocklistSubscribersQuery`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.BlocklistSubscribersQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlocklistSubscribersQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | Arbitrary SQL expression. | 

### Return type

[**GetHealthCheck200Response**](GetHealthCheck200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubscriber

> CreateSubscriber200Response CreateSubscriber(ctx).NewSubscriber(newSubscriber).Execute()





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
    newSubscriber := *openapiclient.NewNewSubscriber() // NewSubscriber | new subscriber info (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscribersApi.CreateSubscriber(context.Background()).NewSubscriber(newSubscriber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.CreateSubscriber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubscriber`: CreateSubscriber200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.CreateSubscriber`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newSubscriber** | [**NewSubscriber**](NewSubscriber.md) | new subscriber info | 

### Return type

[**CreateSubscriber200Response**](CreateSubscriber200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscriberBouncesById

> GetHealthCheck200Response DeleteSubscriberBouncesById(ctx, id).Execute()





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
    id := int32(56) // int32 | subscriber id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscribersApi.DeleteSubscriberBouncesById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.DeleteSubscriberBouncesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubscriberBouncesById`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.DeleteSubscriberBouncesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | subscriber id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriberBouncesByIdRequest struct via the builder pattern


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


## DeleteSubscriberById

> GetHealthCheck200Response DeleteSubscriberById(ctx, id).Execute()





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
    id := int32(56) // int32 | The id value of the subscriber you want to get.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscribersApi.DeleteSubscriberById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.DeleteSubscriberById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubscriberById`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.DeleteSubscriberById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id value of the subscriber you want to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriberByIdRequest struct via the builder pattern


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


## DeleteSubscriberByList

> GetHealthCheck200Response DeleteSubscriberByList(ctx).Id(id).Execute()





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
    id := "id_example" // string | subscriber id/s to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscribersApi.DeleteSubscriberByList(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.DeleteSubscriberByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubscriberByList`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.DeleteSubscriberByList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriberByListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | subscriber id/s to be deleted | 

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


## DeleteSubscriberByQuery

> GetHealthCheck200Response DeleteSubscriberByQuery(ctx).Body(body).Execute()





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
    body := "body_example" // string | Arbitrary SQL expression. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscribersApi.DeleteSubscriberByQuery(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.DeleteSubscriberByQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubscriberByQuery`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.DeleteSubscriberByQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriberByQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | Arbitrary SQL expression. | 

### Return type

[**GetHealthCheck200Response**](GetHealthCheck200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSubscriberDataByID

> SubscriberData ExportSubscriberDataByID(ctx, id).Execute()





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
    id := int32(56) // int32 | The id value of subscriber profile you want to export

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscribersApi.ExportSubscriberDataByID(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.ExportSubscriberDataByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSubscriberDataByID`: SubscriberData
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.ExportSubscriberDataByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id value of subscriber profile you want to export | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportSubscriberDataByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriberData**](SubscriberData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriberBouncesById

> GetSubscriberBouncesById200Response GetSubscriberBouncesById(ctx, id).Execute()





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
    id := int32(56) // int32 | subscriber id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscribersApi.GetSubscriberBouncesById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.GetSubscriberBouncesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriberBouncesById`: GetSubscriberBouncesById200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.GetSubscriberBouncesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | subscriber id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriberBouncesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSubscriberBouncesById200Response**](GetSubscriberBouncesById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriberById

> CreateSubscriber200Response GetSubscriberById(ctx, id).Execute()





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
    id := int32(56) // int32 | The id value of the subscriber you want to get.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscribersApi.GetSubscriberById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.GetSubscriberById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriberById`: CreateSubscriber200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.GetSubscriberById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id value of the subscriber you want to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriberByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateSubscriber200Response**](CreateSubscriber200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscribers

> GetSubscribers200Response GetSubscribers(ctx).Page(page).PerPage(perPage).Query(query).Execute()





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
    page := int32(56) // int32 | number of records to skip (optional)
    perPage := int32(56) // int32 | max number of records to return per page (optional)
    query := "query_example" // string | query subscribers with an SQL expression. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscribersApi.GetSubscribers(context.Background()).Page(page).PerPage(perPage).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.GetSubscribers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscribers`: GetSubscribers200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.GetSubscribers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | number of records to skip | 
 **perPage** | **int32** | max number of records to return per page | 
 **query** | **string** | query subscribers with an SQL expression. | 

### Return type

[**GetSubscribers200Response**](GetSubscribers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageBlocklistBySubscriberList

> GetHealthCheck200Response ManageBlocklistBySubscriberList(ctx).SubscriberQueryRequest(subscriberQueryRequest).Execute()





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
    subscriberQueryRequest := *openapiclient.NewSubscriberQueryRequest() // SubscriberQueryRequest | The list of subscribers to blocklist (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscribersApi.ManageBlocklistBySubscriberList(context.Background()).SubscriberQueryRequest(subscriberQueryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.ManageBlocklistBySubscriberList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageBlocklistBySubscriberList`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.ManageBlocklistBySubscriberList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManageBlocklistBySubscriberListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriberQueryRequest** | [**SubscriberQueryRequest**](SubscriberQueryRequest.md) | The list of subscribers to blocklist | 

### Return type

[**GetHealthCheck200Response**](GetHealthCheck200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageBlocklistSubscribersById

> GetHealthCheck200Response ManageBlocklistSubscribersById(ctx, id).SubscriberQueryRequest(subscriberQueryRequest).Execute()





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
    id := int32(56) // int32 | The id value of the subscriber you want to blocklist.
    subscriberQueryRequest := *openapiclient.NewSubscriberQueryRequest() // SubscriberQueryRequest | The id of subscriber to add or remove (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscribersApi.ManageBlocklistSubscribersById(context.Background(), id).SubscriberQueryRequest(subscriberQueryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.ManageBlocklistSubscribersById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageBlocklistSubscribersById`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.ManageBlocklistSubscribersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id value of the subscriber you want to blocklist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageBlocklistSubscribersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriberQueryRequest** | [**SubscriberQueryRequest**](SubscriberQueryRequest.md) | The id of subscriber to add or remove | 

### Return type

[**GetHealthCheck200Response**](GetHealthCheck200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageSubscriberListById

> GetHealthCheck200Response ManageSubscriberListById(ctx, id).SubscriberQueryRequest(subscriberQueryRequest).Execute()





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
    id := int32(56) // int32 | The id of list you want to update
    subscriberQueryRequest := *openapiclient.NewSubscriberQueryRequest() // SubscriberQueryRequest | The list of subscribers to add or remove (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscribersApi.ManageSubscriberListById(context.Background(), id).SubscriberQueryRequest(subscriberQueryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.ManageSubscriberListById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageSubscriberListById`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.ManageSubscriberListById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id of list you want to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageSubscriberListByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriberQueryRequest** | [**SubscriberQueryRequest**](SubscriberQueryRequest.md) | The list of subscribers to add or remove | 

### Return type

[**GetHealthCheck200Response**](GetHealthCheck200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageSubscriberLists

> GetHealthCheck200Response ManageSubscriberLists(ctx).SubscriberQueryRequest(subscriberQueryRequest).Execute()





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
    subscriberQueryRequest := *openapiclient.NewSubscriberQueryRequest() // SubscriberQueryRequest | The list of subscribers details to add or remove (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscribersApi.ManageSubscriberLists(context.Background()).SubscriberQueryRequest(subscriberQueryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.ManageSubscriberLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageSubscriberLists`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.ManageSubscriberLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManageSubscriberListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriberQueryRequest** | [**SubscriberQueryRequest**](SubscriberQueryRequest.md) | The list of subscribers details to add or remove | 

### Return type

[**GetHealthCheck200Response**](GetHealthCheck200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageSubscriberListsByQuery

> GetHealthCheck200Response ManageSubscriberListsByQuery(ctx).Body(body).Execute()





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
    body := "body_example" // string | Arbitrary SQL expression. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscribersApi.ManageSubscriberListsByQuery(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.ManageSubscriberListsByQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageSubscriberListsByQuery`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.ManageSubscriberListsByQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManageSubscriberListsByQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | Arbitrary SQL expression. | 

### Return type

[**GetHealthCheck200Response**](GetHealthCheck200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriberSendOptinById

> GetHealthCheck200Response SubscriberSendOptinById(ctx, id).Execute()





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
    id := int32(56) // int32 | sends an optin confirmation e-mail to a subscriber

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscribersApi.SubscriberSendOptinById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.SubscriberSendOptinById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriberSendOptinById`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.SubscriberSendOptinById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | sends an optin confirmation e-mail to a subscriber | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriberSendOptinByIdRequest struct via the builder pattern


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


## UpdateSubscriberById

> CreateSubscriber200Response UpdateSubscriberById(ctx, id).UpdateSubscriber(updateSubscriber).Execute()





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
    id := int32(56) // int32 | The id of subscriber to update
    updateSubscriber := *openapiclient.NewUpdateSubscriber() // UpdateSubscriber | new subscriber info (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscribersApi.UpdateSubscriberById(context.Background(), id).UpdateSubscriber(updateSubscriber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.UpdateSubscriberById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubscriberById`: CreateSubscriber200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.UpdateSubscriberById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id of subscriber to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriberByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSubscriber** | [**UpdateSubscriber**](UpdateSubscriber.md) | new subscriber info | 

### Return type

[**CreateSubscriber200Response**](CreateSubscriber200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

