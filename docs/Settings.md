# Settings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppSiteName** | Pointer to **string** |  | [optional] 
**AppRootUrl** | Pointer to **string** |  | [optional] 
**AppLogoUrl** | Pointer to **string** |  | [optional] 
**AppFaviconUrl** | Pointer to **string** |  | [optional] 
**AppFromEmail** | Pointer to **string** |  | [optional] 
**AppNotifyEmails** | Pointer to **[]string** |  | [optional] 
**AppEnablePublicSubscriptionPage** | Pointer to **bool** |  | [optional] 
**AppEnablePublicArchive** | Pointer to **bool** |  | [optional] 
**AppSendOptinConfirmation** | Pointer to **bool** |  | [optional] 
**AppCheckUpdates** | Pointer to **bool** |  | [optional] 
**AppLang** | Pointer to **string** |  | [optional] 
**AppBatchSize** | Pointer to **int32** |  | [optional] 
**AppConcurrency** | Pointer to **int32** |  | [optional] 
**AppMaxSendErrors** | Pointer to **int32** |  | [optional] 
**AppMessageRate** | Pointer to **int32** |  | [optional] 
**AppMessageSlidingWindow** | Pointer to **bool** |  | [optional] 
**AppMessageSlidingWindowDuration** | Pointer to **string** |  | [optional] 
**AppMessageSlidingWindowRate** | Pointer to **int32** |  | [optional] 
**PrivacyIndividualTracking** | Pointer to **bool** |  | [optional] 
**PrivacyUnsubscribeHeader** | Pointer to **bool** |  | [optional] 
**PrivacyAllowBlocklist** | Pointer to **bool** |  | [optional] 
**PrivacyAllowPreferences** | Pointer to **bool** |  | [optional] 
**PrivacyAllowExport** | Pointer to **bool** |  | [optional] 
**PrivacyAllowWipe** | Pointer to **bool** |  | [optional] 
**PrivacyExportable** | Pointer to **[]string** |  | [optional] 
**PrivacyDomainBlocklist** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UploadProvider** | Pointer to **string** |  | [optional] 
**UploadFilesystemUploadPath** | Pointer to **string** |  | [optional] 
**UploadFilesystemUploadUri** | Pointer to **string** |  | [optional] 
**UploadS3Url** | Pointer to **string** |  | [optional] 
**UploadS3PublicUrl** | Pointer to **string** |  | [optional] 
**UploadS3AwsAccessKeyId** | Pointer to **string** |  | [optional] 
**UploadS3AwsDefaultRegion** | Pointer to **string** |  | [optional] 
**UploadS3Bucket** | Pointer to **string** |  | [optional] 
**UploadS3BucketDomain** | Pointer to **string** |  | [optional] 
**UploadS3BucketPath** | Pointer to **string** |  | [optional] 
**UploadS3BucketType** | Pointer to **string** |  | [optional] 
**UploadS3Expiry** | Pointer to **string** |  | [optional] 
**Smtp** | Pointer to [**[]SMTPSettings**](SMTPSettings.md) |  | [optional] 
**Messengers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**BounceEnabled** | Pointer to **bool** |  | [optional] 
**BounceWebhooksEnabled** | Pointer to **bool** |  | [optional] 
**BounceCount** | Pointer to **int32** |  | [optional] 
**BounceAction** | Pointer to **string** |  | [optional] 
**BounceSesEnabled** | Pointer to **bool** |  | [optional] 
**BounceSendgridEnabled** | Pointer to **bool** |  | [optional] 
**BounceSendgridKey** | Pointer to **string** |  | [optional] 
**BounceMailboxes** | Pointer to [**[]MailBoxBounces**](MailBoxBounces.md) |  | [optional] 
**AppearanceAdminCustomCss** | Pointer to **string** |  | [optional] 
**AppearanceAdminCustomJs** | Pointer to **string** |  | [optional] 
**AppearancePublicCustomCss** | Pointer to **string** |  | [optional] 
**AppearancePublicCustomJs** | Pointer to **string** |  | [optional] 

## Methods

### NewSettings

`func NewSettings() *Settings`

NewSettings instantiates a new Settings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsWithDefaults

`func NewSettingsWithDefaults() *Settings`

NewSettingsWithDefaults instantiates a new Settings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppSiteName

`func (o *Settings) GetAppSiteName() string`

GetAppSiteName returns the AppSiteName field if non-nil, zero value otherwise.

### GetAppSiteNameOk

`func (o *Settings) GetAppSiteNameOk() (*string, bool)`

GetAppSiteNameOk returns a tuple with the AppSiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSiteName

`func (o *Settings) SetAppSiteName(v string)`

SetAppSiteName sets AppSiteName field to given value.

### HasAppSiteName

`func (o *Settings) HasAppSiteName() bool`

HasAppSiteName returns a boolean if a field has been set.

### GetAppRootUrl

`func (o *Settings) GetAppRootUrl() string`

GetAppRootUrl returns the AppRootUrl field if non-nil, zero value otherwise.

### GetAppRootUrlOk

`func (o *Settings) GetAppRootUrlOk() (*string, bool)`

GetAppRootUrlOk returns a tuple with the AppRootUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRootUrl

`func (o *Settings) SetAppRootUrl(v string)`

SetAppRootUrl sets AppRootUrl field to given value.

### HasAppRootUrl

