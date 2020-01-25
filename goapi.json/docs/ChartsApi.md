# \ChartsApi

All URIs are relative to *https://misskey.kurume-nct.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChartsActiveUsers**](ChartsApi.md#ChartsActiveUsers) | **Post** /charts/active-users | charts/active-users
[**ChartsDrive**](ChartsApi.md#ChartsDrive) | **Post** /charts/drive | charts/drive
[**ChartsFederation**](ChartsApi.md#ChartsFederation) | **Post** /charts/federation | charts/federation
[**ChartsHashtag**](ChartsApi.md#ChartsHashtag) | **Post** /charts/hashtag | charts/hashtag
[**ChartsInstance**](ChartsApi.md#ChartsInstance) | **Post** /charts/instance | charts/instance
[**ChartsNetwork**](ChartsApi.md#ChartsNetwork) | **Post** /charts/network | charts/network
[**ChartsNotes**](ChartsApi.md#ChartsNotes) | **Post** /charts/notes | charts/notes
[**ChartsUserDrive**](ChartsApi.md#ChartsUserDrive) | **Post** /charts/user/drive | charts/user/drive
[**ChartsUserFollowing**](ChartsApi.md#ChartsUserFollowing) | **Post** /charts/user/following | charts/user/following
[**ChartsUserNotes**](ChartsApi.md#ChartsUserNotes) | **Post** /charts/user/notes | charts/user/notes
[**ChartsUserReactions**](ChartsApi.md#ChartsUserReactions) | **Post** /charts/user/reactions | charts/user/reactions
[**ChartsUsers**](ChartsApi.md#ChartsUsers) | **Post** /charts/users | charts/users



## ChartsActiveUsers

> []map[string]interface{} ChartsActiveUsers(ctx, inlineObject34)

charts/active-users

アクティブユーザーのチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject34** | [**InlineObject34**](InlineObject34.md)|  | 

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChartsDrive

> InlineResponse2002 ChartsDrive(ctx, inlineObject35)

charts/drive

ドライブのチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject35** | [**InlineObject35**](InlineObject35.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChartsFederation

> []map[string]interface{} ChartsFederation(ctx, inlineObject36)

charts/federation

フェデレーションのチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject36** | [**InlineObject36**](InlineObject36.md)|  | 

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChartsHashtag

> []map[string]interface{} ChartsHashtag(ctx, inlineObject37)

charts/hashtag

ハッシュタグごとのチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject37** | [**InlineObject37**](InlineObject37.md)|  | 

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChartsInstance

> []map[string]interface{} ChartsInstance(ctx, inlineObject38)

charts/instance

インスタンスごとのチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject38** | [**InlineObject38**](InlineObject38.md)|  | 

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChartsNetwork

> []map[string]interface{} ChartsNetwork(ctx, inlineObject39)

charts/network

ネットワークのチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject39** | [**InlineObject39**](InlineObject39.md)|  | 

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChartsNotes

> InlineResponse2003 ChartsNotes(ctx, inlineObject40)

charts/notes

投稿のチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject40** | [**InlineObject40**](InlineObject40.md)|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChartsUserDrive

> InlineResponse2002Local ChartsUserDrive(ctx, inlineObject41)

charts/user/drive

ユーザーごとのドライブのチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject41** | [**InlineObject41**](InlineObject41.md)|  | 

### Return type

[**InlineResponse2002Local**](inline_response_200_2_local.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChartsUserFollowing

> InlineResponse2004 ChartsUserFollowing(ctx, inlineObject42)

charts/user/following

ユーザーごとのフォロー/フォロワーのチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject42** | [**InlineObject42**](InlineObject42.md)|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChartsUserNotes

> InlineResponse2003Local ChartsUserNotes(ctx, inlineObject43)

charts/user/notes

ユーザーごとの投稿のチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject43** | [**InlineObject43**](InlineObject43.md)|  | 

### Return type

[**InlineResponse2003Local**](inline_response_200_3_local.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChartsUserReactions

> []map[string]interface{} ChartsUserReactions(ctx, inlineObject44)

charts/user/reactions

ユーザーごとの被リアクション数のチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject44** | [**InlineObject44**](InlineObject44.md)|  | 

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChartsUsers

> InlineResponse2005 ChartsUsers(ctx, inlineObject45)

charts/users

ユーザーのチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject45** | [**InlineObject45**](InlineObject45.md)|  | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

