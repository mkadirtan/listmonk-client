# \SettingsApi

All URIs are relative to *http://localhost:9000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettings**](SettingsApi.md#GetSettings) | **Get** /settings | 
[**TestSMTPSettings**](SettingsApi.md#TestSMTPSettings) | **Post** /settings/smtp/test | 
[**UpdateSettings**](SettingsApi.md#UpdateSettings) | **Put** /settings | 



## GetSettings

> GetSettings200Response GetSettings(ctx).Execute()





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
    resp, r, err := apiClient.SettingsApi.GetSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSettings`: GetSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsRequest struct via the builder pattern


### Return type

[**GetSettings200Response**](GetSettings200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSMTPSettings

> GetHealthCheck200Response TestSMTPSettings(ctx).SMTPTest(sMTPTest).Execute()





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
    sMTPTest := *openapiclient.NewSMTPTest() // SMTPTest | updated SMTP settings field values (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.TestSMTPSettings(context.Background()).SMTPTest(sMTPTest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.TestSMTPSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestSMTPSettings`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.TestSMTPSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestSMTPSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sMTPTest** | [**SMTPTest**](SMTPTest.md) | updated SMTP settings field values | 

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


## UpdateSettings

> GetHealthCheck200Response UpdateSettings(ctx).Settings(settings).Execute()





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
    settings := *openapiclient.NewSettings() // Settings | updated settings field values (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.UpdateSettings(context.Background()).Settings(settings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.UpdateSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSettings`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.UpdateSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settings** | [**Settings**](Settings.md) | updated settings field values | 

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

