# \CampaignsApi

All URIs are relative to *http://localhost:9000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCampaign**](CampaignsApi.md#CreateCampaign) | **Post** /campaigns | 
[**CreateCampaignContentById**](CampaignsApi.md#CreateCampaignContentById) | **Post** /campaigns/{id}/content | 
[**DeleteCampaignById**](CampaignsApi.md#DeleteCampaignById) | **Delete** /campaigns/{id} | 
[**GetCampaignAnalytics**](CampaignsApi.md#GetCampaignAnalytics) | **Get** /campaigns/analytics/{type} | 
[**GetCampaignById**](CampaignsApi.md#GetCampaignById) | **Get** /campaigns/{id} | 
[**GetCampaigns**](CampaignsApi.md#GetCampaigns) | **Get** /campaigns | 
[**GetRunningCampaignStats**](CampaignsApi.md#GetRunningCampaignStats) | **Get** /campaigns/running/stats | 
[**PreviewCampaignById**](CampaignsApi.md#PreviewCampaignById) | **Get** /campaigns/{id}/preview | 
[**PreviewCampaignTextById**](CampaignsApi.md#PreviewCampaignTextById) | **Post** /campaigns/{id}/text | 
[**TestCampaignById**](CampaignsApi.md#TestCampaignById) | **Post** /campaigns/{id}/test | 
[**UpdateCampaignArchiveById**](CampaignsApi.md#UpdateCampaignArchiveById) | **Put** /campaigns/{id}/archive | 
[**UpdateCampaignById**](CampaignsApi.md#UpdateCampaignById) | **Put** /campaigns/{id} | 
[**UpdateCampaignStatusById**](CampaignsApi.md#UpdateCampaignStatusById) | **Put** /campaigns/{id}/status | 
[**UpdatePreviewCampaignById**](CampaignsApi.md#UpdatePreviewCampaignById) | **Post** /campaigns/{id}/preview | 



## CreateCampaign

> CreateCampaign200Response CreateCampaign(ctx).CampaignRequest(campaignRequest).Execute()





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
    campaignRequest := *openapiclient.NewCampaignRequest() // CampaignRequest | new campaign info (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.CreateCampaign(context.Background()).CampaignRequest(campaignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.CreateCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCampaign`: CreateCampaign200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.CreateCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignRequest** | [**CampaignRequest**](CampaignRequest.md) | new campaign info | 

### Return type

[**CreateCampaign200Response**](CreateCampaign200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCampaignContentById

> GetImportSubscriberStats200Response CreateCampaignContentById(ctx, id).CampaignContentRequest(campaignContentRequest).Execute()





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
    id := int32(56) // int32 | ID of campaign that you choose to create content
    campaignContentRequest := *openapiclient.NewCampaignContentRequest() // CampaignContentRequest | updated campaign content (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.CreateCampaignContentById(context.Background(), id).CampaignContentRequest(campaignContentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.CreateCampaignContentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCampaignContentById`: GetImportSubscriberStats200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.CreateCampaignContentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of campaign that you choose to create content | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignContentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignContentRequest** | [**CampaignContentRequest**](CampaignContentRequest.md) | updated campaign content | 

### Return type

[**GetImportSubscriberStats200Response**](GetImportSubscriberStats200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaignById

> GetHealthCheck200Response DeleteCampaignById(ctx, id).Execute()





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
    id := int32(56) // int32 | The id value of the campaign you want to get.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.DeleteCampaignById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.DeleteCampaignById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCampaignById`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.DeleteCampaignById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id value of the campaign you want to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignByIdRequest struct via the builder pattern


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


## GetCampaignAnalytics

> GetCampaignAnalytics200Response GetCampaignAnalytics(ctx, type_).From(from).To(to).Id(id).Execute()





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
    type_ := "type__example" // string | type of stats, either links, view, click or bounce
    from := "from_example" // string | start value of date range
    to := "to_example" // string | end value of date range
    id := "id_example" // string | campaign id/s to retrive view counts (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.GetCampaignAnalytics(context.Background(), type_).From(from).To(to).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.GetCampaignAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaignAnalytics`: GetCampaignAnalytics200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.GetCampaignAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | type of stats, either links, view, click or bounce | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | start value of date range | 
 **to** | **string** | end value of date range | 
 **id** | **string** | campaign id/s to retrive view counts | 

### Return type

[**GetCampaignAnalytics200Response**](GetCampaignAnalytics200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignById

> GetCampaignById200Response GetCampaignById(ctx, id).NoBody(noBody).Execute()





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
    id := int32(56) // int32 | The id value of the campaign you want to get.
    noBody := true // bool | boolean flag for response with/without body (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.GetCampaignById(context.Background(), id).NoBody(noBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.GetCampaignById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaignById`: GetCampaignById200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.GetCampaignById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id value of the campaign you want to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noBody** | **bool** | boolean flag for response with/without body | 

### Return type

[**GetCampaignById200Response**](GetCampaignById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaigns

> GetCampaigns200Response GetCampaigns(ctx).Status(status).NoBody(noBody).Page(page).PerPage(perPage).Query(query).OrderBy(orderBy).Order(order).Execute()





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
    status := []string{"Status_example"} // []string | status flag of campaign (optional)
    noBody := true // bool | boolean flag for response with/without body (optional)
    page := int32(56) // int32 | total number of pages (optional)
    perPage := int32(56) // int32 | number of items per page (optional)
    query := "query_example" // string | Optional string to search a list by name. (optional)
    orderBy := "orderBy_example" // string | Field to sort results by. name|status|created_at|updated_at (optional)
    order := "order_example" // string | ASC|DESC Sort by ascending or descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.GetCampaigns(context.Background()).Status(status).NoBody(noBody).Page(page).PerPage(perPage).Query(query).OrderBy(orderBy).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.GetCampaigns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaigns`: GetCampaigns200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.GetCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **[]string** | status flag of campaign | 
 **noBody** | **bool** | boolean flag for response with/without body | 
 **page** | **int32** | total number of pages | 
 **perPage** | **int32** | number of items per page | 
 **query** | **string** | Optional string to search a list by name. | 
 **orderBy** | **string** | Field to sort results by. name|status|created_at|updated_at | 
 **order** | **string** | ASC|DESC Sort by ascending or descending order. | 

### Return type

[**GetCampaigns200Response**](GetCampaigns200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunningCampaignStats

> GetRunningCampaignStats200Response GetRunningCampaignStats(ctx).Execute()





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
    resp, r, err := apiClient.CampaignsApi.GetRunningCampaignStats(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.GetRunningCampaignStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRunningCampaignStats`: GetRunningCampaignStats200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.GetRunningCampaignStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunningCampaignStatsRequest struct via the builder pattern


### Return type

[**GetRunningCampaignStats200Response**](GetRunningCampaignStats200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewCampaignById

> string PreviewCampaignById(ctx, id).TemplateId(templateId).Execute()





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
    id := int32(56) // int32 | The id value of the campaign you want to get the preview of
    templateId := int32(56) // int32 | template id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.PreviewCampaignById(context.Background(), id).TemplateId(templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.PreviewCampaignById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewCampaignById`: string
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.PreviewCampaignById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id value of the campaign you want to get the preview of | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewCampaignByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateId** | **int32** | template id | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewCampaignTextById

> string PreviewCampaignTextById(ctx, id).TemplateId(templateId).ContentType(contentType).Body(body).Execute()





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
    id := int32(56) // int32 | The id value of the campaign you want to get the preview of
    templateId := int32(56) // int32 | template id (optional)
    contentType := "contentType_example" // string | content type (optional)
    body := "body_example" // string | campaign body (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.PreviewCampaignTextById(context.Background(), id).TemplateId(templateId).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.PreviewCampaignTextById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewCampaignTextById`: string
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.PreviewCampaignTextById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id value of the campaign you want to get the preview of | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewCampaignTextByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateId** | **int32** | template id | 
 **contentType** | **string** | content type | 
 **body** | **string** | campaign body | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestCampaignById

> GetHealthCheck200Response TestCampaignById(ctx, id).TemplateId(templateId).Execute()





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
    id := int32(56) // int32 | ID of campaign that you want to test
    templateId := int32(56) // int32 | template id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.TestCampaignById(context.Background(), id).TemplateId(templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.TestCampaignById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestCampaignById`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.TestCampaignById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of campaign that you want to test | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestCampaignByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateId** | **int32** | template id | 

### Return type

[**GetHealthCheck200Response**](GetHealthCheck200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaignArchiveById

> GetHealthCheck200Response UpdateCampaignArchiveById(ctx, id).UpdateCampaignArchiveByIdRequest(updateCampaignArchiveByIdRequest).Execute()





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
    id := int32(56) // int32 | The id value of the campaign you want to get the preview of
    updateCampaignArchiveByIdRequest := *openapiclient.NewUpdateCampaignArchiveByIdRequest() // UpdateCampaignArchiveByIdRequest | archive campaign related parameters (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.UpdateCampaignArchiveById(context.Background(), id).UpdateCampaignArchiveByIdRequest(updateCampaignArchiveByIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.UpdateCampaignArchiveById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCampaignArchiveById`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.UpdateCampaignArchiveById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id value of the campaign you want to get the preview of | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignArchiveByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCampaignArchiveByIdRequest** | [**UpdateCampaignArchiveByIdRequest**](UpdateCampaignArchiveByIdRequest.md) | archive campaign related parameters | 

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


## UpdateCampaignById

> CreateCampaign200Response UpdateCampaignById(ctx, id).CampaignRequest(campaignRequest).Execute()





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
    id := int32(56) // int32 | the id value of campaign you want to update
    campaignRequest := *openapiclient.NewCampaignRequest() // CampaignRequest | updated campaign fields (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.UpdateCampaignById(context.Background(), id).CampaignRequest(campaignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.UpdateCampaignById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCampaignById`: CreateCampaign200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.UpdateCampaignById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | the id value of campaign you want to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignRequest** | [**CampaignRequest**](CampaignRequest.md) | updated campaign fields | 

### Return type

[**CreateCampaign200Response**](CreateCampaign200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaignStatusById

> GetCampaignById200Response UpdateCampaignStatusById(ctx, id).UpdateCampaignStatusByIdRequest(updateCampaignStatusByIdRequest).Execute()





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
    id := int32(56) // int32 | The id value of the campaign you want to get the preview of
    updateCampaignStatusByIdRequest := *openapiclient.NewUpdateCampaignStatusByIdRequest() // UpdateCampaignStatusByIdRequest | campaign status update (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.UpdateCampaignStatusById(context.Background(), id).UpdateCampaignStatusByIdRequest(updateCampaignStatusByIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.UpdateCampaignStatusById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCampaignStatusById`: GetCampaignById200Response
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.UpdateCampaignStatusById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id value of the campaign you want to get the preview of | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignStatusByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCampaignStatusByIdRequest** | [**UpdateCampaignStatusByIdRequest**](UpdateCampaignStatusByIdRequest.md) | campaign status update | 

### Return type

[**GetCampaignById200Response**](GetCampaignById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePreviewCampaignById

> string UpdatePreviewCampaignById(ctx, id).TemplateId(templateId).ContentType(contentType).Body(body).Execute()





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
    id := int32(56) // int32 | The id value of the campaign you want to get the preview of
    templateId := int32(56) // int32 | template id (optional)
    contentType := "contentType_example" // string | content type (optional)
    body := "body_example" // string | template body (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.UpdatePreviewCampaignById(context.Background(), id).TemplateId(templateId).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.UpdatePreviewCampaignById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePreviewCampaignById`: string
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.UpdatePreviewCampaignById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id value of the campaign you want to get the preview of | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePreviewCampaignByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateId** | **int32** | template id | 
 **contentType** | **string** | content type | 
 **body** | **string** | template body | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