`func (o *Settings) HasAppRootUrl() bool`

HasAppRootUrl returns a boolean if a field has been set.

### GetAppLogoUrl

`func (o *Settings) GetAppLogoUrl() string`

GetAppLogoUrl returns the AppLogoUrl field if non-nil, zero value otherwise.

### GetAppLogoUrlOk

`func (o *Settings) GetAppLogoUrlOk() (*string, bool)`

GetAppLogoUrlOk returns a tuple with the AppLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLogoUrl

`func (o *Settings) SetAppLogoUrl(v string)`

SetAppLogoUrl sets AppLogoUrl field to given value.

### HasAppLogoUrl

`func (o *Settings) HasAppLogoUrl() bool`

HasAppLogoUrl returns a boolean if a field has been set.

### GetAppFaviconUrl

`func (o *Settings) GetAppFaviconUrl() string`

GetAppFaviconUrl returns the AppFaviconUrl field if non-nil, zero value otherwise.

### GetAppFaviconUrlOk

`func (o *Settings) GetAppFaviconUrlOk() (*string, bool)`

GetAppFaviconUrlOk returns a tuple with the AppFaviconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppFaviconUrl

`func (o *Settings) SetAppFaviconUrl(v string)`

SetAppFaviconUrl sets AppFaviconUrl field to given value.

### HasAppFaviconUrl

`func (o *Settings) HasAppFaviconUrl() bool`

HasAppFaviconUrl returns a boolean if a field has been set.

### GetAppFromEmail

`func (o *Settings) GetAppFromEmail() string`

GetAppFromEmail returns the AppFromEmail field if non-nil, zero value otherwise.

### GetAppFromEmailOk

`func (o *Settings) GetAppFromEmailOk() (*string, bool)`

GetAppFromEmailOk returns a tuple with the AppFromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppFromEmail

`func (o *Settings) SetAppFromEmail(v string)`

SetAppFromEmail sets AppFromEmail field to given value.

### HasAppFromEmail

`func (o *Settings) HasAppFromEmail() bool`

HasAppFromEmail returns a boolean if a field has been set.

### GetAppNotifyEmails

`func (o *Settings) GetAppNotifyEmails() []string`

GetAppNotifyEmails returns the AppNotifyEmails field if non-nil, zero value otherwise.

### GetAppNotifyEmailsOk

`func (o *Settings) GetAppNotifyEmailsOk() (*[]string, bool)`

GetAppNotifyEmailsOk returns a tuple with the AppNotifyEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppNotifyEmails

`func (o *Settings) SetAppNotifyEmails(v []string)`

SetAppNotifyEmails sets AppNotifyEmails field to given value.

### HasAppNotifyEmails

`func (o *Settings) HasAppNotifyEmails() bool`

HasAppNotifyEmails returns a boolean if a field has been set.

### GetAppEnablePublicSubscriptionPage

`func (o *Settings) GetAppEnablePublicSubscriptionPage() bool`

GetAppEnablePublicSubscriptionPage returns the AppEnablePublicSubscriptionPage field if non-nil, zero value otherwise.

### GetAppEnablePublicSubscriptionPageOk

`func (o *Settings) GetAppEnablePublicSubscriptionPageOk() (*bool, bool)`

GetAppEnablePublicSubscriptionPageOk returns a tuple with the AppEnablePublicSubscriptionPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEnablePublicSubscriptionPage

`func (o *Settings) SetAppEnablePublicSubscriptionPage(v bool)`

SetAppEnablePublicSubscriptionPage sets AppEnablePublicSubscriptionPage field to given value.

### HasAppEnablePublicSubscriptionPage

`func (o *Settings) HasAppEnablePublicSubscriptionPage() bool`

HasAppEnablePublicSubscriptionPage returns a boolean if a field has been set.

### GetAppEnablePublicArchive

`func (o *Settings) GetAppEnablePublicArchive() bool`

GetAppEnablePublicArchive returns the AppEnablePublicArchive field if non-nil, zero value otherwise.

### GetAppEnablePublicArchiveOk

`func (o *Settings) GetAppEnablePublicArchiveOk() (*bool, bool)`

GetAppEnablePublicArchiveOk returns a tuple with the AppEnablePublicArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEnablePublicArchive

`func (o *Settings) SetAppEnablePublicArchive(v bool)`

SetAppEnablePublicArchive sets AppEnablePublicArchive field to given value.

### HasAppEnablePublicArchive

`func (o *Settings) HasAppEnablePublicArchive() bool`

HasAppEnablePublicArchive returns a boolean if a field has been set.

### GetAppSendOptinConfirmation

`func (o *Settings) GetAppSendOptinConfirmation() bool`

GetAppSendOptinConfirmation returns the AppSendOptinConfirmation field if non-nil, zero value otherwise.

### GetAppSendOptinConfirmationOk

`func (o *Settings) GetAppSendOptinConfirmationOk() (*bool, bool)`

GetAppSendOptinConfirmationOk returns a tuple with the AppSendOptinConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSendOptinConfirmation

`func (o *Settings) SetAppSendOptinConfirmation(v bool)`

SetAppSendOptinConfirmation sets AppSendOptinConfirmation field to given value.

### HasAppSendOptinConfirmation

`func (o *Settings) HasAppSendOptinConfirmation() bool`

HasAppSendOptinConfirmation returns a boolean if a field has been set.

