# \FollowingApi

All URIs are relative to *https://misskey.kurume-nct.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChartsUserFollowing**](FollowingApi.md#ChartsUserFollowing) | **Post** /charts/user/following | charts/user/following
[**FollowingCreate**](FollowingApi.md#FollowingCreate) | **Post** /following/create | following/create
[**FollowingDelete**](FollowingApi.md#FollowingDelete) | **Post** /following/delete | following/delete
[**FollowingRequestsAccept**](FollowingApi.md#FollowingRequestsAccept) | **Post** /following/requests/accept | following/requests/accept
[**FollowingRequestsCancel**](FollowingApi.md#FollowingRequestsCancel) | **Post** /following/requests/cancel | following/requests/cancel
[**FollowingRequestsList**](FollowingApi.md#FollowingRequestsList) | **Post** /following/requests/list | following/requests/list
[**FollowingRequestsReject**](FollowingApi.md#FollowingRequestsReject) | **Post** /following/requests/reject | following/requests/reject
[**IClearFollowRequestNotification**](FollowingApi.md#IClearFollowRequestNotification) | **Post** /i/clear-follow-request-notification | i/clear-follow-request-notification



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


## FollowingCreate

> FollowingCreate(ctx, inlineObject67)

following/create

指定したユーザーをフォローします。  **Credential required**: *Yes* / **Permission**: *following-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject67** | [**InlineObject67**](InlineObject67.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FollowingDelete

> FollowingDelete(ctx, inlineObject68)

following/delete

指定したユーザーのフォローを解除します。  **Credential required**: *Yes* / **Permission**: *following-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject68** | [**InlineObject68**](InlineObject68.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FollowingRequestsAccept

> FollowingRequestsAccept(ctx, inlineObject69)

following/requests/accept

自分に届いた、指定したフォローリクエストを承認します。  **Credential required**: *Yes* / **Permission**: *following-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject69** | [**InlineObject69**](InlineObject69.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FollowingRequestsCancel

> FollowingRequestsCancel(ctx, inlineObject70)

following/requests/cancel

自分が作成した、指定したフォローリクエストをキャンセルします。  **Credential required**: *Yes* / **Permission**: *following-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject70** | [**InlineObject70**](InlineObject70.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FollowingRequestsList

> FollowingRequestsList(ctx, body)

following/requests/list

自分に届いたフォローリクエストの一覧を取得します。  **Credential required**: *Yes* / **Permission**: *following-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **map[string]interface{}**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FollowingRequestsReject

> FollowingRequestsReject(ctx, inlineObject71)

following/requests/reject

自分に届いた、指定したフォローリクエストを拒否します。  **Credential required**: *Yes* / **Permission**: *following-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject71** | [**InlineObject71**](InlineObject71.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IClearFollowRequestNotification

> IClearFollowRequestNotification(ctx, body)

i/clear-follow-request-notification

No description provided.  **Credential required**: *Yes* / **Permission**: *account-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **map[string]interface{}**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

