# \NotificationsApi

All URIs are relative to *https://misskey.kurume-nct.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**INotifications**](NotificationsApi.md#INotifications) | **Post** /i/notifications | i/notifications
[**NotificationsMarkAllAsRead**](NotificationsApi.md#NotificationsMarkAllAsRead) | **Post** /notifications/mark_all_as_read | notifications/mark_all_as_read
[**NotificationsMarkAllAsRead_0**](NotificationsApi.md#NotificationsMarkAllAsRead_0) | **Post** /notifications/mark-all-as-read | notifications/mark-all-as-read



## INotifications

> []Notification INotifications(ctx, inlineObject80)

i/notifications

通知一覧を取得します。  **Credential required**: *Yes* / **Permission**: *account-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject80** | [**InlineObject80**](InlineObject80.md)|  | 

### Return type

[**[]Notification**](Notification.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationsMarkAllAsRead

> NotificationsMarkAllAsRead(ctx, body)

notifications/mark_all_as_read

全ての通知を既読にします。  **Credential required**: *Yes* / **Permission**: *notification-write*

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


## NotificationsMarkAllAsRead_0

> NotificationsMarkAllAsRead_0(ctx, body)

notifications/mark-all-as-read

全ての通知を既読にします。  **Credential required**: *Yes* / **Permission**: *notification-write*

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