### GetAppCheckUpdates

`func (o *Settings) GetAppCheckUpdates() bool`

GetAppCheckUpdates returns the AppCheckUpdates field if non-nil, zero value otherwise.

### GetAppCheckUpdatesOk

`func (o *Settings) GetAppCheckUpdatesOk() (*bool, bool)`

GetAppCheckUpdatesOk returns a tuple with the AppCheckUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCheckUpdates

`func (o *Settings) SetAppCheckUpdates(v bool)`

SetAppCheckUpdates sets AppCheckUpdates field to given value.

### HasAppCheckUpdates

`func (o *Settings) HasAppCheckUpdates() bool`

HasAppCheckUpdates returns a boolean if a field has been set.

### GetAppLang

`func (o *Settings) GetAppLang() string`

GetAppLang returns the AppLang field if non-nil, zero value otherwise.

### GetAppLangOk

`func (o *Settings) GetAppLangOk() (*string, bool)`

GetAppLangOk returns a tuple with the AppLang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLang

`func (o *Settings) SetAppLang(v string)`

SetAppLang sets AppLang field to given value.

### HasAppLang

`func (o *Settings) HasAppLang() bool`

HasAppLang returns a boolean if a field has been set.

### GetAppBatchSize

`func (o *Settings) GetAppBatchSize() int32`

GetAppBatchSize returns the AppBatchSize field if non-nil, zero value otherwise.

### GetAppBatchSizeOk

`func (o *Settings) GetAppBatchSizeOk() (*int32, bool)`

GetAppBatchSizeOk returns a tuple with the AppBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppBatchSize

`func (o *Settings) SetAppBatchSize(v int32)`

SetAppBatchSize sets AppBatchSize field to given value.

### HasAppBatchSize

`func (o *Settings) HasAppBatchSize() bool`

HasAppBatchSize returns a boolean if a field has been set.

### GetAppConcurrency

`func (o *Settings) GetAppConcurrency() int32`

GetAppConcurrency returns the AppConcurrency field if non-nil, zero value otherwise.

### GetAppConcurrencyOk

`func (o *Settings) GetAppConcurrencyOk() (*int32, bool)`

GetAppConcurrencyOk returns a tuple with the AppConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConcurrency

`func (o *Settings) SetAppConcurrency(v int32)`

SetAppConcurrency sets AppConcurrency field to given value.

### HasAppConcurrency

`func (o *Settings) HasAppConcurrency() bool`

HasAppConcurrency returns a boolean if a field has been set.

### GetAppMaxSendErrors

`func (o *Settings) GetAppMaxSendErrors() int32`

GetAppMaxSendErrors returns the AppMaxSendErrors field if non-nil, zero value otherwise.

### GetAppMaxSendErrorsOk

`func (o *Settings) GetAppMaxSendErrorsOk() (*int32, bool)`

GetAppMaxSendErrorsOk returns a tuple with the AppMaxSendErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppMaxSendErrors

`func (o *Settings) SetAppMaxSendErrors(v int32)`

SetAppMaxSendErrors sets AppMaxSendErrors field to given value.

### HasAppMaxSendErrors

`func (o *Settings) HasAppMaxSendErrors() bool`

HasAppMaxSendErrors returns a boolean if a field has been set.

### GetAppMessageRate

`func (o *Settings) GetAppMessageRate() int32`

GetAppMessageRate returns the AppMessageRate field if non-nil, zero value otherwise.

### GetAppMessageRateOk

`func (o *Settings) GetAppMessageRateOk() (*int32, bool)`

GetAppMessageRateOk returns a tuple with the AppMessageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppMessageRate

`func (o *Settings) SetAppMessageRate(v int32)`

SetAppMessageRate sets AppMessageRate field to given value.

### HasAppMessageRate

`func (o *Settings) HasAppMessageRate() bool`

HasAppMessageRate returns a boolean if a field has been set.

### GetAppMessageSlidingWindow

`func (o *Settings) GetAppMessageSlidingWindow() bool`

GetAppMessageSlidingWindow returns the AppMessageSlidingWindow field if non-nil, zero value otherwise.

### GetAppMessageSlidingWindowOk

`func (o *Settings) GetAppMessageSlidingWindowOk() (*bool, bool)`

GetAppMessageSlidingWindowOk returns a tuple with the AppMessageSlidingWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppMessageSlidingWindow

`func (o *Settings) SetAppMessageSlidingWindow(v bool)`

SetAppMessageSlidingWindow sets AppMessageSlidingWindow field to given value.

### HasAppMessageSlidingWindow

`func (o *Settings) HasAppMessageSlidingWindow() bool`

HasAppMessageSlidingWindow returns a boolean if a field has been set.

### GetAppMessageSlidingWindowDuration

`func (o *Settings) GetAppMessageSlidingWindowDuration() string`

GetAppMessageSlidingWindowDuration returns the AppMessageSlidingWindowDuration field if non-nil, zero value otherwise.

### GetAppMessageSlidingWindowDurationOk

`func (o *Settings) GetAppMessageSlidingWindowDurationOk() (*string, bool)`

GetAppMessageSlidingWindowDurationOk returns a tuple with the AppMessageSlidingWindowDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppMessageSlidingWindowDuration

`func (o *Settings) SetAppMessageSlidingWindowDuration(v string)`

