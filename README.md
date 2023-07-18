# Go API client for openapi

The API collection for listmonk

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:9000/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdminApi* | [**ReloadApp**](docs/AdminApi.md#reloadapp) | **Post** /admin/reload | 
*BouncesApi* | [**DeleteBounceById**](docs/BouncesApi.md#deletebouncebyid) | **Delete** /bounces/{id} | 
*BouncesApi* | [**DeleteBounces**](docs/BouncesApi.md#deletebounces) | **Delete** /bounces | 
*BouncesApi* | [**GetBounceById**](docs/BouncesApi.md#getbouncebyid) | **Get** /bounces/{id} | 
*BouncesApi* | [**GetBounces**](docs/BouncesApi.md#getbounces) | **Get** /bounces | 
*CampaignsApi* | [**CreateCampaign**](docs/CampaignsApi.md#createcampaign) | **Post** /campaigns | 
*CampaignsApi* | [**CreateCampaignContentById**](docs/CampaignsApi.md#createcampaigncontentbyid) | **Post** /campaigns/{id}/content | 
*CampaignsApi* | [**DeleteCampaignById**](docs/CampaignsApi.md#deletecampaignbyid) | **Delete** /campaigns/{id} | 
*CampaignsApi* | [**GetCampaignAnalytics**](docs/CampaignsApi.md#getcampaignanalytics) | **Get** /campaigns/analytics/{type} | 
*CampaignsApi* | [**GetCampaignById**](docs/CampaignsApi.md#getcampaignbyid) | **Get** /campaigns/{id} | 
*CampaignsApi* | [**GetCampaigns**](docs/CampaignsApi.md#getcampaigns) | **Get** /campaigns | 
*CampaignsApi* | [**GetRunningCampaignStats**](docs/CampaignsApi.md#getrunningcampaignstats) | **Get** /campaigns/running/stats | 
*CampaignsApi* | [**PreviewCampaignById**](docs/CampaignsApi.md#previewcampaignbyid) | **Get** /campaigns/{id}/preview | 
*CampaignsApi* | [**PreviewCampaignTextById**](docs/CampaignsApi.md#previewcampaigntextbyid) | **Post** /campaigns/{id}/text | 
*CampaignsApi* | [**TestCampaignById**](docs/CampaignsApi.md#testcampaignbyid) | **Post** /campaigns/{id}/test | 
*CampaignsApi* | [**UpdateCampaignArchiveById**](docs/CampaignsApi.md#updatecampaignarchivebyid) | **Put** /campaigns/{id}/archive | 
*CampaignsApi* | [**UpdateCampaignById**](docs/CampaignsApi.md#updatecampaignbyid) | **Put** /campaigns/{id} | 
*CampaignsApi* | [**UpdateCampaignStatusById**](docs/CampaignsApi.md#updatecampaignstatusbyid) | **Put** /campaigns/{id}/status | 
*CampaignsApi* | [**UpdatePreviewCampaignById**](docs/CampaignsApi.md#updatepreviewcampaignbyid) | **Post** /campaigns/{id}/preview | 
*ImportApi* | [**GetImportSubscriberStats**](docs/ImportApi.md#getimportsubscriberstats) | **Get** /import/subscribers/logs | 
*ImportApi* | [**GetImportSubscribers**](docs/ImportApi.md#getimportsubscribers) | **Get** /import/subscribers | 
*ImportApi* | [**ImportSubscribers**](docs/ImportApi.md#importsubscribers) | **Post** /import/subscribers | 
*ImportApi* | [**StopImportSubscribers**](docs/ImportApi.md#stopimportsubscribers) | **Delete** /import/subscribers | 
*ListsApi* | [**CreateList**](docs/ListsApi.md#createlist) | **Post** /lists | 
*ListsApi* | [**DeleteListById**](docs/ListsApi.md#deletelistbyid) | **Delete** /lists/{list_id} | 
*ListsApi* | [**GetListById**](docs/ListsApi.md#getlistbyid) | **Get** /lists/{list_id} | 
*ListsApi* | [**GetLists**](docs/ListsApi.md#getlists) | **Get** /lists | 
*ListsApi* | [**UpdateListById**](docs/ListsApi.md#updatelistbyid) | **Put** /lists/{list_id} | 
*LogsApi* | [**GetLogs**](docs/LogsApi.md#getlogs) | **Get** /logs | 
*MaintenanceApi* | [**DeleteCampaignAnalyticsByType**](docs/MaintenanceApi.md#deletecampaignanalyticsbytype) | **Delete** /maintenance/analytics/{type} | 
*MaintenanceApi* | [**DeleteGCSubscribers**](docs/MaintenanceApi.md#deletegcsubscribers) | **Delete** /maintenance/subscribers/{type} | 
*MaintenanceApi* | [**DeleteUnconfirmedSubscriptions**](docs/MaintenanceApi.md#deleteunconfirmedsubscriptions) | **Delete** /maintenance/subscriptions/unconfirmed | 
*MediaApi* | [**DeleteMediaById**](docs/MediaApi.md#deletemediabyid) | **Delete** /media/{id} | 
*MediaApi* | [**GetMedia**](docs/MediaApi.md#getmedia) | **Get** /media | 
*MediaApi* | [**GetMediaById**](docs/MediaApi.md#getmediabyid) | **Get** /media/{id} | 
*MediaApi* | [**UploadMedia**](docs/MediaApi.md#uploadmedia) | **Post** /media | 
*MiscellaneousApi* | [**GetDashboardCharts**](docs/MiscellaneousApi.md#getdashboardcharts) | **Get** /dashboard/charts | 
*MiscellaneousApi* | [**GetDashboardCounts**](docs/MiscellaneousApi.md#getdashboardcounts) | **Get** /dashboard/counts | 
*MiscellaneousApi* | [**GetHealthCheck**](docs/MiscellaneousApi.md#gethealthcheck) | **Get** /health | 
*MiscellaneousApi* | [**GetI18nLang**](docs/MiscellaneousApi.md#geti18nlang) | **Get** /lang/{lang} | 
*MiscellaneousApi* | [**GetServerConfig**](docs/MiscellaneousApi.md#getserverconfig) | **Get** /config | 
*PublicApi* | [**GetPublicLists**](docs/PublicApi.md#getpubliclists) | **Get** /public/lists | 
*PublicApi* | [**HandlePublicSubscription**](docs/PublicApi.md#handlepublicsubscription) | **Post** /public/subscription | 
*SettingsApi* | [**GetSettings**](docs/SettingsApi.md#getsettings) | **Get** /settings | 
*SettingsApi* | [**TestSMTPSettings**](docs/SettingsApi.md#testsmtpsettings) | **Post** /settings/smtp/test | 
*SettingsApi* | [**UpdateSettings**](docs/SettingsApi.md#updatesettings) | **Put** /settings | 
*SubscribersApi* | [**BlocklistSubscribersQuery**](docs/SubscribersApi.md#blocklistsubscribersquery) | **Put** /subscribers/query/blocklist | 
*SubscribersApi* | [**CreateSubscriber**](docs/SubscribersApi.md#createsubscriber) | **Post** /subscribers | 
*SubscribersApi* | [**DeleteSubscriberBouncesById**](docs/SubscribersApi.md#deletesubscriberbouncesbyid) | **Delete** /subscribers/{id}/bounces | 
*SubscribersApi* | [**DeleteSubscriberById**](docs/SubscribersApi.md#deletesubscriberbyid) | **Delete** /subscribers/{id} | 
*SubscribersApi* | [**DeleteSubscriberByList**](docs/SubscribersApi.md#deletesubscriberbylist) | **Delete** /subscribers | 
*SubscribersApi* | [**DeleteSubscriberByQuery**](docs/SubscribersApi.md#deletesubscriberbyquery) | **Post** /subscribers/query/delete | 
*SubscribersApi* | [**ExportSubscriberDataByID**](docs/SubscribersApi.md#exportsubscriberdatabyid) | **Get** /subscribers/{id}/export | 
*SubscribersApi* | [**GetSubscriberBouncesById**](docs/SubscribersApi.md#getsubscriberbouncesbyid) | **Get** /subscribers/{id}/bounces | 
*SubscribersApi* | [**GetSubscriberById**](docs/SubscribersApi.md#getsubscriberbyid) | **Get** /subscribers/{id} | 
*SubscribersApi* | [**GetSubscribers**](docs/SubscribersApi.md#getsubscribers) | **Get** /subscribers | 
*SubscribersApi* | [**ManageBlocklistBySubscriberList**](docs/SubscribersApi.md#manageblocklistbysubscriberlist) | **Put** /subscribers/blocklist | 
*SubscribersApi* | [**ManageBlocklistSubscribersById**](docs/SubscribersApi.md#manageblocklistsubscribersbyid) | **Put** /subscribers/{id}/blocklist | 
*SubscribersApi* | [**ManageSubscriberListById**](docs/SubscribersApi.md#managesubscriberlistbyid) | **Put** /subscribers/lists/{id} | 
*SubscribersApi* | [**ManageSubscriberLists**](docs/SubscribersApi.md#managesubscriberlists) | **Put** /subscribers/lists | 
*SubscribersApi* | [**ManageSubscriberListsByQuery**](docs/SubscribersApi.md#managesubscriberlistsbyquery) | **Put** /subscribers/query/lists | 
*SubscribersApi* | [**SubscriberSendOptinById**](docs/SubscribersApi.md#subscribersendoptinbyid) | **Post** /subscribers/{id}/optin | 
*SubscribersApi* | [**UpdateSubscriberById**](docs/SubscribersApi.md#updatesubscriberbyid) | **Put** /subscribers/{id} | 
*TemplatesApi* | [**DeleteTemplateById**](docs/TemplatesApi.md#deletetemplatebyid) | **Delete** /templates/{id} | 
*TemplatesApi* | [**GetTemplateById**](docs/TemplatesApi.md#gettemplatebyid) | **Get** /templates/{id} | 
*TemplatesApi* | [**GetTemplates**](docs/TemplatesApi.md#gettemplates) | **Get** /templates | 
*TemplatesApi* | [**PreviewTemplate**](docs/TemplatesApi.md#previewtemplate) | **Post** /templates/preview | 
*TemplatesApi* | [**PreviewTemplateById**](docs/TemplatesApi.md#previewtemplatebyid) | **Get** /templates/{id}/preview | 
*TemplatesApi* | [**UpdateTemplateById**](docs/TemplatesApi.md#updatetemplatebyid) | **Put** /templates/{id}/default | 
*TransactionalApi* | [**TransactWithSubscriber**](docs/TransactionalApi.md#transactwithsubscriber) | **Post** /tx | 


## Documentation For Models

 - [Bounce](docs/Bounce.md)
 - [BounceResultsInner](docs/BounceResultsInner.md)
 - [BounceResultsInnerCampaign](docs/BounceResultsInnerCampaign.md)
 - [Campaign](docs/Campaign.md)
 - [CampaignAnalyticsCount](docs/CampaignAnalyticsCount.md)
 - [CampaignContentRequest](docs/CampaignContentRequest.md)
 - [CampaignRequest](docs/CampaignRequest.md)
 - [CampaignRequestSendAt](docs/CampaignRequestSendAt.md)
 - [CampaignStats](docs/CampaignStats.md)
 - [CampaignUpdate](docs/CampaignUpdate.md)
 - [CreateCampaign200Response](docs/CreateCampaign200Response.md)
 - [CreateList200Response](docs/CreateList200Response.md)
 - [CreateSubscriber200Response](docs/CreateSubscriber200Response.md)
 - [DashboardChart](docs/DashboardChart.md)
 - [DashboardChartLinkClicksInner](docs/DashboardChartLinkClicksInner.md)
 - [DashboardCount](docs/DashboardCount.md)
 - [DashboardCountData](docs/DashboardCountData.md)
 - [DashboardCountDataCampaigns](docs/DashboardCountDataCampaigns.md)
 - [DashboardCountDataCampaignsByStatus](docs/DashboardCountDataCampaignsByStatus.md)
 - [DashboardCountDataLists](docs/DashboardCountDataLists.md)
 - [DashboardCountDataSubscribers](docs/DashboardCountDataSubscribers.md)
 - [DeleteGCSubscribers200Response](docs/DeleteGCSubscribers200Response.md)
 - [DeleteGCSubscribers200ResponseData](docs/DeleteGCSubscribers200ResponseData.md)
 - [GetBounceById200Response](docs/GetBounceById200Response.md)
 - [GetBounces200Response](docs/GetBounces200Response.md)
 - [GetBounces200ResponseData](docs/GetBounces200ResponseData.md)
 - [GetCampaignAnalytics200Response](docs/GetCampaignAnalytics200Response.md)
 - [GetCampaignById200Response](docs/GetCampaignById200Response.md)
 - [GetCampaigns200Response](docs/GetCampaigns200Response.md)
 - [GetCampaigns200ResponseData](docs/GetCampaigns200ResponseData.md)
 - [GetDashboardCharts200Response](docs/GetDashboardCharts200Response.md)
 - [GetDashboardCounts200Response](docs/GetDashboardCounts200Response.md)
 - [GetHealthCheck200Response](docs/GetHealthCheck200Response.md)
 - [GetI18nLang200Response](docs/GetI18nLang200Response.md)
 - [GetImportSubscriberStats200Response](docs/GetImportSubscriberStats200Response.md)
 - [GetImportSubscribers200Response](docs/GetImportSubscribers200Response.md)
 - [GetLists200Response](docs/GetLists200Response.md)
 - [GetLists200ResponseData](docs/GetLists200ResponseData.md)
 - [GetLogs200Response](docs/GetLogs200Response.md)
 - [GetMedia200Response](docs/GetMedia200Response.md)
 - [GetPublicLists200ResponseInner](docs/GetPublicLists200ResponseInner.md)
 - [GetRunningCampaignStats200Response](docs/GetRunningCampaignStats200Response.md)
 - [GetServerConfig200Response](docs/GetServerConfig200Response.md)
 - [GetSettings200Response](docs/GetSettings200Response.md)
 - [GetSubscriberBouncesById200Response](docs/GetSubscriberBouncesById200Response.md)
 - [GetSubscribers200Response](docs/GetSubscribers200Response.md)
 - [GetSubscribers200ResponseData](docs/GetSubscribers200ResponseData.md)
 - [GetTemplateById200Response](docs/GetTemplateById200Response.md)
 - [GetTemplates200Response](docs/GetTemplates200Response.md)
 - [HandlePublicSubscription200Response](docs/HandlePublicSubscription200Response.md)
 - [HandlePublicSubscriptionRequest](docs/HandlePublicSubscriptionRequest.md)
 - [ImportStatus](docs/ImportStatus.md)
 - [ImportStatusData](docs/ImportStatusData.md)
 - [ImportSubscribersRequest](docs/ImportSubscribersRequest.md)
 - [LanguagePack](docs/LanguagePack.md)
 - [LanguagePackData](docs/LanguagePackData.md)
 - [List](docs/List.md)
 - [MailBoxBounces](docs/MailBoxBounces.md)
 - [MediaFileObject](docs/MediaFileObject.md)
 - [NewList](docs/NewList.md)
 - [NewSubscriber](docs/NewSubscriber.md)
 - [NewSubscriberAttribs](docs/NewSubscriberAttribs.md)
 - [NewSubscriberAttribsStack](docs/NewSubscriberAttribsStack.md)
 - [SMTPSettings](docs/SMTPSettings.md)
 - [SMTPTest](docs/SMTPTest.md)
 - [ServerConfig](docs/ServerConfig.md)
 - [ServerConfigData](docs/ServerConfigData.md)
 - [ServerConfigDataLangsInner](docs/ServerConfigDataLangsInner.md)
 - [Settings](docs/Settings.md)
 - [Subscriber](docs/Subscriber.md)
 - [SubscriberData](docs/SubscriberData.md)
 - [SubscriberListsInner](docs/SubscriberListsInner.md)
 - [SubscriberProfile](docs/SubscriberProfile.md)
 - [SubscriberProfileAttribs](docs/SubscriberProfileAttribs.md)
 - [SubscriberQueryRequest](docs/SubscriberQueryRequest.md)
 - [Subscriptions](docs/Subscriptions.md)
 - [Template](docs/Template.md)
 - [TransactionalMessage](docs/TransactionalMessage.md)
 - [UpdateCampaignArchiveByIdRequest](docs/UpdateCampaignArchiveByIdRequest.md)
 - [UpdateCampaignStatusByIdRequest](docs/UpdateCampaignStatusByIdRequest.md)
 - [UpdateSubscriber](docs/UpdateSubscriber.md)
 - [UploadMedia200Response](docs/UploadMedia200Response.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


