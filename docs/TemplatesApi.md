# \TemplatesApi

All URIs are relative to *http://localhost:9000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTemplateById**](TemplatesApi.md#DeleteTemplateById) | **Delete** /templates/{id} | 
[**GetTemplateById**](TemplatesApi.md#GetTemplateById) | **Get** /templates/{id} | 
[**GetTemplates**](TemplatesApi.md#GetTemplates) | **Get** /templates | 
[**PreviewTemplate**](TemplatesApi.md#PreviewTemplate) | **Post** /templates/preview | 
[**PreviewTemplateById**](TemplatesApi.md#PreviewTemplateById) | **Get** /templates/{id}/preview | 
[**UpdateTemplateById**](TemplatesApi.md#UpdateTemplateById) | **Put** /templates/{id}/default | 



## DeleteTemplateById

> GetHealthCheck200Response DeleteTemplateById(ctx, id).Execute()





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
    id := int32(56) // int32 | The id value of the template you want to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.DeleteTemplateById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.DeleteTemplateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTemplateById`: GetHealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.DeleteTemplateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id value of the template you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateByIdRequest struct via the builder pattern


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


## GetTemplateById

> GetTemplateById200Response GetTemplateById(ctx, id).NoBody(noBody).Execute()





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
    id := int32(56) // int32 | The id value of the template you want to get.
    noBody := true // bool | boolean flag for response with/without body (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.GetTemplateById(context.Background(), id).NoBody(noBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.GetTemplateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplateById`: GetTemplateById200Response
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.GetTemplateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id value of the template you want to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noBody** | **bool** | boolean flag for response with/without body | 

### Return type

[**GetTemplateById200Response**](GetTemplateById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplates

> GetTemplates200Response GetTemplates(ctx).NoBody(noBody).Execute()





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
    noBody := true // bool | boolean flag for response with/without body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.GetTemplates(context.Background()).NoBody(noBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.GetTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplates`: GetTemplates200Response
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.GetTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noBody** | **bool** | boolean flag for response with/without body | 

### Return type

[**GetTemplates200Response**](GetTemplates200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewTemplate

> string PreviewTemplate(ctx).TemplateType(templateType).Body(body).Execute()





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
    templateType := "templateType_example" // string | type of template (optional)
    body := "body_example" // string | template body (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.PreviewTemplate(context.Background()).TemplateType(templateType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.PreviewTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewTemplate`: string
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.PreviewTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateType** | **string** | type of template | 
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


## PreviewTemplateById

> string PreviewTemplateById(ctx, id).TemplateType(templateType).Body(body).Execute()





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
    id := int32(56) // int32 | The id value of the template you want to get.
    templateType := "templateType_example" // string | type of template (optional)
    body := "body_example" // string | template body (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.PreviewTemplateById(context.Background(), id).TemplateType(templateType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.PreviewTemplateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewTemplateById`: string
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.PreviewTemplateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id value of the template you want to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewTemplateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateType** | **string** | type of template | 
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


## UpdateTemplateById

> Template UpdateTemplateById(ctx, id).Execute()





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
    id := int32(56) // int32 | The id value of the template you want to set to the default template.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.UpdateTemplateById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.UpdateTemplateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTemplateById`: Template
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.UpdateTemplateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id value of the template you want to set to the default template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Template**](Template.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

