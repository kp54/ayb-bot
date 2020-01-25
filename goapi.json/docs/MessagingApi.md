# \MessagingApi

All URIs are relative to *https://misskey.kurume-nct.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IReadAllMessagingMessages**](MessagingApi.md#IReadAllMessagingMessages) | **Post** /i/read_all_messaging_messages | i/read_all_messaging_messages
[**IReadAllMessagingMessages_0**](MessagingApi.md#IReadAllMessagingMessages_0) | **Post** /i/read-all-messaging-messages | i/read-all-messaging-messages
[**MessagingHistory**](MessagingApi.md#MessagingHistory) | **Post** /messaging/history | messaging/history
[**MessagingMessages**](MessagingApi.md#MessagingMessages) | **Post** /messaging/messages | messaging/messages
[**MessagingMessagesCreate**](MessagingApi.md#MessagingMessagesCreate) | **Post** /messaging/messages/create | messaging/messages/create
[**MessagingMessagesDelete**](MessagingApi.md#MessagingMessagesDelete) | **Post** /messaging/messages/delete | messaging/messages/delete
[**MessagingMessagesRead**](MessagingApi.md#MessagingMessagesRead) | **Post** /messaging/messages/read | messaging/messages/read



## IReadAllMessagingMessages

> IReadAllMessagingMessages(ctx, body)

i/read_all_messaging_messages

トークメッセージをすべて既読にします。  **Credential required**: *Yes* / **Permission**: *account-write*

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


## IReadAllMessagingMessages_0

> IReadAllMessagingMessages_0(ctx, body)

i/read-all-messaging-messages

トークメッセージをすべて既読にします。  **Credential required**: *Yes* / **Permission**: *account-write*

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


## MessagingHistory

> []MessagingMessage MessagingHistory(ctx, inlineObject84)

messaging/history

Messagingの履歴を取得します。  **Credential required**: *Yes* / **Permission**: *messaging-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject84** | [**InlineObject84**](InlineObject84.md)|  | 

### Return type

[**[]MessagingMessage**](MessagingMessage.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MessagingMessages

> []MessagingMessage MessagingMessages(ctx, inlineObject85)

messaging/messages

指定したユーザーとのMessagingのメッセージ一覧を取得します。  **Credential required**: *Yes* / **Permission**: *messaging-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject85** | [**InlineObject85**](InlineObject85.md)|  | 

### Return type

[**[]MessagingMessage**](MessagingMessage.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MessagingMessagesCreate

> MessagingMessage MessagingMessagesCreate(ctx, inlineObject86)

messaging/messages/create

指定したユーザーへMessagingのメッセージを送信します。  **Credential required**: *Yes* / **Permission**: *messaging-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject86** | [**InlineObject86**](InlineObject86.md)|  | 

### Return type

[**MessagingMessage**](MessagingMessage.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MessagingMessagesDelete

> MessagingMessagesDelete(ctx, inlineObject87)

messaging/messages/delete

指定したメッセージを削除します。  **Credential required**: *Yes* / **Permission**: *messaging-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject87** | [**InlineObject87**](InlineObject87.md)|  | 

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


## MessagingMessagesRead

> MessagingMessagesRead(ctx, inlineObject88)

messaging/messages/read

指定した自分宛てのメッセージを既読にします。  **Credential required**: *Yes* / **Permission**: *messaging-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject88** | [**InlineObject88**](InlineObject88.md)|  | 

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

