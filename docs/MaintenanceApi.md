# \MaintenanceApi

All URIs are relative to *http://localhost:9000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCampaignAnalyticsByType**](MaintenanceApi.md#DeleteCampaignAnalyticsByType) | **Delete** /maintenance/analytics/{type} | 
[**DeleteGCSubscribers**](MaintenanceApi.md#DeleteGCSubscribers) | **Delete** /maintenance/subscribers/{type} | 
[**DeleteUnconfirmedSubscriptions**](MaintenanceApi.md#DeleteUnconfirmedSubscriptions) | **Delete** /maintenance/subscriptions/unconfirmed | 



## DeleteCampaignAnalyticsByType

> GetHealthCheck200Response DeleteCampaignAnalyticsByType(ctx, type_).BeforeDate(beforeDate).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    type_ := "type__example" // string | type of GC collected subscribers
    beforeDate := time.Now() // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaintenanceApi.DeleteCampaignAnalyticsByType(context.Background(), type_).BeforeDate(beforeDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceApi.DeleteCampaignAnalyticsByType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCampaignAnalyticsByType`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceApi.DeleteCampaignAnalyticsByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | type of GC collected subscribers | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignAnalyticsByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **beforeDate** | **string** |  | 

### Return type

[**GetHealthCheck200Response**](GetHealthCheck200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGCSubscribers

> DeleteGCSubscribers200Response DeleteGCSubscribers(ctx, type_).Execute()





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
    type_ := "type__example" // string | type of GC collected subscribers

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaintenanceApi.DeleteGCSubscribers(context.Background(), type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceApi.DeleteGCSubscribers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteGCSubscribers`: DeleteGCSubscribers200Response
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceApi.DeleteGCSubscribers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | type of GC collected subscribers | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGCSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteGCSubscribers200Response**](DeleteGCSubscribers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUnconfirmedSubscriptions

> DeleteGCSubscribers200Response DeleteUnconfirmedSubscriptions(ctx).BeforeDate(beforeDate).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    beforeDate := time.Now() // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaintenanceApi.DeleteUnconfirmedSubscriptions(context.Background()).BeforeDate(beforeDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceApi.DeleteUnconfirmedSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUnconfirmedSubscriptions`: DeleteGCSubscribers200Response
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceApi.DeleteUnconfirmedSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUnconfirmedSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **beforeDate** | **string** |  | 

### Return type

[**DeleteGCSubscribers200Response**](DeleteGCSubscribers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

