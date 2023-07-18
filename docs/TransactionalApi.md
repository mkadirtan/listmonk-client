# \TransactionalApi

All URIs are relative to *http://localhost:9000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactWithSubscriber**](TransactionalApi.md#TransactWithSubscriber) | **Post** /tx | 



## TransactWithSubscriber

> GetHealthCheck200Response TransactWithSubscriber(ctx).TransactionalMessage(transactionalMessage).Execute()





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
    transactionalMessage := *openapiclient.NewTransactionalMessage() // TransactionalMessage | email message to a subscriber (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionalApi.TransactWithSubscriber(context.Background()).TransactionalMessage(transactionalMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionalApi.TransactWithSubscriber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactWithSubscriber`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `TransactionalApi.TransactWithSubscriber`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactWithSubscriberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionalMessage** | [**TransactionalMessage**](TransactionalMessage.md) | email message to a subscriber | 

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