SetAppMessageSlidingWindowDuration sets AppMessageSlidingWindowDuration field to given value.

### HasAppMessageSlidingWindowDuration

`func (o *Settings) HasAppMessageSlidingWindowDuration() bool`

HasAppMessageSlidingWindowDuration returns a boolean if a field has been set.

### GetAppMessageSlidingWindowRate

`func (o *Settings) GetAppMessageSlidingWindowRate() int32`

GetAppMessageSlidingWindowRate returns the AppMessageSlidingWindowRate field if non-nil, zero value otherwise.

### GetAppMessageSlidingWindowRateOk

`func (o *Settings) GetAppMessageSlidingWindowRateOk() (*int32, bool)`

GetAppMessageSlidingWindowRateOk returns a tuple with the AppMessageSlidingWindowRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppMessageSlidingWindowRate

`func (o *Settings) SetAppMessageSlidingWindowRate(v int32)`

SetAppMessageSlidingWindowRate sets AppMessageSlidingWindowRate field to given value.

### HasAppMessageSlidingWindowRate

`func (o *Settings) HasAppMessageSlidingWindowRate() bool`

HasAppMessageSlidingWindowRate returns a boolean if a field has been set.

### GetPrivacyIndividualTracking

`func (o *Settings) GetPrivacyIndividualTracking() bool`

GetPrivacyIndividualTracking returns the PrivacyIndividualTracking field if non-nil, zero value otherwise.

### GetPrivacyIndividualTrackingOk

`func (o *Settings) GetPrivacyIndividualTrackingOk() (*bool, bool)`

GetPrivacyIndividualTrackingOk returns a tuple with the PrivacyIndividualTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyIndividualTracking

`func (o *Settings) SetPrivacyIndividualTracking(v bool)`

SetPrivacyIndividualTracking sets PrivacyIndividualTracking field to given value.

### HasPrivacyIndividualTracking

`func (o *Settings) HasPrivacyIndividualTracking() bool`

HasPrivacyIndividualTracking returns a boolean if a field has been set.

### GetPrivacyUnsubscribeHeader

`func (o *Settings) GetPrivacyUnsubscribeHeader() bool`

GetPrivacyUnsubscribeHeader returns the PrivacyUnsubscribeHeader field if non-nil, zero value otherwise.

### GetPrivacyUnsubscribeHeaderOk

`func (o *Settings) GetPrivacyUnsubscribeHeaderOk() (*bool, bool)`

GetPrivacyUnsubscribeHeaderOk returns a tuple with the PrivacyUnsubscribeHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyUnsubscribeHeader

`func (o *Settings) SetPrivacyUnsubscribeHeader(v bool)`

SetPrivacyUnsubscribeHeader sets PrivacyUnsubscribeHeader field to given value.

### HasPrivacyUnsubscribeHeader

`func (o *Settings) HasPrivacyUnsubscribeHeader() bool`

HasPrivacyUnsubscribeHeader returns a boolean if a field has been set.

### GetPrivacyAllowBlocklist

`func (o *Settings) GetPrivacyAllowBlocklist() bool`

GetPrivacyAllowBlocklist returns the PrivacyAllowBlocklist field if non-nil, zero value otherwise.

### GetPrivacyAllowBlocklistOk

`func (o *Settings) GetPrivacyAllowBlocklistOk() (*bool, bool)`

GetPrivacyAllowBlocklistOk returns a tuple with the PrivacyAllowBlocklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyAllowBlocklist

`func (o *Settings) SetPrivacyAllowBlocklist(v bool)`

SetPrivacyAllowBlocklist sets PrivacyAllowBlocklist field to given value.

### HasPrivacyAllowBlocklist

`func (o *Settings) HasPrivacyAllowBlocklist() bool`

HasPrivacyAllowBlocklist returns a boolean if a field has been set.

### GetPrivacyAllowPreferences

`func (o *Settings) GetPrivacyAllowPreferences() bool`

GetPrivacyAllowPreferences returns the PrivacyAllowPreferences field if non-nil, zero value otherwise.

### GetPrivacyAllowPreferencesOk

`func (o *Settings) GetPrivacyAllowPreferencesOk() (*bool, bool)`

GetPrivacyAllowPreferencesOk returns a tuple with the PrivacyAllowPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyAllowPreferences

`func (o *Settings) SetPrivacyAllowPreferences(v bool)`

SetPrivacyAllowPreferences sets PrivacyAllowPreferences field to given value.

### HasPrivacyAllowPreferences

`func (o *Settings) HasPrivacyAllowPreferences() bool`

HasPrivacyAllowPreferences returns a boolean if a field has been set.

### GetPrivacyAllowExport

`func (o *Settings) GetPrivacyAllowExport() bool`

GetPrivacyAllowExport returns the PrivacyAllowExport field if non-nil, zero value otherwise.

### GetPrivacyAllowExportOk

`func (o *Settings) GetPrivacyAllowExportOk() (*bool, bool)`

GetPrivacyAllowExportOk returns a tuple with the PrivacyAllowExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyAllowExport

`func (o *Settings) SetPrivacyAllowExport(v bool)`

SetPrivacyAllowExport sets PrivacyAllowExport field to given value.

### HasPrivacyAllowExport

`func (o *Settings) HasPrivacyAllowExport() bool`

HasPrivacyAllowExport returns a boolean if a field has been set.

