# \MiscellaneousApi

All URIs are relative to *http://localhost:9000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDashboardCharts**](MiscellaneousApi.md#GetDashboardCharts) | **Get** /dashboard/charts | 
[**GetDashboardCounts**](MiscellaneousApi.md#GetDashboardCounts) | **Get** /dashboard/counts | 
[**GetHealthCheck**](MiscellaneousApi.md#GetHealthCheck) | **Get** /health | 
[**GetI18nLang**](MiscellaneousApi.md#GetI18nLang) | **Get** /lang/{lang} | 
[**GetServerConfig**](MiscellaneousApi.md#GetServerConfig) | **Get** /config | 



## GetDashboardCharts

> GetDashboardCharts200Response GetDashboardCharts(ctx).Execute()





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
    resp, r, err := apiClient.MiscellaneousApi.GetDashboardCharts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousApi.GetDashboardCharts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboardCharts`: GetDashboardCharts200Response
    fmt.Fprintf(os.Stdout, "Response from `MiscellaneousApi.GetDashboardCharts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardChartsRequest struct via the builder pattern


### Return type

[**GetDashboardCharts200Response**](GetDashboardCharts200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardCounts

> GetDashboardCounts200Response GetDashboardCounts(ctx).Execute()





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
    resp, r, err := apiClient.MiscellaneousApi.GetDashboardCounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousApi.GetDashboardCounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboardCounts`: GetDashboardCounts200Response
    fmt.Fprintf(os.Stdout, "Response from `MiscellaneousApi.GetDashboardCounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardCountsRequest struct via the builder pattern


### Return type

[**GetDashboardCounts200Response**](GetDashboardCounts200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHealthCheck

> GetHealthCheck200Response GetHealthCheck(ctx).Execute()





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
    resp, r, err := apiClient.MiscellaneousApi.GetHealthCheck(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousApi.GetHealthCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealthCheck`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `MiscellaneousApi.GetHealthCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthCheckRequest struct via the builder pattern


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


## GetI18nLang

> GetI18nLang200Response GetI18nLang(ctx, lang).Execute()





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
    lang := "lang_example" // string | JSON language pack required

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MiscellaneousApi.GetI18nLang(context.Background(), lang).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousApi.GetI18nLang``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetI18nLang`: GetI18nLang200Response
    fmt.Fprintf(os.Stdout, "Response from `MiscellaneousApi.GetI18nLang`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lang** | **string** | JSON language pack required | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetI18nLangRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetI18nLang200Response**](GetI18nLang200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerConfig

> GetServerConfig200Response GetServerConfig(ctx).Execute()





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
    resp, r, err := apiClient.MiscellaneousApi.GetServerConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousApi.GetServerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerConfig`: GetServerConfig200Response
    fmt.Fprintf(os.Stdout, "Response from `MiscellaneousApi.GetServerConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerConfigRequest struct via the builder pattern


### Return type

[**GetServerConfig200Response**](GetServerConfig200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

