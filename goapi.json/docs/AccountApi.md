# \AccountApi

All URIs are relative to *https://misskey.kurume-nct.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlockingList**](AccountApi.md#BlockingList) | **Post** /blocking/list | blocking/list
[**Drive**](AccountApi.md#Drive) | **Post** /drive | drive
[**FollowingRequestsAccept**](AccountApi.md#FollowingRequestsAccept) | **Post** /following/requests/accept | following/requests/accept
[**FollowingRequestsCancel**](AccountApi.md#FollowingRequestsCancel) | **Post** /following/requests/cancel | following/requests/cancel
[**FollowingRequestsList**](AccountApi.md#FollowingRequestsList) | **Post** /following/requests/list | following/requests/list
[**FollowingRequestsReject**](AccountApi.md#FollowingRequestsReject) | **Post** /following/requests/reject | following/requests/reject
[**I**](AccountApi.md#I) | **Post** /i | i
[**IClearFollowRequestNotification**](AccountApi.md#IClearFollowRequestNotification) | **Post** /i/clear-follow-request-notification | i/clear-follow-request-notification
[**IFavorites**](AccountApi.md#IFavorites) | **Post** /i/favorites | i/favorites
[**INotifications**](AccountApi.md#INotifications) | **Post** /i/notifications | i/notifications
[**IPin**](AccountApi.md#IPin) | **Post** /i/pin | i/pin
[**IReadAllMessagingMessages**](AccountApi.md#IReadAllMessagingMessages) | **Post** /i/read_all_messaging_messages | i/read_all_messaging_messages
[**IReadAllMessagingMessages_0**](AccountApi.md#IReadAllMessagingMessages_0) | **Post** /i/read-all-messaging-messages | i/read-all-messaging-messages
[**IReadAllUnreadNotes**](AccountApi.md#IReadAllUnreadNotes) | **Post** /i/read_all_unread_notes | i/read_all_unread_notes
[**IReadAllUnreadNotes_0**](AccountApi.md#IReadAllUnreadNotes_0) | **Post** /i/read-all-unread-notes | i/read-all-unread-notes
[**IUnpin**](AccountApi.md#IUnpin) | **Post** /i/unpin | i/unpin
[**IUpdate**](AccountApi.md#IUpdate) | **Post** /i/update | i/update
[**MuteList**](AccountApi.md#MuteList) | **Post** /mute/list | mute/list
[**MyApps**](AccountApi.md#MyApps) | **Post** /my/apps | my/apps
[**NotificationsMarkAllAsRead**](AccountApi.md#NotificationsMarkAllAsRead) | **Post** /notifications/mark_all_as_read | notifications/mark_all_as_read
[**NotificationsMarkAllAsRead_0**](AccountApi.md#NotificationsMarkAllAsRead_0) | **Post** /notifications/mark-all-as-read | notifications/mark-all-as-read
[**SwRegister**](AccountApi.md#SwRegister) | **Post** /sw/register | sw/register
[**UsersListsList**](AccountApi.md#UsersListsList) | **Post** /users/lists/list | users/lists/list
[**UsersListsShow**](AccountApi.md#UsersListsShow) | **Post** /users/lists/show | users/lists/show



## BlockingList

> []Blocking BlockingList(ctx, inlineObject33)

blocking/list

ブロックしているユーザー一覧を取得します。  **Credential required**: *Yes* / **Permission**: *following-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject33** | [**InlineObject33**](InlineObject33.md)|  | 

### Return type

[**[]Blocking**](Blocking.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Drive

> InlineResponse2006 Drive(ctx, body)

drive

ドライブの情報を取得します。  **Credential required**: *Yes* / **Permission**: *drive-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **map[string]interface{}**|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

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


## I

> User I(ctx, body)

i

自分のアカウント情報を取得します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **map[string]interface{}**|  | 

### Return type

[**User**](User.md)

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


## IFavorites

> IFavorites(ctx, inlineObject79)

i/favorites

お気に入りに登録した投稿一覧を取得します。  **Credential required**: *Yes* / **Permission**: *favorites-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject79** | [**InlineObject79**](InlineObject79.md)|  | 

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


## IPin

> IPin(ctx, inlineObject81)

i/pin

指定した投稿をピン留めします。  **Credential required**: *Yes* / **Permission**: *account-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject81** | [**InlineObject81**](InlineObject81.md)|  | 

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


## IReadAllUnreadNotes

> IReadAllUnreadNotes(ctx, body)

i/read_all_unread_notes

未読の投稿をすべて既読にします。  **Credential required**: *Yes* / **Permission**: *account-write*

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


## IReadAllUnreadNotes_0

> IReadAllUnreadNotes_0(ctx, body)

i/read-all-unread-notes

未読の投稿をすべて既読にします。  **Credential required**: *Yes* / **Permission**: *account-write*

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


## IUnpin

> IUnpin(ctx, inlineObject82)

i/unpin

指定した投稿のピン留めを解除します。  **Credential required**: *Yes* / **Permission**: *account-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject82** | [**InlineObject82**](InlineObject82.md)|  | 

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


## IUpdate

> IUpdate(ctx, inlineObject83)

i/update

アカウント情報を更新します。  **Credential required**: *Yes* / **Permission**: *account-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject83** | [**InlineObject83**](InlineObject83.md)|  | 

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


## MuteList

> []Muting MuteList(ctx, inlineObject92)

mute/list

ミュートしているユーザー一覧を取得します。  **Credential required**: *Yes* / **Permission**: *account/read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject92** | [**InlineObject92**](InlineObject92.md)|  | 

### Return type

[**[]Muting**](Muting.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MyApps

> MyApps(ctx, inlineObject93)

my/apps

自分のアプリケーション一覧を取得します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject93** | [**InlineObject93**](InlineObject93.md)|  | 

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


## SwRegister

> SwRegister(ctx, inlineObject122)

sw/register

No description provided.  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject122** | [**InlineObject122**](InlineObject122.md)|  | 

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


## UsersListsList

> []UserList UsersListsList(ctx, body)

users/lists/list

自分の作成したユーザーリスト一覧を取得します。  **Credential required**: *Yes* / **Permission**: *account-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **map[string]interface{}**|  | 

### Return type

[**[]UserList**](UserList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListsShow

> UserList UsersListsShow(ctx, inlineObject133)

users/lists/show

指定したユーザーリストの情報を取得します。  **Credential required**: *Yes* / **Permission**: *account-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject133** | [**InlineObject133**](InlineObject133.md)|  | 

### Return type

[**UserList**](UserList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