### GetPrivacyAllowWipe

`func (o *Settings) GetPrivacyAllowWipe() bool`

GetPrivacyAllowWipe returns the PrivacyAllowWipe field if non-nil, zero value otherwise.

### GetPrivacyAllowWipeOk

`func (o *Settings) GetPrivacyAllowWipeOk() (*bool, bool)`

GetPrivacyAllowWipeOk returns a tuple with the PrivacyAllowWipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyAllowWipe

`func (o *Settings) SetPrivacyAllowWipe(v bool)`

SetPrivacyAllowWipe sets PrivacyAllowWipe field to given value.

### HasPrivacyAllowWipe

`func (o *Settings) HasPrivacyAllowWipe() bool`

HasPrivacyAllowWipe returns a boolean if a field has been set.

### GetPrivacyExportable

`func (o *Settings) GetPrivacyExportable() []string`

GetPrivacyExportable returns the PrivacyExportable field if non-nil, zero value otherwise.

### GetPrivacyExportableOk

`func (o *Settings) GetPrivacyExportableOk() (*[]string, bool)`

GetPrivacyExportableOk returns a tuple with the PrivacyExportable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyExportable

`func (o *Settings) SetPrivacyExportable(v []string)`

SetPrivacyExportable sets PrivacyExportable field to given value.

### HasPrivacyExportable

`func (o *Settings) HasPrivacyExportable() bool`

HasPrivacyExportable returns a boolean if a field has been set.

### GetPrivacyDomainBlocklist

`func (o *Settings) GetPrivacyDomainBlocklist() []map[string]interface{}`

GetPrivacyDomainBlocklist returns the PrivacyDomainBlocklist field if non-nil, zero value otherwise.

### GetPrivacyDomainBlocklistOk

`func (o *Settings) GetPrivacyDomainBlocklistOk() (*[]map[string]interface{}, bool)`

GetPrivacyDomainBlocklistOk returns a tuple with the PrivacyDomainBlocklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyDomainBlocklist

`func (o *Settings) SetPrivacyDomainBlocklist(v []map[string]interface{})`

SetPrivacyDomainBlocklist sets PrivacyDomainBlocklist field to given value.

### HasPrivacyDomainBlocklist

`func (o *Settings) HasPrivacyDomainBlocklist() bool`

HasPrivacyDomainBlocklist returns a boolean if a field has been set.

### GetUploadProvider

`func (o *Settings) GetUploadProvider() string`

GetUploadProvider returns the UploadProvider field if non-nil, zero value otherwise.

### GetUploadProviderOk

`func (o *Settings) GetUploadProviderOk() (*string, bool)`

GetUploadProviderOk returns a tuple with the UploadProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadProvider

`func (o *Settings) SetUploadProvider(v string)`

SetUploadProvider sets UploadProvider field to given value.

### HasUploadProvider

`func (o *Settings) HasUploadProvider() bool`

HasUploadProvider returns a boolean if a field has been set.

### GetUploadFilesystemUploadPath

`func (o *Settings) GetUploadFilesystemUploadPath() string`

GetUploadFilesystemUploadPath returns the UploadFilesystemUploadPath field if non-nil, zero value otherwise.

### GetUploadFilesystemUploadPathOk

`func (o *Settings) GetUploadFilesystemUploadPathOk() (*string, bool)`

GetUploadFilesystemUploadPathOk returns a tuple with the UploadFilesystemUploadPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadFilesystemUploadPath

`func (o *Settings) SetUploadFilesystemUploadPath(v string)`

SetUploadFilesystemUploadPath sets UploadFilesystemUploadPath field to given value.

### HasUploadFilesystemUploadPath

`func (o *Settings) HasUploadFilesystemUploadPath() bool`

HasUploadFilesystemUploadPath returns a boolean if a field has been set.

### GetUploadFilesystemUploadUri

`func (o *Settings) GetUploadFilesystemUploadUri() string`

GetUploadFilesystemUploadUri returns the UploadFilesystemUploadUri field if non-nil, zero value otherwise.

### GetUploadFilesystemUploadUriOk

`func (o *Settings) GetUploadFilesystemUploadUriOk() (*string, bool)`

GetUploadFilesystemUploadUriOk returns a tuple with the UploadFilesystemUploadUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadFilesystemUploadUri

`func (o *Settings) SetUploadFilesystemUploadUri(v string)`

SetUploadFilesystemUploadUri sets UploadFilesystemUploadUri field to given value.

### HasUploadFilesystemUploadUri

`func (o *Settings) HasUploadFilesystemUploadUri() bool`

HasUploadFilesystemUploadUri returns a boolean if a field has been set.

### GetUploadS3Url

`func (o *Settings) GetUploadS3Url() string`

GetUploadS3Url returns the UploadS3Url field if non-nil, zero value otherwise.

### GetUploadS3UrlOk

`func (o *Settings) GetUploadS3UrlOk() (*string, bool)`

GetUploadS3UrlOk returns a tuple with the UploadS3Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadS3Url

`func (o *Settings) SetUploadS3Url(v string)`

SetUploadS3Url sets UploadS3Url field to given value.

### HasUploadS3Url

`func (o *Settings) HasUploadS3Url() bool`

HasUploadS3Url returns a boolean if a field has been set.

### GetUploadS3PublicUrl

