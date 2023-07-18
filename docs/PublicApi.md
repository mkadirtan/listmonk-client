# \PublicApi

All URIs are relative to *http://localhost:9000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPublicLists**](PublicApi.md#GetPublicLists) | **Get** /public/lists | 
[**HandlePublicSubscription**](PublicApi.md#HandlePublicSubscription) | **Post** /public/subscription | 



## GetPublicLists

> []GetPublicLists200ResponseInner GetPublicLists(ctx).Execute()





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
    resp, r, err := apiClient.PublicApi.GetPublicLists(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.GetPublicLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPublicLists`: []GetPublicLists200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.GetPublicLists`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicListsRequest struct via the builder pattern


### Return type

[**[]GetPublicLists200ResponseInner**](GetPublicLists200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandlePublicSubscription

> HandlePublicSubscription200Response HandlePublicSubscription(ctx).HandlePublicSubscriptionRequest(handlePublicSubscriptionRequest).Execute()





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
    handlePublicSubscriptionRequest := *openapiclient.NewHandlePublicSubscriptionRequest() // HandlePublicSubscriptionRequest | subscription request parameters (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicApi.HandlePublicSubscription(context.Background()).HandlePublicSubscriptionRequest(handlePublicSubscriptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.HandlePublicSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HandlePublicSubscription`: HandlePublicSubscription200Response
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.HandlePublicSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandlePublicSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **handlePublicSubscriptionRequest** | [**HandlePublicSubscriptionRequest**](HandlePublicSubscriptionRequest.md) | subscription request parameters | 

### Return type

[**HandlePublicSubscription200Response**](HandlePublicSubscription200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

