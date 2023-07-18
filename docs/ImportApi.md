# \ImportApi

All URIs are relative to *http://localhost:9000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImportSubscriberStats**](ImportApi.md#GetImportSubscriberStats) | **Get** /import/subscribers/logs | 
[**GetImportSubscribers**](ImportApi.md#GetImportSubscribers) | **Get** /import/subscribers | 
[**ImportSubscribers**](ImportApi.md#ImportSubscribers) | **Post** /import/subscribers | 
[**StopImportSubscribers**](ImportApi.md#StopImportSubscribers) | **Delete** /import/subscribers | 



## GetImportSubscriberStats

> GetImportSubscriberStats200Response GetImportSubscriberStats(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportApi.GetImportSubscriberStats(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportApi.GetImportSubscriberStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImportSubscriberStats`: GetImportSubscriberStats200Response
    fmt.Fprintf(os.Stdout, "Response from `ImportApi.GetImportSubscriberStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportSubscriberStatsRequest struct via the builder pattern


### Return type

[**GetImportSubscriberStats200Response**](GetImportSubscriberStats200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImportSubscribers

> GetImportSubscribers200Response GetImportSubscribers(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportApi.GetImportSubscribers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportApi.GetImportSubscribers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImportSubscribers`: GetImportSubscribers200Response
    fmt.Fprintf(os.Stdout, "Response from `ImportApi.GetImportSubscribers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportSubscribersRequest struct via the builder pattern


### Return type

[**GetImportSubscribers200Response**](GetImportSubscribers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSubscribers

> GetImportSubscribers200Response ImportSubscribers(ctx).ImportSubscribersRequest(importSubscribersRequest).Execute()





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
    importSubscribersRequest := *openapiclient.NewImportSubscribersRequest() // ImportSubscribersRequest | uploads and bulk imports of compressed CSV files (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportApi.ImportSubscribers(context.Background()).ImportSubscribersRequest(importSubscribersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportApi.ImportSubscribers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportSubscribers`: GetImportSubscribers200Response
    fmt.Fprintf(os.Stdout, "Response from `ImportApi.ImportSubscribers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importSubscribersRequest** | [**ImportSubscribersRequest**](ImportSubscribersRequest.md) | uploads and bulk imports of compressed CSV files | 

### Return type

[**GetImportSubscribers200Response**](GetImportSubscribers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopImportSubscribers

> GetImportSubscribers200Response StopImportSubscribers(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportApi.StopImportSubscribers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportApi.StopImportSubscribers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopImportSubscribers`: GetImportSubscribers200Response
    fmt.Fprintf(os.Stdout, "Response from `ImportApi.StopImportSubscribers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStopImportSubscribersRequest struct via the builder pattern


### Return type

[**GetImportSubscribers200Response**](GetImportSubscribers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