`func (o *Settings) GetUploadS3PublicUrl() string`

GetUploadS3PublicUrl returns the UploadS3PublicUrl field if non-nil, zero value otherwise.

### GetUploadS3PublicUrlOk

`func (o *Settings) GetUploadS3PublicUrlOk() (*string, bool)`

GetUploadS3PublicUrlOk returns a tuple with the UploadS3PublicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadS3PublicUrl

`func (o *Settings) SetUploadS3PublicUrl(v string)`

SetUploadS3PublicUrl sets UploadS3PublicUrl field to given value.

### HasUploadS3PublicUrl

`func (o *Settings) HasUploadS3PublicUrl() bool`

HasUploadS3PublicUrl returns a boolean if a field has been set.

### GetUploadS3AwsAccessKeyId

`func (o *Settings) GetUploadS3AwsAccessKeyId() string`

GetUploadS3AwsAccessKeyId returns the UploadS3AwsAccessKeyId field if non-nil, zero value otherwise.

### GetUploadS3AwsAccessKeyIdOk

`func (o *Settings) GetUploadS3AwsAccessKeyIdOk() (*string, bool)`

GetUploadS3AwsAccessKeyIdOk returns a tuple with the UploadS3AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadS3AwsAccessKeyId

`func (o *Settings) SetUploadS3AwsAccessKeyId(v string)`

SetUploadS3AwsAccessKeyId sets UploadS3AwsAccessKeyId field to given value.

### HasUploadS3AwsAccessKeyId

`func (o *Settings) HasUploadS3AwsAccessKeyId() bool`

HasUploadS3AwsAccessKeyId returns a boolean if a field has been set.

### GetUploadS3AwsDefaultRegion

`func (o *Settings) GetUploadS3AwsDefaultRegion() string`

GetUploadS3AwsDefaultRegion returns the UploadS3AwsDefaultRegion field if non-nil, zero value otherwise.

### GetUploadS3AwsDefaultRegionOk

`func (o *Settings) GetUploadS3AwsDefaultRegionOk() (*string, bool)`

GetUploadS3AwsDefaultRegionOk returns a tuple with the UploadS3AwsDefaultRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadS3AwsDefaultRegion

`func (o *Settings) SetUploadS3AwsDefaultRegion(v string)`

SetUploadS3AwsDefaultRegion sets UploadS3AwsDefaultRegion field to given value.

### HasUploadS3AwsDefaultRegion

`func (o *Settings) HasUploadS3AwsDefaultRegion() bool`

HasUploadS3AwsDefaultRegion returns a boolean if a field has been set.

### GetUploadS3Bucket

`func (o *Settings) GetUploadS3Bucket() string`

GetUploadS3Bucket returns the UploadS3Bucket field if non-nil, zero value otherwise.

### GetUploadS3BucketOk

`func (o *Settings) GetUploadS3BucketOk() (*string, bool)`

GetUploadS3BucketOk returns a tuple with the UploadS3Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadS3Bucket

`func (o *Settings) SetUploadS3Bucket(v string)`

SetUploadS3Bucket sets UploadS3Bucket field to given value.

### HasUploadS3Bucket

`func (o *Settings) HasUploadS3Bucket() bool`

HasUploadS3Bucket returns a boolean if a field has been set.

### GetUploadS3BucketDomain

`func (o *Settings) GetUploadS3BucketDomain() string`

GetUploadS3BucketDomain returns the UploadS3BucketDomain field if non-nil, zero value otherwise.

### GetUploadS3BucketDomainOk

`func (o *Settings) GetUploadS3BucketDomainOk() (*string, bool)`

GetUploadS3BucketDomainOk returns a tuple with the UploadS3BucketDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadS3BucketDomain

`func (o *Settings) SetUploadS3BucketDomain(v string)`

SetUploadS3BucketDomain sets UploadS3BucketDomain field to given value.

### HasUploadS3BucketDomain

`func (o *Settings) HasUploadS3BucketDomain() bool`

HasUploadS3BucketDomain returns a boolean if a field has been set.

### GetUploadS3BucketPath

`func (o *Settings) GetUploadS3BucketPath() string`

GetUploadS3BucketPath returns the UploadS3BucketPath field if non-nil, zero value otherwise.

### GetUploadS3BucketPathOk

`func (o *Settings) GetUploadS3BucketPathOk() (*string, bool)`

GetUploadS3BucketPathOk returns a tuple with the UploadS3BucketPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadS3BucketPath

`func (o *Settings) SetUploadS3BucketPath(v string)`

SetUploadS3BucketPath sets UploadS3BucketPath field to given value.

### HasUploadS3BucketPath

`func (o *Settings) HasUploadS3BucketPath() bool`

HasUploadS3BucketPath returns a boolean if a field has been set.

### GetUploadS3BucketType

`func (o *Settings) GetUploadS3BucketType() string`

GetUploadS3BucketType returns the UploadS3BucketType field if non-nil, zero value otherwise.

### GetUploadS3BucketTypeOk

`func (o *Settings) GetUploadS3BucketTypeOk() (*string, bool)`

GetUploadS3BucketTypeOk returns a tuple with the UploadS3BucketType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadS3BucketType

`func (o *Settings) SetUploadS3BucketType(v string)`

SetUploadS3BucketType sets UploadS3BucketType field to given value.

### HasUploadS3BucketType

`func (o *Settings) HasUploadS3BucketType() bool`

HasUploadS3BucketType returns a boolean if a field has been set.

### GetUploadS3Expiry

`func (o *Settings) GetUploadS3Expiry() string`

GetUploadS3Expiry returns the UploadS3Expiry field if non-nil, zero value otherwise.

### GetUploadS3ExpiryOk

`func (o *Settings) GetUploadS3ExpiryOk() (*string, bool)`

GetUploadS3ExpiryOk returns a tuple with the UploadS3Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadS3Expiry

`func (o *Settings) SetUploadS3Expiry(v string)`

SetUploadS3Expiry sets UploadS3Expiry field to given value.

### HasUploadS3Expiry

`func (o *Settings) HasUploadS3Expiry() bool`

HasUploadS3Expiry returns a boolean if a field has been set.

### GetSmtp

`func (o *Settings) GetSmtp() []SMTPSettings`

GetSmtp returns the Smtp field if non-nil, zero value otherwise.

### GetSmtpOk

`func (o *Settings) GetSmtpOk() (*[]SMTPSettings, bool)`

GetSmtpOk returns a tuple with the Smtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtp

`func (o *Settings) SetSmtp(v []SMTPSettings)`

SetSmtp sets Smtp field to given value.

### HasSmtp

`func (o *Settings) HasSmtp() bool`

HasSmtp returns a boolean if a field has been set.

### GetMessengers

`func (o *Settings) GetMessengers() []map[string]interface{}`

GetMessengers returns the Messengers field if non-nil, zero value otherwise.

### GetMessengersOk

`func (o *Settings) GetMessengersOk() (*[]map[string]interface{}, bool)`

GetMessengersOk returns a tuple with the Messengers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessengers

`func (o *Settings) SetMessengers(v []map[string]interface{})`

SetMessengers sets Messengers field to given value.

### HasMessengers

`func (o *Settings) HasMessengers() bool`

HasMessengers returns a boolean if a field has been set.

### GetBounceEnabled

`func (o *Settings) GetBounceEnabled() bool`

GetBounceEnabled returns the BounceEnabled field if non-nil, zero value otherwise.

### GetBounceEnabledOk

`func (o *Settings) GetBounceEnabledOk() (*bool, bool)`

GetBounceEnabledOk returns a tuple with the BounceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceEnabled

`func (o *Settings) SetBounceEnabled(v bool)`

SetBounceEnabled sets BounceEnabled field to given value.

### HasBounceEnabled

`func (o *Settings) HasBounceEnabled() bool`

HasBounceEnabled returns a boolean if a field has been set.

### GetBounceWebhooksEnabled

`func (o *Settings) GetBounceWebhooksEnabled() bool`

GetBounceWebhooksEnabled returns the BounceWebhooksEnabled field if non-nil, zero value otherwise.

### GetBounceWebhooksEnabledOk

`func (o *Settings) GetBounceWebhooksEnabledOk() (*bool, bool)`

GetBounceWebhooksEnabledOk returns a tuple with the BounceWebhooksEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceWebhooksEnabled

`func (o *Settings) SetBounceWebhooksEnabled(v bool)`

SetBounceWebhooksEnabled sets BounceWebhooksEnabled field to given value.

### HasBounceWebhooksEnabled

`func (o *Settings) HasBounceWebhooksEnabled() bool`

HasBounceWebhooksEnabled returns a boolean if a field has been set.

### GetBounceCount

`func (o *Settings) GetBounceCount() int32`

GetBounceCount returns the BounceCount field if non-nil, zero value otherwise.

### GetBounceCountOk

`func (o *Settings) GetBounceCountOk() (*int32, bool)`

GetBounceCountOk returns a tuple with the BounceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceCount

`func (o *Settings) SetBounceCount(v int32)`

SetBounceCount sets BounceCount field to given value.

### HasBounceCount

`func (o *Settings) HasBounceCount() bool`

HasBounceCount returns a boolean if a field has been set.

### GetBounceAction

`func (o *Settings) GetBounceAction() string`

GetBounceAction returns the BounceAction field if non-nil, zero value otherwise.

### GetBounceActionOk

`func (o *Settings) GetBounceActionOk() (*string, bool)`

GetBounceActionOk returns a tuple with the BounceAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceAction

`func (o *Settings) SetBounceAction(v string)`

SetBounceAction sets BounceAction field to given value.

### HasBounceAction

`func (o *Settings) HasBounceAction() bool`

HasBounceAction returns a boolean if a field has been set.

### GetBounceSesEnabled

`func (o *Settings) GetBounceSesEnabled() bool`

GetBounceSesEnabled returns the BounceSesEnabled field if non-nil, zero value otherwise.

### GetBounceSesEnabledOk

`func (o *Settings) GetBounceSesEnabledOk() (*bool, bool)`

GetBounceSesEnabledOk returns a tuple with the BounceSesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceSesEnabled

`func (o *Settings) SetBounceSesEnabled(v bool)`

SetBounceSesEnabled sets BounceSesEnabled field to given value.

### HasBounceSesEnabled

`func (o *Settings) HasBounceSesEnabled() bool`

HasBounceSesEnabled returns a boolean if a field has been set.

### GetBounceSendgridEnabled

`func (o *Settings) GetBounceSendgridEnabled() bool`

GetBounceSendgridEnabled returns the BounceSendgridEnabled field if non-nil, zero value otherwise.

### GetBounceSendgridEnabledOk

`func (o *Settings) GetBounceSendgridEnabledOk() (*bool, bool)`

GetBounceSendgridEnabledOk returns a tuple with the BounceSendgridEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceSendgridEnabled

`func (o *Settings) SetBounceSendgridEnabled(v bool)`

SetBounceSendgridEnabled sets BounceSendgridEnabled field to given value.

### HasBounceSendgridEnabled

`func (o *Settings) HasBounceSendgridEnabled() bool`

HasBounceSendgridEnabled returns a boolean if a field has been set.

### GetBounceSendgridKey

`func (o *Settings) GetBounceSendgridKey() string`

GetBounceSendgridKey returns the BounceSendgridKey field if non-nil, zero value otherwise.

### GetBounceSendgridKeyOk

`func (o *Settings) GetBounceSendgridKeyOk() (*string, bool)`

GetBounceSendgridKeyOk returns a tuple with the BounceSendgridKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceSendgridKey

`func (o *Settings) SetBounceSendgridKey(v string)`

SetBounceSendgridKey sets BounceSendgridKey field to given value.

### HasBounceSendgridKey

`func (o *Settings) HasBounceSendgridKey() bool`

HasBounceSendgridKey returns a boolean if a field has been set.

### GetBounceMailboxes

`func (o *Settings) GetBounceMailboxes() []MailBoxBounces`

GetBounceMailboxes returns the BounceMailboxes field if non-nil, zero value otherwise.

### GetBounceMailboxesOk

`func (o *Settings) GetBounceMailboxesOk() (*[]MailBoxBounces, bool)`

GetBounceMailboxesOk returns a tuple with the BounceMailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceMailboxes

`func (o *Settings) SetBounceMailboxes(v []MailBoxBounces)`

SetBounceMailboxes sets BounceMailboxes field to given value.

### HasBounceMailboxes

`func (o *Settings) HasBounceMailboxes() bool`

HasBounceMailboxes returns a boolean if a field has been set.

### GetAppearanceAdminCustomCss

`func (o *Settings) GetAppearanceAdminCustomCss() string`

GetAppearanceAdminCustomCss returns the AppearanceAdminCustomCss field if non-nil, zero value otherwise.

### GetAppearanceAdminCustomCssOk

`func (o *Settings) GetAppearanceAdminCustomCssOk() (*string, bool)`

GetAppearanceAdminCustomCssOk returns a tuple with the AppearanceAdminCustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearanceAdminCustomCss

`func (o *Settings) SetAppearanceAdminCustomCss(v string)`

SetAppearanceAdminCustomCss sets AppearanceAdminCustomCss field to given value.

### HasAppearanceAdminCustomCss

`func (o *Settings) HasAppearanceAdminCustomCss() bool`

HasAppearanceAdminCustomCss returns a boolean if a field has been set.

### GetAppearanceAdminCustomJs

`func (o *Settings) GetAppearanceAdminCustomJs() string`

GetAppearanceAdminCustomJs returns the AppearanceAdminCustomJs field if non-nil, zero value otherwise.

### GetAppearanceAdminCustomJsOk

`func (o *Settings) GetAppearanceAdminCustomJsOk() (*string, bool)`

GetAppearanceAdminCustomJsOk returns a tuple with the AppearanceAdminCustomJs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearanceAdminCustomJs

`func (o *Settings) SetAppearanceAdminCustomJs(v string)`

SetAppearanceAdminCustomJs sets AppearanceAdminCustomJs field to given value.

### HasAppearanceAdminCustomJs

`func (o *Settings) HasAppearanceAdminCustomJs() bool`

HasAppearanceAdminCustomJs returns a boolean if a field has been set.

### GetAppearancePublicCustomCss

`func (o *Settings) GetAppearancePublicCustomCss() string`

GetAppearancePublicCustomCss returns the AppearancePublicCustomCss field if non-nil, zero value otherwise.

### GetAppearancePublicCustomCssOk

`func (o *Settings) GetAppearancePublicCustomCssOk() (*string, bool)`

GetAppearancePublicCustomCssOk returns a tuple with the AppearancePublicCustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearancePublicCustomCss

`func (o *Settings) SetAppearancePublicCustomCss(v string)`

SetAppearancePublicCustomCss sets AppearancePublicCustomCss field to given value.

### HasAppearancePublicCustomCss

`func (o *Settings) HasAppearancePublicCustomCss() bool`

HasAppearancePublicCustomCss returns a boolean if a field has been set.

### GetAppearancePublicCustomJs

`func (o *Settings) GetAppearancePublicCustomJs() string`

GetAppearancePublicCustomJs returns the AppearancePublicCustomJs field if non-nil, zero value otherwise.

### GetAppearancePublicCustomJsOk

`func (o *Settings) GetAppearancePublicCustomJsOk() (*string, bool)`

GetAppearancePublicCustomJsOk returns a tuple with the AppearancePublicCustomJs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearancePublicCustomJs

`func (o *Settings) SetAppearancePublicCustomJs(v string)`

SetAppearancePublicCustomJs sets AppearancePublicCustomJs field to given value.

### HasAppearancePublicCustomJs

`func (o *Settings) HasAppearancePublicCustomJs() bool`

HasAppearancePublicCustomJs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


